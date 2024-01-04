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

// checks if the ApplicationDeleteRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationDeleteRes{}

// ApplicationDeleteRes struct for ApplicationDeleteRes
type ApplicationDeleteRes struct {
	CreatedAt int64 `json:"created_at"`
	Id string `json:"id"`
	Name string `json:"name"`
	UpdatedAt int64 `json:"updated_at"`
	WsId string `json:"ws_id"`
}

type _ApplicationDeleteRes ApplicationDeleteRes

// NewApplicationDeleteRes instantiates a new ApplicationDeleteRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationDeleteRes(createdAt int64, id string, name string, updatedAt int64, wsId string) *ApplicationDeleteRes {
	this := ApplicationDeleteRes{}
	this.CreatedAt = createdAt
	this.Id = id
	this.Name = name
	this.UpdatedAt = updatedAt
	this.WsId = wsId
	return &this
}

// NewApplicationDeleteResWithDefaults instantiates a new ApplicationDeleteRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationDeleteResWithDefaults() *ApplicationDeleteRes {
	this := ApplicationDeleteRes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *ApplicationDeleteRes) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ApplicationDeleteRes) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ApplicationDeleteRes) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetId returns the Id field value
func (o *ApplicationDeleteRes) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApplicationDeleteRes) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApplicationDeleteRes) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ApplicationDeleteRes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationDeleteRes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplicationDeleteRes) SetName(v string) {
	o.Name = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ApplicationDeleteRes) GetUpdatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ApplicationDeleteRes) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ApplicationDeleteRes) SetUpdatedAt(v int64) {
	o.UpdatedAt = v
}

// GetWsId returns the WsId field value
func (o *ApplicationDeleteRes) GetWsId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WsId
}

// GetWsIdOk returns a tuple with the WsId field value
// and a boolean to check if the value has been set.
func (o *ApplicationDeleteRes) GetWsIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WsId, true
}

// SetWsId sets field value
func (o *ApplicationDeleteRes) SetWsId(v string) {
	o.WsId = v
}

func (o ApplicationDeleteRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationDeleteRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["ws_id"] = o.WsId
	return toSerialize, nil
}

func (o *ApplicationDeleteRes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"created_at",
		"id",
		"name",
		"updated_at",
		"ws_id",
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

	varApplicationDeleteRes := _ApplicationDeleteRes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApplicationDeleteRes)

	if err != nil {
		return err
	}

	*o = ApplicationDeleteRes(varApplicationDeleteRes)

	return err
}

type NullableApplicationDeleteRes struct {
	value *ApplicationDeleteRes
	isSet bool
}

func (v NullableApplicationDeleteRes) Get() *ApplicationDeleteRes {
	return v.value
}

func (v *NullableApplicationDeleteRes) Set(val *ApplicationDeleteRes) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationDeleteRes) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationDeleteRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationDeleteRes(val *ApplicationDeleteRes) *NullableApplicationDeleteRes {
	return &NullableApplicationDeleteRes{value: val, isSet: true}
}

func (v NullableApplicationDeleteRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationDeleteRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


