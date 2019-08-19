// +build !ignore_autogenerated

/*
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2019 Red Hat, Inc.
 *
 */
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineDisruptionBudget) DeepCopyInto(out *MachineDisruptionBudget) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineDisruptionBudget.
func (in *MachineDisruptionBudget) DeepCopy() *MachineDisruptionBudget {
	if in == nil {
		return nil
	}
	out := new(MachineDisruptionBudget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineDisruptionBudget) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineDisruptionBudgetList) DeepCopyInto(out *MachineDisruptionBudgetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MachineDisruptionBudget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineDisruptionBudgetList.
func (in *MachineDisruptionBudgetList) DeepCopy() *MachineDisruptionBudgetList {
	if in == nil {
		return nil
	}
	out := new(MachineDisruptionBudgetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineDisruptionBudgetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineDisruptionBudgetSpec) DeepCopyInto(out *MachineDisruptionBudgetSpec) {
	*out = *in
	if in.MinAvailable != nil {
		in, out := &in.MinAvailable, &out.MinAvailable
		*out = new(int32)
		**out = **in
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.MaxUnavailable != nil {
		in, out := &in.MaxUnavailable, &out.MaxUnavailable
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineDisruptionBudgetSpec.
func (in *MachineDisruptionBudgetSpec) DeepCopy() *MachineDisruptionBudgetSpec {
	if in == nil {
		return nil
	}
	out := new(MachineDisruptionBudgetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineDisruptionBudgetStatus) DeepCopyInto(out *MachineDisruptionBudgetStatus) {
	*out = *in
	if in.DisruptedMachines != nil {
		in, out := &in.DisruptedMachines, &out.DisruptedMachines
		*out = make(map[string]v1.Time, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineDisruptionBudgetStatus.
func (in *MachineDisruptionBudgetStatus) DeepCopy() *MachineDisruptionBudgetStatus {
	if in == nil {
		return nil
	}
	out := new(MachineDisruptionBudgetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineHealthCheck) DeepCopyInto(out *MachineHealthCheck) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineHealthCheck.
func (in *MachineHealthCheck) DeepCopy() *MachineHealthCheck {
	if in == nil {
		return nil
	}
	out := new(MachineHealthCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineHealthCheck) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineHealthCheckList) DeepCopyInto(out *MachineHealthCheckList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MachineHealthCheck, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineHealthCheckList.
func (in *MachineHealthCheckList) DeepCopy() *MachineHealthCheckList {
	if in == nil {
		return nil
	}
	out := new(MachineHealthCheckList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineHealthCheckList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineHealthCheckSpec) DeepCopyInto(out *MachineHealthCheckSpec) {
	*out = *in
	if in.RemediationStrategy != nil {
		in, out := &in.RemediationStrategy, &out.RemediationStrategy
		*out = new(RemediationStrategyType)
		**out = **in
	}
	in.Selector.DeepCopyInto(&out.Selector)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineHealthCheckSpec.
func (in *MachineHealthCheckSpec) DeepCopy() *MachineHealthCheckSpec {
	if in == nil {
		return nil
	}
	out := new(MachineHealthCheckSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineHealthCheckStatus) DeepCopyInto(out *MachineHealthCheckStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineHealthCheckStatus.
func (in *MachineHealthCheckStatus) DeepCopy() *MachineHealthCheckStatus {
	if in == nil {
		return nil
	}
	out := new(MachineHealthCheckStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineRemediation) DeepCopyInto(out *MachineRemediation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineRemediation.
func (in *MachineRemediation) DeepCopy() *MachineRemediation {
	if in == nil {
		return nil
	}
	out := new(MachineRemediation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineRemediation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineRemediationList) DeepCopyInto(out *MachineRemediationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MachineRemediation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineRemediationList.
func (in *MachineRemediationList) DeepCopy() *MachineRemediationList {
	if in == nil {
		return nil
	}
	out := new(MachineRemediationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineRemediationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineRemediationOperator) DeepCopyInto(out *MachineRemediationOperator) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineRemediationOperator.
func (in *MachineRemediationOperator) DeepCopy() *MachineRemediationOperator {
	if in == nil {
		return nil
	}
	out := new(MachineRemediationOperator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineRemediationOperator) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineRemediationOperatorList) DeepCopyInto(out *MachineRemediationOperatorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MachineRemediationOperator, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineRemediationOperatorList.
func (in *MachineRemediationOperatorList) DeepCopy() *MachineRemediationOperatorList {
	if in == nil {
		return nil
	}
	out := new(MachineRemediationOperatorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineRemediationOperatorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineRemediationOperatorSpec) DeepCopyInto(out *MachineRemediationOperatorSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineRemediationOperatorSpec.
func (in *MachineRemediationOperatorSpec) DeepCopy() *MachineRemediationOperatorSpec {
	if in == nil {
		return nil
	}
	out := new(MachineRemediationOperatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineRemediationOperatorStatus) DeepCopyInto(out *MachineRemediationOperatorStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]MachineRemediationOperatorStatusCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineRemediationOperatorStatus.
func (in *MachineRemediationOperatorStatus) DeepCopy() *MachineRemediationOperatorStatus {
	if in == nil {
		return nil
	}
	out := new(MachineRemediationOperatorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineRemediationOperatorStatusCondition) DeepCopyInto(out *MachineRemediationOperatorStatusCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineRemediationOperatorStatusCondition.
func (in *MachineRemediationOperatorStatusCondition) DeepCopy() *MachineRemediationOperatorStatusCondition {
	if in == nil {
		return nil
	}
	out := new(MachineRemediationOperatorStatusCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineRemediationSpec) DeepCopyInto(out *MachineRemediationSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineRemediationSpec.
func (in *MachineRemediationSpec) DeepCopy() *MachineRemediationSpec {
	if in == nil {
		return nil
	}
	out := new(MachineRemediationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineRemediationStatus) DeepCopyInto(out *MachineRemediationStatus) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.EndTime != nil {
		in, out := &in.EndTime, &out.EndTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineRemediationStatus.
func (in *MachineRemediationStatus) DeepCopy() *MachineRemediationStatus {
	if in == nil {
		return nil
	}
	out := new(MachineRemediationStatus)
	in.DeepCopyInto(out)
	return out
}
