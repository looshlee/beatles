// Code generated by go-swagger; DO NOT EDIT.

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/cilium/api/v1/models"
)

// PostIPAMCreatedCode is the HTTP code returned for type PostIPAMCreated
const PostIPAMCreatedCode int = 201

/*PostIPAMCreated Success

swagger:response postIpAMCreated
*/
type PostIPAMCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IPAMResponse `json:"body,omitempty"`
}

// NewPostIPAMCreated creates PostIPAMCreated with default headers values
func NewPostIPAMCreated() *PostIPAMCreated {
	return &PostIPAMCreated{}
}

// WithPayload adds the payload to the post Ip a m created response
func (o *PostIPAMCreated) WithPayload(payload *models.IPAMResponse) *PostIPAMCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post Ip a m created response
func (o *PostIPAMCreated) SetPayload(payload *models.IPAMResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostIPAMCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostIPAMFailureCode is the HTTP code returned for type PostIPAMFailure
const PostIPAMFailureCode int = 502

/*PostIPAMFailure Allocation failure

swagger:response postIpAMFailure
*/
type PostIPAMFailure struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewPostIPAMFailure creates PostIPAMFailure with default headers values
func NewPostIPAMFailure() *PostIPAMFailure {
	return &PostIPAMFailure{}
}

// WithPayload adds the payload to the post Ip a m failure response
func (o *PostIPAMFailure) WithPayload(payload models.Error) *PostIPAMFailure {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post Ip a m failure response
func (o *PostIPAMFailure) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostIPAMFailure) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(502)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}
