/*
Kanthor SDK API

SDK API

API version: 1.0
Contact: support@kanthorlabs.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GatewayError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayError{}

// GatewayError struct for GatewayError
type GatewayError struct {
	Code string `json:"code"`
	Error string `json:"error"`
}

type _GatewayError GatewayError

// NewGatewayError instantiates a new GatewayError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayError(code string, error_ string) *GatewayError {
	this := GatewayError{}
	this.Code = code
	this.Error = error_
	return &this
}

// NewGatewayErrorWithDefaults instantiates a new GatewayError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayErrorWithDefaults() *GatewayError {
	this := GatewayError{}
	return &this
}

// GetCode returns the Code field value
func (o *GatewayError) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *GatewayError) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *GatewayError) SetCode(v string) {
	o.Code = v
}

// GetError returns the Error field value
func (o *GatewayError) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *GatewayError) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *GatewayError) SetError(v string) {
	o.Error = v
}

func (o GatewayError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["error"] = o.Error
	return toSerialize, nil
}

func (o *GatewayError) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"error",
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

	varGatewayError := _GatewayError{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGatewayError)

	if err != nil {
		return err
	}

	*o = GatewayError(varGatewayError)

	return err
}

type NullableGatewayError struct {
	value *GatewayError
	isSet bool
}

func (v NullableGatewayError) Get() *GatewayError {
	return v.value
}

func (v *NullableGatewayError) Set(val *GatewayError) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayError) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayError(val *GatewayError) *NullableGatewayError {
	return &NullableGatewayError{value: val, isSet: true}
}

func (v NullableGatewayError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


