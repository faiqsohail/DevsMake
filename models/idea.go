// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Idea idea
//
// swagger:model Idea
type Idea struct {

	// author id
	AuthorID int64 `json:"author_id,omitempty"`

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// description
	// Required: true
	Description *string `json:"description"`

	// dislikes
	Dislikes int64 `json:"dislikes,omitempty"`

	// likes
	Likes int64 `json:"likes,omitempty"`

	// submissions
	Submissions int64 `json:"submissions,omitempty"`

	// title
	// Required: true
	Title *string `json:"title"`

	// uuid
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this idea
func (m *Idea) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Idea) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Idea) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *Idea) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this idea based on context it is used
func (m *Idea) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Idea) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Idea) UnmarshalBinary(b []byte) error {
	var res Idea
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}