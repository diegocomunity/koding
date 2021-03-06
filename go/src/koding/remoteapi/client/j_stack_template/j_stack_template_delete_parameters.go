package j_stack_template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// NewJStackTemplateDeleteParams creates a new JStackTemplateDeleteParams object
// with the default values initialized.
func NewJStackTemplateDeleteParams() *JStackTemplateDeleteParams {
	var ()
	return &JStackTemplateDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJStackTemplateDeleteParamsWithTimeout creates a new JStackTemplateDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJStackTemplateDeleteParamsWithTimeout(timeout time.Duration) *JStackTemplateDeleteParams {
	var ()
	return &JStackTemplateDeleteParams{

		timeout: timeout,
	}
}

// NewJStackTemplateDeleteParamsWithContext creates a new JStackTemplateDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewJStackTemplateDeleteParamsWithContext(ctx context.Context) *JStackTemplateDeleteParams {
	var ()
	return &JStackTemplateDeleteParams{

		Context: ctx,
	}
}

/*JStackTemplateDeleteParams contains all the parameters to send to the API endpoint
for the j stack template delete operation typically these are written to a http.Request
*/
type JStackTemplateDeleteParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector
	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the j stack template delete params
func (o *JStackTemplateDeleteParams) WithTimeout(timeout time.Duration) *JStackTemplateDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j stack template delete params
func (o *JStackTemplateDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j stack template delete params
func (o *JStackTemplateDeleteParams) WithContext(ctx context.Context) *JStackTemplateDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j stack template delete params
func (o *JStackTemplateDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j stack template delete params
func (o *JStackTemplateDeleteParams) WithBody(body models.DefaultSelector) *JStackTemplateDeleteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j stack template delete params
func (o *JStackTemplateDeleteParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WithID adds the id to the j stack template delete params
func (o *JStackTemplateDeleteParams) WithID(id string) *JStackTemplateDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the j stack template delete params
func (o *JStackTemplateDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *JStackTemplateDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
