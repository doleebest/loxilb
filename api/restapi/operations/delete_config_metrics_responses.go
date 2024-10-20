// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/loxilb-io/loxilb/api/models"
)

// DeleteConfigMetricsNoContentCode is the HTTP code returned for type DeleteConfigMetricsNoContent
const DeleteConfigMetricsNoContentCode int = 204

/*
DeleteConfigMetricsNoContent OK

swagger:response deleteConfigMetricsNoContent
*/
type DeleteConfigMetricsNoContent struct {
}

// NewDeleteConfigMetricsNoContent creates DeleteConfigMetricsNoContent with default headers values
func NewDeleteConfigMetricsNoContent() *DeleteConfigMetricsNoContent {

	return &DeleteConfigMetricsNoContent{}
}

// WriteResponse to the client
func (o *DeleteConfigMetricsNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteConfigMetricsBadRequestCode is the HTTP code returned for type DeleteConfigMetricsBadRequest
const DeleteConfigMetricsBadRequestCode int = 400

/*
DeleteConfigMetricsBadRequest Malformed arguments for API call

swagger:response deleteConfigMetricsBadRequest
*/
type DeleteConfigMetricsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigMetricsBadRequest creates DeleteConfigMetricsBadRequest with default headers values
func NewDeleteConfigMetricsBadRequest() *DeleteConfigMetricsBadRequest {

	return &DeleteConfigMetricsBadRequest{}
}

// WithPayload adds the payload to the delete config metrics bad request response
func (o *DeleteConfigMetricsBadRequest) WithPayload(payload *models.Error) *DeleteConfigMetricsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config metrics bad request response
func (o *DeleteConfigMetricsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigMetricsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigMetricsUnauthorizedCode is the HTTP code returned for type DeleteConfigMetricsUnauthorized
const DeleteConfigMetricsUnauthorizedCode int = 401

/*
DeleteConfigMetricsUnauthorized Invalid authentication credentials

swagger:response deleteConfigMetricsUnauthorized
*/
type DeleteConfigMetricsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigMetricsUnauthorized creates DeleteConfigMetricsUnauthorized with default headers values
func NewDeleteConfigMetricsUnauthorized() *DeleteConfigMetricsUnauthorized {

	return &DeleteConfigMetricsUnauthorized{}
}

// WithPayload adds the payload to the delete config metrics unauthorized response
func (o *DeleteConfigMetricsUnauthorized) WithPayload(payload *models.Error) *DeleteConfigMetricsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config metrics unauthorized response
func (o *DeleteConfigMetricsUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigMetricsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigMetricsInternalServerErrorCode is the HTTP code returned for type DeleteConfigMetricsInternalServerError
const DeleteConfigMetricsInternalServerErrorCode int = 500

/*
DeleteConfigMetricsInternalServerError Internal service error

swagger:response deleteConfigMetricsInternalServerError
*/
type DeleteConfigMetricsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigMetricsInternalServerError creates DeleteConfigMetricsInternalServerError with default headers values
func NewDeleteConfigMetricsInternalServerError() *DeleteConfigMetricsInternalServerError {

	return &DeleteConfigMetricsInternalServerError{}
}

// WithPayload adds the payload to the delete config metrics internal server error response
func (o *DeleteConfigMetricsInternalServerError) WithPayload(payload *models.Error) *DeleteConfigMetricsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config metrics internal server error response
func (o *DeleteConfigMetricsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigMetricsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigMetricsServiceUnavailableCode is the HTTP code returned for type DeleteConfigMetricsServiceUnavailable
const DeleteConfigMetricsServiceUnavailableCode int = 503

/*
DeleteConfigMetricsServiceUnavailable Maintenance mode

swagger:response deleteConfigMetricsServiceUnavailable
*/
type DeleteConfigMetricsServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigMetricsServiceUnavailable creates DeleteConfigMetricsServiceUnavailable with default headers values
func NewDeleteConfigMetricsServiceUnavailable() *DeleteConfigMetricsServiceUnavailable {

	return &DeleteConfigMetricsServiceUnavailable{}
}

// WithPayload adds the payload to the delete config metrics service unavailable response
func (o *DeleteConfigMetricsServiceUnavailable) WithPayload(payload *models.Error) *DeleteConfigMetricsServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config metrics service unavailable response
func (o *DeleteConfigMetricsServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigMetricsServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
