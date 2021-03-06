// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/MihaFriskovec/3fs-assignment/models"
)

// UpdateUserOKCode is the HTTP code returned for type UpdateUserOK
const UpdateUserOKCode int = 200

/*UpdateUserOK Updated

swagger:response updateUserOK
*/
type UpdateUserOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewUpdateUserOK creates UpdateUserOK with default headers values
func NewUpdateUserOK() *UpdateUserOK {

	return &UpdateUserOK{}
}

// WithPayload adds the payload to the update user o k response
func (o *UpdateUserOK) WithPayload(payload *models.User) *UpdateUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user o k response
func (o *UpdateUserOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateUserBadRequestCode is the HTTP code returned for type UpdateUserBadRequest
const UpdateUserBadRequestCode int = 400

/*UpdateUserBadRequest User error

swagger:response updateUserBadRequest
*/
type UpdateUserBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateUserBadRequest creates UpdateUserBadRequest with default headers values
func NewUpdateUserBadRequest() *UpdateUserBadRequest {

	return &UpdateUserBadRequest{}
}

// WithPayload adds the payload to the update user bad request response
func (o *UpdateUserBadRequest) WithPayload(payload *models.Error) *UpdateUserBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user bad request response
func (o *UpdateUserBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateUserInternalServerErrorCode is the HTTP code returned for type UpdateUserInternalServerError
const UpdateUserInternalServerErrorCode int = 500

/*UpdateUserInternalServerError Server error

swagger:response updateUserInternalServerError
*/
type UpdateUserInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateUserInternalServerError creates UpdateUserInternalServerError with default headers values
func NewUpdateUserInternalServerError() *UpdateUserInternalServerError {

	return &UpdateUserInternalServerError{}
}

// WithPayload adds the payload to the update user internal server error response
func (o *UpdateUserInternalServerError) WithPayload(payload *models.Error) *UpdateUserInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user internal server error response
func (o *UpdateUserInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UpdateUserDefault error

swagger:response updateUserDefault
*/
type UpdateUserDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateUserDefault creates UpdateUserDefault with default headers values
func NewUpdateUserDefault(code int) *UpdateUserDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateUserDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update user default response
func (o *UpdateUserDefault) WithStatusCode(code int) *UpdateUserDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update user default response
func (o *UpdateUserDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update user default response
func (o *UpdateUserDefault) WithPayload(payload *models.Error) *UpdateUserDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user default response
func (o *UpdateUserDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUserDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
