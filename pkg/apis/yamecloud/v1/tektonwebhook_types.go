package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type Params struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type Job struct {
	// +optional
	Branch string `json:"branch,omitempty"`
	// +optional
	PipelineRun string `json:"pipeline_run,omitempty"`
	// +optional
	Params []Params `json:"params,omitempty"`
}

// TektonWebHookSpec defines the desired state of TektonWebHook
type TektonWebHookSpec struct {
	// +optional
	Secret string `json:"secret,omitempty"`
	// +optional
	Git string `json:"git,omitempty"`
	// +optional
	Jobs []Job `json:"jobs,omitempty"`
}

// TektonWebHookStatus defines the observed state of TektonWebHook
type TektonWebHookStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TektonWebHook is the Schema for the tektonwebhooks API
type TektonWebHook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TektonWebHookSpec   `json:"spec,omitempty"`
	Status TektonWebHookStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TektonWebHookList contains a list of TektonWebHook
type TektonWebHookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TektonWebHook `json:"items"`
}

func init() {
	register(&TektonWebHook{}, &TektonWebHookList{})
}
