// Code generated by go-swagger; DO NOT EDIT.

// SPDX-FileCopyrightText: Copyright 2025 SAP SE or an SAP affiliate company
//
// SPDX-License-Identifier: Apache-2.0

package pools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/sapcc/andromeda/models"
)

// GetPoolsPoolIDOKCode is the HTTP code returned for type GetPoolsPoolIDOK
const GetPoolsPoolIDOKCode int = 200

/*
GetPoolsPoolIDOK Shows the details of a specific pool.

swagger:response getPoolsPoolIdOK
*/
type GetPoolsPoolIDOK struct {

	/*
	  In: Body
	*/
	Payload *GetPoolsPoolIDOKBody `json:"body,omitempty"`
}

// NewGetPoolsPoolIDOK creates GetPoolsPoolIDOK with default headers values
func NewGetPoolsPoolIDOK() *GetPoolsPoolIDOK {

	return &GetPoolsPoolIDOK{}
}

// WithPayload adds the payload to the get pools pool Id o k response
func (o *GetPoolsPoolIDOK) WithPayload(payload *GetPoolsPoolIDOKBody) *GetPoolsPoolIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get pools pool Id o k response
func (o *GetPoolsPoolIDOK) SetPayload(payload *GetPoolsPoolIDOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPoolsPoolIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPoolsPoolIDNotFoundCode is the HTTP code returned for type GetPoolsPoolIDNotFound
const GetPoolsPoolIDNotFoundCode int = 404

/*
GetPoolsPoolIDNotFound Not Found

swagger:response getPoolsPoolIdNotFound
*/
type GetPoolsPoolIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPoolsPoolIDNotFound creates GetPoolsPoolIDNotFound with default headers values
func NewGetPoolsPoolIDNotFound() *GetPoolsPoolIDNotFound {

	return &GetPoolsPoolIDNotFound{}
}

// WithPayload adds the payload to the get pools pool Id not found response
func (o *GetPoolsPoolIDNotFound) WithPayload(payload *models.Error) *GetPoolsPoolIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get pools pool Id not found response
func (o *GetPoolsPoolIDNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPoolsPoolIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
GetPoolsPoolIDDefault Unexpected Error

swagger:response getPoolsPoolIdDefault
*/
type GetPoolsPoolIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPoolsPoolIDDefault creates GetPoolsPoolIDDefault with default headers values
func NewGetPoolsPoolIDDefault(code int) *GetPoolsPoolIDDefault {
	if code <= 0 {
		code = 500
	}

	return &GetPoolsPoolIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get pools pool ID default response
func (o *GetPoolsPoolIDDefault) WithStatusCode(code int) *GetPoolsPoolIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get pools pool ID default response
func (o *GetPoolsPoolIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get pools pool ID default response
func (o *GetPoolsPoolIDDefault) WithPayload(payload *models.Error) *GetPoolsPoolIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get pools pool ID default response
func (o *GetPoolsPoolIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPoolsPoolIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
