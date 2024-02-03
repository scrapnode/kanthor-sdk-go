/*
Kanthor SDK API

SDK API

API version: 2022.1201.1701
Contact: support@kanthorlabs.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ApplicationCreateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationCreateReq{}

// ApplicationCreateReq struct for ApplicationCreateReq
type ApplicationCreateReq struct {
	Name string `json:"name"`
}

type _ApplicationCreateReq ApplicationCreateReq

// NewApplicationCreateReq instantiates a new ApplicationCreateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationCreateReq(name string) *ApplicationCreateReq {
	this := ApplicationCreateReq{}
	this.Name = name
	return &this
}

// NewApplicationCreateReqWithDefaults instantiates a new ApplicationCreateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationCreateReqWithDefaults() *ApplicationCreateReq {
	this := ApplicationCreateReq{}
	return &this
}

// GetName returns the Name field value
func (o *ApplicationCreateReq) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationCreateReq) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplicationCreateReq) SetName(v string) {
	o.Name = v
}

func (o ApplicationCreateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationCreateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *ApplicationCreateReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varApplicationCreateReq := _ApplicationCreateReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApplicationCreateReq)

	if err != nil {
		return err
	}

	*o = ApplicationCreateReq(varApplicationCreateReq)

	return err
}

type NullableApplicationCreateReq struct {
	value *ApplicationCreateReq
	isSet bool
}

func (v NullableApplicationCreateReq) Get() *ApplicationCreateReq {
	return v.value
}

func (v *NullableApplicationCreateReq) Set(val *ApplicationCreateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationCreateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationCreateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationCreateReq(val *ApplicationCreateReq) *NullableApplicationCreateReq {
	return &NullableApplicationCreateReq{value: val, isSet: true}
}

func (v NullableApplicationCreateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationCreateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


