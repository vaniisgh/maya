// +build !ignore_autogenerated

/*
Copyright 2019 The OpenEBS Authors

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataItem) DeepCopyInto(out *DataItem) {
	*out = *in
	if in.Entries != nil {
		in, out := &in.Entries, &out.Entries
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataItem.
func (in *DataItem) DeepCopy() *DataItem {
	if in == nil {
		return nil
	}
	out := new(DataItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceDetails) DeepCopyInto(out *ResourceDetails) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceDetails.
func (in *ResourceDetails) DeepCopy() *ResourceDetails {
	if in == nil {
		return nil
	}
	out := new(ResourceDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceState) DeepCopyInto(out *ResourceState) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceState.
func (in *ResourceState) DeepCopy() *ResourceState {
	if in == nil {
		return nil
	}
	out := new(ResourceState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeConfig) DeepCopyInto(out *UpgradeConfig) {
	*out = *in
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make([]DataItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ResourceDetails, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeConfig.
func (in *UpgradeConfig) DeepCopy() *UpgradeConfig {
	if in == nil {
		return nil
	}
	out := new(UpgradeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeResource) DeepCopyInto(out *UpgradeResource) {
	*out = *in
	out.ResourceDetails = in.ResourceDetails
	in.PreState.DeepCopyInto(&out.PreState)
	in.PostState.DeepCopyInto(&out.PostState)
	if in.SubResources != nil {
		in, out := &in.SubResources, &out.SubResources
		*out = make([]UpgradeSubResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeResource.
func (in *UpgradeResource) DeepCopy() *UpgradeResource {
	if in == nil {
		return nil
	}
	out := new(UpgradeResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeResult) DeepCopyInto(out *UpgradeResult) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Config.DeepCopyInto(&out.Config)
	if in.Tasks != nil {
		in, out := &in.Tasks, &out.Tasks
		*out = make([]UpgradeResultTask, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeResult.
func (in *UpgradeResult) DeepCopy() *UpgradeResult {
	if in == nil {
		return nil
	}
	out := new(UpgradeResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UpgradeResult) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeResultConfig) DeepCopyInto(out *UpgradeResultConfig) {
	*out = *in
	out.ResourceDetails = in.ResourceDetails
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make([]DataItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeResultConfig.
func (in *UpgradeResultConfig) DeepCopy() *UpgradeResultConfig {
	if in == nil {
		return nil
	}
	out := new(UpgradeResultConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeResultList) DeepCopyInto(out *UpgradeResultList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]UpgradeResult, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeResultList.
func (in *UpgradeResultList) DeepCopy() *UpgradeResultList {
	if in == nil {
		return nil
	}
	out := new(UpgradeResultList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UpgradeResultList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeResultStatus) DeepCopyInto(out *UpgradeResultStatus) {
	*out = *in
	in.Resource.DeepCopyInto(&out.Resource)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeResultStatus.
func (in *UpgradeResultStatus) DeepCopy() *UpgradeResultStatus {
	if in == nil {
		return nil
	}
	out := new(UpgradeResultStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeResultTask) DeepCopyInto(out *UpgradeResultTask) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeResultTask.
func (in *UpgradeResultTask) DeepCopy() *UpgradeResultTask {
	if in == nil {
		return nil
	}
	out := new(UpgradeResultTask)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeSubResource) DeepCopyInto(out *UpgradeSubResource) {
	*out = *in
	out.ResourceDetails = in.ResourceDetails
	in.PreState.DeepCopyInto(&out.PreState)
	in.PostState.DeepCopyInto(&out.PostState)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeSubResource.
func (in *UpgradeSubResource) DeepCopy() *UpgradeSubResource {
	if in == nil {
		return nil
	}
	out := new(UpgradeSubResource)
	in.DeepCopyInto(out)
	return out
}