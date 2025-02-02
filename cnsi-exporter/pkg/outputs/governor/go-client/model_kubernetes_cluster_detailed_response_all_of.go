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

// KubernetesClusterDetailedResponseAllOf struct for KubernetesClusterDetailedResponseAllOf
type KubernetesClusterDetailedResponseAllOf struct {
	Telemetry            *KubernetesTelemetryResponse `json:"telemetry,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesClusterDetailedResponseAllOf KubernetesClusterDetailedResponseAllOf

// NewKubernetesClusterDetailedResponseAllOf instantiates a new KubernetesClusterDetailedResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesClusterDetailedResponseAllOf() *KubernetesClusterDetailedResponseAllOf {
	this := KubernetesClusterDetailedResponseAllOf{}
	return &this
}

// NewKubernetesClusterDetailedResponseAllOfWithDefaults instantiates a new KubernetesClusterDetailedResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesClusterDetailedResponseAllOfWithDefaults() *KubernetesClusterDetailedResponseAllOf {
	this := KubernetesClusterDetailedResponseAllOf{}
	return &this
}

// GetTelemetry returns the Telemetry field value if set, zero value otherwise.
func (o *KubernetesClusterDetailedResponseAllOf) GetTelemetry() KubernetesTelemetryResponse {
	if o == nil || o.Telemetry == nil {
		var ret KubernetesTelemetryResponse
		return ret
	}
	return *o.Telemetry
}

// GetTelemetryOk returns a tuple with the Telemetry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesClusterDetailedResponseAllOf) GetTelemetryOk() (*KubernetesTelemetryResponse, bool) {
	if o == nil || o.Telemetry == nil {
		return nil, false
	}
	return o.Telemetry, true
}

// HasTelemetry returns a boolean if a field has been set.
func (o *KubernetesClusterDetailedResponseAllOf) HasTelemetry() bool {
	if o != nil && o.Telemetry != nil {
		return true
	}

	return false
}

// SetTelemetry gets a reference to the given KubernetesTelemetryResponse and assigns it to the Telemetry field.
func (o *KubernetesClusterDetailedResponseAllOf) SetTelemetry(v KubernetesTelemetryResponse) {
	o.Telemetry = &v
}

func (o KubernetesClusterDetailedResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Telemetry != nil {
		toSerialize["telemetry"] = o.Telemetry
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesClusterDetailedResponseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKubernetesClusterDetailedResponseAllOf := _KubernetesClusterDetailedResponseAllOf{}

	if err = json.Unmarshal(bytes, &varKubernetesClusterDetailedResponseAllOf); err == nil {
		*o = KubernetesClusterDetailedResponseAllOf(varKubernetesClusterDetailedResponseAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "telemetry")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesClusterDetailedResponseAllOf struct {
	value *KubernetesClusterDetailedResponseAllOf
	isSet bool
}

func (v NullableKubernetesClusterDetailedResponseAllOf) Get() *KubernetesClusterDetailedResponseAllOf {
	return v.value
}

func (v *NullableKubernetesClusterDetailedResponseAllOf) Set(val *KubernetesClusterDetailedResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesClusterDetailedResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesClusterDetailedResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesClusterDetailedResponseAllOf(val *KubernetesClusterDetailedResponseAllOf) *NullableKubernetesClusterDetailedResponseAllOf {
	return &NullableKubernetesClusterDetailedResponseAllOf{value: val, isSet: true}
}

func (v NullableKubernetesClusterDetailedResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesClusterDetailedResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
