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

// checks if the MessageCreateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageCreateReq{}

// MessageCreateReq struct for MessageCreateReq
type MessageCreateReq struct {
	AppId string `json:"app_id"`
	Body string `json:"body"`
	Headers map[string]string `json:"headers"`
	Type string `json:"type"`
}

type _MessageCreateReq MessageCreateReq

// NewMessageCreateReq instantiates a new MessageCreateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageCreateReq(appId string, body string, headers map[string]string, type_ string) *MessageCreateReq {
	this := MessageCreateReq{}
	this.AppId = appId
	this.Body = body
	this.Headers = headers
	this.Type = type_
	return &this
}

// NewMessageCreateReqWithDefaults instantiates a new MessageCreateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageCreateReqWithDefaults() *MessageCreateReq {
	this := MessageCreateReq{}
	return &this
}

// GetAppId returns the AppId field value
func (o *MessageCreateReq) GetAppId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
func (o *MessageCreateReq) GetAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value
func (o *MessageCreateReq) SetAppId(v string) {
	o.AppId = v
}

// GetBody returns the Body field value
func (o *MessageCreateReq) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *MessageCreateReq) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *MessageCreateReq) SetBody(v string) {
	o.Body = v
}

// GetHeaders returns the Headers field value
func (o *MessageCreateReq) GetHeaders() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value
// and a boolean to check if the value has been set.
func (o *MessageCreateReq) GetHeadersOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Headers, true
}

// SetHeaders sets field value
func (o *MessageCreateReq) SetHeaders(v map[string]string) {
	o.Headers = v
}

// GetType returns the Type field value
func (o *MessageCreateReq) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MessageCreateReq) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MessageCreateReq) SetType(v string) {
	o.Type = v
}

func (o MessageCreateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageCreateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["app_id"] = o.AppId
	toSerialize["body"] = o.Body
	toSerialize["headers"] = o.Headers
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *MessageCreateReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"app_id",
		"body",
		"headers",
		"type",
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

	varMessageCreateReq := _MessageCreateReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageCreateReq)

	if err != nil {
		return err
	}

	*o = MessageCreateReq(varMessageCreateReq)

	return err
}

type NullableMessageCreateReq struct {
	value *MessageCreateReq
	isSet bool
}

func (v NullableMessageCreateReq) Get() *MessageCreateReq {
	return v.value
}

func (v *NullableMessageCreateReq) Set(val *MessageCreateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageCreateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageCreateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageCreateReq(val *MessageCreateReq) *NullableMessageCreateReq {
	return &NullableMessageCreateReq{value: val, isSet: true}
}

func (v NullableMessageCreateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageCreateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


