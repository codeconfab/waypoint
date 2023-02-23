// Code generated by go-swagger; DO NOT EDIT.

package waypoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/waypoint/pkg/client/gen/models"
)

// NewWaypointRunPipelineParams creates a new WaypointRunPipelineParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointRunPipelineParams() *WaypointRunPipelineParams {
	return &WaypointRunPipelineParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointRunPipelineParamsWithTimeout creates a new WaypointRunPipelineParams object
// with the ability to set a timeout on a request.
func NewWaypointRunPipelineParamsWithTimeout(timeout time.Duration) *WaypointRunPipelineParams {
	return &WaypointRunPipelineParams{
		timeout: timeout,
	}
}

// NewWaypointRunPipelineParamsWithContext creates a new WaypointRunPipelineParams object
// with the ability to set a context for a request.
func NewWaypointRunPipelineParamsWithContext(ctx context.Context) *WaypointRunPipelineParams {
	return &WaypointRunPipelineParams{
		Context: ctx,
	}
}

// NewWaypointRunPipelineParamsWithHTTPClient creates a new WaypointRunPipelineParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointRunPipelineParamsWithHTTPClient(client *http.Client) *WaypointRunPipelineParams {
	return &WaypointRunPipelineParams{
		HTTPClient: client,
	}
}

/*
WaypointRunPipelineParams contains all the parameters to send to the API endpoint

	for the waypoint run pipeline operation.

	Typically these are written to a http.Request.
*/
type WaypointRunPipelineParams struct {

	// Body.
	Body *models.HashicorpWaypointRunPipelineRequest

	/* PipelineOwnerPipelineName.

	   the name of the defined pipeline config
	*/
	PipelineOwnerPipelineName string

	// PipelineOwnerProjectProject.
	PipelineOwnerProjectProject string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint run pipeline params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointRunPipelineParams) WithDefaults() *WaypointRunPipelineParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint run pipeline params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointRunPipelineParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint run pipeline params
func (o *WaypointRunPipelineParams) WithTimeout(timeout time.Duration) *WaypointRunPipelineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint run pipeline params
func (o *WaypointRunPipelineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint run pipeline params
func (o *WaypointRunPipelineParams) WithContext(ctx context.Context) *WaypointRunPipelineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint run pipeline params
func (o *WaypointRunPipelineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint run pipeline params
func (o *WaypointRunPipelineParams) WithHTTPClient(client *http.Client) *WaypointRunPipelineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint run pipeline params
func (o *WaypointRunPipelineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the waypoint run pipeline params
func (o *WaypointRunPipelineParams) WithBody(body *models.HashicorpWaypointRunPipelineRequest) *WaypointRunPipelineParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint run pipeline params
func (o *WaypointRunPipelineParams) SetBody(body *models.HashicorpWaypointRunPipelineRequest) {
	o.Body = body
}

// WithPipelineOwnerPipelineName adds the pipelineOwnerPipelineName to the waypoint run pipeline params
func (o *WaypointRunPipelineParams) WithPipelineOwnerPipelineName(pipelineOwnerPipelineName string) *WaypointRunPipelineParams {
	o.SetPipelineOwnerPipelineName(pipelineOwnerPipelineName)
	return o
}

// SetPipelineOwnerPipelineName adds the pipelineOwnerPipelineName to the waypoint run pipeline params
func (o *WaypointRunPipelineParams) SetPipelineOwnerPipelineName(pipelineOwnerPipelineName string) {
	o.PipelineOwnerPipelineName = pipelineOwnerPipelineName
}

// WithPipelineOwnerProjectProject adds the pipelineOwnerProjectProject to the waypoint run pipeline params
func (o *WaypointRunPipelineParams) WithPipelineOwnerProjectProject(pipelineOwnerProjectProject string) *WaypointRunPipelineParams {
	o.SetPipelineOwnerProjectProject(pipelineOwnerProjectProject)
	return o
}

// SetPipelineOwnerProjectProject adds the pipelineOwnerProjectProject to the waypoint run pipeline params
func (o *WaypointRunPipelineParams) SetPipelineOwnerProjectProject(pipelineOwnerProjectProject string) {
	o.PipelineOwnerProjectProject = pipelineOwnerProjectProject
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointRunPipelineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param pipeline.owner.pipeline_name
	if err := r.SetPathParam("pipeline.owner.pipeline_name", o.PipelineOwnerPipelineName); err != nil {
		return err
	}

	// path param pipeline.owner.project.project
	if err := r.SetPathParam("pipeline.owner.project.project", o.PipelineOwnerProjectProject); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}