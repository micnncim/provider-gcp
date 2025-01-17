//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2019 The Crossplane Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CryptoKey) DeepCopyInto(out *CryptoKey) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CryptoKey.
func (in *CryptoKey) DeepCopy() *CryptoKey {
	if in == nil {
		return nil
	}
	out := new(CryptoKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CryptoKey) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CryptoKeyList) DeepCopyInto(out *CryptoKeyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CryptoKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CryptoKeyList.
func (in *CryptoKeyList) DeepCopy() *CryptoKeyList {
	if in == nil {
		return nil
	}
	out := new(CryptoKeyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CryptoKeyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CryptoKeyObservation) DeepCopyInto(out *CryptoKeyObservation) {
	*out = *in
	if in.Primary != nil {
		in, out := &in.Primary, &out.Primary
		*out = new(CryptoKeyVersion)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CryptoKeyObservation.
func (in *CryptoKeyObservation) DeepCopy() *CryptoKeyObservation {
	if in == nil {
		return nil
	}
	out := new(CryptoKeyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CryptoKeyParameters) DeepCopyInto(out *CryptoKeyParameters) {
	*out = *in
	if in.KeyRing != nil {
		in, out := &in.KeyRing, &out.KeyRing
		*out = new(string)
		**out = **in
	}
	if in.KeyRingRef != nil {
		in, out := &in.KeyRingRef, &out.KeyRingRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.KeyRingSelector != nil {
		in, out := &in.KeyRingSelector, &out.KeyRingSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.RotationPeriod != nil {
		in, out := &in.RotationPeriod, &out.RotationPeriod
		*out = new(string)
		**out = **in
	}
	if in.NextRotationTime != nil {
		in, out := &in.NextRotationTime, &out.NextRotationTime
		*out = new(string)
		**out = **in
	}
	if in.VersionTemplate != nil {
		in, out := &in.VersionTemplate, &out.VersionTemplate
		*out = new(CryptoKeyVersionTemplate)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CryptoKeyParameters.
func (in *CryptoKeyParameters) DeepCopy() *CryptoKeyParameters {
	if in == nil {
		return nil
	}
	out := new(CryptoKeyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CryptoKeyPolicy) DeepCopyInto(out *CryptoKeyPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CryptoKeyPolicy.
func (in *CryptoKeyPolicy) DeepCopy() *CryptoKeyPolicy {
	if in == nil {
		return nil
	}
	out := new(CryptoKeyPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CryptoKeyPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CryptoKeyPolicyList) DeepCopyInto(out *CryptoKeyPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CryptoKeyPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CryptoKeyPolicyList.
func (in *CryptoKeyPolicyList) DeepCopy() *CryptoKeyPolicyList {
	if in == nil {
		return nil
	}
	out := new(CryptoKeyPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CryptoKeyPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CryptoKeyPolicyParameters) DeepCopyInto(out *CryptoKeyPolicyParameters) {
	*out = *in
	if in.CryptoKey != nil {
		in, out := &in.CryptoKey, &out.CryptoKey
		*out = new(string)
		**out = **in
	}
	if in.CryptoKeyRef != nil {
		in, out := &in.CryptoKeyRef, &out.CryptoKeyRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.CryptoKeySelector != nil {
		in, out := &in.CryptoKeySelector, &out.CryptoKeySelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	in.Policy.DeepCopyInto(&out.Policy)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CryptoKeyPolicyParameters.
func (in *CryptoKeyPolicyParameters) DeepCopy() *CryptoKeyPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(CryptoKeyPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CryptoKeyPolicySpec) DeepCopyInto(out *CryptoKeyPolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CryptoKeyPolicySpec.
func (in *CryptoKeyPolicySpec) DeepCopy() *CryptoKeyPolicySpec {
	if in == nil {
		return nil
	}
	out := new(CryptoKeyPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CryptoKeyPolicyStatus) DeepCopyInto(out *CryptoKeyPolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CryptoKeyPolicyStatus.
func (in *CryptoKeyPolicyStatus) DeepCopy() *CryptoKeyPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(CryptoKeyPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CryptoKeySpec) DeepCopyInto(out *CryptoKeySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CryptoKeySpec.
func (in *CryptoKeySpec) DeepCopy() *CryptoKeySpec {
	if in == nil {
		return nil
	}
	out := new(CryptoKeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CryptoKeyStatus) DeepCopyInto(out *CryptoKeyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CryptoKeyStatus.
func (in *CryptoKeyStatus) DeepCopy() *CryptoKeyStatus {
	if in == nil {
		return nil
	}
	out := new(CryptoKeyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CryptoKeyVersion) DeepCopyInto(out *CryptoKeyVersion) {
	*out = *in
	if in.Attestation != nil {
		in, out := &in.Attestation, &out.Attestation
		*out = new(KeyOperationAttestation)
		**out = **in
	}
	if in.ExternalProtectionLevelOptions != nil {
		in, out := &in.ExternalProtectionLevelOptions, &out.ExternalProtectionLevelOptions
		*out = new(ExternalProtectionLevelOptions)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CryptoKeyVersion.
func (in *CryptoKeyVersion) DeepCopy() *CryptoKeyVersion {
	if in == nil {
		return nil
	}
	out := new(CryptoKeyVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CryptoKeyVersionTemplate) DeepCopyInto(out *CryptoKeyVersionTemplate) {
	*out = *in
	if in.Algorithm != nil {
		in, out := &in.Algorithm, &out.Algorithm
		*out = new(string)
		**out = **in
	}
	if in.ProtectionLevel != nil {
		in, out := &in.ProtectionLevel, &out.ProtectionLevel
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CryptoKeyVersionTemplate.
func (in *CryptoKeyVersionTemplate) DeepCopy() *CryptoKeyVersionTemplate {
	if in == nil {
		return nil
	}
	out := new(CryptoKeyVersionTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalProtectionLevelOptions) DeepCopyInto(out *ExternalProtectionLevelOptions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalProtectionLevelOptions.
func (in *ExternalProtectionLevelOptions) DeepCopy() *ExternalProtectionLevelOptions {
	if in == nil {
		return nil
	}
	out := new(ExternalProtectionLevelOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyOperationAttestation) DeepCopyInto(out *KeyOperationAttestation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyOperationAttestation.
func (in *KeyOperationAttestation) DeepCopy() *KeyOperationAttestation {
	if in == nil {
		return nil
	}
	out := new(KeyOperationAttestation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyRing) DeepCopyInto(out *KeyRing) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyRing.
func (in *KeyRing) DeepCopy() *KeyRing {
	if in == nil {
		return nil
	}
	out := new(KeyRing)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeyRing) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyRingList) DeepCopyInto(out *KeyRingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KeyRing, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyRingList.
func (in *KeyRingList) DeepCopy() *KeyRingList {
	if in == nil {
		return nil
	}
	out := new(KeyRingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeyRingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyRingObservation) DeepCopyInto(out *KeyRingObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyRingObservation.
func (in *KeyRingObservation) DeepCopy() *KeyRingObservation {
	if in == nil {
		return nil
	}
	out := new(KeyRingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyRingParameters) DeepCopyInto(out *KeyRingParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyRingParameters.
func (in *KeyRingParameters) DeepCopy() *KeyRingParameters {
	if in == nil {
		return nil
	}
	out := new(KeyRingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyRingSpec) DeepCopyInto(out *KeyRingSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyRingSpec.
func (in *KeyRingSpec) DeepCopy() *KeyRingSpec {
	if in == nil {
		return nil
	}
	out := new(KeyRingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyRingStatus) DeepCopyInto(out *KeyRingStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyRingStatus.
func (in *KeyRingStatus) DeepCopy() *KeyRingStatus {
	if in == nil {
		return nil
	}
	out := new(KeyRingStatus)
	in.DeepCopyInto(out)
	return out
}
