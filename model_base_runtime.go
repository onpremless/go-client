/*
onpremless

onpremless api

API version: 1.0.3
Contact: ilya.sumb@gmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package opless

import (
	"encoding/json"
)

// checks if the BaseRuntime type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseRuntime{}

// BaseRuntime struct for BaseRuntime
type BaseRuntime struct {
	Name string `json:"name"`
}

// NewBaseRuntime instantiates a new BaseRuntime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseRuntime(name string) *BaseRuntime {
	this := BaseRuntime{}
	this.Name = name
	return &this
}

// NewBaseRuntimeWithDefaults instantiates a new BaseRuntime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseRuntimeWithDefaults() *BaseRuntime {
	this := BaseRuntime{}
	return &this
}

// GetName returns the Name field value
func (o *BaseRuntime) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BaseRuntime) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BaseRuntime) SetName(v string) {
	o.Name = v
}

func (o BaseRuntime) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseRuntime) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableBaseRuntime struct {
	value *BaseRuntime
	isSet bool
}

func (v NullableBaseRuntime) Get() *BaseRuntime {
	return v.value
}

func (v *NullableBaseRuntime) Set(val *BaseRuntime) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseRuntime) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseRuntime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseRuntime(val *BaseRuntime) *NullableBaseRuntime {
	return &NullableBaseRuntime{value: val, isSet: true}
}

func (v NullableBaseRuntime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseRuntime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


