/*
Catalog Governor Service REST API

This is the service to track assets deployed in customer clusters

API version: ${project.version}
Contact: content-building-ecosystem@vmware.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// MatchLabel Resources with matching label on which agent can scan workloads
type MatchLabel struct {
	// MatchLabel key
	Key string `json:"key"`
	// MatchLabel value
	Value                string `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _MatchLabel MatchLabel

// NewMatchLabel instantiates a new MatchLabel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMatchLabel(key string, value string) *MatchLabel {
	this := MatchLabel{}
	this.Key = key
	this.Value = value
	return &this
}

// NewMatchLabelWithDefaults instantiates a new MatchLabel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMatchLabelWithDefaults() *MatchLabel {
	this := MatchLabel{}
	return &this
}

// GetKey returns the Key field value
func (o *MatchLabel) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *MatchLabel) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *MatchLabel) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *MatchLabel) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *MatchLabel) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *MatchLabel) SetValue(v string) {
	o.Value = v
}

func (o MatchLabel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MatchLabel) UnmarshalJSON(bytes []byte) (err error) {
	varMatchLabel := _MatchLabel{}

	if err = json.Unmarshal(bytes, &varMatchLabel); err == nil {
		*o = MatchLabel(varMatchLabel)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "key")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMatchLabel struct {
	value *MatchLabel
	isSet bool
}

func (v NullableMatchLabel) Get() *MatchLabel {
	return v.value
}

func (v *NullableMatchLabel) Set(val *MatchLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableMatchLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableMatchLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMatchLabel(val *MatchLabel) *NullableMatchLabel {
	return &NullableMatchLabel{value: val, isSet: true}
}

func (v NullableMatchLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMatchLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
