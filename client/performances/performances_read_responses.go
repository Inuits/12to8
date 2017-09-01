// Code generated by go-swagger; DO NOT EDIT.

package performances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PerformancesReadReader is a Reader for the PerformancesRead structure.
type PerformancesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PerformancesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPerformancesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPerformancesReadOK creates a PerformancesReadOK with default headers values
func NewPerformancesReadOK() *PerformancesReadOK {
	return &PerformancesReadOK{}
}

/*PerformancesReadOK handles this case with default header values.

PerformancesReadOK performances read o k
*/
type PerformancesReadOK struct {
}

func (o *PerformancesReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/performances/{id}/][%d] performancesReadOK ", 200)
}

func (o *PerformancesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}