// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SpringCloudAcceleratorInitParameters struct {

	// The name which should be used for this Spring Cloud Accelerator. Changing this forces a new Spring Cloud Accelerator to be created. The only possible value is default.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the Spring Cloud Service. Changing this forces a new Spring Cloud Accelerator to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/appplatform/v1beta2.SpringCloudService
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	SpringCloudServiceID *string `json:"springCloudServiceId,omitempty" tf:"spring_cloud_service_id,omitempty"`

	// Reference to a SpringCloudService in appplatform to populate springCloudServiceId.
	// +kubebuilder:validation:Optional
	SpringCloudServiceIDRef *v1.Reference `json:"springCloudServiceIdRef,omitempty" tf:"-"`

	// Selector for a SpringCloudService in appplatform to populate springCloudServiceId.
	// +kubebuilder:validation:Optional
	SpringCloudServiceIDSelector *v1.Selector `json:"springCloudServiceIdSelector,omitempty" tf:"-"`
}

type SpringCloudAcceleratorObservation struct {

	// The ID of the Spring Cloud Accelerator.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name which should be used for this Spring Cloud Accelerator. Changing this forces a new Spring Cloud Accelerator to be created. The only possible value is default.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the Spring Cloud Service. Changing this forces a new Spring Cloud Accelerator to be created.
	SpringCloudServiceID *string `json:"springCloudServiceId,omitempty" tf:"spring_cloud_service_id,omitempty"`
}

type SpringCloudAcceleratorParameters struct {

	// The name which should be used for this Spring Cloud Accelerator. Changing this forces a new Spring Cloud Accelerator to be created. The only possible value is default.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the Spring Cloud Service. Changing this forces a new Spring Cloud Accelerator to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/appplatform/v1beta2.SpringCloudService
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SpringCloudServiceID *string `json:"springCloudServiceId,omitempty" tf:"spring_cloud_service_id,omitempty"`

	// Reference to a SpringCloudService in appplatform to populate springCloudServiceId.
	// +kubebuilder:validation:Optional
	SpringCloudServiceIDRef *v1.Reference `json:"springCloudServiceIdRef,omitempty" tf:"-"`

	// Selector for a SpringCloudService in appplatform to populate springCloudServiceId.
	// +kubebuilder:validation:Optional
	SpringCloudServiceIDSelector *v1.Selector `json:"springCloudServiceIdSelector,omitempty" tf:"-"`
}

// SpringCloudAcceleratorSpec defines the desired state of SpringCloudAccelerator
type SpringCloudAcceleratorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SpringCloudAcceleratorParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider SpringCloudAcceleratorInitParameters `json:"initProvider,omitempty"`
}

// SpringCloudAcceleratorStatus defines the observed state of SpringCloudAccelerator.
type SpringCloudAcceleratorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SpringCloudAcceleratorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SpringCloudAccelerator is the Schema for the SpringCloudAccelerators API. Manages a Spring Cloud Accelerator.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SpringCloudAccelerator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   SpringCloudAcceleratorSpec   `json:"spec"`
	Status SpringCloudAcceleratorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpringCloudAcceleratorList contains a list of SpringCloudAccelerators
type SpringCloudAcceleratorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpringCloudAccelerator `json:"items"`
}

// Repository type metadata.
var (
	SpringCloudAccelerator_Kind             = "SpringCloudAccelerator"
	SpringCloudAccelerator_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SpringCloudAccelerator_Kind}.String()
	SpringCloudAccelerator_KindAPIVersion   = SpringCloudAccelerator_Kind + "." + CRDGroupVersion.String()
	SpringCloudAccelerator_GroupVersionKind = CRDGroupVersion.WithKind(SpringCloudAccelerator_Kind)
)

func init() {
	SchemeBuilder.Register(&SpringCloudAccelerator{}, &SpringCloudAcceleratorList{})
}
