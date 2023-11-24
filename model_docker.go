/*
onpremless

onpremless api

API version: 1.0.2
Contact: ilya.sumb@gmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package opless

import (
	"encoding/json"
)

// checks if the Docker type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Docker{}

// Docker struct for Docker
type Docker struct {
	Image *string `json:"image,omitempty"`
	Container *string `json:"container,omitempty"`
	ContainerId *string `json:"container_id,omitempty"`
	Status string `json:"status"`
}

// NewDocker instantiates a new Docker object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocker(status string) *Docker {
	this := Docker{}
	this.Status = status
	return &this
}

// NewDockerWithDefaults instantiates a new Docker object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDockerWithDefaults() *Docker {
	this := Docker{}
	return &this
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *Docker) GetImage() string {
	if o == nil || IsNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Docker) GetImageOk() (*string, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *Docker) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *Docker) SetImage(v string) {
	o.Image = &v
}

// GetContainer returns the Container field value if set, zero value otherwise.
func (o *Docker) GetContainer() string {
	if o == nil || IsNil(o.Container) {
		var ret string
		return ret
	}
	return *o.Container
}

// GetContainerOk returns a tuple with the Container field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Docker) GetContainerOk() (*string, bool) {
	if o == nil || IsNil(o.Container) {
		return nil, false
	}
	return o.Container, true
}

// HasContainer returns a boolean if a field has been set.
func (o *Docker) HasContainer() bool {
	if o != nil && !IsNil(o.Container) {
		return true
	}

	return false
}

// SetContainer gets a reference to the given string and assigns it to the Container field.
func (o *Docker) SetContainer(v string) {
	o.Container = &v
}

// GetContainerId returns the ContainerId field value if set, zero value otherwise.
func (o *Docker) GetContainerId() string {
	if o == nil || IsNil(o.ContainerId) {
		var ret string
		return ret
	}
	return *o.ContainerId
}

// GetContainerIdOk returns a tuple with the ContainerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Docker) GetContainerIdOk() (*string, bool) {
	if o == nil || IsNil(o.ContainerId) {
		return nil, false
	}
	return o.ContainerId, true
}

// HasContainerId returns a boolean if a field has been set.
func (o *Docker) HasContainerId() bool {
	if o != nil && !IsNil(o.ContainerId) {
		return true
	}

	return false
}

// SetContainerId gets a reference to the given string and assigns it to the ContainerId field.
func (o *Docker) SetContainerId(v string) {
	o.ContainerId = &v
}

// GetStatus returns the Status field value
func (o *Docker) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Docker) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Docker) SetStatus(v string) {
	o.Status = v
}

func (o Docker) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Docker) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.Container) {
		toSerialize["container"] = o.Container
	}
	if !IsNil(o.ContainerId) {
		toSerialize["container_id"] = o.ContainerId
	}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullableDocker struct {
	value *Docker
	isSet bool
}

func (v NullableDocker) Get() *Docker {
	return v.value
}

func (v *NullableDocker) Set(val *Docker) {
	v.value = val
	v.isSet = true
}

func (v NullableDocker) IsSet() bool {
	return v.isSet
}

func (v *NullableDocker) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocker(val *Docker) *NullableDocker {
	return &NullableDocker{value: val, isSet: true}
}

func (v NullableDocker) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocker) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


