// Code generated by go-swagger; DO NOT EDIT.

package vendor_swagger

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListVendorHandlerFunc turns a function with the right signature into a list vendor handler
type ListVendorHandlerFunc func(ListVendorParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListVendorHandlerFunc) Handle(params ListVendorParams) middleware.Responder {
	return fn(params)
}

// ListVendorHandler interface for that can handle valid list vendor params
type ListVendorHandler interface {
	Handle(ListVendorParams) middleware.Responder
}

// NewListVendor creates a new http.Handler for the list vendor operation
func NewListVendor(ctx *middleware.Context, handler ListVendorHandler) *ListVendor {
	return &ListVendor{Context: ctx, Handler: handler}
}

/* ListVendor swagger:route GET /vendors vendor listVendor

list vendors

*/
type ListVendor struct {
	Context *middleware.Context
	Handler ListVendorHandler
}

func (o *ListVendor) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListVendorParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// ListVendorOKBodyItems0 list vendor o k body items0
//
// swagger:model ListVendorOKBodyItems0
type ListVendorOKBodyItems0 struct {

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// url
	URL string `json:"url,omitempty"`

	// vendor type
	VendorType string `json:"vendorType,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this list vendor o k body items0
func (o *ListVendorOKBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list vendor o k body items0 based on context it is used
func (o *ListVendorOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListVendorOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListVendorOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res ListVendorOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}