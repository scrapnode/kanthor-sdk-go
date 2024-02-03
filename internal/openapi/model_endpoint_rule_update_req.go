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

// checks if the EndpointRuleUpdateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndpointRuleUpdateReq{}

// EndpointRuleUpdateReq struct for EndpointRuleUpdateReq
type EndpointRuleUpdateReq struct {
	ConditionExpression string `json:"condition_expression"`
	ConditionSource string `json:"condition_source"`
	Exclusionary bool `json:"exclusionary"`
	Name string `json:"name"`
	Priority int64 `json:"priority"`
}

type _EndpointRuleUpdateReq EndpointRuleUpdateReq

// NewEndpointRuleUpdateReq instantiates a new EndpointRuleUpdateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointRuleUpdateReq(conditionExpression string, conditionSource string, exclusionary bool, name string, priority int64) *EndpointRuleUpdateReq {
	this := EndpointRuleUpdateReq{}
	this.ConditionExpression = conditionExpression
	this.ConditionSource = conditionSource
	this.Exclusionary = exclusionary
	this.Name = name
	this.Priority = priority
	return &this
}

// NewEndpointRuleUpdateReqWithDefaults instantiates a new EndpointRuleUpdateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointRuleUpdateReqWithDefaults() *EndpointRuleUpdateReq {
	this := EndpointRuleUpdateReq{}
	return &this
}

// GetConditionExpression returns the ConditionExpression field value
func (o *EndpointRuleUpdateReq) GetConditionExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConditionExpression
}

// GetConditionExpressionOk returns a tuple with the ConditionExpression field value
// and a boolean to check if the value has been set.
func (o *EndpointRuleUpdateReq) GetConditionExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConditionExpression, true
}

// SetConditionExpression sets field value
func (o *EndpointRuleUpdateReq) SetConditionExpression(v string) {
	o.ConditionExpression = v
}

// GetConditionSource returns the ConditionSource field value
func (o *EndpointRuleUpdateReq) GetConditionSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConditionSource
}

// GetConditionSourceOk returns a tuple with the ConditionSource field value
// and a boolean to check if the value has been set.
func (o *EndpointRuleUpdateReq) GetConditionSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConditionSource, true
}

// SetConditionSource sets field value
func (o *EndpointRuleUpdateReq) SetConditionSource(v string) {
	o.ConditionSource = v
}

// GetExclusionary returns the Exclusionary field value
func (o *EndpointRuleUpdateReq) GetExclusionary() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Exclusionary
}

// GetExclusionaryOk returns a tuple with the Exclusionary field value
// and a boolean to check if the value has been set.
func (o *EndpointRuleUpdateReq) GetExclusionaryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exclusionary, true
}

// SetExclusionary sets field value
func (o *EndpointRuleUpdateReq) SetExclusionary(v bool) {
	o.Exclusionary = v
}

// GetName returns the Name field value
func (o *EndpointRuleUpdateReq) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EndpointRuleUpdateReq) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EndpointRuleUpdateReq) SetName(v string) {
	o.Name = v
}

// GetPriority returns the Priority field value
func (o *EndpointRuleUpdateReq) GetPriority() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *EndpointRuleUpdateReq) GetPriorityOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *EndpointRuleUpdateReq) SetPriority(v int64) {
	o.Priority = v
}

func (o EndpointRuleUpdateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointRuleUpdateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["condition_expression"] = o.ConditionExpression
	toSerialize["condition_source"] = o.ConditionSource
	toSerialize["exclusionary"] = o.Exclusionary
	toSerialize["name"] = o.Name
	toSerialize["priority"] = o.Priority
	return toSerialize, nil
}

func (o *EndpointRuleUpdateReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"condition_expression",
		"condition_source",
		"exclusionary",
		"name",
		"priority",
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

	varEndpointRuleUpdateReq := _EndpointRuleUpdateReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEndpointRuleUpdateReq)

	if err != nil {
		return err
	}

	*o = EndpointRuleUpdateReq(varEndpointRuleUpdateReq)

	return err
}

type NullableEndpointRuleUpdateReq struct {
	value *EndpointRuleUpdateReq
	isSet bool
}

func (v NullableEndpointRuleUpdateReq) Get() *EndpointRuleUpdateReq {
	return v.value
}

func (v *NullableEndpointRuleUpdateReq) Set(val *EndpointRuleUpdateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointRuleUpdateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointRuleUpdateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointRuleUpdateReq(val *EndpointRuleUpdateReq) *NullableEndpointRuleUpdateReq {
	return &NullableEndpointRuleUpdateReq{value: val, isSet: true}
}

func (v NullableEndpointRuleUpdateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointRuleUpdateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


