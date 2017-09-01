// Code generated by go-swagger; DO NOT EDIT.

package my_performances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// MyPerformancesStandbyDeleteReader is a Reader for the MyPerformancesStandbyDelete structure.
type MyPerformancesStandbyDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MyPerformancesStandbyDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewMyPerformancesStandbyDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMyPerformancesStandbyDeleteNoContent creates a MyPerformancesStandbyDeleteNoContent with default headers values
func NewMyPerformancesStandbyDeleteNoContent() *MyPerformancesStandbyDeleteNoContent {
	return &MyPerformancesStandbyDeleteNoContent{}
}

/*MyPerformancesStandbyDeleteNoContent handles this case with default header values.

MyPerformancesStandbyDeleteNoContent my performances standby delete no content
*/
type MyPerformancesStandbyDeleteNoContent struct {
}

func (o *MyPerformancesStandbyDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/my_performances/standby/{id}/][%d] myPerformancesStandbyDeleteNoContent ", 204)
}

func (o *MyPerformancesStandbyDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}