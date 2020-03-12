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
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteBucketHandlerFunc turns a function with the right signature into a delete bucket handler
type DeleteBucketHandlerFunc func(DeleteBucketParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteBucketHandlerFunc) Handle(params DeleteBucketParams) middleware.Responder {
	return fn(params)
}

// DeleteBucketHandler interface for that can handle valid delete bucket params
type DeleteBucketHandler interface {
	Handle(DeleteBucketParams) middleware.Responder
}

// NewDeleteBucket creates a new http.Handler for the delete bucket operation
func NewDeleteBucket(ctx *middleware.Context, handler DeleteBucketHandler) *DeleteBucket {
	return &DeleteBucket{Context: ctx, Handler: handler}
}

/*DeleteBucket swagger:route DELETE /api/v1/buckets/{name} UserAPI deleteBucket

Delete Bucket

*/
type DeleteBucket struct {
	Context *middleware.Context
	Handler DeleteBucketHandler
}

func (o *DeleteBucket) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteBucketParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
