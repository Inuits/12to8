// Code generated by go-swagger; DO NOT EDIT.

package user_relatives

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UserRelativesReadReader is a Reader for the UserRelativesRead structure.
type UserRelativesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserRelativesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUserRelativesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserRelativesReadOK creates a UserRelativesReadOK with default headers values
func NewUserRelativesReadOK() *UserRelativesReadOK {
	return &UserRelativesReadOK{}
}

/*UserRelativesReadOK handles this case with default header values.

UserRelativesReadOK user relatives read o k
*/
type UserRelativesReadOK struct {
}

func (o *UserRelativesReadOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/user_relatives/{id}/][%d] userRelativesReadOK ", 200)
}

func (o *UserRelativesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
