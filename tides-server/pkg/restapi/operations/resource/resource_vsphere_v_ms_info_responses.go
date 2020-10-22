// Code generated by go-swagger; DO NOT EDIT.

package resource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ResourceVsphereVMsInfoOKCode is the HTTP code returned for type ResourceVsphereVMsInfoOK
const ResourceVsphereVMsInfoOKCode int = 200

/*ResourceVsphereVMsInfoOK get detailed info of VMs belonging to a user

swagger:response resourceVsphereVMsInfoOK
*/
type ResourceVsphereVMsInfoOK struct {

	/*
	  In: Body
	*/
	Payload *ResourceVsphereVMsInfoOKBody `json:"body,omitempty"`
}

// NewResourceVsphereVMsInfoOK creates ResourceVsphereVMsInfoOK with default headers values
func NewResourceVsphereVMsInfoOK() *ResourceVsphereVMsInfoOK {

	return &ResourceVsphereVMsInfoOK{}
}

// WithPayload adds the payload to the resource vsphere v ms info o k response
func (o *ResourceVsphereVMsInfoOK) WithPayload(payload *ResourceVsphereVMsInfoOKBody) *ResourceVsphereVMsInfoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the resource vsphere v ms info o k response
func (o *ResourceVsphereVMsInfoOK) SetPayload(payload *ResourceVsphereVMsInfoOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ResourceVsphereVMsInfoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ResourceVsphereVMsInfoUnauthorizedCode is the HTTP code returned for type ResourceVsphereVMsInfoUnauthorized
const ResourceVsphereVMsInfoUnauthorizedCode int = 401

/*ResourceVsphereVMsInfoUnauthorized Unauthorized

swagger:response resourceVsphereVMsInfoUnauthorized
*/
type ResourceVsphereVMsInfoUnauthorized struct {
}

// NewResourceVsphereVMsInfoUnauthorized creates ResourceVsphereVMsInfoUnauthorized with default headers values
func NewResourceVsphereVMsInfoUnauthorized() *ResourceVsphereVMsInfoUnauthorized {

	return &ResourceVsphereVMsInfoUnauthorized{}
}

// WriteResponse to the client
func (o *ResourceVsphereVMsInfoUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ResourceVsphereVMsInfoNotFoundCode is the HTTP code returned for type ResourceVsphereVMsInfoNotFound
const ResourceVsphereVMsInfoNotFoundCode int = 404

/*ResourceVsphereVMsInfoNotFound resource not found

swagger:response resourceVsphereVMsInfoNotFound
*/
type ResourceVsphereVMsInfoNotFound struct {

	/*
	  In: Body
	*/
	Payload *ResourceVsphereVMsInfoNotFoundBody `json:"body,omitempty"`
}

// NewResourceVsphereVMsInfoNotFound creates ResourceVsphereVMsInfoNotFound with default headers values
func NewResourceVsphereVMsInfoNotFound() *ResourceVsphereVMsInfoNotFound {

	return &ResourceVsphereVMsInfoNotFound{}
}

// WithPayload adds the payload to the resource vsphere v ms info not found response
func (o *ResourceVsphereVMsInfoNotFound) WithPayload(payload *ResourceVsphereVMsInfoNotFoundBody) *ResourceVsphereVMsInfoNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the resource vsphere v ms info not found response
func (o *ResourceVsphereVMsInfoNotFound) SetPayload(payload *ResourceVsphereVMsInfoNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ResourceVsphereVMsInfoNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}