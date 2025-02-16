// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewListUserOfOrgParams creates a new ListUserOfOrgParams object
//
// There are no default values defined in the spec.
func NewListUserOfOrgParams() ListUserOfOrgParams {

	return ListUserOfOrgParams{}
}

// ListUserOfOrgParams contains all the bound params for the list user of org operation
// typically these are obtained from a http.Request
//
// swagger:parameters list User of org
type ListUserOfOrgParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	OrgName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewListUserOfOrgParams() beforehand.
func (o *ListUserOfOrgParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rOrgName, rhkOrgName, _ := route.Params.GetOK("orgName")
	if err := o.bindOrgName(rOrgName, rhkOrgName, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindOrgName binds and validates parameter OrgName from path.
func (o *ListUserOfOrgParams) bindOrgName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.OrgName = raw

	return nil
}
