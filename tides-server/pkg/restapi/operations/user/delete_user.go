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

// DeleteUserHandlerFunc turns a function with the right signature into a delete user handler
type DeleteUserHandlerFunc func(DeleteUserParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteUserHandlerFunc) Handle(params DeleteUserParams) middleware.Responder {
	return fn(params)
}

// DeleteUserHandler interface for that can handle valid delete user params
type DeleteUserHandler interface {
	Handle(DeleteUserParams) middleware.Responder
}

// NewDeleteUser creates a new http.Handler for the delete user operation
func NewDeleteUser(ctx *middleware.Context, handler DeleteUserHandler) *DeleteUser {
	return &DeleteUser{Context: ctx, Handler: handler}
}

/* DeleteUser swagger:route DELETE /user/{id} user deleteUser

delete User

*/
type DeleteUser struct {
	Context *middleware.Context
	Handler DeleteUserHandler
}

func (o *DeleteUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteUserParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// DeleteUserNotFoundBody delete user not found body
//
// swagger:model DeleteUserNotFoundBody
type DeleteUserNotFoundBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this delete user not found body
func (o *DeleteUserNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete user not found body based on context it is used
func (o *DeleteUserNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteUserNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteUserNotFoundBody) UnmarshalBinary(b []byte) error {
	var res DeleteUserNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// DeleteUserOKBody delete user o k body
//
// swagger:model DeleteUserOKBody
type DeleteUserOKBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this delete user o k body
func (o *DeleteUserOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete user o k body based on context it is used
func (o *DeleteUserOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteUserOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteUserOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteUserOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
