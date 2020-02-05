package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RanHelloWorldSpec defines the desired state of RanHelloWorld
type RanHelloWorldSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
        Message string `json:"message"`
}

// RanHelloWorldStatus defines the observed state of RanHelloWorld
type RanHelloWorldStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
        Message string `json:"message"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RanHelloWorld is the Schema for the ranhelloworlds API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=ranhelloworlds,scope=Namespaced
type RanHelloWorld struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RanHelloWorldSpec   `json:"spec,omitempty"`
	Status RanHelloWorldStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RanHelloWorldList contains a list of RanHelloWorld
type RanHelloWorldList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RanHelloWorld `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RanHelloWorld{}, &RanHelloWorldList{})
}
