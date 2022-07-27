// Code generated by go-swagger; DO NOT EDIT.

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/cilium/cilium/api/v1/models"
)

// PatchEndpointIDConfigOKCode is the HTTP code returned for type PatchEndpointIDConfigOK
const PatchEndpointIDConfigOKCode int = 200

/*PatchEndpointIDConfigOK Success

swagger:response patchEndpointIdConfigOK
*/
type PatchEndpointIDConfigOK struct {
}

// NewPatchEndpointIDConfigOK creates PatchEndpointIDConfigOK with default headers values
func NewPatchEndpointIDConfigOK() *PatchEndpointIDConfigOK {

	return &PatchEndpointIDConfigOK{}
}

// WriteResponse to the client
func (o *PatchEndpointIDConfigOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PatchEndpointIDConfigInvalidCode is the HTTP code returned for type PatchEndpointIDConfigInvalid
const PatchEndpointIDConfigInvalidCode int = 400

/*PatchEndpointIDConfigInvalid Invalid configuration request

swagger:response patchEndpointIdConfigInvalid
*/
type PatchEndpointIDConfigInvalid struct {
}

// NewPatchEndpointIDConfigInvalid creates PatchEndpointIDConfigInvalid with default headers values
func NewPatchEndpointIDConfigInvalid() *PatchEndpointIDConfigInvalid {

	return &PatchEndpointIDConfigInvalid{}
}

// WriteResponse to the client
func (o *PatchEndpointIDConfigInvalid) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// PatchEndpointIDConfigNotFoundCode is the HTTP code returned for type PatchEndpointIDConfigNotFound
const PatchEndpointIDConfigNotFoundCode int = 404

/*PatchEndpointIDConfigNotFound Endpoint not found

swagger:response patchEndpointIdConfigNotFound
*/
type PatchEndpointIDConfigNotFound struct {
}

// NewPatchEndpointIDConfigNotFound creates PatchEndpointIDConfigNotFound with default headers values
func NewPatchEndpointIDConfigNotFound() *PatchEndpointIDConfigNotFound {

	return &PatchEndpointIDConfigNotFound{}
}

// WriteResponse to the client
func (o *PatchEndpointIDConfigNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// PatchEndpointIDConfigFailedCode is the HTTP code returned for type PatchEndpointIDConfigFailed
const PatchEndpointIDConfigFailedCode int = 500

/*PatchEndpointIDConfigFailed Update failed. Details in message.

swagger:response patchEndpointIdConfigFailed
*/
type PatchEndpointIDConfigFailed struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewPatchEndpointIDConfigFailed creates PatchEndpointIDConfigFailed with default headers values
func NewPatchEndpointIDConfigFailed() *PatchEndpointIDConfigFailed {

	return &PatchEndpointIDConfigFailed{}
}

// WithPayload adds the payload to the patch endpoint Id config failed response
func (o *PatchEndpointIDConfigFailed) WithPayload(payload models.Error) *PatchEndpointIDConfigFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch endpoint Id config failed response
func (o *PatchEndpointIDConfigFailed) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchEndpointIDConfigFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
