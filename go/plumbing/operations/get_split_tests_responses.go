// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netlify/open-api/go/models"
)

// GetSplitTestsReader is a Reader for the GetSplitTests structure.
type GetSplitTestsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSplitTestsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSplitTestsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetSplitTestsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSplitTestsOK creates a GetSplitTestsOK with default headers values
func NewGetSplitTestsOK() *GetSplitTestsOK {
	return &GetSplitTestsOK{}
}

/*GetSplitTestsOK handles this case with default header values.

split_tests
*/
type GetSplitTestsOK struct {
	Payload models.SplitTests
}

func (o *GetSplitTestsOK) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/traffic_splits][%d] getSplitTestsOK  %+v", 200, o.Payload)
}

func (o *GetSplitTestsOK) GetPayload() models.SplitTests {
	return o.Payload
}

func (o *GetSplitTestsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSplitTestsDefault creates a GetSplitTestsDefault with default headers values
func NewGetSplitTestsDefault(code int) *GetSplitTestsDefault {
	return &GetSplitTestsDefault{
		_statusCode: code,
	}
}

/*GetSplitTestsDefault handles this case with default header values.

error
*/
type GetSplitTestsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get split tests default response
func (o *GetSplitTestsDefault) Code() int {
	return o._statusCode
}

func (o *GetSplitTestsDefault) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/traffic_splits][%d] getSplitTests default  %+v", o._statusCode, o.Payload)
}

func (o *GetSplitTestsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSplitTestsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
