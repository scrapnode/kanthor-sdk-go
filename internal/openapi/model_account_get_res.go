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

// checks if the AccountGetRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountGetRes{}

// AccountGetRes struct for AccountGetRes
type AccountGetRes struct {
	Account Account `json:"account"`
}

type _AccountGetRes AccountGetRes

// NewAccountGetRes instantiates a new AccountGetRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountGetRes(account Account) *AccountGetRes {
	this := AccountGetRes{}
	this.Account = account
	return &this
}

// NewAccountGetResWithDefaults instantiates a new AccountGetRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountGetResWithDefaults() *AccountGetRes {
	this := AccountGetRes{}
	return &this
}

// GetAccount returns the Account field value
func (o *AccountGetRes) GetAccount() Account {
	if o == nil {
		var ret Account
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *AccountGetRes) GetAccountOk() (*Account, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *AccountGetRes) SetAccount(v Account) {
	o.Account = v
}

func (o AccountGetRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountGetRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account"] = o.Account
	return toSerialize, nil
}

func (o *AccountGetRes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account",
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

	varAccountGetRes := _AccountGetRes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccountGetRes)

	if err != nil {
		return err
	}

	*o = AccountGetRes(varAccountGetRes)

	return err
}

type NullableAccountGetRes struct {
	value *AccountGetRes
	isSet bool
}

func (v NullableAccountGetRes) Get() *AccountGetRes {
	return v.value
}

func (v *NullableAccountGetRes) Set(val *AccountGetRes) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountGetRes) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountGetRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountGetRes(val *AccountGetRes) *NullableAccountGetRes {
	return &NullableAccountGetRes{value: val, isSet: true}
}

func (v NullableAccountGetRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountGetRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


