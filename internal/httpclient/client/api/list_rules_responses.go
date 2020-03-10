// Passcode generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/ory/oathkeeper/internal/httpclient/models"
)

// ListRulesReader is a Reader for the ListRules structure.
type ListRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewListRulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListRulesOK creates a ListRulesOK with default headers values
func NewListRulesOK() *ListRulesOK {
	return &ListRulesOK{}
}

/*ListRulesOK handles this case with default header values.

A list of rules
*/
type ListRulesOK struct {
	Payload []*models.Rule
}

func (o *ListRulesOK) Error() string {
	return fmt.Sprintf("[GET /rules][%d] listRulesOK  %+v", 200, o.Payload)
}

func (o *ListRulesOK) GetPayload() []*models.Rule {
	return o.Payload
}

func (o *ListRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRulesInternalServerError creates a ListRulesInternalServerError with default headers values
func NewListRulesInternalServerError() *ListRulesInternalServerError {
	return &ListRulesInternalServerError{}
}

/*ListRulesInternalServerError handles this case with default header values.

The standard error format
*/
type ListRulesInternalServerError struct {
	Payload *ListRulesInternalServerErrorBody
}

func (o *ListRulesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /rules][%d] listRulesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListRulesInternalServerError) GetPayload() *ListRulesInternalServerErrorBody {
	return o.Payload
}

func (o *ListRulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListRulesInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ListRulesInternalServerErrorBody list rules internal server error body
swagger:model ListRulesInternalServerErrorBody
*/
type ListRulesInternalServerErrorBody struct {

	// code
	Code int64 `json:"code,omitempty"`

	// details
	Details []map[string]interface{} `json:"details"`

	// message
	Message string `json:"message,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// request
	Request string `json:"request,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this list rules internal server error body
func (o *ListRulesInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListRulesInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListRulesInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res ListRulesInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
