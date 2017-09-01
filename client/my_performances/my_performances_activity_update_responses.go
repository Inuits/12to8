// Code generated by go-swagger; DO NOT EDIT.

package my_performances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// MyPerformancesActivityUpdateReader is a Reader for the MyPerformancesActivityUpdate structure.
type MyPerformancesActivityUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MyPerformancesActivityUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewMyPerformancesActivityUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMyPerformancesActivityUpdateOK creates a MyPerformancesActivityUpdateOK with default headers values
func NewMyPerformancesActivityUpdateOK() *MyPerformancesActivityUpdateOK {
	return &MyPerformancesActivityUpdateOK{}
}

/*MyPerformancesActivityUpdateOK handles this case with default header values.

MyPerformancesActivityUpdateOK my performances activity update o k
*/
type MyPerformancesActivityUpdateOK struct {
}

func (o *MyPerformancesActivityUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/my_performances/activity/{id}/][%d] myPerformancesActivityUpdateOK ", 200)
}

func (o *MyPerformancesActivityUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*MyPerformancesActivityUpdateBody my performances activity update body
swagger:model MyPerformancesActivityUpdateBody
*/

type MyPerformancesActivityUpdateBody struct {

	// contract
	// Required: true
	Contract *string `json:"contract"`

	// contract role
	ContractRole string `json:"contract_role,omitempty"`

	// day
	// Required: true
	Day *int64 `json:"day"`

	// description
	Description string `json:"description,omitempty"`

	// duration
	Duration float64 `json:"duration,omitempty"`

	// performance type
	// Required: true
	PerformanceType *string `json:"performance_type"`

	// redmine id
	RedmineID string `json:"redmine_id,omitempty"`

	// timesheet
	// Required: true
	Timesheet *string `json:"timesheet"`
}

/* polymorph MyPerformancesActivityUpdateBody contract false */

/* polymorph MyPerformancesActivityUpdateBody contract_role false */

/* polymorph MyPerformancesActivityUpdateBody day false */

/* polymorph MyPerformancesActivityUpdateBody description false */

/* polymorph MyPerformancesActivityUpdateBody duration false */

/* polymorph MyPerformancesActivityUpdateBody performance_type false */

/* polymorph MyPerformancesActivityUpdateBody redmine_id false */

/* polymorph MyPerformancesActivityUpdateBody timesheet false */

// MarshalBinary interface implementation
func (o *MyPerformancesActivityUpdateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *MyPerformancesActivityUpdateBody) UnmarshalBinary(b []byte) error {
	var res MyPerformancesActivityUpdateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
