package controllers

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	proomptengv1alpha1 "github.com/proompteng/proompteng/operator/api/v1alpha1"
)

// MemoryReconciler updates Memory status fields.
type MemoryReconciler struct {
	client.Client
	Scheme   *runtime.Scheme
	Recorder record.EventRecorder
}

func (r *MemoryReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	var memory proomptengv1alpha1.Memory
	if err := r.Get(ctx, req.NamespacedName, &memory); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if memory.Spec.Type == "" {
		r.Recorder.Event(&memory, corev1.EventTypeWarning, "InvalidSpec", "spec.type must be provided")
		logger.Info("memory spec.type missing", "name", req.NamespacedName.String())
		return ctrl.Result{}, nil
	}

	memory.Status.Phase = "Ready"
	memory.Status.Message = fmt.Sprintf("Validated memory target %s", truncate(memory.Spec.URI, 32))
	memory.Status.ObservedGeneration = memory.Generation

	if err := r.Status().Update(ctx, &memory); err != nil {
		logger.Error(err, "unable to update Memory status")
		return ctrl.Result{}, err
	}

	r.Recorder.Event(&memory, corev1.EventTypeNormal, "Synchronized", "Memory configuration validated")
	logger.Info("Memory reconciled", "name", req.NamespacedName.String())
	return ctrl.Result{}, nil
}

func (r *MemoryReconciler) SetupWithManager(mgr ctrl.Manager) error {
	if r.Recorder == nil {
		r.Recorder = mgr.GetEventRecorderFor("proompteng-memory-controller")
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(&proomptengv1alpha1.Memory{}).
		Complete(r)
}

func truncate(value string, length int) string {
	if len(value) <= length {
		return value
	}
	return value[:length] + "..."
}
