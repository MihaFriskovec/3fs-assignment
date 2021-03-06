// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/MihaFriskovec/3fs-assignment/models"
)

// CreateGroupCreatedCode is the HTTP code returned for type CreateGroupCreated
const CreateGroupCreatedCode int = 201

/*CreateGroupCreated Created

swagger:response createGroupCreated
*/
type CreateGroupCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Group `json:"body,omitempty"`
}

// NewCreateGroupCreated creates CreateGroupCreated with default headers values
func NewCreateGroupCreated() *CreateGroupCreated {

	return &CreateGroupCreated{}
}

// WithPayload adds the payload to the create group created response
func (o *CreateGroupCreated) WithPayload(payload *models.Group) *CreateGroupCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create group created response
func (o *CreateGroupCreated) SetPayload(payload *models.Group) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateGroupCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateGroupBadRequestCode is the HTTP code returned for type CreateGroupBadRequest
const CreateGroupBadRequestCode int = 400

/*CreateGroupBadRequest User error

swagger:response createGroupBadRequest
*/
type CreateGroupBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateGroupBadRequest creates CreateGroupBadRequest with default headers values
func NewCreateGroupBadRequest() *CreateGroupBadRequest {

	return &CreateGroupBadRequest{}
}

// WithPayload adds the payload to the create group bad request response
func (o *CreateGroupBadRequest) WithPayload(payload *models.Error) *CreateGroupBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create group bad request response
func (o *CreateGroupBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateGroupBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateGroupInternalServerErrorCode is the HTTP code returned for type CreateGroupInternalServerError
const CreateGroupInternalServerErrorCode int = 500

/*CreateGroupInternalServerError Server error

swagger:response createGroupInternalServerError
*/
type CreateGroupInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateGroupInternalServerError creates CreateGroupInternalServerError with default headers values
func NewCreateGroupInternalServerError() *CreateGroupInternalServerError {

	return &CreateGroupInternalServerError{}
}

// WithPayload adds the payload to the create group internal server error response
func (o *CreateGroupInternalServerError) WithPayload(payload *models.Error) *CreateGroupInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create group internal server error response
func (o *CreateGroupInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateGroupInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateGroupDefault Unknown error

swagger:response createGroupDefault
*/
type CreateGroupDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateGroupDefault creates CreateGroupDefault with default headers values
func NewCreateGroupDefault(code int) *CreateGroupDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateGroupDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create group default response
func (o *CreateGroupDefault) WithStatusCode(code int) *CreateGroupDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create group default response
func (o *CreateGroupDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create group default response
func (o *CreateGroupDefault) WithPayload(payload *models.Error) *CreateGroupDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create group default response
func (o *CreateGroupDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateGroupDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
