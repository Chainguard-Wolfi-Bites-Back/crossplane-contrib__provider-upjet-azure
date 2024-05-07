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

type LogAnalyticsDataExportRuleInitParameters struct {

	// The destination resource ID. It should be a storage account, an event hub namespace or an event hub. If the destination is an event hub namespace, an event hub would be created for each table automatically.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta2.Account
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	DestinationResourceID *string `json:"destinationResourceId,omitempty" tf:"destination_resource_id,omitempty"`

	// Reference to a Account in storage to populate destinationResourceId.
	// +kubebuilder:validation:Optional
	DestinationResourceIDRef *v1.Reference `json:"destinationResourceIdRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate destinationResourceId.
	// +kubebuilder:validation:Optional
	DestinationResourceIDSelector *v1.Selector `json:"destinationResourceIdSelector,omitempty" tf:"-"`

	// Is this Log Analytics Data Export Rule enabled? Possible values include true or false. Defaults to false.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The name of the Resource Group where the Log Analytics Data Export should exist. Changing this forces a new Log Analytics Data Export Rule to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A list of table names to export to the destination resource, for example: ["Heartbeat", "SecurityEvent"].
	// +listType=set
	TableNames []*string `json:"tableNames,omitempty" tf:"table_names,omitempty"`
}

type LogAnalyticsDataExportRuleObservation struct {

	// The destination resource ID. It should be a storage account, an event hub namespace or an event hub. If the destination is an event hub namespace, an event hub would be created for each table automatically.
	DestinationResourceID *string `json:"destinationResourceId,omitempty" tf:"destination_resource_id,omitempty"`

	// Is this Log Analytics Data Export Rule enabled? Possible values include true or false. Defaults to false.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The ID of the created Data Export Rule.
	ExportRuleID *string `json:"exportRuleId,omitempty" tf:"export_rule_id,omitempty"`

	// The ID of the Log Analytics Data Export Rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the Resource Group where the Log Analytics Data Export should exist. Changing this forces a new Log Analytics Data Export Rule to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A list of table names to export to the destination resource, for example: ["Heartbeat", "SecurityEvent"].
	// +listType=set
	TableNames []*string `json:"tableNames,omitempty" tf:"table_names,omitempty"`

	// The resource ID of the workspace. Changing this forces a new Log Analytics Data Export Rule to be created.
	WorkspaceResourceID *string `json:"workspaceResourceId,omitempty" tf:"workspace_resource_id,omitempty"`
}

type LogAnalyticsDataExportRuleParameters struct {

	// The destination resource ID. It should be a storage account, an event hub namespace or an event hub. If the destination is an event hub namespace, an event hub would be created for each table automatically.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta2.Account
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DestinationResourceID *string `json:"destinationResourceId,omitempty" tf:"destination_resource_id,omitempty"`

	// Reference to a Account in storage to populate destinationResourceId.
	// +kubebuilder:validation:Optional
	DestinationResourceIDRef *v1.Reference `json:"destinationResourceIdRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate destinationResourceId.
	// +kubebuilder:validation:Optional
	DestinationResourceIDSelector *v1.Selector `json:"destinationResourceIdSelector,omitempty" tf:"-"`

	// Is this Log Analytics Data Export Rule enabled? Possible values include true or false. Defaults to false.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The name of the Resource Group where the Log Analytics Data Export should exist. Changing this forces a new Log Analytics Data Export Rule to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A list of table names to export to the destination resource, for example: ["Heartbeat", "SecurityEvent"].
	// +kubebuilder:validation:Optional
	// +listType=set
	TableNames []*string `json:"tableNames,omitempty" tf:"table_names,omitempty"`

	// The resource ID of the workspace. Changing this forces a new Log Analytics Data Export Rule to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/operationalinsights/v1beta2.Workspace
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	WorkspaceResourceID *string `json:"workspaceResourceId,omitempty" tf:"workspace_resource_id,omitempty"`

	// Reference to a Workspace in operationalinsights to populate workspaceResourceId.
	// +kubebuilder:validation:Optional
	WorkspaceResourceIDRef *v1.Reference `json:"workspaceResourceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in operationalinsights to populate workspaceResourceId.
	// +kubebuilder:validation:Optional
	WorkspaceResourceIDSelector *v1.Selector `json:"workspaceResourceIdSelector,omitempty" tf:"-"`
}

// LogAnalyticsDataExportRuleSpec defines the desired state of LogAnalyticsDataExportRule
type LogAnalyticsDataExportRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LogAnalyticsDataExportRuleParameters `json:"forProvider"`
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
	InitProvider LogAnalyticsDataExportRuleInitParameters `json:"initProvider,omitempty"`
}

// LogAnalyticsDataExportRuleStatus defines the observed state of LogAnalyticsDataExportRule.
type LogAnalyticsDataExportRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LogAnalyticsDataExportRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// LogAnalyticsDataExportRule is the Schema for the LogAnalyticsDataExportRules API. Manages a log analytics Data Export Rule.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LogAnalyticsDataExportRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.tableNames) || (has(self.initProvider) && has(self.initProvider.tableNames))",message="spec.forProvider.tableNames is a required parameter"
	Spec   LogAnalyticsDataExportRuleSpec   `json:"spec"`
	Status LogAnalyticsDataExportRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LogAnalyticsDataExportRuleList contains a list of LogAnalyticsDataExportRules
type LogAnalyticsDataExportRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LogAnalyticsDataExportRule `json:"items"`
}

// Repository type metadata.
var (
	LogAnalyticsDataExportRule_Kind             = "LogAnalyticsDataExportRule"
	LogAnalyticsDataExportRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LogAnalyticsDataExportRule_Kind}.String()
	LogAnalyticsDataExportRule_KindAPIVersion   = LogAnalyticsDataExportRule_Kind + "." + CRDGroupVersion.String()
	LogAnalyticsDataExportRule_GroupVersionKind = CRDGroupVersion.WithKind(LogAnalyticsDataExportRule_Kind)
)

func init() {
	SchemeBuilder.Register(&LogAnalyticsDataExportRule{}, &LogAnalyticsDataExportRuleList{})
}
