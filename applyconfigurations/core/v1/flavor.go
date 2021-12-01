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

// FlavorApplyConfiguration represents an declarative configuration of the Flavor type for use
// with apply.
type FlavorApplyConfiguration struct {
	ID           *string `json:"id,omitempty"`
	Ram          *int32  `json:"ram,omitempty"`
	CPU          *int32  `json:"vcpus,omitempty"`
	Disk         *int32  `json:"disk,omitempty"`
	IsPublic     *bool   `json:"osFlavorAccessIsPublic,omitempty"`
	ExtEphemeral *int32  `json:"osFLVExtDataEphemeral,omitempty"`
}

// FlavorApplyConfiguration constructs an declarative configuration of the Flavor type for use with
// apply.
func Flavor() *FlavorApplyConfiguration {
	return &FlavorApplyConfiguration{}
}

// WithID sets the ID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ID field is set to the value of the last call.
func (b *FlavorApplyConfiguration) WithID(value string) *FlavorApplyConfiguration {
	b.ID = &value
	return b
}

// WithRam sets the Ram field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Ram field is set to the value of the last call.
func (b *FlavorApplyConfiguration) WithRam(value int32) *FlavorApplyConfiguration {
	b.Ram = &value
	return b
}

// WithCPU sets the CPU field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CPU field is set to the value of the last call.
func (b *FlavorApplyConfiguration) WithCPU(value int32) *FlavorApplyConfiguration {
	b.CPU = &value
	return b
}

// WithDisk sets the Disk field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Disk field is set to the value of the last call.
func (b *FlavorApplyConfiguration) WithDisk(value int32) *FlavorApplyConfiguration {
	b.Disk = &value
	return b
}

// WithIsPublic sets the IsPublic field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IsPublic field is set to the value of the last call.
func (b *FlavorApplyConfiguration) WithIsPublic(value bool) *FlavorApplyConfiguration {
	b.IsPublic = &value
	return b
}

// WithExtEphemeral sets the ExtEphemeral field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ExtEphemeral field is set to the value of the last call.
func (b *FlavorApplyConfiguration) WithExtEphemeral(value int32) *FlavorApplyConfiguration {
	b.ExtEphemeral = &value
	return b
}
