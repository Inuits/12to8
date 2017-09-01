// Code generated by go-swagger; DO NOT EDIT.

package contracts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ContractsReadReader is a Reader for the ContractsRead structure.
type ContractsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContractsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewContractsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewContractsReadOK creates a ContractsReadOK with default headers values
func NewContractsReadOK() *ContractsReadOK {
	return &ContractsReadOK{}
}

/*ContractsReadOK handles this case with default header values.

ContractsReadOK contracts read o k
*/
type ContractsReadOK struct {
}

func (o *ContractsReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/contracts/{id}/][%d] contractsReadOK ", 200)
}

func (o *ContractsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}