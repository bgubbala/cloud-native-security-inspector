// Copyright 2022 VMware, Inc.
// SPDX-License-Identifier: Apache-2.0
// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"goharbor.io/k8s-security-inspector/pkg/rest/models"
)

// HealthZReader is a Reader for the HealthZ structure.
type HealthZReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HealthZReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHealthZOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHealthZOK creates a HealthZOK with default headers values
func NewHealthZOK() *HealthZOK {
	return &HealthZOK{}
}

/*HealthZOK handles this case with default header values.

successful operation
*/
type HealthZOK struct {
	Payload *models.StatusResponse
}

func (o *HealthZOK) Error() string {
	return fmt.Sprintf("[GET /][%d] healthZOK  %+v", 200, o.Payload)
}

func (o *HealthZOK) GetPayload() *models.StatusResponse {
	return o.Payload
}

func (o *HealthZOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}