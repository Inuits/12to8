// Code generated by go-swagger; DO NOT EDIT.

package my_workschedules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// MyWorkschedulesCreateReader is a Reader for the MyWorkschedulesCreate structure.
type MyWorkschedulesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MyWorkschedulesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewMyWorkschedulesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMyWorkschedulesCreateCreated creates a MyWorkschedulesCreateCreated with default headers values
func NewMyWorkschedulesCreateCreated() *MyWorkschedulesCreateCreated {
	return &MyWorkschedulesCreateCreated{}
}

/*MyWorkschedulesCreateCreated handles this case with default header values.

MyWorkschedulesCreateCreated my workschedules create created
*/
type MyWorkschedulesCreateCreated struct {
}

func (o *MyWorkschedulesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v1/my_workschedules/][%d] myWorkschedulesCreateCreated ", 201)
}

func (o *MyWorkschedulesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*MyWorkschedulesCreateBody my workschedules create body
swagger:model MyWorkschedulesCreateBody
*/

type MyWorkschedulesCreateBody struct {

	// friday
	Friday float64 `json:"friday,omitempty"`

	// label
	// Required: true
	Label *string `json:"label"`

	// monday
	Monday float64 `json:"monday,omitempty"`

	// saturday
	Saturday float64 `json:"saturday,omitempty"`

	// sunday
	Sunday float64 `json:"sunday,omitempty"`

	// thursday
	Thursday float64 `json:"thursday,omitempty"`

	// tuesday
	Tuesday float64 `json:"tuesday,omitempty"`

	// wednesday
	Wednesday float64 `json:"wednesday,omitempty"`
}

/* polymorph MyWorkschedulesCreateBody friday false */

/* polymorph MyWorkschedulesCreateBody label false */

/* polymorph MyWorkschedulesCreateBody monday false */

/* polymorph MyWorkschedulesCreateBody saturday false */

/* polymorph MyWorkschedulesCreateBody sunday false */

/* polymorph MyWorkschedulesCreateBody thursday false */

/* polymorph MyWorkschedulesCreateBody tuesday false */

/* polymorph MyWorkschedulesCreateBody wednesday false */

// MarshalBinary interface implementation
func (o *MyWorkschedulesCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *MyWorkschedulesCreateBody) UnmarshalBinary(b []byte) error {
	var res MyWorkschedulesCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}