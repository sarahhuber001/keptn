// Code generated by go-swagger; DO NOT EDIT.

package stage_resource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/keptn/keptn/new-configuration-service/models"
)

// PostProjectProjectNameStageStageNameResourceCreatedCode is the HTTP code returned for type PostProjectProjectNameStageStageNameResourceCreated
const PostProjectProjectNameStageStageNameResourceCreatedCode int = 201

/*PostProjectProjectNameStageStageNameResourceCreated Success. Stage resource has been created. The version of the new configuration is returned.

swagger:response postProjectProjectNameStageStageNameResourceCreated
*/
type PostProjectProjectNameStageStageNameResourceCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Version `json:"body,omitempty"`
}

// NewPostProjectProjectNameStageStageNameResourceCreated creates PostProjectProjectNameStageStageNameResourceCreated with default headers values
func NewPostProjectProjectNameStageStageNameResourceCreated() *PostProjectProjectNameStageStageNameResourceCreated {

	return &PostProjectProjectNameStageStageNameResourceCreated{}
}

// WithPayload adds the payload to the post project project name stage stage name resource created response
func (o *PostProjectProjectNameStageStageNameResourceCreated) WithPayload(payload *models.Version) *PostProjectProjectNameStageStageNameResourceCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post project project name stage stage name resource created response
func (o *PostProjectProjectNameStageStageNameResourceCreated) SetPayload(payload *models.Version) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostProjectProjectNameStageStageNameResourceCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostProjectProjectNameStageStageNameResourceBadRequestCode is the HTTP code returned for type PostProjectProjectNameStageStageNameResourceBadRequest
const PostProjectProjectNameStageStageNameResourceBadRequestCode int = 400

/*PostProjectProjectNameStageStageNameResourceBadRequest Failed. Stage resource could not be created.

swagger:response postProjectProjectNameStageStageNameResourceBadRequest
*/
type PostProjectProjectNameStageStageNameResourceBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostProjectProjectNameStageStageNameResourceBadRequest creates PostProjectProjectNameStageStageNameResourceBadRequest with default headers values
func NewPostProjectProjectNameStageStageNameResourceBadRequest() *PostProjectProjectNameStageStageNameResourceBadRequest {

	return &PostProjectProjectNameStageStageNameResourceBadRequest{}
}

// WithPayload adds the payload to the post project project name stage stage name resource bad request response
func (o *PostProjectProjectNameStageStageNameResourceBadRequest) WithPayload(payload *models.Error) *PostProjectProjectNameStageStageNameResourceBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post project project name stage stage name resource bad request response
func (o *PostProjectProjectNameStageStageNameResourceBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostProjectProjectNameStageStageNameResourceBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PostProjectProjectNameStageStageNameResourceDefault Error

swagger:response postProjectProjectNameStageStageNameResourceDefault
*/
type PostProjectProjectNameStageStageNameResourceDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostProjectProjectNameStageStageNameResourceDefault creates PostProjectProjectNameStageStageNameResourceDefault with default headers values
func NewPostProjectProjectNameStageStageNameResourceDefault(code int) *PostProjectProjectNameStageStageNameResourceDefault {
	if code <= 0 {
		code = 500
	}

	return &PostProjectProjectNameStageStageNameResourceDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post project project name stage stage name resource default response
func (o *PostProjectProjectNameStageStageNameResourceDefault) WithStatusCode(code int) *PostProjectProjectNameStageStageNameResourceDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post project project name stage stage name resource default response
func (o *PostProjectProjectNameStageStageNameResourceDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post project project name stage stage name resource default response
func (o *PostProjectProjectNameStageStageNameResourceDefault) WithPayload(payload *models.Error) *PostProjectProjectNameStageStageNameResourceDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post project project name stage stage name resource default response
func (o *PostProjectProjectNameStageStageNameResourceDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostProjectProjectNameStageStageNameResourceDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
