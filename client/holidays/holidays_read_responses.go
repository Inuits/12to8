// Code generated by go-swagger; DO NOT EDIT.

package holidays

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// HolidaysReadReader is a Reader for the HolidaysRead structure.
type HolidaysReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HolidaysReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewHolidaysReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHolidaysReadOK creates a HolidaysReadOK with default headers values
func NewHolidaysReadOK() *HolidaysReadOK {
	return &HolidaysReadOK{}
}

/*HolidaysReadOK handles this case with default header values.

HolidaysReadOK holidays read o k
*/
type HolidaysReadOK struct {
}

func (o *HolidaysReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/holidays/{id}/][%d] holidaysReadOK ", 200)
}

func (o *HolidaysReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
