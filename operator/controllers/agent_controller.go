package controllers

import (
	"context"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"

	proomptengv1alpha1 "github.com/proompteng/proompteng/operator/api/v1alpha1"
)

const (
	defaultAgentPort = 8000
	metricsPort      = 9090
)

// AgentReconciler reconciles Agent resources.
type AgentReconciler struct {
	client.Client
	Scheme   *runtime.Scheme
	Recorder record.EventRecorder
}

// Reconcile ensures a Deployment and Service exist for each Agent resource.
func (r *AgentReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	var agent proomptengv1alpha1.Agent
	if err := r.Get(ctx, req.NamespacedName, &agent); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if agent.Spec.Runtime.Image == "" {
		r.Recorder.Event(&agent, corev1.EventTypeWarning, "InvalidSpec", "runtime.image must be provided")
		logger.Info("agent runtime.image missing", "name", req.NamespacedName.String())
		return ctrl.Result{}, nil
	}

	deploymentName := fmt.Sprintf("%s-runtime", agent.Name)
	deployment := &appsv1.Deployment{ObjectMeta: metav1.ObjectMeta{
		Name:      deploymentName,
		Namespace: agent.Namespace,
	}}

	_, err := controllerutil.CreateOrUpdate(ctx, r.Client, deployment, func() error {
		labels := agentRuntimeLabels(agent.Name)
		if deployment.Labels == nil {
			deployment.Labels = map[string]string{}
		}
		for k, v := range agentMetadataLabels(&agent) {
			deployment.Labels[k] = v
		}

		replicas := int32(1)
		deployment.Spec.Replicas = &replicas
		deployment.Spec.Selector = &metav1.LabelSelector{MatchLabels: labels}
		deployment.Spec.Template.ObjectMeta.Labels = labels
		deployment.Spec.Template.Spec.Containers = []corev1.Container{buildAgentContainer(&agent)}

		return controllerutil.SetControllerReference(&agent, deployment, r.Scheme)
	})
	if err != nil {
		logger.Error(err, "unable to ensure Deployment")
		return ctrl.Result{}, err
	}

	serviceName := fmt.Sprintf("%s-svc", agent.Name)
	service := &corev1.Service{ObjectMeta: metav1.ObjectMeta{
		Name:      serviceName,
		Namespace: agent.Namespace,
	}}

	_, err = controllerutil.CreateOrUpdate(ctx, r.Client, service, func() error {
		service.Spec.Selector = agentRuntimeLabels(agent.Name)
		service.Spec.Ports = []corev1.ServicePort{
			{
				Name:       "http",
				Port:       agentServicePort(agent.Spec.Runtime.Service),
				TargetPort: intstr.FromString("http"),
			},
			{
				Name:       "metrics",
				Port:       metricsPort,
				TargetPort: intstr.FromString("metrics"),
			},
		}
		if service.Labels == nil {
			service.Labels = map[string]string{}
		}
		for k, v := range agentMetadataLabels(&agent) {
			service.Labels[k] = v
		}
		return controllerutil.SetControllerReference(&agent, service, r.Scheme)
	})
	if err != nil {
		logger.Error(err, "unable to ensure Service")
		return ctrl.Result{}, err
	}

	// Re-fetch deployment to obtain status (CreateOrUpdate may not have populated it).
	if err := r.Get(ctx, types.NamespacedName{Name: deploymentName, Namespace: agent.Namespace}, deployment); err != nil {
		logger.Error(err, "unable to fetch Deployment status")
		return ctrl.Result{}, err
	}

	agent.Status.ReadyReplicas = deployment.Status.ReadyReplicas
	agent.Status.ServiceName = service.Name
	agent.Status.Phase = "Deployed"
	agent.Status.Conditions = appendOrUpdateCondition(agent.Status.Conditions, metav1.Condition{
		Type:               "Ready",
		Status:             metav1.ConditionTrue,
		Reason:             "RuntimeAvailable",
		Message:            "Agent runtime deployment is available",
		ObservedGeneration: agent.Generation,
		LastTransitionTime: metav1.Now(),
	})

	if err := r.Status().Update(ctx, &agent); err != nil {
		logger.Error(err, "unable to update Agent status")
		return ctrl.Result{}, err
	}

	r.Recorder.Event(&agent, corev1.EventTypeNormal, "Reconciled", fmt.Sprintf("Agent runtime deployed via %s", deploymentName))
	logger.Info("Agent reconciled", "name", req.NamespacedName.String())
	return ctrl.Result{}, nil
}

func (r *AgentReconciler) SetupWithManager(mgr ctrl.Manager) error {
	if r.Recorder == nil {
		r.Recorder = mgr.GetEventRecorderFor("proompteng-agent-controller")
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(&proomptengv1alpha1.Agent{}).
		Owns(&appsv1.Deployment{}).
		Owns(&corev1.Service{}).
		Complete(r)
}

func agentRuntimeLabels(name string) map[string]string {
	return map[string]string{
		"app.kubernetes.io/name":      name,
		"app.kubernetes.io/component": "agent-runtime",
		"app.kubernetes.io/part-of":   "proompteng",
	}
}

func agentMetadataLabels(agent *proomptengv1alpha1.Agent) map[string]string {
	labels := map[string]string{
		"app.kubernetes.io/name":       agent.Name,
		"app.kubernetes.io/component":  "agent-runtime",
		"app.kubernetes.io/instance":   agent.Name,
		"app.kubernetes.io/managed-by": "proompteng-operator",
		"app.kubernetes.io/part-of":    "proompteng",
	}
	for k, v := range agent.Labels {
		if _, exists := labels[k]; !exists {
			labels[k] = v
		}
	}
	return labels
}

func buildAgentContainer(agent *proomptengv1alpha1.Agent) corev1.Container {
	port := agentServicePort(agent.Spec.Runtime.Service)
	env := []corev1.EnvVar{
		{Name: "AGENT_NAME", Value: agent.Name},
		{Name: "MODEL_PROVIDER", Value: agent.Spec.Model.Provider},
		{Name: "MODEL_NAME", Value: agent.Spec.Model.Name},
	}
	for _, e := range agent.Spec.Runtime.Env {
		env = append(env, corev1.EnvVar{Name: e.Name, Value: e.Value})
	}

	return corev1.Container{
		Name:  "agent",
		Image: agent.Spec.Runtime.Image,
		Ports: []corev1.ContainerPort{
			{Name: "http", ContainerPort: port},
			{Name: "metrics", ContainerPort: metricsPort},
		},
		Env: env,
		LivenessProbe: &corev1.Probe{
			ProbeHandler: corev1.ProbeHandler{
				HTTPGet: &corev1.HTTPGetAction{Path: "/healthz", Port: intstr.FromString("http")},
			},
			InitialDelaySeconds: 5,
			PeriodSeconds:       10,
		},
		ReadinessProbe: &corev1.Probe{
			ProbeHandler: corev1.ProbeHandler{
				HTTPGet: &corev1.HTTPGetAction{Path: "/readyz", Port: intstr.FromString("http")},
			},
			InitialDelaySeconds: 3,
			PeriodSeconds:       10,
		},
	}
}

func agentServicePort(cfg *proomptengv1alpha1.AgentServiceConfig) int32 {
	if cfg != nil && cfg.Port != 0 {
		return cfg.Port
	}
	return defaultAgentPort
}

func appendOrUpdateCondition(conditions []metav1.Condition, newCond metav1.Condition) []metav1.Condition {
	for idx, cond := range conditions {
		if cond.Type == newCond.Type {
			conditions[idx] = newCond
			return conditions
		}
	}
	return append(conditions, newCond)
}
