package kv

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPutEntryParams creates a new PutEntryParams object
// with the default values initialized.
func NewPutEntryParams() *PutEntryParams {
	var ()
	return &PutEntryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutEntryParamsWithTimeout creates a new PutEntryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutEntryParamsWithTimeout(timeout time.Duration) *PutEntryParams {
	var ()
	return &PutEntryParams{

		timeout: timeout,
	}
}

// NewPutEntryParamsWithContext creates a new PutEntryParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutEntryParamsWithContext(ctx context.Context) *PutEntryParams {
	var ()
	return &PutEntryParams{

		Context: ctx,
	}
}

// NewPutEntryParamsWithHTTPClient creates a new PutEntryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutEntryParamsWithHTTPClient(client *http.Client) *PutEntryParams {
	var ()
	return &PutEntryParams{
		HTTPClient: client,
	}
}

/*PutEntryParams contains all the parameters to send to the API endpoint
for the put entry operation typically these are written to a http.Request
*/
type PutEntryParams struct {

	/*IfMatch
	  when this is an update to an entry, then this field needs to be present

	*/
	IfMatch *string
	/*XRequestID
	  A unique UUID for the request

	*/
	XRequestID *string
	/*Body*/
	Body io.ReadCloser
	/*Key
	  The key for a given entry

	*/
	Key string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put entry params
func (o *PutEntryParams) WithTimeout(timeout time.Duration) *PutEntryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put entry params
func (o *PutEntryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put entry params
func (o *PutEntryParams) WithContext(ctx context.Context) *PutEntryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put entry params
func (o *PutEntryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put entry params
func (o *PutEntryParams) WithHTTPClient(client *http.Client) *PutEntryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put entry params
func (o *PutEntryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfMatch adds the ifMatch to the put entry params
func (o *PutEntryParams) WithIfMatch(ifMatch *string) *PutEntryParams {
	o.SetIfMatch(ifMatch)
	return o
}

// SetIfMatch adds the ifMatch to the put entry params
func (o *PutEntryParams) SetIfMatch(ifMatch *string) {
	o.IfMatch = ifMatch
}

// WithXRequestID adds the xRequestID to the put entry params
func (o *PutEntryParams) WithXRequestID(xRequestID *string) *PutEntryParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the put entry params
func (o *PutEntryParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the put entry params
func (o *PutEntryParams) WithBody(body io.ReadCloser) *PutEntryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put entry params
func (o *PutEntryParams) SetBody(body io.ReadCloser) {
	o.Body = body
}

// WithKey adds the key to the put entry params
func (o *PutEntryParams) WithKey(key string) *PutEntryParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the put entry params
func (o *PutEntryParams) SetKey(key string) {
	o.Key = key
}

// WriteToRequest writes these params to a swagger request
func (o *PutEntryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IfMatch != nil {

		// header param If-Match
		if err := r.SetHeaderParam("If-Match", *o.IfMatch); err != nil {
			return err
		}

	}

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}

	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param key
	if err := r.SetPathParam("key", o.Key); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
