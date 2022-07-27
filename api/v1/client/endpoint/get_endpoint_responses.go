// Code generated by go-swagger; DO NOT EDIT.

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cilium/cilium/api/v1/models"
)

// GetEndpointReader is a Reader for the GetEndpoint structure.
type GetEndpointReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEndpointReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEndpointOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetEndpointNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEndpointOK creates a GetEndpointOK with default headers values
func NewGetEndpointOK() *GetEndpointOK {
	return &GetEndpointOK{}
}

/*GetEndpointOK handles this case with default header values.

Success
*/
type GetEndpointOK struct {
	Payload []*models.Endpoint
}

func (o *GetEndpointOK) Error() string {
	return fmt.Sprintf("[GET /endpoint][%d] getEndpointOK  %+v", 200, o.Payload)
}

func (o *GetEndpointOK) GetPayload() []*models.Endpoint {
	return o.Payload
}

func (o *GetEndpointOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEndpointNotFound creates a GetEndpointNotFound with default headers values
func NewGetEndpointNotFound() *GetEndpointNotFound {
	return &GetEndpointNotFound{}
}

/*GetEndpointNotFound handles this case with default header values.

Endpoints with provided parameters not found
*/
type GetEndpointNotFound struct {
}

func (o *GetEndpointNotFound) Error() string {
	return fmt.Sprintf("[GET /endpoint][%d] getEndpointNotFound ", 404)
}

func (o *GetEndpointNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
