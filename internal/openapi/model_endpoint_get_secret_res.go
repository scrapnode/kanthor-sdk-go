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

// checks if the EndpointGetSecretRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndpointGetSecretRes{}

// EndpointGetSecretRes struct for EndpointGetSecretRes
type EndpointGetSecretRes struct {
	SecretKey string `json:"secret_key"`
}

type _EndpointGetSecretRes EndpointGetSecretRes

// NewEndpointGetSecretRes instantiates a new EndpointGetSecretRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointGetSecretRes(secretKey string) *EndpointGetSecretRes {
	this := EndpointGetSecretRes{}
	this.SecretKey = secretKey
	return &this
}

// NewEndpointGetSecretResWithDefaults instantiates a new EndpointGetSecretRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointGetSecretResWithDefaults() *EndpointGetSecretRes {
	this := EndpointGetSecretRes{}
	return &this
}

// GetSecretKey returns the SecretKey field value
func (o *EndpointGetSecretRes) GetSecretKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value
// and a boolean to check if the value has been set.
func (o *EndpointGetSecretRes) GetSecretKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretKey, true
}

// SetSecretKey sets field value
func (o *EndpointGetSecretRes) SetSecretKey(v string) {
	o.SecretKey = v
}

func (o EndpointGetSecretRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointGetSecretRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["secret_key"] = o.SecretKey
	return toSerialize, nil
}

func (o *EndpointGetSecretRes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"secret_key",
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

	varEndpointGetSecretRes := _EndpointGetSecretRes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEndpointGetSecretRes)

	if err != nil {
		return err
	}

	*o = EndpointGetSecretRes(varEndpointGetSecretRes)

	return err
}

type NullableEndpointGetSecretRes struct {
	value *EndpointGetSecretRes
	isSet bool
}

func (v NullableEndpointGetSecretRes) Get() *EndpointGetSecretRes {
	return v.value
}

func (v *NullableEndpointGetSecretRes) Set(val *EndpointGetSecretRes) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointGetSecretRes) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointGetSecretRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointGetSecretRes(val *EndpointGetSecretRes) *NullableEndpointGetSecretRes {
	return &NullableEndpointGetSecretRes{value: val, isSet: true}
}

func (v NullableEndpointGetSecretRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointGetSecretRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


