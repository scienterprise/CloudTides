// Code generated by go-swagger; DO NOT EDIT.

package vendor_swagger

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ListVendorOKCode is the HTTP code returned for type ListVendorOK
const ListVendorOKCode int = 200

/*ListVendorOK OK

swagger:response listVendorOK
*/
type ListVendorOK struct {

	/*
	  In: Body
	*/
	Payload []*ListVendorOKBodyItems0 `json:"body,omitempty"`
}

// NewListVendorOK creates ListVendorOK with default headers values
func NewListVendorOK() *ListVendorOK {

	return &ListVendorOK{}
}

// WithPayload adds the payload to the list vendor o k response
func (o *ListVendorOK) WithPayload(payload []*ListVendorOKBodyItems0) *ListVendorOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list vendor o k response
func (o *ListVendorOK) SetPayload(payload []*ListVendorOKBodyItems0) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListVendorOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*ListVendorOKBodyItems0, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListVendorUnauthorizedCode is the HTTP code returned for type ListVendorUnauthorized
const ListVendorUnauthorizedCode int = 401

/*ListVendorUnauthorized Unauthorized

swagger:response listVendorUnauthorized
*/
type ListVendorUnauthorized struct {
}

// NewListVendorUnauthorized creates ListVendorUnauthorized with default headers values
func NewListVendorUnauthorized() *ListVendorUnauthorized {

	return &ListVendorUnauthorized{}
}

// WriteResponse to the client
func (o *ListVendorUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
