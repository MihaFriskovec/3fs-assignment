// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/MihaFriskovec/3fs-assignment/models"
)

// UpdateGroupOKCode is the HTTP code returned for type UpdateGroupOK
const UpdateGroupOKCode int = 200

/*UpdateGroupOK Updated

swagger:response updateGroupOK
*/
type UpdateGroupOK struct {

	/*
	  In: Body
	*/
	Payload *models.Group `json:"body,omitempty"`
}

// NewUpdateGroupOK creates UpdateGroupOK with default headers values
func NewUpdateGroupOK() *UpdateGroupOK {

	return &UpdateGroupOK{}
}

// WithPayload adds the payload to the update group o k response
func (o *UpdateGroupOK) WithPayload(payload *models.Group) *UpdateGroupOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update group o k response
func (o *UpdateGroupOK) SetPayload(payload *models.Group) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateGroupOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateGroupBadRequestCode is the HTTP code returned for type UpdateGroupBadRequest
const UpdateGroupBadRequestCode int = 400

/*UpdateGroupBadRequest User error

swagger:response updateGroupBadRequest
*/
type UpdateGroupBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateGroupBadRequest creates UpdateGroupBadRequest with default headers values
func NewUpdateGroupBadRequest() *UpdateGroupBadRequest {

	return &UpdateGroupBadRequest{}
}

// WithPayload adds the payload to the update group bad request response
func (o *UpdateGroupBadRequest) WithPayload(payload *models.Error) *UpdateGroupBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update group bad request response
func (o *UpdateGroupBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateGroupBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateGroupInternalServerErrorCode is the HTTP code returned for type UpdateGroupInternalServerError
const UpdateGroupInternalServerErrorCode int = 500

/*UpdateGroupInternalServerError Server error

swagger:response updateGroupInternalServerError
*/
type UpdateGroupInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateGroupInternalServerError creates UpdateGroupInternalServerError with default headers values
func NewUpdateGroupInternalServerError() *UpdateGroupInternalServerError {

	return &UpdateGroupInternalServerError{}
}

// WithPayload adds the payload to the update group internal server error response
func (o *UpdateGroupInternalServerError) WithPayload(payload *models.Error) *UpdateGroupInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update group internal server error response
func (o *UpdateGroupInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateGroupInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UpdateGroupDefault error

swagger:response updateGroupDefault
*/
type UpdateGroupDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateGroupDefault creates UpdateGroupDefault with default headers values
func NewUpdateGroupDefault(code int) *UpdateGroupDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateGroupDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update group default response
func (o *UpdateGroupDefault) WithStatusCode(code int) *UpdateGroupDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update group default response
func (o *UpdateGroupDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update group default response
func (o *UpdateGroupDefault) WithPayload(payload *models.Error) *UpdateGroupDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update group default response
func (o *UpdateGroupDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateGroupDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}