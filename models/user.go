// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// User user
// swagger:model user
type User struct {

	// email
	// Required: true
	// Max Length: 50
	// Min Length: 1
	Email *string `json:"email"`

	// group
	Group primitive.ObjectID `json:"group,omitempty" bson:"group"`

	// id
	// Read Only: true
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`

	// name
	// Required: true
	// Max Length: 50
	// Min Length: 1
	Name *string `json:"name"`

	// password
	// Required: true
	// Max Length: 100
	// Min Length: 6
	Password *string `json:"password"`
}

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *User) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	if err := validate.MinLength("email", "body", string(*m.Email), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("email", "body", string(*m.Email), 50); err != nil {
		return err
	}

	return nil
}

func (m *User) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 50); err != nil {
		return err
	}

	return nil
}

func (m *User) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	if err := validate.MinLength("password", "body", string(*m.Password), 6); err != nil {
		return err
	}

	if err := validate.MaxLength("password", "body", string(*m.Password), 100); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}