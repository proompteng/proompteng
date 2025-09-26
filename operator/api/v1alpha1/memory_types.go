package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// Memory describes a memory backend for agents.
type Memory struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MemorySpec   `json:"spec,omitempty"`
	Status MemoryStatus `json:"status,omitempty"`
}

// MemorySpec holds backend configuration fields.
type MemorySpec struct {
	Type                 string            `json:"type"`
	URI                  string            `json:"uri,omitempty"`
	CredentialsSecretRef string            `json:"credentialsSecretRef,omitempty"`
	Config               map[string]string `json:"config,omitempty"`
}

// MemoryStatus reports the readiness of the memory backend.
type MemoryStatus struct {
	Phase              string `json:"phase,omitempty"`
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	Message            string `json:"message,omitempty"`
}

// +kubebuilder:object:root=true
// MemoryList contains a list of Memory.
type MemoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Memory `json:"items"`
}

func (in *Memory) DeepCopyInto(out *Memory) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

func (in *Memory) DeepCopy() *Memory {
	if in == nil {
		return nil
	}
	out := new(Memory)
	in.DeepCopyInto(out)
	return out
}

func (in *Memory) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *MemoryList) DeepCopyInto(out *MemoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Memory, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

func (in *MemoryList) DeepCopy() *MemoryList {
	if in == nil {
		return nil
	}
	out := new(MemoryList)
	in.DeepCopyInto(out)
	return out
}

func (in *MemoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *MemorySpec) DeepCopyInto(out *MemorySpec) {
	*out = *in
	if in.Config != nil {
		out.Config = make(map[string]string, len(in.Config))
		for k, v := range in.Config {
			out.Config[k] = v
		}
	}
}

func (in *MemorySpec) DeepCopy() *MemorySpec {
	if in == nil {
		return nil
	}
	out := new(MemorySpec)
	in.DeepCopyInto(out)
	return out
}
