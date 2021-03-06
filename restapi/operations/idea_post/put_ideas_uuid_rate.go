// Code generated by go-swagger; DO NOT EDIT.

package idea_post

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

	"devsmake/models"
)

// PutIdeasUUIDRateHandlerFunc turns a function with the right signature into a put ideas UUID rate handler
type PutIdeasUUIDRateHandlerFunc func(PutIdeasUUIDRateParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn PutIdeasUUIDRateHandlerFunc) Handle(params PutIdeasUUIDRateParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// PutIdeasUUIDRateHandler interface for that can handle valid put ideas UUID rate params
type PutIdeasUUIDRateHandler interface {
	Handle(PutIdeasUUIDRateParams, *models.Principal) middleware.Responder
}

// NewPutIdeasUUIDRate creates a new http.Handler for the put ideas UUID rate operation
func NewPutIdeasUUIDRate(ctx *middleware.Context, handler PutIdeasUUIDRateHandler) *PutIdeasUUIDRate {
	return &PutIdeasUUIDRate{Context: ctx, Handler: handler}
}

/* PutIdeasUUIDRate swagger:route PUT /ideas/{uuid}/rate idea post putIdeasUuidRate

rate an idea post

*/
type PutIdeasUUIDRate struct {
	Context *middleware.Context
	Handler PutIdeasUUIDRateHandler
}

func (o *PutIdeasUUIDRate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPutIdeasUUIDRateParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PutIdeasUUIDRateBody put ideas UUID rate body
//
// swagger:model PutIdeasUUIDRateBody
type PutIdeasUUIDRateBody struct {

	// rating
	// Required: true
	// Enum: [like neutral dislike]
	Rating *string `json:"rating"`
}

// Validate validates this put ideas UUID rate body
func (o *PutIdeasUUIDRateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRating(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var putIdeasUuidRateBodyTypeRatingPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["like","neutral","dislike"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		putIdeasUuidRateBodyTypeRatingPropEnum = append(putIdeasUuidRateBodyTypeRatingPropEnum, v)
	}
}

const (

	// PutIdeasUUIDRateBodyRatingLike captures enum value "like"
	PutIdeasUUIDRateBodyRatingLike string = "like"

	// PutIdeasUUIDRateBodyRatingNeutral captures enum value "neutral"
	PutIdeasUUIDRateBodyRatingNeutral string = "neutral"

	// PutIdeasUUIDRateBodyRatingDislike captures enum value "dislike"
	PutIdeasUUIDRateBodyRatingDislike string = "dislike"
)

// prop value enum
func (o *PutIdeasUUIDRateBody) validateRatingEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, putIdeasUuidRateBodyTypeRatingPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PutIdeasUUIDRateBody) validateRating(formats strfmt.Registry) error {

	if err := validate.Required("rating"+"."+"rating", "body", o.Rating); err != nil {
		return err
	}

	// value enum
	if err := o.validateRatingEnum("rating"+"."+"rating", "body", *o.Rating); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this put ideas UUID rate body based on context it is used
func (o *PutIdeasUUIDRateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutIdeasUUIDRateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutIdeasUUIDRateBody) UnmarshalBinary(b []byte) error {
	var res PutIdeasUUIDRateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
