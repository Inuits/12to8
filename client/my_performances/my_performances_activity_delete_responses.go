// Code generated by go-swagger; DO NOT EDIT.

package my_performances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// MyPerformancesActivityDeleteReader is a Reader for the MyPerformancesActivityDelete structure.
type MyPerformancesActivityDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MyPerformancesActivityDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewMyPerformancesActivityDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMyPerformancesActivityDeleteNoContent creates a MyPerformancesActivityDeleteNoContent with default headers values
func NewMyPerformancesActivityDeleteNoContent() *MyPerformancesActivityDeleteNoContent {
	return &MyPerformancesActivityDeleteNoContent{}
}

/*MyPerformancesActivityDeleteNoContent handles this case with default header values.

MyPerformancesActivityDeleteNoContent my performances activity delete no content
*/
type MyPerformancesActivityDeleteNoContent struct {
}

func (o *MyPerformancesActivityDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/my_performances/activity/{id}/][%d] myPerformancesActivityDeleteNoContent ", 204)
}

func (o *MyPerformancesActivityDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}