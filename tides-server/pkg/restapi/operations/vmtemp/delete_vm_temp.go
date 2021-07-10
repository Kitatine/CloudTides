// Code generated by go-swagger; DO NOT EDIT.

package vmtemp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteVMTempHandlerFunc turns a function with the right signature into a delete VM temp handler
type DeleteVMTempHandlerFunc func(DeleteVMTempParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteVMTempHandlerFunc) Handle(params DeleteVMTempParams) middleware.Responder {
	return fn(params)
}

// DeleteVMTempHandler interface for that can handle valid delete VM temp params
type DeleteVMTempHandler interface {
	Handle(DeleteVMTempParams) middleware.Responder
}

// NewDeleteVMTemp creates a new http.Handler for the delete VM temp operation
func NewDeleteVMTemp(ctx *middleware.Context, handler DeleteVMTempHandler) *DeleteVMTemp {
	return &DeleteVMTemp{Context: ctx, Handler: handler}
}

/* DeleteVMTemp swagger:route DELETE /vmtemp/{id} vmtemp deleteVmTemp

delete VMTemplate

*/
type DeleteVMTemp struct {
	Context *middleware.Context
	Handler DeleteVMTempHandler
}

func (o *DeleteVMTemp) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteVMTempParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// DeleteVMTempNotFoundBody delete VM temp not found body
//
// swagger:model DeleteVMTempNotFoundBody
type DeleteVMTempNotFoundBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this delete VM temp not found body
func (o *DeleteVMTempNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete VM temp not found body based on context it is used
func (o *DeleteVMTempNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteVMTempNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteVMTempNotFoundBody) UnmarshalBinary(b []byte) error {
	var res DeleteVMTempNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// DeleteVMTempOKBody delete VM temp o k body
//
// swagger:model DeleteVMTempOKBody
type DeleteVMTempOKBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this delete VM temp o k body
func (o *DeleteVMTempOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete VM temp o k body based on context it is used
func (o *DeleteVMTempOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteVMTempOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteVMTempOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteVMTempOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}