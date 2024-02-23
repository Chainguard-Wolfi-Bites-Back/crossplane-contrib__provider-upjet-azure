// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type NetworkInterfaceNatRuleAssociationInitParameters struct {

	// The Name of the IP Configuration within the Network Interface which should be connected to the NAT Rule. Changing this forces a new resource to be created.
	IPConfigurationName *string `json:"ipConfigurationName,omitempty" tf:"ip_configuration_name,omitempty"`

	// The ID of the Load Balancer NAT Rule which this Network Interface which should be connected to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=LoadBalancerNatRule
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	NATRuleID *string `json:"natRuleId,omitempty" tf:"nat_rule_id,omitempty"`

	// Reference to a LoadBalancerNatRule to populate natRuleId.
	// +kubebuilder:validation:Optional
	NATRuleIDRef *v1.Reference `json:"natRuleIdRef,omitempty" tf:"-"`

	// Selector for a LoadBalancerNatRule to populate natRuleId.
	// +kubebuilder:validation:Optional
	NATRuleIDSelector *v1.Selector `json:"natRuleIdSelector,omitempty" tf:"-"`

	// The ID of the Network Interface. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=NetworkInterface
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	// Reference to a NetworkInterface to populate networkInterfaceId.
	// +kubebuilder:validation:Optional
	NetworkInterfaceIDRef *v1.Reference `json:"networkInterfaceIdRef,omitempty" tf:"-"`

	// Selector for a NetworkInterface to populate networkInterfaceId.
	// +kubebuilder:validation:Optional
	NetworkInterfaceIDSelector *v1.Selector `json:"networkInterfaceIdSelector,omitempty" tf:"-"`
}

type NetworkInterfaceNatRuleAssociationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Name of the IP Configuration within the Network Interface which should be connected to the NAT Rule. Changing this forces a new resource to be created.
	IPConfigurationName *string `json:"ipConfigurationName,omitempty" tf:"ip_configuration_name,omitempty"`

	// The ID of the Load Balancer NAT Rule which this Network Interface which should be connected to. Changing this forces a new resource to be created.
	NATRuleID *string `json:"natRuleId,omitempty" tf:"nat_rule_id,omitempty"`

	// The ID of the Network Interface. Changing this forces a new resource to be created.
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`
}

type NetworkInterfaceNatRuleAssociationParameters struct {

	// The Name of the IP Configuration within the Network Interface which should be connected to the NAT Rule. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	IPConfigurationName *string `json:"ipConfigurationName,omitempty" tf:"ip_configuration_name,omitempty"`

	// The ID of the Load Balancer NAT Rule which this Network Interface which should be connected to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=LoadBalancerNatRule
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	NATRuleID *string `json:"natRuleId,omitempty" tf:"nat_rule_id,omitempty"`

	// Reference to a LoadBalancerNatRule to populate natRuleId.
	// +kubebuilder:validation:Optional
	NATRuleIDRef *v1.Reference `json:"natRuleIdRef,omitempty" tf:"-"`

	// Selector for a LoadBalancerNatRule to populate natRuleId.
	// +kubebuilder:validation:Optional
	NATRuleIDSelector *v1.Selector `json:"natRuleIdSelector,omitempty" tf:"-"`

	// The ID of the Network Interface. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=NetworkInterface
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`

	// Reference to a NetworkInterface to populate networkInterfaceId.
	// +kubebuilder:validation:Optional
	NetworkInterfaceIDRef *v1.Reference `json:"networkInterfaceIdRef,omitempty" tf:"-"`

	// Selector for a NetworkInterface to populate networkInterfaceId.
	// +kubebuilder:validation:Optional
	NetworkInterfaceIDSelector *v1.Selector `json:"networkInterfaceIdSelector,omitempty" tf:"-"`
}

// NetworkInterfaceNatRuleAssociationSpec defines the desired state of NetworkInterfaceNatRuleAssociation
type NetworkInterfaceNatRuleAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkInterfaceNatRuleAssociationParameters `json:"forProvider"`
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
	InitProvider NetworkInterfaceNatRuleAssociationInitParameters `json:"initProvider,omitempty"`
}

// NetworkInterfaceNatRuleAssociationStatus defines the observed state of NetworkInterfaceNatRuleAssociation.
type NetworkInterfaceNatRuleAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkInterfaceNatRuleAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// NetworkInterfaceNatRuleAssociation is the Schema for the NetworkInterfaceNatRuleAssociations API. Manages the association between a Network Interface and a Load Balancer's NAT Rule.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type NetworkInterfaceNatRuleAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ipConfigurationName) || (has(self.initProvider) && has(self.initProvider.ipConfigurationName))",message="spec.forProvider.ipConfigurationName is a required parameter"
	Spec   NetworkInterfaceNatRuleAssociationSpec   `json:"spec"`
	Status NetworkInterfaceNatRuleAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkInterfaceNatRuleAssociationList contains a list of NetworkInterfaceNatRuleAssociations
type NetworkInterfaceNatRuleAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkInterfaceNatRuleAssociation `json:"items"`
}

// Repository type metadata.
var (
	NetworkInterfaceNatRuleAssociation_Kind             = "NetworkInterfaceNatRuleAssociation"
	NetworkInterfaceNatRuleAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NetworkInterfaceNatRuleAssociation_Kind}.String()
	NetworkInterfaceNatRuleAssociation_KindAPIVersion   = NetworkInterfaceNatRuleAssociation_Kind + "." + CRDGroupVersion.String()
	NetworkInterfaceNatRuleAssociation_GroupVersionKind = CRDGroupVersion.WithKind(NetworkInterfaceNatRuleAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&NetworkInterfaceNatRuleAssociation{}, &NetworkInterfaceNatRuleAssociationList{})
}
