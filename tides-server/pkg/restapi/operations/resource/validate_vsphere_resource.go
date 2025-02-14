// Code generated by go-swagger; DO NOT EDIT.

package resource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ValidateVsphereResourceHandlerFunc turns a function with the right signature into a validate vsphere resource handler
type ValidateVsphereResourceHandlerFunc func(ValidateVsphereResourceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ValidateVsphereResourceHandlerFunc) Handle(params ValidateVsphereResourceParams) middleware.Responder {
	return fn(params)
}

// ValidateVsphereResourceHandler interface for that can handle valid validate vsphere resource params
type ValidateVsphereResourceHandler interface {
	Handle(ValidateVsphereResourceParams) middleware.Responder
}

// NewValidateVsphereResource creates a new http.Handler for the validate vsphere resource operation
func NewValidateVsphereResource(ctx *middleware.Context, handler ValidateVsphereResourceHandler) *ValidateVsphereResource {
	return &ValidateVsphereResource{Context: ctx, Handler: handler}
}

/* ValidateVsphereResource swagger:route GET /resource/vsphere/validate resource validateVsphereResource

returns the list of data centers belonging to the host

*/
type ValidateVsphereResource struct {
	Context *middleware.Context
	Handler ValidateVsphereResourceHandler
}

func (o *ValidateVsphereResource) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewValidateVsphereResourceParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// ValidateVsphereResourceBody validate vsphere resource body
//
// swagger:model ValidateVsphereResourceBody
type ValidateVsphereResourceBody struct {

	// host
	Host string `json:"host,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this validate vsphere resource body
func (o *ValidateVsphereResourceBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this validate vsphere resource body based on context it is used
func (o *ValidateVsphereResourceBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ValidateVsphereResourceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ValidateVsphereResourceBody) UnmarshalBinary(b []byte) error {
	var res ValidateVsphereResourceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// ValidateVsphereResourceNotFoundBody validate vsphere resource not found body
//
// swagger:model ValidateVsphereResourceNotFoundBody
type ValidateVsphereResourceNotFoundBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this validate vsphere resource not found body
func (o *ValidateVsphereResourceNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this validate vsphere resource not found body based on context it is used
func (o *ValidateVsphereResourceNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ValidateVsphereResourceNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ValidateVsphereResourceNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ValidateVsphereResourceNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// ValidateVsphereResourceOKBody list of data centers belonging to the host
//
// swagger:model ValidateVsphereResourceOKBody
type ValidateVsphereResourceOKBody struct {

	// message
	Message string `json:"message,omitempty"`

	// results
	Results []string `json:"results"`
}

// Validate validates this validate vsphere resource o k body
func (o *ValidateVsphereResourceOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this validate vsphere resource o k body based on context it is used
func (o *ValidateVsphereResourceOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ValidateVsphereResourceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ValidateVsphereResourceOKBody) UnmarshalBinary(b []byte) error {
	var res ValidateVsphereResourceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
