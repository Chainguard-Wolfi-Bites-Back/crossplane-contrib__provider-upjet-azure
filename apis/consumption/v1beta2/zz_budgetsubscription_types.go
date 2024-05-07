// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BudgetSubscriptionFilterDimensionInitParameters struct {

	// The name of the tag to use for the filter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The operator to use for comparison. The allowed values are In. Defaults to In.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Specifies a list of values for the tag.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type BudgetSubscriptionFilterDimensionObservation struct {

	// The name of the tag to use for the filter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The operator to use for comparison. The allowed values are In. Defaults to In.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Specifies a list of values for the tag.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type BudgetSubscriptionFilterDimensionParameters struct {

	// The name of the tag to use for the filter.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The operator to use for comparison. The allowed values are In. Defaults to In.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Specifies a list of values for the tag.
	// +kubebuilder:validation:Optional
	Values []*string `json:"values" tf:"values,omitempty"`
}

type BudgetSubscriptionFilterInitParameters struct {

	// One or more dimension blocks as defined below to filter the budget on.
	Dimension []BudgetSubscriptionFilterDimensionInitParameters `json:"dimension,omitempty" tf:"dimension,omitempty"`

	// A not block as defined below to filter the budget on. This is deprecated as the API no longer supports it and will be removed in version 4.0 of the provider.
	Not *BudgetSubscriptionFilterNotInitParameters `json:"not,omitempty" tf:"not,omitempty"`

	// One or more tag blocks as defined below to filter the budget on.
	Tag []BudgetSubscriptionFilterTagInitParameters `json:"tag,omitempty" tf:"tag,omitempty"`
}

type BudgetSubscriptionFilterNotDimensionInitParameters struct {

	// The name of the tag to use for the filter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The operator to use for comparison. The allowed values are In. Defaults to In.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Specifies a list of values for the tag.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type BudgetSubscriptionFilterNotDimensionObservation struct {

	// The name of the tag to use for the filter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The operator to use for comparison. The allowed values are In. Defaults to In.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Specifies a list of values for the tag.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type BudgetSubscriptionFilterNotDimensionParameters struct {

	// The name of the tag to use for the filter.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The operator to use for comparison. The allowed values are In. Defaults to In.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Specifies a list of values for the tag.
	// +kubebuilder:validation:Optional
	Values []*string `json:"values" tf:"values,omitempty"`
}

type BudgetSubscriptionFilterNotInitParameters struct {

	// One dimension block as defined below to filter the budget on. Conflicts with tag.
	Dimension *BudgetSubscriptionFilterNotDimensionInitParameters `json:"dimension,omitempty" tf:"dimension,omitempty"`

	// One tag block as defined below to filter the budget on. Conflicts with dimension.
	Tag *FilterNotTagInitParameters `json:"tag,omitempty" tf:"tag,omitempty"`
}

type BudgetSubscriptionFilterNotObservation struct {

	// One dimension block as defined below to filter the budget on. Conflicts with tag.
	Dimension *BudgetSubscriptionFilterNotDimensionObservation `json:"dimension,omitempty" tf:"dimension,omitempty"`

	// One tag block as defined below to filter the budget on. Conflicts with dimension.
	Tag *FilterNotTagObservation `json:"tag,omitempty" tf:"tag,omitempty"`
}

type BudgetSubscriptionFilterNotParameters struct {

	// One dimension block as defined below to filter the budget on. Conflicts with tag.
	// +kubebuilder:validation:Optional
	Dimension *BudgetSubscriptionFilterNotDimensionParameters `json:"dimension,omitempty" tf:"dimension,omitempty"`

	// One tag block as defined below to filter the budget on. Conflicts with dimension.
	// +kubebuilder:validation:Optional
	Tag *FilterNotTagParameters `json:"tag,omitempty" tf:"tag,omitempty"`
}

type BudgetSubscriptionFilterObservation struct {

	// One or more dimension blocks as defined below to filter the budget on.
	Dimension []BudgetSubscriptionFilterDimensionObservation `json:"dimension,omitempty" tf:"dimension,omitempty"`

	// A not block as defined below to filter the budget on. This is deprecated as the API no longer supports it and will be removed in version 4.0 of the provider.
	Not *BudgetSubscriptionFilterNotObservation `json:"not,omitempty" tf:"not,omitempty"`

	// One or more tag blocks as defined below to filter the budget on.
	Tag []BudgetSubscriptionFilterTagObservation `json:"tag,omitempty" tf:"tag,omitempty"`
}

type BudgetSubscriptionFilterParameters struct {

	// One or more dimension blocks as defined below to filter the budget on.
	// +kubebuilder:validation:Optional
	Dimension []BudgetSubscriptionFilterDimensionParameters `json:"dimension,omitempty" tf:"dimension,omitempty"`

	// A not block as defined below to filter the budget on. This is deprecated as the API no longer supports it and will be removed in version 4.0 of the provider.
	// +kubebuilder:validation:Optional
	Not *BudgetSubscriptionFilterNotParameters `json:"not,omitempty" tf:"not,omitempty"`

	// One or more tag blocks as defined below to filter the budget on.
	// +kubebuilder:validation:Optional
	Tag []BudgetSubscriptionFilterTagParameters `json:"tag,omitempty" tf:"tag,omitempty"`
}

type BudgetSubscriptionFilterTagInitParameters struct {

	// The name of the tag to use for the filter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The operator to use for comparison. The allowed values are In. Defaults to In.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Specifies a list of values for the tag.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type BudgetSubscriptionFilterTagObservation struct {

	// The name of the tag to use for the filter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The operator to use for comparison. The allowed values are In. Defaults to In.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Specifies a list of values for the tag.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type BudgetSubscriptionFilterTagParameters struct {

	// The name of the tag to use for the filter.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The operator to use for comparison. The allowed values are In. Defaults to In.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Specifies a list of values for the tag.
	// +kubebuilder:validation:Optional
	Values []*string `json:"values" tf:"values,omitempty"`
}

type BudgetSubscriptionInitParameters struct {

	// The total amount of cost to track with the budget.
	Amount *float64 `json:"amount,omitempty" tf:"amount,omitempty"`

	// The ETag of the Subscription Consumption Budget.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// A filter block as defined below.
	Filter *BudgetSubscriptionFilterInitParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// One or more notification blocks as defined below.
	Notification []BudgetSubscriptionNotificationInitParameters `json:"notification,omitempty" tf:"notification,omitempty"`

	// The ID of the Subscription for which to create a Consumption Budget. Changing this forces a new resource to be created.
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`

	// The time covered by a budget. Tracking of the amount will be reset based on the time grain. Must be one of BillingAnnual, BillingMonth, BillingQuarter, Annually, Monthly and Quarterly. Defaults to Monthly. Changing this forces a new resource to be created.
	TimeGrain *string `json:"timeGrain,omitempty" tf:"time_grain,omitempty"`

	// A time_period block as defined below.
	TimePeriod *BudgetSubscriptionTimePeriodInitParameters `json:"timePeriod,omitempty" tf:"time_period,omitempty"`
}

type BudgetSubscriptionNotificationInitParameters struct {

	// Specifies a list of email addresses to send the budget notification to when the threshold is exceeded.
	ContactEmails []*string `json:"contactEmails,omitempty" tf:"contact_emails,omitempty"`

	// Specifies a list of Action Group IDs to send the budget notification to when the threshold is exceeded.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/insights/v1beta1.MonitorActionGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	ContactGroups []*string `json:"contactGroups,omitempty" tf:"contact_groups,omitempty"`

	// References to MonitorActionGroup in insights to populate contactGroups.
	// +kubebuilder:validation:Optional
	ContactGroupsRefs []v1.Reference `json:"contactGroupsRefs,omitempty" tf:"-"`

	// Selector for a list of MonitorActionGroup in insights to populate contactGroups.
	// +kubebuilder:validation:Optional
	ContactGroupsSelector *v1.Selector `json:"contactGroupsSelector,omitempty" tf:"-"`

	// Specifies a list of contact roles to send the budget notification to when the threshold is exceeded.
	ContactRoles []*string `json:"contactRoles,omitempty" tf:"contact_roles,omitempty"`

	// Should the notification be enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The comparison operator for the notification. Must be one of EqualTo, GreaterThan, or GreaterThanOrEqualTo.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Threshold value associated with a notification. Notification is sent when the cost exceeded the threshold. It is always percent and has to be between 0 and 1000.
	Threshold *float64 `json:"threshold,omitempty" tf:"threshold,omitempty"`

	// The type of threshold for the notification. This determines whether the notification is triggered by forecasted costs or actual costs. The allowed values are Actual and Forecasted. Default is Actual. Changing this forces a new resource to be created.
	ThresholdType *string `json:"thresholdType,omitempty" tf:"threshold_type,omitempty"`
}

type BudgetSubscriptionNotificationObservation struct {

	// Specifies a list of email addresses to send the budget notification to when the threshold is exceeded.
	ContactEmails []*string `json:"contactEmails,omitempty" tf:"contact_emails,omitempty"`

	// Specifies a list of Action Group IDs to send the budget notification to when the threshold is exceeded.
	ContactGroups []*string `json:"contactGroups,omitempty" tf:"contact_groups,omitempty"`

	// Specifies a list of contact roles to send the budget notification to when the threshold is exceeded.
	ContactRoles []*string `json:"contactRoles,omitempty" tf:"contact_roles,omitempty"`

	// Should the notification be enabled? Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The comparison operator for the notification. Must be one of EqualTo, GreaterThan, or GreaterThanOrEqualTo.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Threshold value associated with a notification. Notification is sent when the cost exceeded the threshold. It is always percent and has to be between 0 and 1000.
	Threshold *float64 `json:"threshold,omitempty" tf:"threshold,omitempty"`

	// The type of threshold for the notification. This determines whether the notification is triggered by forecasted costs or actual costs. The allowed values are Actual and Forecasted. Default is Actual. Changing this forces a new resource to be created.
	ThresholdType *string `json:"thresholdType,omitempty" tf:"threshold_type,omitempty"`
}

type BudgetSubscriptionNotificationParameters struct {

	// Specifies a list of email addresses to send the budget notification to when the threshold is exceeded.
	// +kubebuilder:validation:Optional
	ContactEmails []*string `json:"contactEmails,omitempty" tf:"contact_emails,omitempty"`

	// Specifies a list of Action Group IDs to send the budget notification to when the threshold is exceeded.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/insights/v1beta1.MonitorActionGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ContactGroups []*string `json:"contactGroups,omitempty" tf:"contact_groups,omitempty"`

	// References to MonitorActionGroup in insights to populate contactGroups.
	// +kubebuilder:validation:Optional
	ContactGroupsRefs []v1.Reference `json:"contactGroupsRefs,omitempty" tf:"-"`

	// Selector for a list of MonitorActionGroup in insights to populate contactGroups.
	// +kubebuilder:validation:Optional
	ContactGroupsSelector *v1.Selector `json:"contactGroupsSelector,omitempty" tf:"-"`

	// Specifies a list of contact roles to send the budget notification to when the threshold is exceeded.
	// +kubebuilder:validation:Optional
	ContactRoles []*string `json:"contactRoles,omitempty" tf:"contact_roles,omitempty"`

	// Should the notification be enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The comparison operator for the notification. Must be one of EqualTo, GreaterThan, or GreaterThanOrEqualTo.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// Threshold value associated with a notification. Notification is sent when the cost exceeded the threshold. It is always percent and has to be between 0 and 1000.
	// +kubebuilder:validation:Optional
	Threshold *float64 `json:"threshold" tf:"threshold,omitempty"`

	// The type of threshold for the notification. This determines whether the notification is triggered by forecasted costs or actual costs. The allowed values are Actual and Forecasted. Default is Actual. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ThresholdType *string `json:"thresholdType,omitempty" tf:"threshold_type,omitempty"`
}

type BudgetSubscriptionObservation struct {

	// The total amount of cost to track with the budget.
	Amount *float64 `json:"amount,omitempty" tf:"amount,omitempty"`

	// The ETag of the Subscription Consumption Budget.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// A filter block as defined below.
	Filter *BudgetSubscriptionFilterObservation `json:"filter,omitempty" tf:"filter,omitempty"`

	// The ID of the Subscription Consumption Budget.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// One or more notification blocks as defined below.
	Notification []BudgetSubscriptionNotificationObservation `json:"notification,omitempty" tf:"notification,omitempty"`

	// The ID of the Subscription for which to create a Consumption Budget. Changing this forces a new resource to be created.
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`

	// The time covered by a budget. Tracking of the amount will be reset based on the time grain. Must be one of BillingAnnual, BillingMonth, BillingQuarter, Annually, Monthly and Quarterly. Defaults to Monthly. Changing this forces a new resource to be created.
	TimeGrain *string `json:"timeGrain,omitempty" tf:"time_grain,omitempty"`

	// A time_period block as defined below.
	TimePeriod *BudgetSubscriptionTimePeriodObservation `json:"timePeriod,omitempty" tf:"time_period,omitempty"`
}

type BudgetSubscriptionParameters struct {

	// The total amount of cost to track with the budget.
	// +kubebuilder:validation:Optional
	Amount *float64 `json:"amount,omitempty" tf:"amount,omitempty"`

	// The ETag of the Subscription Consumption Budget.
	// +kubebuilder:validation:Optional
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// A filter block as defined below.
	// +kubebuilder:validation:Optional
	Filter *BudgetSubscriptionFilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// One or more notification blocks as defined below.
	// +kubebuilder:validation:Optional
	Notification []BudgetSubscriptionNotificationParameters `json:"notification,omitempty" tf:"notification,omitempty"`

	// The ID of the Subscription for which to create a Consumption Budget. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`

	// The time covered by a budget. Tracking of the amount will be reset based on the time grain. Must be one of BillingAnnual, BillingMonth, BillingQuarter, Annually, Monthly and Quarterly. Defaults to Monthly. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	TimeGrain *string `json:"timeGrain,omitempty" tf:"time_grain,omitempty"`

	// A time_period block as defined below.
	// +kubebuilder:validation:Optional
	TimePeriod *BudgetSubscriptionTimePeriodParameters `json:"timePeriod,omitempty" tf:"time_period,omitempty"`
}

type BudgetSubscriptionTimePeriodInitParameters struct {

	// The end date for the budget. If not set this will be 10 years after the start date.
	EndDate *string `json:"endDate,omitempty" tf:"end_date,omitempty"`

	// The start date for the budget. The start date must be first of the month and should be less than the end date. Budget start date must be on or after June 1, 2017. Future start date should not be more than twelve months. Past start date should be selected within the timegrain period. Changing this forces a new Subscription Consumption Budget to be created.
	StartDate *string `json:"startDate,omitempty" tf:"start_date,omitempty"`
}

type BudgetSubscriptionTimePeriodObservation struct {

	// The end date for the budget. If not set this will be 10 years after the start date.
	EndDate *string `json:"endDate,omitempty" tf:"end_date,omitempty"`

	// The start date for the budget. The start date must be first of the month and should be less than the end date. Budget start date must be on or after June 1, 2017. Future start date should not be more than twelve months. Past start date should be selected within the timegrain period. Changing this forces a new Subscription Consumption Budget to be created.
	StartDate *string `json:"startDate,omitempty" tf:"start_date,omitempty"`
}

type BudgetSubscriptionTimePeriodParameters struct {

	// The end date for the budget. If not set this will be 10 years after the start date.
	// +kubebuilder:validation:Optional
	EndDate *string `json:"endDate,omitempty" tf:"end_date,omitempty"`

	// The start date for the budget. The start date must be first of the month and should be less than the end date. Budget start date must be on or after June 1, 2017. Future start date should not be more than twelve months. Past start date should be selected within the timegrain period. Changing this forces a new Subscription Consumption Budget to be created.
	// +kubebuilder:validation:Optional
	StartDate *string `json:"startDate" tf:"start_date,omitempty"`
}

type FilterNotTagInitParameters struct {

	// The name of the tag to use for the filter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The operator to use for comparison. The allowed values are In. Defaults to In.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Specifies a list of values for the tag.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type FilterNotTagObservation struct {

	// The name of the tag to use for the filter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The operator to use for comparison. The allowed values are In. Defaults to In.
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Specifies a list of values for the tag.
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type FilterNotTagParameters struct {

	// The name of the tag to use for the filter.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The operator to use for comparison. The allowed values are In. Defaults to In.
	// +kubebuilder:validation:Optional
	Operator *string `json:"operator,omitempty" tf:"operator,omitempty"`

	// Specifies a list of values for the tag.
	// +kubebuilder:validation:Optional
	Values []*string `json:"values" tf:"values,omitempty"`
}

// BudgetSubscriptionSpec defines the desired state of BudgetSubscription
type BudgetSubscriptionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BudgetSubscriptionParameters `json:"forProvider"`
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
	InitProvider BudgetSubscriptionInitParameters `json:"initProvider,omitempty"`
}

// BudgetSubscriptionStatus defines the observed state of BudgetSubscription.
type BudgetSubscriptionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BudgetSubscriptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// BudgetSubscription is the Schema for the BudgetSubscriptions API. Manages a Subscription Consumption Budget.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type BudgetSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.amount) || (has(self.initProvider) && has(self.initProvider.amount))",message="spec.forProvider.amount is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.notification) || (has(self.initProvider) && has(self.initProvider.notification))",message="spec.forProvider.notification is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.subscriptionId) || (has(self.initProvider) && has(self.initProvider.subscriptionId))",message="spec.forProvider.subscriptionId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.timePeriod) || (has(self.initProvider) && has(self.initProvider.timePeriod))",message="spec.forProvider.timePeriod is a required parameter"
	Spec   BudgetSubscriptionSpec   `json:"spec"`
	Status BudgetSubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BudgetSubscriptionList contains a list of BudgetSubscriptions
type BudgetSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BudgetSubscription `json:"items"`
}

// Repository type metadata.
var (
	BudgetSubscription_Kind             = "BudgetSubscription"
	BudgetSubscription_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BudgetSubscription_Kind}.String()
	BudgetSubscription_KindAPIVersion   = BudgetSubscription_Kind + "." + CRDGroupVersion.String()
	BudgetSubscription_GroupVersionKind = CRDGroupVersion.WithKind(BudgetSubscription_Kind)
)

func init() {
	SchemeBuilder.Register(&BudgetSubscription{}, &BudgetSubscriptionList{})
}
