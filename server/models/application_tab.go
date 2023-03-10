// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ApplicationTab application tab
//
// swagger:model application_tab
type ApplicationTab struct {

	// application
	Application *Application `json:"Application,omitempty"`

	// application Id
	ApplicationID int64 `json:"ApplicationId,omitempty"`

	// display order
	DisplayOrder int64 `json:"DisplayOrder,omitempty"`

	// Id
	ID int64 `json:"Id,omitempty"`

	// tab Id
	TabID int64 `json:"TabId,omitempty"`
}

// Validate validates this application tab
func (m *ApplicationTab) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplication(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationTab) validateApplication(formats strfmt.Registry) error {
	if swag.IsZero(m.Application) { // not required
		return nil
	}

	if m.Application != nil {
		if err := m.Application.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Application")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Application")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this application tab based on the context it is used
func (m *ApplicationTab) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApplication(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationTab) contextValidateApplication(ctx context.Context, formats strfmt.Registry) error {

	if m.Application != nil {
		if err := m.Application.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Application")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Application")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationTab) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationTab) UnmarshalBinary(b []byte) error {
	var res ApplicationTab
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
