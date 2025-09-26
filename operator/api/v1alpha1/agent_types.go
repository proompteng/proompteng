package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// Agent defines a deployable AI agent managed by the operator.
type Agent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AgentSpec   `json:"spec,omitempty"`
	Status AgentStatus `json:"status,omitempty"`
}

// AgentSpec captures desired agent configuration.
type AgentSpec struct {
	DisplayName string           `json:"displayName,omitempty"`
	Model       AgentModelSpec   `json:"model"`
	Runtime     AgentRuntimeSpec `json:"runtime"`
	Transport   *AgentTransport  `json:"transport,omitempty"`
	MemoryRefs  []string         `json:"memoryRefs,omitempty"`
	ToolRefs    []string         `json:"toolRefs,omitempty"`
}

// AgentModelSpec identifies the backing model.
type AgentModelSpec struct {
	Provider   string            `json:"provider"`
	Name       string            `json:"name"`
	Parameters map[string]string `json:"parameters,omitempty"`
}

// AgentRuntimeSpec contains runtime image hints.
type AgentRuntimeSpec struct {
	Image   string              `json:"image"`
	Env     []AgentEnvVar       `json:"env,omitempty"`
	Service *AgentServiceConfig `json:"service,omitempty"`
}

// AgentEnvVar models a simple env var key/value.
type AgentEnvVar struct {
	Name  string `json:"name"`
	Value string `json:"value,omitempty"`
}

// AgentServiceConfig describes optional service port info.
type AgentServiceConfig struct {
	Port int32  `json:"port,omitempty"`
	Path string `json:"path,omitempty"`
}

// AgentTransport describes how the agent is exposed.
type AgentTransport struct {
	Protocol string `json:"protocol,omitempty"`
	Endpoint string `json:"endpoint,omitempty"`
}

// AgentStatus reports basic reconciliation state.
type AgentStatus struct {
	Phase         string             `json:"phase,omitempty"`
	ServiceName   string             `json:"serviceName,omitempty"`
	ReadyReplicas int32              `json:"readyReplicas,omitempty"`
	Conditions    []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// AgentList contains a list of Agent.
type AgentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Agent `json:"items"`
}

func (in *Agent) DeepCopyInto(out *Agent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

func (in *Agent) DeepCopy() *Agent {
	if in == nil {
		return nil
	}
	out := new(Agent)
	in.DeepCopyInto(out)
	return out
}

func (in *Agent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *AgentList) DeepCopyInto(out *AgentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Agent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

func (in *AgentList) DeepCopy() *AgentList {
	if in == nil {
		return nil
	}
	out := new(AgentList)
	in.DeepCopyInto(out)
	return out
}

func (in *AgentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *AgentSpec) DeepCopyInto(out *AgentSpec) {
	*out = *in
	if in.Model.Parameters != nil {
		out.Model.Parameters = make(map[string]string, len(in.Model.Parameters))
		for k, v := range in.Model.Parameters {
			out.Model.Parameters[k] = v
		}
	}
	if in.Runtime.Env != nil {
		out.Runtime.Env = make([]AgentEnvVar, len(in.Runtime.Env))
		copy(out.Runtime.Env, in.Runtime.Env)
	}
	if in.MemoryRefs != nil {
		out.MemoryRefs = make([]string, len(in.MemoryRefs))
		copy(out.MemoryRefs, in.MemoryRefs)
	}
	if in.ToolRefs != nil {
		out.ToolRefs = make([]string, len(in.ToolRefs))
		copy(out.ToolRefs, in.ToolRefs)
	}
	if in.Transport != nil {
		transport := *in.Transport
		out.Transport = &transport
	}
	if in.Runtime.Service != nil {
		service := *in.Runtime.Service
		out.Runtime.Service = &service
	}
}

func (in *AgentSpec) DeepCopy() *AgentSpec {
	if in == nil {
		return nil
	}
	out := new(AgentSpec)
	in.DeepCopyInto(out)
	return out
}

func (in *AgentStatus) DeepCopyInto(out *AgentStatus) {
	*out = *in
	if in.Conditions != nil {
		out.Conditions = make([]metav1.Condition, len(in.Conditions))
		copy(out.Conditions, in.Conditions)
	}
}

func (in *AgentStatus) DeepCopy() *AgentStatus {
	if in == nil {
		return nil
	}
	out := new(AgentStatus)
	in.DeepCopyInto(out)
	return out
}
