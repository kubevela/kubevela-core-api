/*
Copyright 2023 The KubeVela Authors.

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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// AppPolicyApplyConfiguration represents an declarative configuration of the AppPolicy type for use
// with apply.
type AppPolicyApplyConfiguration struct {
	Name       *string               `json:"name,omitempty"`
	Type       *string               `json:"type,omitempty"`
	Properties *runtime.RawExtension `json:"properties,omitempty"`
}

// AppPolicyApplyConfiguration constructs an declarative configuration of the AppPolicy type for use with
// apply.
func AppPolicy() *AppPolicyApplyConfiguration {
	return &AppPolicyApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *AppPolicyApplyConfiguration) WithName(value string) *AppPolicyApplyConfiguration {
	b.Name = &value
	return b
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *AppPolicyApplyConfiguration) WithType(value string) *AppPolicyApplyConfiguration {
	b.Type = &value
	return b
}

// WithProperties sets the Properties field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Properties field is set to the value of the last call.
func (b *AppPolicyApplyConfiguration) WithProperties(value runtime.RawExtension) *AppPolicyApplyConfiguration {
	b.Properties = &value
	return b
}
