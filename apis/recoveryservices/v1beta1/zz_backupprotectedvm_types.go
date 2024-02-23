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

type BackupProtectedVMInitParameters struct {

	// Specifies the id of the backup policy to use.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/recoveryservices/v1beta1.BackupPolicyVM
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	BackupPolicyID *string `json:"backupPolicyId,omitempty" tf:"backup_policy_id,omitempty"`

	// Reference to a BackupPolicyVM in recoveryservices to populate backupPolicyId.
	// +kubebuilder:validation:Optional
	BackupPolicyIDRef *v1.Reference `json:"backupPolicyIdRef,omitempty" tf:"-"`

	// Selector for a BackupPolicyVM in recoveryservices to populate backupPolicyId.
	// +kubebuilder:validation:Optional
	BackupPolicyIDSelector *v1.Selector `json:"backupPolicyIdSelector,omitempty" tf:"-"`

	// A list of Disks' Logical Unit Numbers(LUN) to be excluded for VM Protection.
	// +listType=set
	ExcludeDiskLuns []*float64 `json:"excludeDiskLuns,omitempty" tf:"exclude_disk_luns,omitempty"`

	// A list of Disks' Logical Unit Numbers(LUN) to be included for VM Protection.
	// +listType=set
	IncludeDiskLuns []*float64 `json:"includeDiskLuns,omitempty" tf:"include_disk_luns,omitempty"`

	// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/recoveryservices/v1beta1.Vault
	RecoveryVaultName *string `json:"recoveryVaultName,omitempty" tf:"recovery_vault_name,omitempty"`

	// Reference to a Vault in recoveryservices to populate recoveryVaultName.
	// +kubebuilder:validation:Optional
	RecoveryVaultNameRef *v1.Reference `json:"recoveryVaultNameRef,omitempty" tf:"-"`

	// Selector for a Vault in recoveryservices to populate recoveryVaultName.
	// +kubebuilder:validation:Optional
	RecoveryVaultNameSelector *v1.Selector `json:"recoveryVaultNameSelector,omitempty" tf:"-"`

	// The name of the resource group in which to create the Recovery Services Vault. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Specifies the ID of the VM to backup. Changing this forces a new resource to be created.
	SourceVMID *string `json:"sourceVmId,omitempty" tf:"source_vm_id,omitempty"`
}

type BackupProtectedVMObservation struct {

	// Specifies the id of the backup policy to use.
	BackupPolicyID *string `json:"backupPolicyId,omitempty" tf:"backup_policy_id,omitempty"`

	// A list of Disks' Logical Unit Numbers(LUN) to be excluded for VM Protection.
	// +listType=set
	ExcludeDiskLuns []*float64 `json:"excludeDiskLuns,omitempty" tf:"exclude_disk_luns,omitempty"`

	// The ID of the Backup Protected Virtual Machine.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A list of Disks' Logical Unit Numbers(LUN) to be included for VM Protection.
	// +listType=set
	IncludeDiskLuns []*float64 `json:"includeDiskLuns,omitempty" tf:"include_disk_luns,omitempty"`

	// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
	RecoveryVaultName *string `json:"recoveryVaultName,omitempty" tf:"recovery_vault_name,omitempty"`

	// The name of the resource group in which to create the Recovery Services Vault. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Specifies the ID of the VM to backup. Changing this forces a new resource to be created.
	SourceVMID *string `json:"sourceVmId,omitempty" tf:"source_vm_id,omitempty"`
}

type BackupProtectedVMParameters struct {

	// Specifies the id of the backup policy to use.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/recoveryservices/v1beta1.BackupPolicyVM
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	BackupPolicyID *string `json:"backupPolicyId,omitempty" tf:"backup_policy_id,omitempty"`

	// Reference to a BackupPolicyVM in recoveryservices to populate backupPolicyId.
	// +kubebuilder:validation:Optional
	BackupPolicyIDRef *v1.Reference `json:"backupPolicyIdRef,omitempty" tf:"-"`

	// Selector for a BackupPolicyVM in recoveryservices to populate backupPolicyId.
	// +kubebuilder:validation:Optional
	BackupPolicyIDSelector *v1.Selector `json:"backupPolicyIdSelector,omitempty" tf:"-"`

	// A list of Disks' Logical Unit Numbers(LUN) to be excluded for VM Protection.
	// +kubebuilder:validation:Optional
	// +listType=set
	ExcludeDiskLuns []*float64 `json:"excludeDiskLuns,omitempty" tf:"exclude_disk_luns,omitempty"`

	// A list of Disks' Logical Unit Numbers(LUN) to be included for VM Protection.
	// +kubebuilder:validation:Optional
	// +listType=set
	IncludeDiskLuns []*float64 `json:"includeDiskLuns,omitempty" tf:"include_disk_luns,omitempty"`

	// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/recoveryservices/v1beta1.Vault
	// +kubebuilder:validation:Optional
	RecoveryVaultName *string `json:"recoveryVaultName,omitempty" tf:"recovery_vault_name,omitempty"`

	// Reference to a Vault in recoveryservices to populate recoveryVaultName.
	// +kubebuilder:validation:Optional
	RecoveryVaultNameRef *v1.Reference `json:"recoveryVaultNameRef,omitempty" tf:"-"`

	// Selector for a Vault in recoveryservices to populate recoveryVaultName.
	// +kubebuilder:validation:Optional
	RecoveryVaultNameSelector *v1.Selector `json:"recoveryVaultNameSelector,omitempty" tf:"-"`

	// The name of the resource group in which to create the Recovery Services Vault. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Specifies the ID of the VM to backup. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SourceVMID *string `json:"sourceVmId,omitempty" tf:"source_vm_id,omitempty"`
}

// BackupProtectedVMSpec defines the desired state of BackupProtectedVM
type BackupProtectedVMSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackupProtectedVMParameters `json:"forProvider"`
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
	InitProvider BackupProtectedVMInitParameters `json:"initProvider,omitempty"`
}

// BackupProtectedVMStatus defines the observed state of BackupProtectedVM.
type BackupProtectedVMStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackupProtectedVMObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BackupProtectedVM is the Schema for the BackupProtectedVMs API. Manages an Azure Backup Protected VM.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type BackupProtectedVM struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackupProtectedVMSpec   `json:"spec"`
	Status            BackupProtectedVMStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupProtectedVMList contains a list of BackupProtectedVMs
type BackupProtectedVMList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupProtectedVM `json:"items"`
}

// Repository type metadata.
var (
	BackupProtectedVM_Kind             = "BackupProtectedVM"
	BackupProtectedVM_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BackupProtectedVM_Kind}.String()
	BackupProtectedVM_KindAPIVersion   = BackupProtectedVM_Kind + "." + CRDGroupVersion.String()
	BackupProtectedVM_GroupVersionKind = CRDGroupVersion.WithKind(BackupProtectedVM_Kind)
)

func init() {
	SchemeBuilder.Register(&BackupProtectedVM{}, &BackupProtectedVMList{})
}
