// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DevServer dev server
//
// swagger:model devServer
type DevServer struct {

	// branch
	Branch string `json:"branch,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// done at
	DoneAt string `json:"done_at,omitempty"`

	// error at
	ErrorAt string `json:"error_at,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// live at
	LiveAt string `json:"live_at,omitempty"`

	// site id
	SiteID string `json:"site_id,omitempty"`

	// starting at
	StartingAt string `json:"starting_at,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this dev server
func (m *DevServer) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DevServer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DevServer) UnmarshalBinary(b []byte) error {
	var res DevServer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
