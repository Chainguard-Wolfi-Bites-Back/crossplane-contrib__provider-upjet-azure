// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	apisresolver "github.com/upbound/provider-azure/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *IOTSecurityDeviceGroup) ResolveReferences( // ResolveReferences of this IOTSecurityDeviceGroup.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("devices.azure.upbound.io", "v1beta1", "IOTHub", "IOTHubList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IOTHubID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.IOTHubIDRef,
			Selector:     mg.Spec.ForProvider.IOTHubIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IOTHubID")
	}
	mg.Spec.ForProvider.IOTHubID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.IOTHubIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("devices.azure.upbound.io", "v1beta1", "IOTHub", "IOTHubList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.IOTHubID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.IOTHubIDRef,
			Selector:     mg.Spec.InitProvider.IOTHubIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.IOTHubID")
	}
	mg.Spec.InitProvider.IOTHubID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.IOTHubIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this IOTSecuritySolution.
func (mg *IOTSecuritySolution) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("devices.azure.upbound.io", "v1beta1", "IOTHub", "IOTHubList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.IOTHubIds),
			Extract:       resource.ExtractParamPath("id", true),
			References:    mg.Spec.ForProvider.IOTHubIdsRefs,
			Selector:      mg.Spec.ForProvider.IOTHubIdsSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IOTHubIds")
	}
	mg.Spec.ForProvider.IOTHubIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.IOTHubIdsRefs = mrsp.ResolvedReferences
	{
		m, l, err = apisresolver.GetManagedResource("azure.upbound.io", "v1beta1", "ResourceGroup", "ResourceGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
			Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("devices.azure.upbound.io", "v1beta1", "IOTHub", "IOTHubList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.IOTHubIds),
			Extract:       resource.ExtractParamPath("id", true),
			References:    mg.Spec.InitProvider.IOTHubIdsRefs,
			Selector:      mg.Spec.InitProvider.IOTHubIdsSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.IOTHubIds")
	}
	mg.Spec.InitProvider.IOTHubIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.IOTHubIdsRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this SecurityCenterAssessment.
func (mg *SecurityCenterAssessment) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("security.azure.upbound.io", "v1beta1", "SecurityCenterAssessmentPolicy", "SecurityCenterAssessmentPolicyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AssessmentPolicyID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.AssessmentPolicyIDRef,
			Selector:     mg.Spec.ForProvider.AssessmentPolicyIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AssessmentPolicyID")
	}
	mg.Spec.ForProvider.AssessmentPolicyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AssessmentPolicyIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("compute.azure.upbound.io", "v1beta1", "LinuxVirtualMachineScaleSet", "LinuxVirtualMachineScaleSetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TargetResourceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.TargetResourceIDRef,
			Selector:     mg.Spec.ForProvider.TargetResourceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TargetResourceID")
	}
	mg.Spec.ForProvider.TargetResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TargetResourceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("security.azure.upbound.io", "v1beta1", "SecurityCenterAssessmentPolicy", "SecurityCenterAssessmentPolicyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.AssessmentPolicyID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.AssessmentPolicyIDRef,
			Selector:     mg.Spec.InitProvider.AssessmentPolicyIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.AssessmentPolicyID")
	}
	mg.Spec.InitProvider.AssessmentPolicyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.AssessmentPolicyIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("compute.azure.upbound.io", "v1beta1", "LinuxVirtualMachineScaleSet", "LinuxVirtualMachineScaleSetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.TargetResourceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.TargetResourceIDRef,
			Selector:     mg.Spec.InitProvider.TargetResourceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.TargetResourceID")
	}
	mg.Spec.InitProvider.TargetResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.TargetResourceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SecurityCenterServerVulnerabilityAssessment.
func (mg *SecurityCenterServerVulnerabilityAssessment) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("compute.azure.upbound.io", "v1beta2", "LinuxVirtualMachine", "LinuxVirtualMachineList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VirtualMachineID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.VirtualMachineIDRef,
			Selector:     mg.Spec.ForProvider.VirtualMachineIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VirtualMachineID")
	}
	mg.Spec.ForProvider.VirtualMachineID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VirtualMachineIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("compute.azure.upbound.io", "v1beta2", "LinuxVirtualMachine", "LinuxVirtualMachineList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.VirtualMachineID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.VirtualMachineIDRef,
			Selector:     mg.Spec.InitProvider.VirtualMachineIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.VirtualMachineID")
	}
	mg.Spec.InitProvider.VirtualMachineID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.VirtualMachineIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SecurityCenterServerVulnerabilityAssessmentVirtualMachine.
func (mg *SecurityCenterServerVulnerabilityAssessmentVirtualMachine) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("compute.azure.upbound.io", "v1beta2", "LinuxVirtualMachine", "LinuxVirtualMachineList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VirtualMachineID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.VirtualMachineIDRef,
			Selector:     mg.Spec.ForProvider.VirtualMachineIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VirtualMachineID")
	}
	mg.Spec.ForProvider.VirtualMachineID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VirtualMachineIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SecurityCenterWorkspace.
func (mg *SecurityCenterWorkspace) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("operationalinsights.azure.upbound.io", "v1beta2", "Workspace", "WorkspaceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.WorkspaceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.WorkspaceIDRef,
			Selector:     mg.Spec.ForProvider.WorkspaceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.WorkspaceID")
	}
	mg.Spec.ForProvider.WorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.WorkspaceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("operationalinsights.azure.upbound.io", "v1beta2", "Workspace", "WorkspaceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.WorkspaceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.WorkspaceIDRef,
			Selector:     mg.Spec.InitProvider.WorkspaceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.WorkspaceID")
	}
	mg.Spec.InitProvider.WorkspaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.WorkspaceIDRef = rsp.ResolvedReference

	return nil
}
