// Code generated by go-swagger; DO NOT EDIT.

// // This file is part of MinIO Kubernetes Cloud
// // Copyright (c) 2019 MinIO, Inc.
// //
// // This program is free software: you can redistribute it and/or modify
// // it under the terms of the GNU Affero General Public License as published by
// // the Free Software Foundation, either version 3 of the License, or
// // (at your option) any later version.
// //
// // This program is distributed in the hope that it will be useful,
// // but WITHOUT ANY WARRANTY; without even the implied warranty of
// // MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// // GNU Affero General Public License for more details.
// //
// // You should have received a copy of the GNU Affero General Public License
// // along with this program.  If not, see <http://www.gnu.org/licenses/>.

package user_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/minio/m3/mcs/models"
)

// ListBucketsOKCode is the HTTP code returned for type ListBucketsOK
const ListBucketsOKCode int = 200

/*ListBucketsOK A successful response.

swagger:response listBucketsOK
*/
type ListBucketsOK struct {

	/*
	  In: Body
	*/
	Payload *models.ListBucketsResponse `json:"body,omitempty"`
}

// NewListBucketsOK creates ListBucketsOK with default headers values
func NewListBucketsOK() *ListBucketsOK {

	return &ListBucketsOK{}
}

// WithPayload adds the payload to the list buckets o k response
func (o *ListBucketsOK) WithPayload(payload *models.ListBucketsResponse) *ListBucketsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list buckets o k response
func (o *ListBucketsOK) SetPayload(payload *models.ListBucketsResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListBucketsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ListBucketsDefault Generic error response.

swagger:response listBucketsDefault
*/
type ListBucketsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListBucketsDefault creates ListBucketsDefault with default headers values
func NewListBucketsDefault(code int) *ListBucketsDefault {
	if code <= 0 {
		code = 500
	}

	return &ListBucketsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the list buckets default response
func (o *ListBucketsDefault) WithStatusCode(code int) *ListBucketsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the list buckets default response
func (o *ListBucketsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the list buckets default response
func (o *ListBucketsDefault) WithPayload(payload *models.Error) *ListBucketsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list buckets default response
func (o *ListBucketsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListBucketsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
