// Code generated by go-swagger; DO NOT EDIT.

package timesheets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// TimesheetsReadReader is a Reader for the TimesheetsRead structure.
type TimesheetsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TimesheetsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTimesheetsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTimesheetsReadOK creates a TimesheetsReadOK with default headers values
func NewTimesheetsReadOK() *TimesheetsReadOK {
	return &TimesheetsReadOK{}
}

/*TimesheetsReadOK handles this case with default header values.

TimesheetsReadOK timesheets read o k
*/
type TimesheetsReadOK struct {
}

func (o *TimesheetsReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/timesheets/{id}/][%d] timesheetsReadOK ", 200)
}

func (o *TimesheetsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
