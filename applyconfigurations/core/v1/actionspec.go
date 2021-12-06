/*
Copyright The Kubernetes Authors.

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

package v1

// ActionSpecApplyConfiguration represents an declarative configuration of the ActionSpec type for use
// with apply.
type ActionSpecApplyConfiguration struct {
	ResourceID   *string                         `json:"resourceID,omitempty"`
	ResourceType *string                         `json:"resourceType,omitempty"`
	Action       *ActionParamsApplyConfiguration `json:"action,omitempty"`
	NodeName     *string                         `json:"nodeName,omitempty"`
	PodName      *string                         `json:"podName,omitempty"`
	PodID        *string                         `json:"podID,omitempty"`
}

// ActionSpecApplyConfiguration constructs an declarative configuration of the ActionSpec type for use with
// apply.
func ActionSpec() *ActionSpecApplyConfiguration {
	return &ActionSpecApplyConfiguration{}
}

// WithResourceID sets the ResourceID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceID field is set to the value of the last call.
func (b *ActionSpecApplyConfiguration) WithResourceID(value string) *ActionSpecApplyConfiguration {
	b.ResourceID = &value
	return b
}

// WithResourceType sets the ResourceType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceType field is set to the value of the last call.
func (b *ActionSpecApplyConfiguration) WithResourceType(value string) *ActionSpecApplyConfiguration {
	b.ResourceType = &value
	return b
}

// WithAction sets the Action field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Action field is set to the value of the last call.
func (b *ActionSpecApplyConfiguration) WithAction(value *ActionParamsApplyConfiguration) *ActionSpecApplyConfiguration {
	b.Action = value
	return b
}

// WithNodeName sets the NodeName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodeName field is set to the value of the last call.
func (b *ActionSpecApplyConfiguration) WithNodeName(value string) *ActionSpecApplyConfiguration {
	b.NodeName = &value
	return b
}

// WithPodName sets the PodName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PodName field is set to the value of the last call.
func (b *ActionSpecApplyConfiguration) WithPodName(value string) *ActionSpecApplyConfiguration {
	b.PodName = &value
	return b
}

// WithPodID sets the PodID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PodID field is set to the value of the last call.
func (b *ActionSpecApplyConfiguration) WithPodID(value string) *ActionSpecApplyConfiguration {
	b.PodID = &value
	return b
}
