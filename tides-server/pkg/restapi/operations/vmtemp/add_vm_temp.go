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

// AddVMTempHandlerFunc turns a function with the right signature into a add VM temp handler
type AddVMTempHandlerFunc func(AddVMTempParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddVMTempHandlerFunc) Handle(params AddVMTempParams) middleware.Responder {
	return fn(params)
}

// AddVMTempHandler interface for that can handle valid add VM temp params
type AddVMTempHandler interface {
	Handle(AddVMTempParams) middleware.Responder
}

// NewAddVMTemp creates a new http.Handler for the add VM temp operation
func NewAddVMTemp(ctx *middleware.Context, handler AddVMTempHandler) *AddVMTemp {
	return &AddVMTemp{Context: ctx, Handler: handler}
}

/* AddVMTemp swagger:route POST /vmtemp vmtemp addVmTemp

add VMTemp

*/
type AddVMTemp struct {
	Context *middleware.Context
	Handler AddVMTempHandler
}

func (o *AddVMTemp) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddVMTempParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// AddVMTempBody add VM temp body
//
// swagger:model AddVMTempBody
type AddVMTempBody struct {

	// disk
	Disk int64 `json:"disk,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// ports
	Ports string `json:"ports,omitempty"`

	// template ID
	TemplateID int64 `json:"templateID,omitempty"`

	// vcpu
	Vcpu int64 `json:"vcpu,omitempty"`

	// vmem
	Vmem int64 `json:"vmem,omitempty"`
}

// Validate validates this add VM temp body
func (o *AddVMTempBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add VM temp body based on context it is used
func (o *AddVMTempBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddVMTempBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddVMTempBody) UnmarshalBinary(b []byte) error {
	var res AddVMTempBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// AddVMTempNotFoundBody add VM temp not found body
//
// swagger:model AddVMTempNotFoundBody
type AddVMTempNotFoundBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this add VM temp not found body
func (o *AddVMTempNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add VM temp not found body based on context it is used
func (o *AddVMTempNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddVMTempNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddVMTempNotFoundBody) UnmarshalBinary(b []byte) error {
	var res AddVMTempNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// AddVMTempOKBody add VM temp o k body
//
// swagger:model AddVMTempOKBody
type AddVMTempOKBody struct {

	// id
	ID int64 `json:"id,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this add VM temp o k body
func (o *AddVMTempOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add VM temp o k body based on context it is used
func (o *AddVMTempOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddVMTempOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddVMTempOKBody) UnmarshalBinary(b []byte) error {
	var res AddVMTempOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
