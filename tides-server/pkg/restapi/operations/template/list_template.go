// Code generated by go-swagger; DO NOT EDIT.

package template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ListTemplateHandlerFunc turns a function with the right signature into a list template handler
type ListTemplateHandlerFunc func(ListTemplateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListTemplateHandlerFunc) Handle(params ListTemplateParams) middleware.Responder {
	return fn(params)
}

// ListTemplateHandler interface for that can handle valid list template params
type ListTemplateHandler interface {
	Handle(ListTemplateParams) middleware.Responder
}

// NewListTemplate creates a new http.Handler for the list template operation
func NewListTemplate(ctx *middleware.Context, handler ListTemplateHandler) *ListTemplate {
	return &ListTemplate{Context: ctx, Handler: handler}
}

/* ListTemplate swagger:route GET /template template listTemplate

list all available VM templates

*/
type ListTemplate struct {
	Context *middleware.Context
	Handler ListTemplateHandler
}

func (o *ListTemplate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListTemplateParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// ListTemplateOKBodyItems0 list template o k body items0
//
// swagger:model ListTemplateOKBodyItems0
type ListTemplateOKBodyItems0 struct {

	// compatibility
	Compatibility string `json:"compatibility,omitempty"`

	// date added
	DateAdded string `json:"dateAdded,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// guest o s
	GuestOS string `json:"guestOS,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// memory size
	MemorySize float64 `json:"memorySize,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// provisioned space
	ProvisionedSpace float64 `json:"provisionedSpace,omitempty"`

	// resource ID
	ResourceID int64 `json:"resourceID,omitempty"`

	// tag
	Tag string `json:"tag,omitempty"`

	// template type
	// Enum: [datastore upload]
	TemplateType string `json:"templateType,omitempty"`

	// vcpu
	Vcpu int64 `json:"vcpu,omitempty"`
}

// Validate validates this list template o k body items0
func (o *ListTemplateOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTemplateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var listTemplateOKBodyItems0TypeTemplateTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["datastore","upload"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		listTemplateOKBodyItems0TypeTemplateTypePropEnum = append(listTemplateOKBodyItems0TypeTemplateTypePropEnum, v)
	}
}

const (

	// ListTemplateOKBodyItems0TemplateTypeDatastore captures enum value "datastore"
	ListTemplateOKBodyItems0TemplateTypeDatastore string = "datastore"

	// ListTemplateOKBodyItems0TemplateTypeUpload captures enum value "upload"
	ListTemplateOKBodyItems0TemplateTypeUpload string = "upload"
)

// prop value enum
func (o *ListTemplateOKBodyItems0) validateTemplateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, listTemplateOKBodyItems0TypeTemplateTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ListTemplateOKBodyItems0) validateTemplateType(formats strfmt.Registry) error {
	if swag.IsZero(o.TemplateType) { // not required
		return nil
	}

	// value enum
	if err := o.validateTemplateTypeEnum("templateType", "body", o.TemplateType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this list template o k body items0 based on context it is used
func (o *ListTemplateOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListTemplateOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListTemplateOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res ListTemplateOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
