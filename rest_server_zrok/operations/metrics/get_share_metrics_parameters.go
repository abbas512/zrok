// Code generated by go-swagger; DO NOT EDIT.

package metrics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetShareMetricsParams creates a new GetShareMetricsParams object
//
// There are no default values defined in the spec.
func NewGetShareMetricsParams() GetShareMetricsParams {

	return GetShareMetricsParams{}
}

// GetShareMetricsParams contains all the bound params for the get share metrics operation
// typically these are obtained from a http.Request
//
// swagger:parameters getShareMetrics
type GetShareMetricsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	ShrToken string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetShareMetricsParams() beforehand.
func (o *GetShareMetricsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rShrToken, rhkShrToken, _ := route.Params.GetOK("shrToken")
	if err := o.bindShrToken(rShrToken, rhkShrToken, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindShrToken binds and validates parameter ShrToken from path.
func (o *GetShareMetricsParams) bindShrToken(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ShrToken = raw

	return nil
}
