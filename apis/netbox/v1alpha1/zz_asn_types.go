/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AsnObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AsnParameters struct {

	// +kubebuilder:validation:Required
	Asn *float64 `json:"asn" tf:"asn,omitempty"`

	// +kubebuilder:validation:Required
	RirID *float64 `json:"rirId" tf:"rir_id,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// AsnSpec defines the desired state of Asn
type AsnSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AsnParameters `json:"forProvider"`
}

// AsnStatus defines the observed state of Asn.
type AsnStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AsnObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Asn is the Schema for the Asns API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,netboxjet}
type Asn struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AsnSpec   `json:"spec"`
	Status            AsnStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AsnList contains a list of Asns
type AsnList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Asn `json:"items"`
}

// Repository type metadata.
var (
	Asn_Kind             = "Asn"
	Asn_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Asn_Kind}.String()
	Asn_KindAPIVersion   = Asn_Kind + "." + CRDGroupVersion.String()
	Asn_GroupVersionKind = CRDGroupVersion.WithKind(Asn_Kind)
)

func init() {
	SchemeBuilder.Register(&Asn{}, &AsnList{})
}
