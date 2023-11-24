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

// checks if the CreateRuntime type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRuntime{}

// CreateRuntime struct for CreateRuntime
type CreateRuntime struct {
	Dockerfile string `json:"dockerfile"`
	Name string `json:"name"`
}

// NewCreateRuntime instantiates a new CreateRuntime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRuntime(dockerfile string, name string) *CreateRuntime {
	this := CreateRuntime{}
	this.Name = name
	return &this
}

// NewCreateRuntimeWithDefaults instantiates a new CreateRuntime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRuntimeWithDefaults() *CreateRuntime {
	this := CreateRuntime{}
	return &this
}

// GetDockerfile returns the Dockerfile field value
func (o *CreateRuntime) GetDockerfile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dockerfile
}

// GetDockerfileOk returns a tuple with the Dockerfile field value
// and a boolean to check if the value has been set.
func (o *CreateRuntime) GetDockerfileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dockerfile, true
}

// SetDockerfile sets field value
func (o *CreateRuntime) SetDockerfile(v string) {
	o.Dockerfile = v
}

// GetName returns the Name field value
func (o *CreateRuntime) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateRuntime) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateRuntime) SetName(v string) {
	o.Name = v
}

func (o CreateRuntime) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRuntime) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dockerfile"] = o.Dockerfile
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableCreateRuntime struct {
	value *CreateRuntime
	isSet bool
}

func (v NullableCreateRuntime) Get() *CreateRuntime {
	return v.value
}

func (v *NullableCreateRuntime) Set(val *CreateRuntime) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRuntime) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRuntime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRuntime(val *CreateRuntime) *NullableCreateRuntime {
	return &NullableCreateRuntime{value: val, isSet: true}
}

func (v NullableCreateRuntime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRuntime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


