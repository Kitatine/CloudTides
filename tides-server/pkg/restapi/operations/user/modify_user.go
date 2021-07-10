// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ModifyUserHandlerFunc turns a function with the right signature into a modify user handler
type ModifyUserHandlerFunc func(ModifyUserParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ModifyUserHandlerFunc) Handle(params ModifyUserParams) middleware.Responder {
	return fn(params)
}

// ModifyUserHandler interface for that can handle valid modify user params
type ModifyUserHandler interface {
	Handle(ModifyUserParams) middleware.Responder
}

// NewModifyUser creates a new http.Handler for the modify user operation
func NewModifyUser(ctx *middleware.Context, handler ModifyUserHandler) *ModifyUser {
	return &ModifyUser{Context: ctx, Handler: handler}
}

/* ModifyUser swagger:route PUT /user/{id} user modifyUser

Modify User

*/
type ModifyUser struct {
	Context *middleware.Context
	Handler ModifyUserHandler
}

func (o *ModifyUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewModifyUserParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// ModifyUserBody modify user body
//
// swagger:model ModifyUserBody
type ModifyUserBody struct {

	// email
	Email string `json:"email,omitempty"`

	// org
	Org string `json:"org,omitempty"`

	// phone
	Phone string `json:"phone,omitempty"`

	// role
	Role string `json:"role,omitempty"`
}

// Validate validates this modify user body
func (o *ModifyUserBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this modify user body based on context it is used
func (o *ModifyUserBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ModifyUserBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ModifyUserBody) UnmarshalBinary(b []byte) error {
	var res ModifyUserBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// ModifyUserNotFoundBody modify user not found body
//
// swagger:model ModifyUserNotFoundBody
type ModifyUserNotFoundBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this modify user not found body
func (o *ModifyUserNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this modify user not found body based on context it is used
func (o *ModifyUserNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ModifyUserNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ModifyUserNotFoundBody) UnmarshalBinary(b []byte) error {
	var res ModifyUserNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// ModifyUserOKBody modify user o k body
//
// swagger:model ModifyUserOKBody
type ModifyUserOKBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this modify user o k body
func (o *ModifyUserOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this modify user o k body based on context it is used
func (o *ModifyUserOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ModifyUserOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ModifyUserOKBody) UnmarshalBinary(b []byte) error {
	var res ModifyUserOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}