// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/hydra/internal/httpclient/models"
)

// DynamicClientRegistrationCreateOAuth2ClientReader is a Reader for the DynamicClientRegistrationCreateOAuth2Client structure.
type DynamicClientRegistrationCreateOAuth2ClientReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DynamicClientRegistrationCreateOAuth2ClientReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDynamicClientRegistrationCreateOAuth2ClientCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDynamicClientRegistrationCreateOAuth2ClientDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDynamicClientRegistrationCreateOAuth2ClientCreated creates a DynamicClientRegistrationCreateOAuth2ClientCreated with default headers values
func NewDynamicClientRegistrationCreateOAuth2ClientCreated() *DynamicClientRegistrationCreateOAuth2ClientCreated {
	return &DynamicClientRegistrationCreateOAuth2ClientCreated{}
}

/*DynamicClientRegistrationCreateOAuth2ClientCreated handles this case with default header values.

oAuth2Client
*/
type DynamicClientRegistrationCreateOAuth2ClientCreated struct {
	Payload *models.OAuth2Client
}

func (o *DynamicClientRegistrationCreateOAuth2ClientCreated) Error() string {
	return fmt.Sprintf("[POST /connect/register][%d] dynamicClientRegistrationCreateOAuth2ClientCreated  %+v", 201, o.Payload)
}

func (o *DynamicClientRegistrationCreateOAuth2ClientCreated) GetPayload() *models.OAuth2Client {
	return o.Payload
}

func (o *DynamicClientRegistrationCreateOAuth2ClientCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OAuth2Client)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDynamicClientRegistrationCreateOAuth2ClientDefault creates a DynamicClientRegistrationCreateOAuth2ClientDefault with default headers values
func NewDynamicClientRegistrationCreateOAuth2ClientDefault(code int) *DynamicClientRegistrationCreateOAuth2ClientDefault {
	return &DynamicClientRegistrationCreateOAuth2ClientDefault{
		_statusCode: code,
	}
}

/*DynamicClientRegistrationCreateOAuth2ClientDefault handles this case with default header values.

jsonError
*/
type DynamicClientRegistrationCreateOAuth2ClientDefault struct {
	_statusCode int

	Payload *models.JSONError
}

// Code gets the status code for the dynamic client registration create o auth2 client default response
func (o *DynamicClientRegistrationCreateOAuth2ClientDefault) Code() int {
	return o._statusCode
}

func (o *DynamicClientRegistrationCreateOAuth2ClientDefault) Error() string {
	return fmt.Sprintf("[POST /connect/register][%d] dynamicClientRegistrationCreateOAuth2Client default  %+v", o._statusCode, o.Payload)
}

func (o *DynamicClientRegistrationCreateOAuth2ClientDefault) GetPayload() *models.JSONError {
	return o.Payload
}

func (o *DynamicClientRegistrationCreateOAuth2ClientDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
