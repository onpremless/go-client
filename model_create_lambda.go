/*
onpremless

onpremless api

API version: 1.0.1
Contact: ilya.sumb@gmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package opless

import (
	"encoding/json"
)

// checks if the CreateLambda type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLambda{}

// CreateLambda struct for CreateLambda
type CreateLambda struct {
	Archive string `json:"archive"`
	Name string `json:"name"`
	Runtime string `json:"runtime"`
	LambdaType string `json:"lambda_type"`
}

// NewCreateLambda instantiates a new CreateLambda object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLambda(archive string, name string, runtime string, lambdaType string) *CreateLambda {
	this := CreateLambda{}
	this.Name = name
	this.Runtime = runtime
	this.LambdaType = lambdaType
	return &this
}

// NewCreateLambdaWithDefaults instantiates a new CreateLambda object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLambdaWithDefaults() *CreateLambda {
	this := CreateLambda{}
	return &this
}

// GetArchive returns the Archive field value
func (o *CreateLambda) GetArchive() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Archive
}

// GetArchiveOk returns a tuple with the Archive field value
// and a boolean to check if the value has been set.
func (o *CreateLambda) GetArchiveOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Archive, true
}

// SetArchive sets field value
func (o *CreateLambda) SetArchive(v string) {
	o.Archive = v
}

// GetName returns the Name field value
func (o *CreateLambda) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateLambda) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateLambda) SetName(v string) {
	o.Name = v
}

// GetRuntime returns the Runtime field value
func (o *CreateLambda) GetRuntime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Runtime
}

// GetRuntimeOk returns a tuple with the Runtime field value
// and a boolean to check if the value has been set.
func (o *CreateLambda) GetRuntimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Runtime, true
}

// SetRuntime sets field value
func (o *CreateLambda) SetRuntime(v string) {
	o.Runtime = v
}

// GetLambdaType returns the LambdaType field value
func (o *CreateLambda) GetLambdaType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LambdaType
}

// GetLambdaTypeOk returns a tuple with the LambdaType field value
// and a boolean to check if the value has been set.
func (o *CreateLambda) GetLambdaTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LambdaType, true
}

// SetLambdaType sets field value
func (o *CreateLambda) SetLambdaType(v string) {
	o.LambdaType = v
}

func (o CreateLambda) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateLambda) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["archive"] = o.Archive
	toSerialize["name"] = o.Name
	toSerialize["runtime"] = o.Runtime
	toSerialize["lambda_type"] = o.LambdaType
	return toSerialize, nil
}

type NullableCreateLambda struct {
	value *CreateLambda
	isSet bool
}

func (v NullableCreateLambda) Get() *CreateLambda {
	return v.value
}

func (v *NullableCreateLambda) Set(val *CreateLambda) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLambda) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLambda) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLambda(val *CreateLambda) *NullableCreateLambda {
	return &NullableCreateLambda{value: val, isSet: true}
}

func (v NullableCreateLambda) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLambda) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


