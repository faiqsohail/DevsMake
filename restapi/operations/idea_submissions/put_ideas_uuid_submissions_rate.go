// Code generated by go-swagger; DO NOT EDIT.

package idea_submissions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"devsmake/models"
)

// PutIdeasUUIDSubmissionsRateHandlerFunc turns a function with the right signature into a put ideas UUID submissions rate handler
type PutIdeasUUIDSubmissionsRateHandlerFunc func(PutIdeasUUIDSubmissionsRateParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn PutIdeasUUIDSubmissionsRateHandlerFunc) Handle(params PutIdeasUUIDSubmissionsRateParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// PutIdeasUUIDSubmissionsRateHandler interface for that can handle valid put ideas UUID submissions rate params
type PutIdeasUUIDSubmissionsRateHandler interface {
	Handle(PutIdeasUUIDSubmissionsRateParams, *models.Principal) middleware.Responder
}

// NewPutIdeasUUIDSubmissionsRate creates a new http.Handler for the put ideas UUID submissions rate operation
func NewPutIdeasUUIDSubmissionsRate(ctx *middleware.Context, handler PutIdeasUUIDSubmissionsRateHandler) *PutIdeasUUIDSubmissionsRate {
	return &PutIdeasUUIDSubmissionsRate{Context: ctx, Handler: handler}
}

/* PutIdeasUUIDSubmissionsRate swagger:route PUT /ideas/{uuid}/submissions/rate idea submissions putIdeasUuidSubmissionsRate

rate an idea post submission

*/
type PutIdeasUUIDSubmissionsRate struct {
	Context *middleware.Context
	Handler PutIdeasUUIDSubmissionsRateHandler
}

func (o *PutIdeasUUIDSubmissionsRate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPutIdeasUUIDSubmissionsRateParams()
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

// PutIdeasUUIDSubmissionsRateBody put ideas UUID submissions rate body
//
// swagger:model PutIdeasUUIDSubmissionsRateBody
type PutIdeasUUIDSubmissionsRateBody struct {

	// rating
	// Example: 10
	// Required: true
	// Maximum: 10
	// Minimum: 1
	Rating *int64 `json:"rating"`

	// submission uuid
	// Example: e2c6b170-eb62-4006-8a26-8a57bc36a4ae
	// Required: true
	SubmissionUUID *string `json:"submission_uuid"`
}

// Validate validates this put ideas UUID submissions rate body
func (o *PutIdeasUUIDSubmissionsRateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRating(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSubmissionUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutIdeasUUIDSubmissionsRateBody) validateRating(formats strfmt.Registry) error {

	if err := validate.Required("rating"+"."+"rating", "body", o.Rating); err != nil {
		return err
	}

	if err := validate.MinimumInt("rating"+"."+"rating", "body", *o.Rating, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("rating"+"."+"rating", "body", *o.Rating, 10, false); err != nil {
		return err
	}

	return nil
}

func (o *PutIdeasUUIDSubmissionsRateBody) validateSubmissionUUID(formats strfmt.Registry) error {

	if err := validate.Required("rating"+"."+"submission_uuid", "body", o.SubmissionUUID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this put ideas UUID submissions rate body based on context it is used
func (o *PutIdeasUUIDSubmissionsRateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutIdeasUUIDSubmissionsRateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutIdeasUUIDSubmissionsRateBody) UnmarshalBinary(b []byte) error {
	var res PutIdeasUUIDSubmissionsRateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
