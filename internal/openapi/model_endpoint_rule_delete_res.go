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

// checks if the EndpointRuleDeleteRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndpointRuleDeleteRes{}

// EndpointRuleDeleteRes struct for EndpointRuleDeleteRes
type EndpointRuleDeleteRes struct {
	ConditionExpression *string `json:"condition_expression,omitempty"`
	ConditionSource *string `json:"condition_source,omitempty"`
	CreatedAt *int64 `json:"created_at,omitempty"`
	EpId *string `json:"ep_id,omitempty"`
	Exclusionary *bool `json:"exclusionary,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Priority *int64 `json:"priority,omitempty"`
	UpdatedAt *int64 `json:"updated_at,omitempty"`
}

// NewEndpointRuleDeleteRes instantiates a new EndpointRuleDeleteRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointRuleDeleteRes() *EndpointRuleDeleteRes {
	this := EndpointRuleDeleteRes{}
	return &this
}

// NewEndpointRuleDeleteResWithDefaults instantiates a new EndpointRuleDeleteRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointRuleDeleteResWithDefaults() *EndpointRuleDeleteRes {
	this := EndpointRuleDeleteRes{}
	return &this
}

// GetConditionExpression returns the ConditionExpression field value if set, zero value otherwise.
func (o *EndpointRuleDeleteRes) GetConditionExpression() string {
	if o == nil || IsNil(o.ConditionExpression) {
		var ret string
		return ret
	}
	return *o.ConditionExpression
}

// GetConditionExpressionOk returns a tuple with the ConditionExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointRuleDeleteRes) GetConditionExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.ConditionExpression) {
		return nil, false
	}
	return o.ConditionExpression, true
}

// HasConditionExpression returns a boolean if a field has been set.
func (o *EndpointRuleDeleteRes) HasConditionExpression() bool {
	if o != nil && !IsNil(o.ConditionExpression) {
		return true
	}

	return false
}

// SetConditionExpression gets a reference to the given string and assigns it to the ConditionExpression field.
func (o *EndpointRuleDeleteRes) SetConditionExpression(v string) {
	o.ConditionExpression = &v
}

// GetConditionSource returns the ConditionSource field value if set, zero value otherwise.
func (o *EndpointRuleDeleteRes) GetConditionSource() string {
	if o == nil || IsNil(o.ConditionSource) {
		var ret string
		return ret
	}
	return *o.ConditionSource
}

// GetConditionSourceOk returns a tuple with the ConditionSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointRuleDeleteRes) GetConditionSourceOk() (*string, bool) {
	if o == nil || IsNil(o.ConditionSource) {
		return nil, false
	}
	return o.ConditionSource, true
}

// HasConditionSource returns a boolean if a field has been set.
func (o *EndpointRuleDeleteRes) HasConditionSource() bool {
	if o != nil && !IsNil(o.ConditionSource) {
		return true
	}

	return false
}

// SetConditionSource gets a reference to the given string and assigns it to the ConditionSource field.
func (o *EndpointRuleDeleteRes) SetConditionSource(v string) {
	o.ConditionSource = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *EndpointRuleDeleteRes) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointRuleDeleteRes) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EndpointRuleDeleteRes) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *EndpointRuleDeleteRes) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetEpId returns the EpId field value if set, zero value otherwise.
func (o *EndpointRuleDeleteRes) GetEpId() string {
	if o == nil || IsNil(o.EpId) {
		var ret string
		return ret
	}
	return *o.EpId
}

// GetEpIdOk returns a tuple with the EpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointRuleDeleteRes) GetEpIdOk() (*string, bool) {
	if o == nil || IsNil(o.EpId) {
		return nil, false
	}
	return o.EpId, true
}

// HasEpId returns a boolean if a field has been set.
func (o *EndpointRuleDeleteRes) HasEpId() bool {
	if o != nil && !IsNil(o.EpId) {
		return true
	}

	return false
}

// SetEpId gets a reference to the given string and assigns it to the EpId field.
func (o *EndpointRuleDeleteRes) SetEpId(v string) {
	o.EpId = &v
}

// GetExclusionary returns the Exclusionary field value if set, zero value otherwise.
func (o *EndpointRuleDeleteRes) GetExclusionary() bool {
	if o == nil || IsNil(o.Exclusionary) {
		var ret bool
		return ret
	}
	return *o.Exclusionary
}

// GetExclusionaryOk returns a tuple with the Exclusionary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointRuleDeleteRes) GetExclusionaryOk() (*bool, bool) {
	if o == nil || IsNil(o.Exclusionary) {
		return nil, false
	}
	return o.Exclusionary, true
}

// HasExclusionary returns a boolean if a field has been set.
func (o *EndpointRuleDeleteRes) HasExclusionary() bool {
	if o != nil && !IsNil(o.Exclusionary) {
		return true
	}

	return false
}

// SetExclusionary gets a reference to the given bool and assigns it to the Exclusionary field.
func (o *EndpointRuleDeleteRes) SetExclusionary(v bool) {
	o.Exclusionary = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EndpointRuleDeleteRes) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointRuleDeleteRes) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EndpointRuleDeleteRes) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EndpointRuleDeleteRes) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EndpointRuleDeleteRes) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointRuleDeleteRes) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EndpointRuleDeleteRes) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EndpointRuleDeleteRes) SetName(v string) {
	o.Name = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *EndpointRuleDeleteRes) GetPriority() int64 {
	if o == nil || IsNil(o.Priority) {
		var ret int64
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointRuleDeleteRes) GetPriorityOk() (*int64, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *EndpointRuleDeleteRes) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int64 and assigns it to the Priority field.
func (o *EndpointRuleDeleteRes) SetPriority(v int64) {
	o.Priority = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *EndpointRuleDeleteRes) GetUpdatedAt() int64 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointRuleDeleteRes) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *EndpointRuleDeleteRes) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *EndpointRuleDeleteRes) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

func (o EndpointRuleDeleteRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointRuleDeleteRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConditionExpression) {
		toSerialize["condition_expression"] = o.ConditionExpression
	}
	if !IsNil(o.ConditionSource) {
		toSerialize["condition_source"] = o.ConditionSource
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.EpId) {
		toSerialize["ep_id"] = o.EpId
	}
	if !IsNil(o.Exclusionary) {
		toSerialize["exclusionary"] = o.Exclusionary
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableEndpointRuleDeleteRes struct {
	value *EndpointRuleDeleteRes
	isSet bool
}

func (v NullableEndpointRuleDeleteRes) Get() *EndpointRuleDeleteRes {
	return v.value
}

func (v *NullableEndpointRuleDeleteRes) Set(val *EndpointRuleDeleteRes) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointRuleDeleteRes) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointRuleDeleteRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointRuleDeleteRes(val *EndpointRuleDeleteRes) *NullableEndpointRuleDeleteRes {
	return &NullableEndpointRuleDeleteRes{value: val, isSet: true}
}

func (v NullableEndpointRuleDeleteRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointRuleDeleteRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


