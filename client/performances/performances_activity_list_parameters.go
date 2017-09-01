// Code generated by go-swagger; DO NOT EDIT.

package performances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPerformancesActivityListParams creates a new PerformancesActivityListParams object
// with the default values initialized.
func NewPerformancesActivityListParams() *PerformancesActivityListParams {
	var ()
	return &PerformancesActivityListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPerformancesActivityListParamsWithTimeout creates a new PerformancesActivityListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPerformancesActivityListParamsWithTimeout(timeout time.Duration) *PerformancesActivityListParams {
	var ()
	return &PerformancesActivityListParams{

		timeout: timeout,
	}
}

// NewPerformancesActivityListParamsWithContext creates a new PerformancesActivityListParams object
// with the default values initialized, and the ability to set a context for a request
func NewPerformancesActivityListParamsWithContext(ctx context.Context) *PerformancesActivityListParams {
	var ()
	return &PerformancesActivityListParams{

		Context: ctx,
	}
}

// NewPerformancesActivityListParamsWithHTTPClient creates a new PerformancesActivityListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPerformancesActivityListParamsWithHTTPClient(client *http.Client) *PerformancesActivityListParams {
	var ()
	return &PerformancesActivityListParams{
		HTTPClient: client,
	}
}

/*PerformancesActivityListParams contains all the parameters to send to the API endpoint
for the performances activity list operation typically these are written to a http.Request
*/
type PerformancesActivityListParams struct {

	/*Contract*/
	Contract *string
	/*Day*/
	Day *string
	/*DayGt*/
	DayGt *string
	/*DayGte*/
	DayGte *string
	/*DayLt*/
	DayLt *string
	/*DayLte*/
	DayLte *string
	/*Description*/
	Description *string
	/*DescriptionContains*/
	DescriptionContains *string
	/*DescriptionIcontains*/
	DescriptionIcontains *string
	/*Duration*/
	Duration *string
	/*DurationGt*/
	DurationGt *string
	/*DurationGte*/
	DurationGte *string
	/*DurationLt*/
	DurationLt *string
	/*DurationLte*/
	DurationLte *string
	/*OrderBy*/
	OrderBy *string
	/*Page*/
	Page *string
	/*PageSize*/
	PageSize *string
	/*Timesheet*/
	Timesheet *string
	/*TimesheetMonth*/
	TimesheetMonth *string
	/*TimesheetMonthGte*/
	TimesheetMonthGte *string
	/*TimesheetMonthLte*/
	TimesheetMonthLte *string
	/*TimesheetUserID*/
	TimesheetUserID *string
	/*TimesheetYear*/
	TimesheetYear *string
	/*TimesheetYearGte*/
	TimesheetYearGte *string
	/*TimesheetYearLte*/
	TimesheetYearLte *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the performances activity list params
func (o *PerformancesActivityListParams) WithTimeout(timeout time.Duration) *PerformancesActivityListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the performances activity list params
func (o *PerformancesActivityListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the performances activity list params
func (o *PerformancesActivityListParams) WithContext(ctx context.Context) *PerformancesActivityListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the performances activity list params
func (o *PerformancesActivityListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the performances activity list params
func (o *PerformancesActivityListParams) WithHTTPClient(client *http.Client) *PerformancesActivityListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the performances activity list params
func (o *PerformancesActivityListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContract adds the contract to the performances activity list params
func (o *PerformancesActivityListParams) WithContract(contract *string) *PerformancesActivityListParams {
	o.SetContract(contract)
	return o
}

// SetContract adds the contract to the performances activity list params
func (o *PerformancesActivityListParams) SetContract(contract *string) {
	o.Contract = contract
}

// WithDay adds the day to the performances activity list params
func (o *PerformancesActivityListParams) WithDay(day *string) *PerformancesActivityListParams {
	o.SetDay(day)
	return o
}

// SetDay adds the day to the performances activity list params
func (o *PerformancesActivityListParams) SetDay(day *string) {
	o.Day = day
}

// WithDayGt adds the dayGt to the performances activity list params
func (o *PerformancesActivityListParams) WithDayGt(dayGt *string) *PerformancesActivityListParams {
	o.SetDayGt(dayGt)
	return o
}

// SetDayGt adds the dayGt to the performances activity list params
func (o *PerformancesActivityListParams) SetDayGt(dayGt *string) {
	o.DayGt = dayGt
}

// WithDayGte adds the dayGte to the performances activity list params
func (o *PerformancesActivityListParams) WithDayGte(dayGte *string) *PerformancesActivityListParams {
	o.SetDayGte(dayGte)
	return o
}

// SetDayGte adds the dayGte to the performances activity list params
func (o *PerformancesActivityListParams) SetDayGte(dayGte *string) {
	o.DayGte = dayGte
}

// WithDayLt adds the dayLt to the performances activity list params
func (o *PerformancesActivityListParams) WithDayLt(dayLt *string) *PerformancesActivityListParams {
	o.SetDayLt(dayLt)
	return o
}

// SetDayLt adds the dayLt to the performances activity list params
func (o *PerformancesActivityListParams) SetDayLt(dayLt *string) {
	o.DayLt = dayLt
}

// WithDayLte adds the dayLte to the performances activity list params
func (o *PerformancesActivityListParams) WithDayLte(dayLte *string) *PerformancesActivityListParams {
	o.SetDayLte(dayLte)
	return o
}

// SetDayLte adds the dayLte to the performances activity list params
func (o *PerformancesActivityListParams) SetDayLte(dayLte *string) {
	o.DayLte = dayLte
}

// WithDescription adds the description to the performances activity list params
func (o *PerformancesActivityListParams) WithDescription(description *string) *PerformancesActivityListParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the performances activity list params
func (o *PerformancesActivityListParams) SetDescription(description *string) {
	o.Description = description
}

// WithDescriptionContains adds the descriptionContains to the performances activity list params
func (o *PerformancesActivityListParams) WithDescriptionContains(descriptionContains *string) *PerformancesActivityListParams {
	o.SetDescriptionContains(descriptionContains)
	return o
}

// SetDescriptionContains adds the descriptionContains to the performances activity list params
func (o *PerformancesActivityListParams) SetDescriptionContains(descriptionContains *string) {
	o.DescriptionContains = descriptionContains
}

// WithDescriptionIcontains adds the descriptionIcontains to the performances activity list params
func (o *PerformancesActivityListParams) WithDescriptionIcontains(descriptionIcontains *string) *PerformancesActivityListParams {
	o.SetDescriptionIcontains(descriptionIcontains)
	return o
}

// SetDescriptionIcontains adds the descriptionIcontains to the performances activity list params
func (o *PerformancesActivityListParams) SetDescriptionIcontains(descriptionIcontains *string) {
	o.DescriptionIcontains = descriptionIcontains
}

// WithDuration adds the duration to the performances activity list params
func (o *PerformancesActivityListParams) WithDuration(duration *string) *PerformancesActivityListParams {
	o.SetDuration(duration)
	return o
}

// SetDuration adds the duration to the performances activity list params
func (o *PerformancesActivityListParams) SetDuration(duration *string) {
	o.Duration = duration
}

// WithDurationGt adds the durationGt to the performances activity list params
func (o *PerformancesActivityListParams) WithDurationGt(durationGt *string) *PerformancesActivityListParams {
	o.SetDurationGt(durationGt)
	return o
}

// SetDurationGt adds the durationGt to the performances activity list params
func (o *PerformancesActivityListParams) SetDurationGt(durationGt *string) {
	o.DurationGt = durationGt
}

// WithDurationGte adds the durationGte to the performances activity list params
func (o *PerformancesActivityListParams) WithDurationGte(durationGte *string) *PerformancesActivityListParams {
	o.SetDurationGte(durationGte)
	return o
}

// SetDurationGte adds the durationGte to the performances activity list params
func (o *PerformancesActivityListParams) SetDurationGte(durationGte *string) {
	o.DurationGte = durationGte
}

// WithDurationLt adds the durationLt to the performances activity list params
func (o *PerformancesActivityListParams) WithDurationLt(durationLt *string) *PerformancesActivityListParams {
	o.SetDurationLt(durationLt)
	return o
}

// SetDurationLt adds the durationLt to the performances activity list params
func (o *PerformancesActivityListParams) SetDurationLt(durationLt *string) {
	o.DurationLt = durationLt
}

// WithDurationLte adds the durationLte to the performances activity list params
func (o *PerformancesActivityListParams) WithDurationLte(durationLte *string) *PerformancesActivityListParams {
	o.SetDurationLte(durationLte)
	return o
}

// SetDurationLte adds the durationLte to the performances activity list params
func (o *PerformancesActivityListParams) SetDurationLte(durationLte *string) {
	o.DurationLte = durationLte
}

// WithOrderBy adds the orderBy to the performances activity list params
func (o *PerformancesActivityListParams) WithOrderBy(orderBy *string) *PerformancesActivityListParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the performances activity list params
func (o *PerformancesActivityListParams) SetOrderBy(orderBy *string) {
	o.OrderBy = orderBy
}

// WithPage adds the page to the performances activity list params
func (o *PerformancesActivityListParams) WithPage(page *string) *PerformancesActivityListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the performances activity list params
func (o *PerformancesActivityListParams) SetPage(page *string) {
	o.Page = page
}

// WithPageSize adds the pageSize to the performances activity list params
func (o *PerformancesActivityListParams) WithPageSize(pageSize *string) *PerformancesActivityListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the performances activity list params
func (o *PerformancesActivityListParams) SetPageSize(pageSize *string) {
	o.PageSize = pageSize
}

// WithTimesheet adds the timesheet to the performances activity list params
func (o *PerformancesActivityListParams) WithTimesheet(timesheet *string) *PerformancesActivityListParams {
	o.SetTimesheet(timesheet)
	return o
}

// SetTimesheet adds the timesheet to the performances activity list params
func (o *PerformancesActivityListParams) SetTimesheet(timesheet *string) {
	o.Timesheet = timesheet
}

// WithTimesheetMonth adds the timesheetMonth to the performances activity list params
func (o *PerformancesActivityListParams) WithTimesheetMonth(timesheetMonth *string) *PerformancesActivityListParams {
	o.SetTimesheetMonth(timesheetMonth)
	return o
}

// SetTimesheetMonth adds the timesheetMonth to the performances activity list params
func (o *PerformancesActivityListParams) SetTimesheetMonth(timesheetMonth *string) {
	o.TimesheetMonth = timesheetMonth
}

// WithTimesheetMonthGte adds the timesheetMonthGte to the performances activity list params
func (o *PerformancesActivityListParams) WithTimesheetMonthGte(timesheetMonthGte *string) *PerformancesActivityListParams {
	o.SetTimesheetMonthGte(timesheetMonthGte)
	return o
}

// SetTimesheetMonthGte adds the timesheetMonthGte to the performances activity list params
func (o *PerformancesActivityListParams) SetTimesheetMonthGte(timesheetMonthGte *string) {
	o.TimesheetMonthGte = timesheetMonthGte
}

// WithTimesheetMonthLte adds the timesheetMonthLte to the performances activity list params
func (o *PerformancesActivityListParams) WithTimesheetMonthLte(timesheetMonthLte *string) *PerformancesActivityListParams {
	o.SetTimesheetMonthLte(timesheetMonthLte)
	return o
}

// SetTimesheetMonthLte adds the timesheetMonthLte to the performances activity list params
func (o *PerformancesActivityListParams) SetTimesheetMonthLte(timesheetMonthLte *string) {
	o.TimesheetMonthLte = timesheetMonthLte
}

// WithTimesheetUserID adds the timesheetUserID to the performances activity list params
func (o *PerformancesActivityListParams) WithTimesheetUserID(timesheetUserID *string) *PerformancesActivityListParams {
	o.SetTimesheetUserID(timesheetUserID)
	return o
}

// SetTimesheetUserID adds the timesheetUserId to the performances activity list params
func (o *PerformancesActivityListParams) SetTimesheetUserID(timesheetUserID *string) {
	o.TimesheetUserID = timesheetUserID
}

// WithTimesheetYear adds the timesheetYear to the performances activity list params
func (o *PerformancesActivityListParams) WithTimesheetYear(timesheetYear *string) *PerformancesActivityListParams {
	o.SetTimesheetYear(timesheetYear)
	return o
}

// SetTimesheetYear adds the timesheetYear to the performances activity list params
func (o *PerformancesActivityListParams) SetTimesheetYear(timesheetYear *string) {
	o.TimesheetYear = timesheetYear
}

// WithTimesheetYearGte adds the timesheetYearGte to the performances activity list params
func (o *PerformancesActivityListParams) WithTimesheetYearGte(timesheetYearGte *string) *PerformancesActivityListParams {
	o.SetTimesheetYearGte(timesheetYearGte)
	return o
}

// SetTimesheetYearGte adds the timesheetYearGte to the performances activity list params
func (o *PerformancesActivityListParams) SetTimesheetYearGte(timesheetYearGte *string) {
	o.TimesheetYearGte = timesheetYearGte
}

// WithTimesheetYearLte adds the timesheetYearLte to the performances activity list params
func (o *PerformancesActivityListParams) WithTimesheetYearLte(timesheetYearLte *string) *PerformancesActivityListParams {
	o.SetTimesheetYearLte(timesheetYearLte)
	return o
}

// SetTimesheetYearLte adds the timesheetYearLte to the performances activity list params
func (o *PerformancesActivityListParams) SetTimesheetYearLte(timesheetYearLte *string) {
	o.TimesheetYearLte = timesheetYearLte
}

// WriteToRequest writes these params to a swagger request
func (o *PerformancesActivityListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Contract != nil {

		// query param contract
		var qrContract string
		if o.Contract != nil {
			qrContract = *o.Contract
		}
		qContract := qrContract
		if qContract != "" {
			if err := r.SetQueryParam("contract", qContract); err != nil {
				return err
			}
		}

	}

	if o.Day != nil {

		// query param day
		var qrDay string
		if o.Day != nil {
			qrDay = *o.Day
		}
		qDay := qrDay
		if qDay != "" {
			if err := r.SetQueryParam("day", qDay); err != nil {
				return err
			}
		}

	}

	if o.DayGt != nil {

		// query param day__gt
		var qrDayGt string
		if o.DayGt != nil {
			qrDayGt = *o.DayGt
		}
		qDayGt := qrDayGt
		if qDayGt != "" {
			if err := r.SetQueryParam("day__gt", qDayGt); err != nil {
				return err
			}
		}

	}

	if o.DayGte != nil {

		// query param day__gte
		var qrDayGte string
		if o.DayGte != nil {
			qrDayGte = *o.DayGte
		}
		qDayGte := qrDayGte
		if qDayGte != "" {
			if err := r.SetQueryParam("day__gte", qDayGte); err != nil {
				return err
			}
		}

	}

	if o.DayLt != nil {

		// query param day__lt
		var qrDayLt string
		if o.DayLt != nil {
			qrDayLt = *o.DayLt
		}
		qDayLt := qrDayLt
		if qDayLt != "" {
			if err := r.SetQueryParam("day__lt", qDayLt); err != nil {
				return err
			}
		}

	}

	if o.DayLte != nil {

		// query param day__lte
		var qrDayLte string
		if o.DayLte != nil {
			qrDayLte = *o.DayLte
		}
		qDayLte := qrDayLte
		if qDayLte != "" {
			if err := r.SetQueryParam("day__lte", qDayLte); err != nil {
				return err
			}
		}

	}

	if o.Description != nil {

		// query param description
		var qrDescription string
		if o.Description != nil {
			qrDescription = *o.Description
		}
		qDescription := qrDescription
		if qDescription != "" {
			if err := r.SetQueryParam("description", qDescription); err != nil {
				return err
			}
		}

	}

	if o.DescriptionContains != nil {

		// query param description__contains
		var qrDescriptionContains string
		if o.DescriptionContains != nil {
			qrDescriptionContains = *o.DescriptionContains
		}
		qDescriptionContains := qrDescriptionContains
		if qDescriptionContains != "" {
			if err := r.SetQueryParam("description__contains", qDescriptionContains); err != nil {
				return err
			}
		}

	}

	if o.DescriptionIcontains != nil {

		// query param description__icontains
		var qrDescriptionIcontains string
		if o.DescriptionIcontains != nil {
			qrDescriptionIcontains = *o.DescriptionIcontains
		}
		qDescriptionIcontains := qrDescriptionIcontains
		if qDescriptionIcontains != "" {
			if err := r.SetQueryParam("description__icontains", qDescriptionIcontains); err != nil {
				return err
			}
		}

	}

	if o.Duration != nil {

		// query param duration
		var qrDuration string
		if o.Duration != nil {
			qrDuration = *o.Duration
		}
		qDuration := qrDuration
		if qDuration != "" {
			if err := r.SetQueryParam("duration", qDuration); err != nil {
				return err
			}
		}

	}

	if o.DurationGt != nil {

		// query param duration__gt
		var qrDurationGt string
		if o.DurationGt != nil {
			qrDurationGt = *o.DurationGt
		}
		qDurationGt := qrDurationGt
		if qDurationGt != "" {
			if err := r.SetQueryParam("duration__gt", qDurationGt); err != nil {
				return err
			}
		}

	}

	if o.DurationGte != nil {

		// query param duration__gte
		var qrDurationGte string
		if o.DurationGte != nil {
			qrDurationGte = *o.DurationGte
		}
		qDurationGte := qrDurationGte
		if qDurationGte != "" {
			if err := r.SetQueryParam("duration__gte", qDurationGte); err != nil {
				return err
			}
		}

	}

	if o.DurationLt != nil {

		// query param duration__lt
		var qrDurationLt string
		if o.DurationLt != nil {
			qrDurationLt = *o.DurationLt
		}
		qDurationLt := qrDurationLt
		if qDurationLt != "" {
			if err := r.SetQueryParam("duration__lt", qDurationLt); err != nil {
				return err
			}
		}

	}

	if o.DurationLte != nil {

		// query param duration__lte
		var qrDurationLte string
		if o.DurationLte != nil {
			qrDurationLte = *o.DurationLte
		}
		qDurationLte := qrDurationLte
		if qDurationLte != "" {
			if err := r.SetQueryParam("duration__lte", qDurationLte); err != nil {
				return err
			}
		}

	}

	if o.OrderBy != nil {

		// query param order_by
		var qrOrderBy string
		if o.OrderBy != nil {
			qrOrderBy = *o.OrderBy
		}
		qOrderBy := qrOrderBy
		if qOrderBy != "" {
			if err := r.SetQueryParam("order_by", qOrderBy); err != nil {
				return err
			}
		}

	}

	if o.Page != nil {

		// query param page
		var qrPage string
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := qrPage
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.PageSize != nil {

		// query param page_size
		var qrPageSize string
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := qrPageSize
		if qPageSize != "" {
			if err := r.SetQueryParam("page_size", qPageSize); err != nil {
				return err
			}
		}

	}

	if o.Timesheet != nil {

		// query param timesheet
		var qrTimesheet string
		if o.Timesheet != nil {
			qrTimesheet = *o.Timesheet
		}
		qTimesheet := qrTimesheet
		if qTimesheet != "" {
			if err := r.SetQueryParam("timesheet", qTimesheet); err != nil {
				return err
			}
		}

	}

	if o.TimesheetMonth != nil {

		// query param timesheet__month
		var qrTimesheetMonth string
		if o.TimesheetMonth != nil {
			qrTimesheetMonth = *o.TimesheetMonth
		}
		qTimesheetMonth := qrTimesheetMonth
		if qTimesheetMonth != "" {
			if err := r.SetQueryParam("timesheet__month", qTimesheetMonth); err != nil {
				return err
			}
		}

	}

	if o.TimesheetMonthGte != nil {

		// query param timesheet__month__gte
		var qrTimesheetMonthGte string
		if o.TimesheetMonthGte != nil {
			qrTimesheetMonthGte = *o.TimesheetMonthGte
		}
		qTimesheetMonthGte := qrTimesheetMonthGte
		if qTimesheetMonthGte != "" {
			if err := r.SetQueryParam("timesheet__month__gte", qTimesheetMonthGte); err != nil {
				return err
			}
		}

	}

	if o.TimesheetMonthLte != nil {

		// query param timesheet__month__lte
		var qrTimesheetMonthLte string
		if o.TimesheetMonthLte != nil {
			qrTimesheetMonthLte = *o.TimesheetMonthLte
		}
		qTimesheetMonthLte := qrTimesheetMonthLte
		if qTimesheetMonthLte != "" {
			if err := r.SetQueryParam("timesheet__month__lte", qTimesheetMonthLte); err != nil {
				return err
			}
		}

	}

	if o.TimesheetUserID != nil {

		// query param timesheet__user_id
		var qrTimesheetUserID string
		if o.TimesheetUserID != nil {
			qrTimesheetUserID = *o.TimesheetUserID
		}
		qTimesheetUserID := qrTimesheetUserID
		if qTimesheetUserID != "" {
			if err := r.SetQueryParam("timesheet__user_id", qTimesheetUserID); err != nil {
				return err
			}
		}

	}

	if o.TimesheetYear != nil {

		// query param timesheet__year
		var qrTimesheetYear string
		if o.TimesheetYear != nil {
			qrTimesheetYear = *o.TimesheetYear
		}
		qTimesheetYear := qrTimesheetYear
		if qTimesheetYear != "" {
			if err := r.SetQueryParam("timesheet__year", qTimesheetYear); err != nil {
				return err
			}
		}

	}

	if o.TimesheetYearGte != nil {

		// query param timesheet__year__gte
		var qrTimesheetYearGte string
		if o.TimesheetYearGte != nil {
			qrTimesheetYearGte = *o.TimesheetYearGte
		}
		qTimesheetYearGte := qrTimesheetYearGte
		if qTimesheetYearGte != "" {
			if err := r.SetQueryParam("timesheet__year__gte", qTimesheetYearGte); err != nil {
				return err
			}
		}

	}

	if o.TimesheetYearLte != nil {

		// query param timesheet__year__lte
		var qrTimesheetYearLte string
		if o.TimesheetYearLte != nil {
			qrTimesheetYearLte = *o.TimesheetYearLte
		}
		qTimesheetYearLte := qrTimesheetYearLte
		if qTimesheetYearLte != "" {
			if err := r.SetQueryParam("timesheet__year__lte", qTimesheetYearLte); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}