// Code generated by go-swagger; DO NOT EDIT.

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/cilium/api/v1/models"
)

// PutEndpointIDLabelsOKCode is the HTTP code returned for type PutEndpointIDLabelsOK
const PutEndpointIDLabelsOKCode int = 200

/*PutEndpointIDLabelsOK Success

swagger:response putEndpointIdLabelsOK
*/
type PutEndpointIDLabelsOK struct {
}

// NewPutEndpointIDLabelsOK creates PutEndpointIDLabelsOK with default headers values
func NewPutEndpointIDLabelsOK() *PutEndpointIDLabelsOK {
	return &PutEndpointIDLabelsOK{}
}

// WriteResponse to the client
func (o *PutEndpointIDLabelsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
}

// PutEndpointIDLabelsNotFoundCode is the HTTP code returned for type PutEndpointIDLabelsNotFound
const PutEndpointIDLabelsNotFoundCode int = 404

/*PutEndpointIDLabelsNotFound Endpoint not found

swagger:response putEndpointIdLabelsNotFound
*/
type PutEndpointIDLabelsNotFound struct {
}

// NewPutEndpointIDLabelsNotFound creates PutEndpointIDLabelsNotFound with default headers values
func NewPutEndpointIDLabelsNotFound() *PutEndpointIDLabelsNotFound {
	return &PutEndpointIDLabelsNotFound{}
}

// WriteResponse to the client
func (o *PutEndpointIDLabelsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}

// PutEndpointIDLabelsLabelNotFoundCode is the HTTP code returned for type PutEndpointIDLabelsLabelNotFound
const PutEndpointIDLabelsLabelNotFoundCode int = 460

/*PutEndpointIDLabelsLabelNotFound Label to be deleted not found

swagger:response putEndpointIdLabelsLabelNotFound
*/
type PutEndpointIDLabelsLabelNotFound struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewPutEndpointIDLabelsLabelNotFound creates PutEndpointIDLabelsLabelNotFound with default headers values
func NewPutEndpointIDLabelsLabelNotFound() *PutEndpointIDLabelsLabelNotFound {
	return &PutEndpointIDLabelsLabelNotFound{}
}

// WithPayload adds the payload to the put endpoint Id labels label not found response
func (o *PutEndpointIDLabelsLabelNotFound) WithPayload(payload models.Error) *PutEndpointIDLabelsLabelNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put endpoint Id labels label not found response
func (o *PutEndpointIDLabelsLabelNotFound) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutEndpointIDLabelsLabelNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(460)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// PutEndpointIDLabelsUpdateFailedCode is the HTTP code returned for type PutEndpointIDLabelsUpdateFailed
const PutEndpointIDLabelsUpdateFailedCode int = 500

/*PutEndpointIDLabelsUpdateFailed Error while updating labels

swagger:response putEndpointIdLabelsUpdateFailed
*/
type PutEndpointIDLabelsUpdateFailed struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewPutEndpointIDLabelsUpdateFailed creates PutEndpointIDLabelsUpdateFailed with default headers values
func NewPutEndpointIDLabelsUpdateFailed() *PutEndpointIDLabelsUpdateFailed {
	return &PutEndpointIDLabelsUpdateFailed{}
}

// WithPayload adds the payload to the put endpoint Id labels update failed response
func (o *PutEndpointIDLabelsUpdateFailed) WithPayload(payload models.Error) *PutEndpointIDLabelsUpdateFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put endpoint Id labels update failed response
func (o *PutEndpointIDLabelsUpdateFailed) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutEndpointIDLabelsUpdateFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}
