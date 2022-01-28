// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_v_p_n_connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPcloudVpnconnectionsPeersubnetsGetParams creates a new PcloudVpnconnectionsPeersubnetsGetParams object
// with the default values initialized.
func NewPcloudVpnconnectionsPeersubnetsGetParams() *PcloudVpnconnectionsPeersubnetsGetParams {
	var ()
	return &PcloudVpnconnectionsPeersubnetsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudVpnconnectionsPeersubnetsGetParamsWithTimeout creates a new PcloudVpnconnectionsPeersubnetsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPcloudVpnconnectionsPeersubnetsGetParamsWithTimeout(timeout time.Duration) *PcloudVpnconnectionsPeersubnetsGetParams {
	var ()
	return &PcloudVpnconnectionsPeersubnetsGetParams{

		timeout: timeout,
	}
}

// NewPcloudVpnconnectionsPeersubnetsGetParamsWithContext creates a new PcloudVpnconnectionsPeersubnetsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewPcloudVpnconnectionsPeersubnetsGetParamsWithContext(ctx context.Context) *PcloudVpnconnectionsPeersubnetsGetParams {
	var ()
	return &PcloudVpnconnectionsPeersubnetsGetParams{

		Context: ctx,
	}
}

// NewPcloudVpnconnectionsPeersubnetsGetParamsWithHTTPClient creates a new PcloudVpnconnectionsPeersubnetsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPcloudVpnconnectionsPeersubnetsGetParamsWithHTTPClient(client *http.Client) *PcloudVpnconnectionsPeersubnetsGetParams {
	var ()
	return &PcloudVpnconnectionsPeersubnetsGetParams{
		HTTPClient: client,
	}
}

/*PcloudVpnconnectionsPeersubnetsGetParams contains all the parameters to send to the API endpoint
for the pcloud vpnconnections peersubnets get operation typically these are written to a http.Request
*/
type PcloudVpnconnectionsPeersubnetsGetParams struct {

	/*CloudInstanceID
	  Cloud Instance ID of a PCloud Instance

	*/
	CloudInstanceID string
	/*VpnConnectionID
	  ID of a VPN connection

	*/
	VpnConnectionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pcloud vpnconnections peersubnets get params
func (o *PcloudVpnconnectionsPeersubnetsGetParams) WithTimeout(timeout time.Duration) *PcloudVpnconnectionsPeersubnetsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud vpnconnections peersubnets get params
func (o *PcloudVpnconnectionsPeersubnetsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud vpnconnections peersubnets get params
func (o *PcloudVpnconnectionsPeersubnetsGetParams) WithContext(ctx context.Context) *PcloudVpnconnectionsPeersubnetsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud vpnconnections peersubnets get params
func (o *PcloudVpnconnectionsPeersubnetsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud vpnconnections peersubnets get params
func (o *PcloudVpnconnectionsPeersubnetsGetParams) WithHTTPClient(client *http.Client) *PcloudVpnconnectionsPeersubnetsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud vpnconnections peersubnets get params
func (o *PcloudVpnconnectionsPeersubnetsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud vpnconnections peersubnets get params
func (o *PcloudVpnconnectionsPeersubnetsGetParams) WithCloudInstanceID(cloudInstanceID string) *PcloudVpnconnectionsPeersubnetsGetParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud vpnconnections peersubnets get params
func (o *PcloudVpnconnectionsPeersubnetsGetParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithVpnConnectionID adds the vpnConnectionID to the pcloud vpnconnections peersubnets get params
func (o *PcloudVpnconnectionsPeersubnetsGetParams) WithVpnConnectionID(vpnConnectionID string) *PcloudVpnconnectionsPeersubnetsGetParams {
	o.SetVpnConnectionID(vpnConnectionID)
	return o
}

// SetVpnConnectionID adds the vpnConnectionId to the pcloud vpnconnections peersubnets get params
func (o *PcloudVpnconnectionsPeersubnetsGetParams) SetVpnConnectionID(vpnConnectionID string) {
	o.VpnConnectionID = vpnConnectionID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudVpnconnectionsPeersubnetsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	// path param vpn_connection_id
	if err := r.SetPathParam("vpn_connection_id", o.VpnConnectionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
