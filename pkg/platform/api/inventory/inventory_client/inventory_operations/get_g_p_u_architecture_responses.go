// Code generated by go-swagger; DO NOT EDIT.

package inventory_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ActiveState/cli/pkg/platform/api/inventory/inventory_models"
)

// GetGPUArchitectureReader is a Reader for the GetGPUArchitecture structure.
type GetGPUArchitectureReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGPUArchitectureReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGPUArchitectureOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetGPUArchitectureDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetGPUArchitectureOK creates a GetGPUArchitectureOK with default headers values
func NewGetGPUArchitectureOK() *GetGPUArchitectureOK {
	return &GetGPUArchitectureOK{}
}

/* GetGPUArchitectureOK describes a response with status code 200, with default header values.

The retrieved GPU architecture
*/
type GetGPUArchitectureOK struct {
	Payload *inventory_models.GpuArchitecture
}

func (o *GetGPUArchitectureOK) Error() string {
	return fmt.Sprintf("[GET /v1/gpu-architectures/{gpu_architecture_id}][%d] getGPUArchitectureOK  %+v", 200, o.Payload)
}
func (o *GetGPUArchitectureOK) GetPayload() *inventory_models.GpuArchitecture {
	return o.Payload
}

func (o *GetGPUArchitectureOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.GpuArchitecture)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGPUArchitectureDefault creates a GetGPUArchitectureDefault with default headers values
func NewGetGPUArchitectureDefault(code int) *GetGPUArchitectureDefault {
	return &GetGPUArchitectureDefault{
		_statusCode: code,
	}
}

/* GetGPUArchitectureDefault describes a response with status code -1, with default header values.

generic error response
*/
type GetGPUArchitectureDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the get g p u architecture default response
func (o *GetGPUArchitectureDefault) Code() int {
	return o._statusCode
}

func (o *GetGPUArchitectureDefault) Error() string {
	return fmt.Sprintf("[GET /v1/gpu-architectures/{gpu_architecture_id}][%d] getGPUArchitecture default  %+v", o._statusCode, o.Payload)
}
func (o *GetGPUArchitectureDefault) GetPayload() *inventory_models.RestAPIError {
	return o.Payload
}

func (o *GetGPUArchitectureDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
