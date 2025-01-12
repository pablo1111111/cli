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

// GetNamespaceIngredientVersionsReader is a Reader for the GetNamespaceIngredientVersions structure.
type GetNamespaceIngredientVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNamespaceIngredientVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNamespaceIngredientVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetNamespaceIngredientVersionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetNamespaceIngredientVersionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNamespaceIngredientVersionsOK creates a GetNamespaceIngredientVersionsOK with default headers values
func NewGetNamespaceIngredientVersionsOK() *GetNamespaceIngredientVersionsOK {
	return &GetNamespaceIngredientVersionsOK{}
}

/* GetNamespaceIngredientVersionsOK describes a response with status code 200, with default header values.

A paginated list of ingredient versions
*/
type GetNamespaceIngredientVersionsOK struct {
	Payload *inventory_models.IngredientVersionPagedList
}

func (o *GetNamespaceIngredientVersionsOK) Error() string {
	return fmt.Sprintf("[GET /v1/namespaces/ingredient/versions][%d] getNamespaceIngredientVersionsOK  %+v", 200, o.Payload)
}
func (o *GetNamespaceIngredientVersionsOK) GetPayload() *inventory_models.IngredientVersionPagedList {
	return o.Payload
}

func (o *GetNamespaceIngredientVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.IngredientVersionPagedList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNamespaceIngredientVersionsNotFound creates a GetNamespaceIngredientVersionsNotFound with default headers values
func NewGetNamespaceIngredientVersionsNotFound() *GetNamespaceIngredientVersionsNotFound {
	return &GetNamespaceIngredientVersionsNotFound{}
}

/* GetNamespaceIngredientVersionsNotFound describes a response with status code 404, with default header values.

There is no ingredient with the given namespace and name
*/
type GetNamespaceIngredientVersionsNotFound struct {
	Payload *inventory_models.RestAPIError
}

func (o *GetNamespaceIngredientVersionsNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/namespaces/ingredient/versions][%d] getNamespaceIngredientVersionsNotFound  %+v", 404, o.Payload)
}
func (o *GetNamespaceIngredientVersionsNotFound) GetPayload() *inventory_models.RestAPIError {
	return o.Payload
}

func (o *GetNamespaceIngredientVersionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNamespaceIngredientVersionsDefault creates a GetNamespaceIngredientVersionsDefault with default headers values
func NewGetNamespaceIngredientVersionsDefault(code int) *GetNamespaceIngredientVersionsDefault {
	return &GetNamespaceIngredientVersionsDefault{
		_statusCode: code,
	}
}

/* GetNamespaceIngredientVersionsDefault describes a response with status code -1, with default header values.

generic error response
*/
type GetNamespaceIngredientVersionsDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the get namespace ingredient versions default response
func (o *GetNamespaceIngredientVersionsDefault) Code() int {
	return o._statusCode
}

func (o *GetNamespaceIngredientVersionsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/namespaces/ingredient/versions][%d] getNamespaceIngredientVersions default  %+v", o._statusCode, o.Payload)
}
func (o *GetNamespaceIngredientVersionsDefault) GetPayload() *inventory_models.RestAPIError {
	return o.Payload
}

func (o *GetNamespaceIngredientVersionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
