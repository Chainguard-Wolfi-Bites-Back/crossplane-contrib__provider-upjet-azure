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

type SpringCloudAPIPortalCustomDomainInitParameters struct {

	// The name which should be used for this Spring Cloud API Portal Domain. Changing this forces a new Spring Cloud API Portal Domain to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the Spring Cloud API Portal. Changing this forces a new Spring Cloud API Portal Domain to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/appplatform/v1beta2.SpringCloudAPIPortal
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	SpringCloudAPIPortalID *string `json:"springCloudApiPortalId,omitempty" tf:"spring_cloud_api_portal_id,omitempty"`

	// Reference to a SpringCloudAPIPortal in appplatform to populate springCloudApiPortalId.
	// +kubebuilder:validation:Optional
	SpringCloudAPIPortalIDRef *v1.Reference `json:"springCloudApiPortalIdRef,omitempty" tf:"-"`

	// Selector for a SpringCloudAPIPortal in appplatform to populate springCloudApiPortalId.
	// +kubebuilder:validation:Optional
	SpringCloudAPIPortalIDSelector *v1.Selector `json:"springCloudApiPortalIdSelector,omitempty" tf:"-"`

	// Specifies the thumbprint of the Spring Cloud Certificate that binds to the Spring Cloud API Portal Domain.
	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`
}

type SpringCloudAPIPortalCustomDomainObservation struct {

	// The ID of the Spring Cloud API Portal Domain.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name which should be used for this Spring Cloud API Portal Domain. Changing this forces a new Spring Cloud API Portal Domain to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the Spring Cloud API Portal. Changing this forces a new Spring Cloud API Portal Domain to be created.
	SpringCloudAPIPortalID *string `json:"springCloudApiPortalId,omitempty" tf:"spring_cloud_api_portal_id,omitempty"`

	// Specifies the thumbprint of the Spring Cloud Certificate that binds to the Spring Cloud API Portal Domain.
	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`
}

type SpringCloudAPIPortalCustomDomainParameters struct {

	// The name which should be used for this Spring Cloud API Portal Domain. Changing this forces a new Spring Cloud API Portal Domain to be created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the Spring Cloud API Portal. Changing this forces a new Spring Cloud API Portal Domain to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/appplatform/v1beta2.SpringCloudAPIPortal
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SpringCloudAPIPortalID *string `json:"springCloudApiPortalId,omitempty" tf:"spring_cloud_api_portal_id,omitempty"`

	// Reference to a SpringCloudAPIPortal in appplatform to populate springCloudApiPortalId.
	// +kubebuilder:validation:Optional
	SpringCloudAPIPortalIDRef *v1.Reference `json:"springCloudApiPortalIdRef,omitempty" tf:"-"`

	// Selector for a SpringCloudAPIPortal in appplatform to populate springCloudApiPortalId.
	// +kubebuilder:validation:Optional
	SpringCloudAPIPortalIDSelector *v1.Selector `json:"springCloudApiPortalIdSelector,omitempty" tf:"-"`

	// Specifies the thumbprint of the Spring Cloud Certificate that binds to the Spring Cloud API Portal Domain.
	// +kubebuilder:validation:Optional
	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`
}

// SpringCloudAPIPortalCustomDomainSpec defines the desired state of SpringCloudAPIPortalCustomDomain
type SpringCloudAPIPortalCustomDomainSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SpringCloudAPIPortalCustomDomainParameters `json:"forProvider"`
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
	InitProvider SpringCloudAPIPortalCustomDomainInitParameters `json:"initProvider,omitempty"`
}

// SpringCloudAPIPortalCustomDomainStatus defines the observed state of SpringCloudAPIPortalCustomDomain.
type SpringCloudAPIPortalCustomDomainStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SpringCloudAPIPortalCustomDomainObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SpringCloudAPIPortalCustomDomain is the Schema for the SpringCloudAPIPortalCustomDomains API. Manages a Spring Cloud API Portal Domain.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SpringCloudAPIPortalCustomDomain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   SpringCloudAPIPortalCustomDomainSpec   `json:"spec"`
	Status SpringCloudAPIPortalCustomDomainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpringCloudAPIPortalCustomDomainList contains a list of SpringCloudAPIPortalCustomDomains
type SpringCloudAPIPortalCustomDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpringCloudAPIPortalCustomDomain `json:"items"`
}

// Repository type metadata.
var (
	SpringCloudAPIPortalCustomDomain_Kind             = "SpringCloudAPIPortalCustomDomain"
	SpringCloudAPIPortalCustomDomain_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SpringCloudAPIPortalCustomDomain_Kind}.String()
	SpringCloudAPIPortalCustomDomain_KindAPIVersion   = SpringCloudAPIPortalCustomDomain_Kind + "." + CRDGroupVersion.String()
	SpringCloudAPIPortalCustomDomain_GroupVersionKind = CRDGroupVersion.WithKind(SpringCloudAPIPortalCustomDomain_Kind)
)

func init() {
	SchemeBuilder.Register(&SpringCloudAPIPortalCustomDomain{}, &SpringCloudAPIPortalCustomDomainList{})
}
