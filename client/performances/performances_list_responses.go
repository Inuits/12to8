// Code generated by go-swagger; DO NOT EDIT.

package performances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PerformancesListReader is a Reader for the PerformancesList structure.
type PerformancesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PerformancesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPerformancesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPerformancesListOK creates a PerformancesListOK with default headers values
func NewPerformancesListOK() *PerformancesListOK {
	return &PerformancesListOK{}
}

/*PerformancesListOK handles this case with default header values.

PerformancesListOK performances list o k
*/
type PerformancesListOK struct {
}

func (o *PerformancesListOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/performances/][%d] performancesListOK ", 200)
}

func (o *PerformancesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}