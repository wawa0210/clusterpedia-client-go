// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// TODO: if the apis package of clusterpedia is separated, the file will be remove.

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterGroupStatus) DeepCopyInto(out *ClusterGroupStatus) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ClusterResourceStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterGroupStatus.
func (in *ClusterGroupStatus) DeepCopy() *ClusterGroupStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterResource) DeepCopyInto(out *ClusterResource) {
	*out = *in
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterResource.
func (in *ClusterResource) DeepCopy() *ClusterResource {
	if in == nil {
		return nil
	}
	out := new(ClusterResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterResourceStatus) DeepCopyInto(out *ClusterResourceStatus) {
	*out = *in
	if in.SyncConditions != nil {
		in, out := &in.SyncConditions, &out.SyncConditions
		*out = make([]ClusterResourceSyncCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterResourceStatus.
func (in *ClusterResourceStatus) DeepCopy() *ClusterResourceStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterResourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterResourceSyncCondition) DeepCopyInto(out *ClusterResourceSyncCondition) {
	*out = *in
	if in.StorageResource != nil {
		in, out := &in.StorageResource, &out.StorageResource
		*out = new(string)
		**out = **in
	}
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterResourceSyncCondition.
func (in *ClusterResourceSyncCondition) DeepCopy() *ClusterResourceSyncCondition {
	if in == nil {
		return nil
	}
	out := new(ClusterResourceSyncCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	if in.TokenData != nil {
		in, out := &in.TokenData, &out.TokenData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.CAData != nil {
		in, out := &in.CAData, &out.CAData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.CertData != nil {
		in, out := &in.CertData, &out.CertData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.KeyData != nil {
		in, out := &in.KeyData, &out.KeyData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ClusterResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ClusterGroupStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PediaCluster) DeepCopyInto(out *PediaCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PediaCluster.
func (in *PediaCluster) DeepCopy() *PediaCluster {
	if in == nil {
		return nil
	}
	out := new(PediaCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PediaCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PediaClusterList) DeepCopyInto(out *PediaClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PediaCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PediaClusterList.
func (in *PediaClusterList) DeepCopy() *PediaClusterList {
	if in == nil {
		return nil
	}
	out := new(PediaClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PediaClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
