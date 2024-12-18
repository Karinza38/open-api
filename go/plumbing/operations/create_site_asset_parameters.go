// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewCreateSiteAssetParams creates a new CreateSiteAssetParams object
// with the default values initialized.
func NewCreateSiteAssetParams() *CreateSiteAssetParams {
	var ()
	return &CreateSiteAssetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSiteAssetParamsWithTimeout creates a new CreateSiteAssetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSiteAssetParamsWithTimeout(timeout time.Duration) *CreateSiteAssetParams {
	var ()
	return &CreateSiteAssetParams{

		timeout: timeout,
	}
}

// NewCreateSiteAssetParamsWithContext creates a new CreateSiteAssetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSiteAssetParamsWithContext(ctx context.Context) *CreateSiteAssetParams {
	var ()
	return &CreateSiteAssetParams{

		Context: ctx,
	}
}

// NewCreateSiteAssetParamsWithHTTPClient creates a new CreateSiteAssetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSiteAssetParamsWithHTTPClient(client *http.Client) *CreateSiteAssetParams {
	var ()
	return &CreateSiteAssetParams{
		HTTPClient: client,
	}
}

/*
CreateSiteAssetParams contains all the parameters to send to the API endpoint
for the create site asset operation typically these are written to a http.Request
*/
type CreateSiteAssetParams struct {

	/*ContentType*/
	ContentType string
	/*Name*/
	Name string
	/*SiteID*/
	SiteID string
	/*Size*/
	Size int64
	/*Visibility*/
	Visibility *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create site asset params
func (o *CreateSiteAssetParams) WithTimeout(timeout time.Duration) *CreateSiteAssetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create site asset params
func (o *CreateSiteAssetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create site asset params
func (o *CreateSiteAssetParams) WithContext(ctx context.Context) *CreateSiteAssetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create site asset params
func (o *CreateSiteAssetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create site asset params
func (o *CreateSiteAssetParams) WithHTTPClient(client *http.Client) *CreateSiteAssetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create site asset params
func (o *CreateSiteAssetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentType adds the contentType to the create site asset params
func (o *CreateSiteAssetParams) WithContentType(contentType string) *CreateSiteAssetParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the create site asset params
func (o *CreateSiteAssetParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithName adds the name to the create site asset params
func (o *CreateSiteAssetParams) WithName(name string) *CreateSiteAssetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the create site asset params
func (o *CreateSiteAssetParams) SetName(name string) {
	o.Name = name
}

// WithSiteID adds the siteID to the create site asset params
func (o *CreateSiteAssetParams) WithSiteID(siteID string) *CreateSiteAssetParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the create site asset params
func (o *CreateSiteAssetParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WithSize adds the size to the create site asset params
func (o *CreateSiteAssetParams) WithSize(size int64) *CreateSiteAssetParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the create site asset params
func (o *CreateSiteAssetParams) SetSize(size int64) {
	o.Size = size
}

// WithVisibility adds the visibility to the create site asset params
func (o *CreateSiteAssetParams) WithVisibility(visibility *string) *CreateSiteAssetParams {
	o.SetVisibility(visibility)
	return o
}

// SetVisibility adds the visibility to the create site asset params
func (o *CreateSiteAssetParams) SetVisibility(visibility *string) {
	o.Visibility = visibility
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSiteAssetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param content_type
	qrContentType := o.ContentType
	qContentType := qrContentType
	if qContentType != "" {
		if err := r.SetQueryParam("content_type", qContentType); err != nil {
			return err
		}
	}

	// query param name
	qrName := o.Name
	qName := qrName
	if qName != "" {
		if err := r.SetQueryParam("name", qName); err != nil {
			return err
		}
	}

	// path param site_id
	if err := r.SetPathParam("site_id", o.SiteID); err != nil {
		return err
	}

	// query param size
	qrSize := o.Size
	qSize := swag.FormatInt64(qrSize)
	if qSize != "" {
		if err := r.SetQueryParam("size", qSize); err != nil {
			return err
		}
	}

	if o.Visibility != nil {

		// query param visibility
		var qrVisibility string
		if o.Visibility != nil {
			qrVisibility = *o.Visibility
		}
		qVisibility := qrVisibility
		if qVisibility != "" {
			if err := r.SetQueryParam("visibility", qVisibility); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
