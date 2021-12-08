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

// AddressApplyConfiguration represents an declarative configuration of the Address type for use
// with apply.
type AddressApplyConfiguration struct {
	NetworkPlane *string                         `json:"networkPlane,omitempty"`
	Addresses    []AddressInfoApplyConfiguration `json:"addresses,omitempty"`
}

// AddressApplyConfiguration constructs an declarative configuration of the Address type for use with
// apply.
func Address() *AddressApplyConfiguration {
	return &AddressApplyConfiguration{}
}

// WithNetworkPlane sets the NetworkPlane field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NetworkPlane field is set to the value of the last call.
func (b *AddressApplyConfiguration) WithNetworkPlane(value string) *AddressApplyConfiguration {
	b.NetworkPlane = &value
	return b
}

// WithAddresses adds the given value to the Addresses field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Addresses field.
func (b *AddressApplyConfiguration) WithAddresses(values ...*AddressInfoApplyConfiguration) *AddressApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithAddresses")
		}
		b.Addresses = append(b.Addresses, *values[i])
	}
	return b
}