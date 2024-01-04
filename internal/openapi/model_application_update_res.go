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
)

// checks if the ApplicationUpdateRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationUpdateRes{}

// ApplicationUpdateRes struct for ApplicationUpdateRes
type ApplicationUpdateRes struct {
	CreatedAt *int64 `json:"created_at,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	WsId *string `json:"ws_id,omitempty"`
}

// NewApplicationUpdateRes instantiates a new ApplicationUpdateRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationUpdateRes() *ApplicationUpdateRes {
	this := ApplicationUpdateRes{}
	return &this
}

// NewApplicationUpdateResWithDefaults instantiates a new ApplicationUpdateRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationUpdateResWithDefaults() *ApplicationUpdateRes {
	this := ApplicationUpdateRes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ApplicationUpdateRes) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateRes) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ApplicationUpdateRes) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *ApplicationUpdateRes) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApplicationUpdateRes) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateRes) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApplicationUpdateRes) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApplicationUpdateRes) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApplicationUpdateRes) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateRes) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApplicationUpdateRes) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApplicationUpdateRes) SetName(v string) {
	o.Name = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ApplicationUpdateRes) GetUpdatedAt() int64 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateRes) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ApplicationUpdateRes) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *ApplicationUpdateRes) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

// GetWsId returns the WsId field value if set, zero value otherwise.
func (o *ApplicationUpdateRes) GetWsId() string {
	if o == nil || IsNil(o.WsId) {
		var ret string
		return ret
	}
	return *o.WsId
}

// GetWsIdOk returns a tuple with the WsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationUpdateRes) GetWsIdOk() (*string, bool) {
	if o == nil || IsNil(o.WsId) {
		return nil, false
	}
	return o.WsId, true
}

// HasWsId returns a boolean if a field has been set.
func (o *ApplicationUpdateRes) HasWsId() bool {
	if o != nil && !IsNil(o.WsId) {
		return true
	}

	return false
}

// SetWsId gets a reference to the given string and assigns it to the WsId field.
func (o *ApplicationUpdateRes) SetWsId(v string) {
	o.WsId = &v
}

func (o ApplicationUpdateRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationUpdateRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.WsId) {
		toSerialize["ws_id"] = o.WsId
	}
	return toSerialize, nil
}

type NullableApplicationUpdateRes struct {
	value *ApplicationUpdateRes
	isSet bool
}

func (v NullableApplicationUpdateRes) Get() *ApplicationUpdateRes {
	return v.value
}

func (v *NullableApplicationUpdateRes) Set(val *ApplicationUpdateRes) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationUpdateRes) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationUpdateRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationUpdateRes(val *ApplicationUpdateRes) *NullableApplicationUpdateRes {
	return &NullableApplicationUpdateRes{value: val, isSet: true}
}

func (v NullableApplicationUpdateRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationUpdateRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


