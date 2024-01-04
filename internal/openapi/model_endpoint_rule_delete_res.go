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

// checks if the EndpointRuleDeleteRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndpointRuleDeleteRes{}

// EndpointRuleDeleteRes struct for EndpointRuleDeleteRes
type EndpointRuleDeleteRes struct {
	ConditionExpression string `json:"condition_expression"`
	ConditionSource string `json:"condition_source"`
	CreatedAt int64 `json:"created_at"`
	EpId string `json:"ep_id"`
	Exclusionary bool `json:"exclusionary"`
	Id string `json:"id"`
	Name string `json:"name"`
	Priority int64 `json:"priority"`
	UpdatedAt int64 `json:"updated_at"`
}

type _EndpointRuleDeleteRes EndpointRuleDeleteRes

// NewEndpointRuleDeleteRes instantiates a new EndpointRuleDeleteRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointRuleDeleteRes(conditionExpression string, conditionSource string, createdAt int64, epId string, exclusionary bool, id string, name string, priority int64, updatedAt int64) *EndpointRuleDeleteRes {
	this := EndpointRuleDeleteRes{}
	this.ConditionExpression = conditionExpression
	this.ConditionSource = conditionSource
	this.CreatedAt = createdAt
	this.EpId = epId
	this.Exclusionary = exclusionary
	this.Id = id
	this.Name = name
	this.Priority = priority
	this.UpdatedAt = updatedAt
	return &this
}

// NewEndpointRuleDeleteResWithDefaults instantiates a new EndpointRuleDeleteRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointRuleDeleteResWithDefaults() *EndpointRuleDeleteRes {
	this := EndpointRuleDeleteRes{}
	return &this
}

// GetConditionExpression returns the ConditionExpression field value
func (o *EndpointRuleDeleteRes) GetConditionExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConditionExpression
}

// GetConditionExpressionOk returns a tuple with the ConditionExpression field value
// and a boolean to check if the value has been set.
func (o *EndpointRuleDeleteRes) GetConditionExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConditionExpression, true
}

// SetConditionExpression sets field value
func (o *EndpointRuleDeleteRes) SetConditionExpression(v string) {
	o.ConditionExpression = v
}

// GetConditionSource returns the ConditionSource field value
func (o *EndpointRuleDeleteRes) GetConditionSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConditionSource
}

// GetConditionSourceOk returns a tuple with the ConditionSource field value
// and a boolean to check if the value has been set.
func (o *EndpointRuleDeleteRes) GetConditionSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConditionSource, true
}

// SetConditionSource sets field value
func (o *EndpointRuleDeleteRes) SetConditionSource(v string) {
	o.ConditionSource = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *EndpointRuleDeleteRes) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *EndpointRuleDeleteRes) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *EndpointRuleDeleteRes) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetEpId returns the EpId field value
func (o *EndpointRuleDeleteRes) GetEpId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EpId
}

// GetEpIdOk returns a tuple with the EpId field value
// and a boolean to check if the value has been set.
func (o *EndpointRuleDeleteRes) GetEpIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EpId, true
}

// SetEpId sets field value
func (o *EndpointRuleDeleteRes) SetEpId(v string) {
	o.EpId = v
}

// GetExclusionary returns the Exclusionary field value
func (o *EndpointRuleDeleteRes) GetExclusionary() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Exclusionary
}

// GetExclusionaryOk returns a tuple with the Exclusionary field value
// and a boolean to check if the value has been set.
func (o *EndpointRuleDeleteRes) GetExclusionaryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exclusionary, true
}

// SetExclusionary sets field value
func (o *EndpointRuleDeleteRes) SetExclusionary(v bool) {
	o.Exclusionary = v
}

// GetId returns the Id field value
func (o *EndpointRuleDeleteRes) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EndpointRuleDeleteRes) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EndpointRuleDeleteRes) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *EndpointRuleDeleteRes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EndpointRuleDeleteRes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EndpointRuleDeleteRes) SetName(v string) {
	o.Name = v
}

// GetPriority returns the Priority field value
func (o *EndpointRuleDeleteRes) GetPriority() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *EndpointRuleDeleteRes) GetPriorityOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *EndpointRuleDeleteRes) SetPriority(v int64) {
	o.Priority = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *EndpointRuleDeleteRes) GetUpdatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *EndpointRuleDeleteRes) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *EndpointRuleDeleteRes) SetUpdatedAt(v int64) {
	o.UpdatedAt = v
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
	toSerialize["condition_expression"] = o.ConditionExpression
	toSerialize["condition_source"] = o.ConditionSource
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["ep_id"] = o.EpId
	toSerialize["exclusionary"] = o.Exclusionary
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["priority"] = o.Priority
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *EndpointRuleDeleteRes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"condition_expression",
		"condition_source",
		"created_at",
		"ep_id",
		"exclusionary",
		"id",
		"name",
		"priority",
		"updated_at",
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

	varEndpointRuleDeleteRes := _EndpointRuleDeleteRes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEndpointRuleDeleteRes)

	if err != nil {
		return err
	}

	*o = EndpointRuleDeleteRes(varEndpointRuleDeleteRes)

	return err
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


