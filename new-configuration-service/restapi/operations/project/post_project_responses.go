// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/keptn/keptn/new-configuration-service/models"
)

// PostProjectNoContentCode is the HTTP code returned for type PostProjectNoContent
const PostProjectNoContentCode int = 204

/*PostProjectNoContent Success. Project has been created. Response does not have a body.

swagger:response postProjectNoContent
*/
type PostProjectNoContent struct {
}

// NewPostProjectNoContent creates PostProjectNoContent with default headers values
func NewPostProjectNoContent() *PostProjectNoContent {

	return &PostProjectNoContent{}
}

// WriteResponse to the client
func (o *PostProjectNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// PostProjectBadRequestCode is the HTTP code returned for type PostProjectBadRequest
const PostProjectBadRequestCode int = 400

/*PostProjectBadRequest Failed. Project could not be created.

swagger:response postProjectBadRequest
*/
type PostProjectBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostProjectBadRequest creates PostProjectBadRequest with default headers values
func NewPostProjectBadRequest() *PostProjectBadRequest {

	return &PostProjectBadRequest{}
}

// WithPayload adds the payload to the post project bad request response
func (o *PostProjectBadRequest) WithPayload(payload *models.Error) *PostProjectBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post project bad request response
func (o *PostProjectBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostProjectBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PostProjectDefault Error

swagger:response postProjectDefault
*/
type PostProjectDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostProjectDefault creates PostProjectDefault with default headers values
func NewPostProjectDefault(code int) *PostProjectDefault {
	if code <= 0 {
		code = 500
	}

	return &PostProjectDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post project default response
func (o *PostProjectDefault) WithStatusCode(code int) *PostProjectDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post project default response
func (o *PostProjectDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post project default response
func (o *PostProjectDefault) WithPayload(payload *models.Error) *PostProjectDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post project default response
func (o *PostProjectDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostProjectDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
