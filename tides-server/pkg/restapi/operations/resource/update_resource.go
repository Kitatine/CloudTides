// Code generated by go-swagger; DO NOT EDIT.

package resource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateResourceHandlerFunc turns a function with the right signature into a update resource handler
type UpdateResourceHandlerFunc func(UpdateResourceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateResourceHandlerFunc) Handle(params UpdateResourceParams) middleware.Responder {
	return fn(params)
}

// UpdateResourceHandler interface for that can handle valid update resource params
type UpdateResourceHandler interface {
	Handle(UpdateResourceParams) middleware.Responder
}

// NewUpdateResource creates a new http.Handler for the update resource operation
func NewUpdateResource(ctx *middleware.Context, handler UpdateResourceHandler) *UpdateResource {
	return &UpdateResource{Context: ctx, Handler: handler}
}

/*UpdateResource swagger:route PUT /resource/update resource updateResource

update usage info of resource

*/
type UpdateResource struct {
	Context *middleware.Context
	Handler UpdateResourceHandler
}

func (o *UpdateResource) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateResourceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// UpdateResourceBody update resource body
//
// swagger:model UpdateResourceBody
type UpdateResourceBody struct {

	// current CPU
	CurrentCPU float64 `json:"currentCPU,omitempty"`

	// current RAM
	CurrentRAM float64 `json:"currentRAM,omitempty"`

	// host address
	HostAddress string `json:"hostAddress,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this update resource body
func (o *UpdateResourceBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateResourceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateResourceBody) UnmarshalBinary(b []byte) error {
	var res UpdateResourceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// UpdateResourceNotFoundBody update resource not found body
//
// swagger:model UpdateResourceNotFoundBody
type UpdateResourceNotFoundBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this update resource not found body
func (o *UpdateResourceNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateResourceNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateResourceNotFoundBody) UnmarshalBinary(b []byte) error {
	var res UpdateResourceNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// UpdateResourceOKBody update resource o k body
//
// swagger:model UpdateResourceOKBody
type UpdateResourceOKBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this update resource o k body
func (o *UpdateResourceOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateResourceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateResourceOKBody) UnmarshalBinary(b []byte) error {
	var res UpdateResourceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
