/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this MSSQLDatabase.
func (mg *MSSQLDatabase) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MSSQLDatabase.
func (mg *MSSQLDatabase) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MSSQLDatabase.
func (mg *MSSQLDatabase) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MSSQLDatabase.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MSSQLDatabase) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this MSSQLDatabase.
func (mg *MSSQLDatabase) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this MSSQLDatabase.
func (mg *MSSQLDatabase) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MSSQLDatabase.
func (mg *MSSQLDatabase) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MSSQLDatabase.
func (mg *MSSQLDatabase) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MSSQLDatabase.
func (mg *MSSQLDatabase) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MSSQLDatabase.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MSSQLDatabase) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this MSSQLDatabase.
func (mg *MSSQLDatabase) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this MSSQLDatabase.
func (mg *MSSQLDatabase) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MSSQLFailoverGroup.
func (mg *MSSQLFailoverGroup) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MSSQLFailoverGroup.
func (mg *MSSQLFailoverGroup) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MSSQLFailoverGroup.
func (mg *MSSQLFailoverGroup) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MSSQLFailoverGroup.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MSSQLFailoverGroup) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this MSSQLFailoverGroup.
func (mg *MSSQLFailoverGroup) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this MSSQLFailoverGroup.
func (mg *MSSQLFailoverGroup) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MSSQLFailoverGroup.
func (mg *MSSQLFailoverGroup) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MSSQLFailoverGroup.
func (mg *MSSQLFailoverGroup) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MSSQLFailoverGroup.
func (mg *MSSQLFailoverGroup) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MSSQLFailoverGroup.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MSSQLFailoverGroup) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this MSSQLFailoverGroup.
func (mg *MSSQLFailoverGroup) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this MSSQLFailoverGroup.
func (mg *MSSQLFailoverGroup) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MSSQLManagedDatabase.
func (mg *MSSQLManagedDatabase) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MSSQLManagedDatabase.
func (mg *MSSQLManagedDatabase) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MSSQLManagedDatabase.
func (mg *MSSQLManagedDatabase) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MSSQLManagedDatabase.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MSSQLManagedDatabase) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this MSSQLManagedDatabase.
func (mg *MSSQLManagedDatabase) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this MSSQLManagedDatabase.
func (mg *MSSQLManagedDatabase) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MSSQLManagedDatabase.
func (mg *MSSQLManagedDatabase) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MSSQLManagedDatabase.
func (mg *MSSQLManagedDatabase) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MSSQLManagedDatabase.
func (mg *MSSQLManagedDatabase) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MSSQLManagedDatabase.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MSSQLManagedDatabase) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this MSSQLManagedDatabase.
func (mg *MSSQLManagedDatabase) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this MSSQLManagedDatabase.
func (mg *MSSQLManagedDatabase) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MSSQLManagedInstance.
func (mg *MSSQLManagedInstance) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MSSQLManagedInstance.
func (mg *MSSQLManagedInstance) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MSSQLManagedInstance.
func (mg *MSSQLManagedInstance) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MSSQLManagedInstance.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MSSQLManagedInstance) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this MSSQLManagedInstance.
func (mg *MSSQLManagedInstance) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this MSSQLManagedInstance.
func (mg *MSSQLManagedInstance) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MSSQLManagedInstance.
func (mg *MSSQLManagedInstance) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MSSQLManagedInstance.
func (mg *MSSQLManagedInstance) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MSSQLManagedInstance.
func (mg *MSSQLManagedInstance) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MSSQLManagedInstance.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MSSQLManagedInstance) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this MSSQLManagedInstance.
func (mg *MSSQLManagedInstance) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this MSSQLManagedInstance.
func (mg *MSSQLManagedInstance) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MSSQLManagedInstanceActiveDirectoryAdministrator.
func (mg *MSSQLManagedInstanceActiveDirectoryAdministrator) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MSSQLManagedInstanceActiveDirectoryAdministrator.
func (mg *MSSQLManagedInstanceActiveDirectoryAdministrator) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MSSQLManagedInstanceActiveDirectoryAdministrator.
func (mg *MSSQLManagedInstanceActiveDirectoryAdministrator) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MSSQLManagedInstanceActiveDirectoryAdministrator.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MSSQLManagedInstanceActiveDirectoryAdministrator) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this MSSQLManagedInstanceActiveDirectoryAdministrator.
func (mg *MSSQLManagedInstanceActiveDirectoryAdministrator) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this MSSQLManagedInstanceActiveDirectoryAdministrator.
func (mg *MSSQLManagedInstanceActiveDirectoryAdministrator) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MSSQLManagedInstanceActiveDirectoryAdministrator.
func (mg *MSSQLManagedInstanceActiveDirectoryAdministrator) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MSSQLManagedInstanceActiveDirectoryAdministrator.
func (mg *MSSQLManagedInstanceActiveDirectoryAdministrator) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MSSQLManagedInstanceActiveDirectoryAdministrator.
func (mg *MSSQLManagedInstanceActiveDirectoryAdministrator) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MSSQLManagedInstanceActiveDirectoryAdministrator.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MSSQLManagedInstanceActiveDirectoryAdministrator) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this MSSQLManagedInstanceActiveDirectoryAdministrator.
func (mg *MSSQLManagedInstanceActiveDirectoryAdministrator) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this MSSQLManagedInstanceActiveDirectoryAdministrator.
func (mg *MSSQLManagedInstanceActiveDirectoryAdministrator) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MSSQLManagedInstanceFailoverGroup.
func (mg *MSSQLManagedInstanceFailoverGroup) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MSSQLManagedInstanceFailoverGroup.
func (mg *MSSQLManagedInstanceFailoverGroup) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MSSQLManagedInstanceFailoverGroup.
func (mg *MSSQLManagedInstanceFailoverGroup) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MSSQLManagedInstanceFailoverGroup.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MSSQLManagedInstanceFailoverGroup) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this MSSQLManagedInstanceFailoverGroup.
func (mg *MSSQLManagedInstanceFailoverGroup) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this MSSQLManagedInstanceFailoverGroup.
func (mg *MSSQLManagedInstanceFailoverGroup) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MSSQLManagedInstanceFailoverGroup.
func (mg *MSSQLManagedInstanceFailoverGroup) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MSSQLManagedInstanceFailoverGroup.
func (mg *MSSQLManagedInstanceFailoverGroup) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MSSQLManagedInstanceFailoverGroup.
func (mg *MSSQLManagedInstanceFailoverGroup) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MSSQLManagedInstanceFailoverGroup.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MSSQLManagedInstanceFailoverGroup) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this MSSQLManagedInstanceFailoverGroup.
func (mg *MSSQLManagedInstanceFailoverGroup) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this MSSQLManagedInstanceFailoverGroup.
func (mg *MSSQLManagedInstanceFailoverGroup) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MSSQLManagedInstanceVulnerabilityAssessment.
func (mg *MSSQLManagedInstanceVulnerabilityAssessment) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MSSQLManagedInstanceVulnerabilityAssessment.
func (mg *MSSQLManagedInstanceVulnerabilityAssessment) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MSSQLManagedInstanceVulnerabilityAssessment.
func (mg *MSSQLManagedInstanceVulnerabilityAssessment) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MSSQLManagedInstanceVulnerabilityAssessment.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MSSQLManagedInstanceVulnerabilityAssessment) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this MSSQLManagedInstanceVulnerabilityAssessment.
func (mg *MSSQLManagedInstanceVulnerabilityAssessment) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this MSSQLManagedInstanceVulnerabilityAssessment.
func (mg *MSSQLManagedInstanceVulnerabilityAssessment) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MSSQLManagedInstanceVulnerabilityAssessment.
func (mg *MSSQLManagedInstanceVulnerabilityAssessment) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MSSQLManagedInstanceVulnerabilityAssessment.
func (mg *MSSQLManagedInstanceVulnerabilityAssessment) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MSSQLManagedInstanceVulnerabilityAssessment.
func (mg *MSSQLManagedInstanceVulnerabilityAssessment) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MSSQLManagedInstanceVulnerabilityAssessment.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MSSQLManagedInstanceVulnerabilityAssessment) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this MSSQLManagedInstanceVulnerabilityAssessment.
func (mg *MSSQLManagedInstanceVulnerabilityAssessment) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this MSSQLManagedInstanceVulnerabilityAssessment.
func (mg *MSSQLManagedInstanceVulnerabilityAssessment) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MSSQLOutboundFirewallRule.
func (mg *MSSQLOutboundFirewallRule) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MSSQLOutboundFirewallRule.
func (mg *MSSQLOutboundFirewallRule) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MSSQLOutboundFirewallRule.
func (mg *MSSQLOutboundFirewallRule) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MSSQLOutboundFirewallRule.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MSSQLOutboundFirewallRule) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this MSSQLOutboundFirewallRule.
func (mg *MSSQLOutboundFirewallRule) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this MSSQLOutboundFirewallRule.
func (mg *MSSQLOutboundFirewallRule) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MSSQLOutboundFirewallRule.
func (mg *MSSQLOutboundFirewallRule) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MSSQLOutboundFirewallRule.
func (mg *MSSQLOutboundFirewallRule) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MSSQLOutboundFirewallRule.
func (mg *MSSQLOutboundFirewallRule) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MSSQLOutboundFirewallRule.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MSSQLOutboundFirewallRule) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this MSSQLOutboundFirewallRule.
func (mg *MSSQLOutboundFirewallRule) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this MSSQLOutboundFirewallRule.
func (mg *MSSQLOutboundFirewallRule) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MSSQLServer.
func (mg *MSSQLServer) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MSSQLServer.
func (mg *MSSQLServer) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MSSQLServer.
func (mg *MSSQLServer) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MSSQLServer.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MSSQLServer) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this MSSQLServer.
func (mg *MSSQLServer) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this MSSQLServer.
func (mg *MSSQLServer) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MSSQLServer.
func (mg *MSSQLServer) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MSSQLServer.
func (mg *MSSQLServer) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MSSQLServer.
func (mg *MSSQLServer) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MSSQLServer.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MSSQLServer) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this MSSQLServer.
func (mg *MSSQLServer) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this MSSQLServer.
func (mg *MSSQLServer) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MSSQLServerDNSAlias.
func (mg *MSSQLServerDNSAlias) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MSSQLServerDNSAlias.
func (mg *MSSQLServerDNSAlias) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MSSQLServerDNSAlias.
func (mg *MSSQLServerDNSAlias) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MSSQLServerDNSAlias.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MSSQLServerDNSAlias) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this MSSQLServerDNSAlias.
func (mg *MSSQLServerDNSAlias) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this MSSQLServerDNSAlias.
func (mg *MSSQLServerDNSAlias) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MSSQLServerDNSAlias.
func (mg *MSSQLServerDNSAlias) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MSSQLServerDNSAlias.
func (mg *MSSQLServerDNSAlias) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MSSQLServerDNSAlias.
func (mg *MSSQLServerDNSAlias) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MSSQLServerDNSAlias.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MSSQLServerDNSAlias) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this MSSQLServerDNSAlias.
func (mg *MSSQLServerDNSAlias) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this MSSQLServerDNSAlias.
func (mg *MSSQLServerDNSAlias) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MSSQLServerTransparentDataEncryption.
func (mg *MSSQLServerTransparentDataEncryption) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MSSQLServerTransparentDataEncryption.
func (mg *MSSQLServerTransparentDataEncryption) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MSSQLServerTransparentDataEncryption.
func (mg *MSSQLServerTransparentDataEncryption) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MSSQLServerTransparentDataEncryption.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MSSQLServerTransparentDataEncryption) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this MSSQLServerTransparentDataEncryption.
func (mg *MSSQLServerTransparentDataEncryption) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this MSSQLServerTransparentDataEncryption.
func (mg *MSSQLServerTransparentDataEncryption) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MSSQLServerTransparentDataEncryption.
func (mg *MSSQLServerTransparentDataEncryption) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MSSQLServerTransparentDataEncryption.
func (mg *MSSQLServerTransparentDataEncryption) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MSSQLServerTransparentDataEncryption.
func (mg *MSSQLServerTransparentDataEncryption) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MSSQLServerTransparentDataEncryption.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MSSQLServerTransparentDataEncryption) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this MSSQLServerTransparentDataEncryption.
func (mg *MSSQLServerTransparentDataEncryption) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this MSSQLServerTransparentDataEncryption.
func (mg *MSSQLServerTransparentDataEncryption) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MSSQLVirtualNetworkRule.
func (mg *MSSQLVirtualNetworkRule) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MSSQLVirtualNetworkRule.
func (mg *MSSQLVirtualNetworkRule) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MSSQLVirtualNetworkRule.
func (mg *MSSQLVirtualNetworkRule) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MSSQLVirtualNetworkRule.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MSSQLVirtualNetworkRule) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this MSSQLVirtualNetworkRule.
func (mg *MSSQLVirtualNetworkRule) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this MSSQLVirtualNetworkRule.
func (mg *MSSQLVirtualNetworkRule) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MSSQLVirtualNetworkRule.
func (mg *MSSQLVirtualNetworkRule) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MSSQLVirtualNetworkRule.
func (mg *MSSQLVirtualNetworkRule) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MSSQLVirtualNetworkRule.
func (mg *MSSQLVirtualNetworkRule) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MSSQLVirtualNetworkRule.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MSSQLVirtualNetworkRule) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this MSSQLVirtualNetworkRule.
func (mg *MSSQLVirtualNetworkRule) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this MSSQLVirtualNetworkRule.
func (mg *MSSQLVirtualNetworkRule) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
