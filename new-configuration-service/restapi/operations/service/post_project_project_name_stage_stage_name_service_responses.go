// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/keptn/keptn/new-configuration-service/models"
)

// PostProjectProjectNameStageStageNameServiceNoContentCode is the HTTP code returned for type PostProjectProjectNameStageStageNameServiceNoContent
const PostProjectProjectNameStageStageNameServiceNoContentCode int = 204

/*PostProjectProjectNameStageStageNameServiceNoContent Success. Stage has been created. Response does not have a body.

swagger:response postProjectProjectNameStageStageNameServiceNoContent
*/
type PostProjectProjectNameStageStageNameServiceNoContent struct {
}

// NewPostProjectProjectNameStageStageNameServiceNoContent creates PostProjectProjectNameStageStageNameServiceNoContent with default headers values
func NewPostProjectProjectNameStageStageNameServiceNoContent() *PostProjectProjectNameStageStageNameServiceNoContent {

	return &PostProjectProjectNameStageStageNameServiceNoContent{}
}

// WriteResponse to the client
func (o *PostProjectProjectNameStageStageNameServiceNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// PostProjectProjectNameStageStageNameServiceBadRequestCode is the HTTP code returned for type PostProjectProjectNameStageStageNameServiceBadRequest
const PostProjectProjectNameStageStageNameServiceBadRequestCode int = 400

/*PostProjectProjectNameStageStageNameServiceBadRequest Failed. Service could not be created.

swagger:response postProjectProjectNameStageStageNameServiceBadRequest
*/
type PostProjectProjectNameStageStageNameServiceBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostProjectProjectNameStageStageNameServiceBadRequest creates PostProjectProjectNameStageStageNameServiceBadRequest with default headers values
func NewPostProjectProjectNameStageStageNameServiceBadRequest() *PostProjectProjectNameStageStageNameServiceBadRequest {

	return &PostProjectProjectNameStageStageNameServiceBadRequest{}
}

// WithPayload adds the payload to the post project project name stage stage name service bad request response
func (o *PostProjectProjectNameStageStageNameServiceBadRequest) WithPayload(payload *models.Error) *PostProjectProjectNameStageStageNameServiceBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post project project name stage stage name service bad request response
func (o *PostProjectProjectNameStageStageNameServiceBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostProjectProjectNameStageStageNameServiceBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PostProjectProjectNameStageStageNameServiceDefault Error

swagger:response postProjectProjectNameStageStageNameServiceDefault
*/
type PostProjectProjectNameStageStageNameServiceDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostProjectProjectNameStageStageNameServiceDefault creates PostProjectProjectNameStageStageNameServiceDefault with default headers values
func NewPostProjectProjectNameStageStageNameServiceDefault(code int) *PostProjectProjectNameStageStageNameServiceDefault {
	if code <= 0 {
		code = 500
	}

	return &PostProjectProjectNameStageStageNameServiceDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post project project name stage stage name service default response
func (o *PostProjectProjectNameStageStageNameServiceDefault) WithStatusCode(code int) *PostProjectProjectNameStageStageNameServiceDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post project project name stage stage name service default response
func (o *PostProjectProjectNameStageStageNameServiceDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post project project name stage stage name service default response
func (o *PostProjectProjectNameStageStageNameServiceDefault) WithPayload(payload *models.Error) *PostProjectProjectNameStageStageNameServiceDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post project project name stage stage name service default response
func (o *PostProjectProjectNameStageStageNameServiceDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostProjectProjectNameStageStageNameServiceDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
