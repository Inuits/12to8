// Code generated by go-swagger; DO NOT EDIT.

package companies

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

// NewCompaniesListParams creates a new CompaniesListParams object
// with the default values initialized.
func NewCompaniesListParams() *CompaniesListParams {
	var ()
	return &CompaniesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCompaniesListParamsWithTimeout creates a new CompaniesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCompaniesListParamsWithTimeout(timeout time.Duration) *CompaniesListParams {
	var ()
	return &CompaniesListParams{

		timeout: timeout,
	}
}

// NewCompaniesListParamsWithContext creates a new CompaniesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewCompaniesListParamsWithContext(ctx context.Context) *CompaniesListParams {
	var ()
	return &CompaniesListParams{

		Context: ctx,
	}
}

// NewCompaniesListParamsWithHTTPClient creates a new CompaniesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCompaniesListParamsWithHTTPClient(client *http.Client) *CompaniesListParams {
	var ()
	return &CompaniesListParams{
		HTTPClient: client,
	}
}

/*CompaniesListParams contains all the parameters to send to the API endpoint
for the companies list operation typically these are written to a http.Request
*/
type CompaniesListParams struct {

	/*Address*/
	Address *string
	/*AddressContains*/
	AddressContains *string
	/*AddressIcontains*/
	AddressIcontains *string
	/*Country*/
	Country *string
	/*Internal*/
	Internal *string
	/*Name*/
	Name *string
	/*NameContains*/
	NameContains *string
	/*NameIcontains*/
	NameIcontains *string
	/*OrderBy*/
	OrderBy *string
	/*Page*/
	Page *string
	/*PageSize*/
	PageSize *string
	/*VatIdentificationNumber*/
	VatIdentificationNumber *string
	/*VatIdentificationNumberContains*/
	VatIdentificationNumberContains *string
	/*VatIdentificationNumberIcontains*/
	VatIdentificationNumberIcontains *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the companies list params
func (o *CompaniesListParams) WithTimeout(timeout time.Duration) *CompaniesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the companies list params
func (o *CompaniesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the companies list params
func (o *CompaniesListParams) WithContext(ctx context.Context) *CompaniesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the companies list params
func (o *CompaniesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the companies list params
func (o *CompaniesListParams) WithHTTPClient(client *http.Client) *CompaniesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the companies list params
func (o *CompaniesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddress adds the address to the companies list params
func (o *CompaniesListParams) WithAddress(address *string) *CompaniesListParams {
	o.SetAddress(address)
	return o
}

// SetAddress adds the address to the companies list params
func (o *CompaniesListParams) SetAddress(address *string) {
	o.Address = address
}

// WithAddressContains adds the addressContains to the companies list params
func (o *CompaniesListParams) WithAddressContains(addressContains *string) *CompaniesListParams {
	o.SetAddressContains(addressContains)
	return o
}

// SetAddressContains adds the addressContains to the companies list params
func (o *CompaniesListParams) SetAddressContains(addressContains *string) {
	o.AddressContains = addressContains
}

// WithAddressIcontains adds the addressIcontains to the companies list params
func (o *CompaniesListParams) WithAddressIcontains(addressIcontains *string) *CompaniesListParams {
	o.SetAddressIcontains(addressIcontains)
	return o
}

// SetAddressIcontains adds the addressIcontains to the companies list params
func (o *CompaniesListParams) SetAddressIcontains(addressIcontains *string) {
	o.AddressIcontains = addressIcontains
}

// WithCountry adds the country to the companies list params
func (o *CompaniesListParams) WithCountry(country *string) *CompaniesListParams {
	o.SetCountry(country)
	return o
}

// SetCountry adds the country to the companies list params
func (o *CompaniesListParams) SetCountry(country *string) {
	o.Country = country
}

// WithInternal adds the internal to the companies list params
func (o *CompaniesListParams) WithInternal(internal *string) *CompaniesListParams {
	o.SetInternal(internal)
	return o
}

// SetInternal adds the internal to the companies list params
func (o *CompaniesListParams) SetInternal(internal *string) {
	o.Internal = internal
}

// WithName adds the name to the companies list params
func (o *CompaniesListParams) WithName(name *string) *CompaniesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the companies list params
func (o *CompaniesListParams) SetName(name *string) {
	o.Name = name
}

// WithNameContains adds the nameContains to the companies list params
func (o *CompaniesListParams) WithNameContains(nameContains *string) *CompaniesListParams {
	o.SetNameContains(nameContains)
	return o
}

// SetNameContains adds the nameContains to the companies list params
func (o *CompaniesListParams) SetNameContains(nameContains *string) {
	o.NameContains = nameContains
}

// WithNameIcontains adds the nameIcontains to the companies list params
func (o *CompaniesListParams) WithNameIcontains(nameIcontains *string) *CompaniesListParams {
	o.SetNameIcontains(nameIcontains)
	return o
}

// SetNameIcontains adds the nameIcontains to the companies list params
func (o *CompaniesListParams) SetNameIcontains(nameIcontains *string) {
	o.NameIcontains = nameIcontains
}

// WithOrderBy adds the orderBy to the companies list params
func (o *CompaniesListParams) WithOrderBy(orderBy *string) *CompaniesListParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the companies list params
func (o *CompaniesListParams) SetOrderBy(orderBy *string) {
	o.OrderBy = orderBy
}

// WithPage adds the page to the companies list params
func (o *CompaniesListParams) WithPage(page *string) *CompaniesListParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the companies list params
func (o *CompaniesListParams) SetPage(page *string) {
	o.Page = page
}

// WithPageSize adds the pageSize to the companies list params
func (o *CompaniesListParams) WithPageSize(pageSize *string) *CompaniesListParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the companies list params
func (o *CompaniesListParams) SetPageSize(pageSize *string) {
	o.PageSize = pageSize
}

// WithVatIdentificationNumber adds the vatIdentificationNumber to the companies list params
func (o *CompaniesListParams) WithVatIdentificationNumber(vatIdentificationNumber *string) *CompaniesListParams {
	o.SetVatIdentificationNumber(vatIdentificationNumber)
	return o
}

// SetVatIdentificationNumber adds the vatIdentificationNumber to the companies list params
func (o *CompaniesListParams) SetVatIdentificationNumber(vatIdentificationNumber *string) {
	o.VatIdentificationNumber = vatIdentificationNumber
}

// WithVatIdentificationNumberContains adds the vatIdentificationNumberContains to the companies list params
func (o *CompaniesListParams) WithVatIdentificationNumberContains(vatIdentificationNumberContains *string) *CompaniesListParams {
	o.SetVatIdentificationNumberContains(vatIdentificationNumberContains)
	return o
}

// SetVatIdentificationNumberContains adds the vatIdentificationNumberContains to the companies list params
func (o *CompaniesListParams) SetVatIdentificationNumberContains(vatIdentificationNumberContains *string) {
	o.VatIdentificationNumberContains = vatIdentificationNumberContains
}

// WithVatIdentificationNumberIcontains adds the vatIdentificationNumberIcontains to the companies list params
func (o *CompaniesListParams) WithVatIdentificationNumberIcontains(vatIdentificationNumberIcontains *string) *CompaniesListParams {
	o.SetVatIdentificationNumberIcontains(vatIdentificationNumberIcontains)
	return o
}

// SetVatIdentificationNumberIcontains adds the vatIdentificationNumberIcontains to the companies list params
func (o *CompaniesListParams) SetVatIdentificationNumberIcontains(vatIdentificationNumberIcontains *string) {
	o.VatIdentificationNumberIcontains = vatIdentificationNumberIcontains
}

// WriteToRequest writes these params to a swagger request
func (o *CompaniesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Address != nil {

		// query param address
		var qrAddress string
		if o.Address != nil {
			qrAddress = *o.Address
		}
		qAddress := qrAddress
		if qAddress != "" {
			if err := r.SetQueryParam("address", qAddress); err != nil {
				return err
			}
		}

	}

	if o.AddressContains != nil {

		// query param address__contains
		var qrAddressContains string
		if o.AddressContains != nil {
			qrAddressContains = *o.AddressContains
		}
		qAddressContains := qrAddressContains
		if qAddressContains != "" {
			if err := r.SetQueryParam("address__contains", qAddressContains); err != nil {
				return err
			}
		}

	}

	if o.AddressIcontains != nil {

		// query param address__icontains
		var qrAddressIcontains string
		if o.AddressIcontains != nil {
			qrAddressIcontains = *o.AddressIcontains
		}
		qAddressIcontains := qrAddressIcontains
		if qAddressIcontains != "" {
			if err := r.SetQueryParam("address__icontains", qAddressIcontains); err != nil {
				return err
			}
		}

	}

	if o.Country != nil {

		// query param country
		var qrCountry string
		if o.Country != nil {
			qrCountry = *o.Country
		}
		qCountry := qrCountry
		if qCountry != "" {
			if err := r.SetQueryParam("country", qCountry); err != nil {
				return err
			}
		}

	}

	if o.Internal != nil {

		// query param internal
		var qrInternal string
		if o.Internal != nil {
			qrInternal = *o.Internal
		}
		qInternal := qrInternal
		if qInternal != "" {
			if err := r.SetQueryParam("internal", qInternal); err != nil {
				return err
			}
		}

	}

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.NameContains != nil {

		// query param name__contains
		var qrNameContains string
		if o.NameContains != nil {
			qrNameContains = *o.NameContains
		}
		qNameContains := qrNameContains
		if qNameContains != "" {
			if err := r.SetQueryParam("name__contains", qNameContains); err != nil {
				return err
			}
		}

	}

	if o.NameIcontains != nil {

		// query param name__icontains
		var qrNameIcontains string
		if o.NameIcontains != nil {
			qrNameIcontains = *o.NameIcontains
		}
		qNameIcontains := qrNameIcontains
		if qNameIcontains != "" {
			if err := r.SetQueryParam("name__icontains", qNameIcontains); err != nil {
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

	if o.VatIdentificationNumber != nil {

		// query param vat_identification_number
		var qrVatIdentificationNumber string
		if o.VatIdentificationNumber != nil {
			qrVatIdentificationNumber = *o.VatIdentificationNumber
		}
		qVatIdentificationNumber := qrVatIdentificationNumber
		if qVatIdentificationNumber != "" {
			if err := r.SetQueryParam("vat_identification_number", qVatIdentificationNumber); err != nil {
				return err
			}
		}

	}

	if o.VatIdentificationNumberContains != nil {

		// query param vat_identification_number__contains
		var qrVatIdentificationNumberContains string
		if o.VatIdentificationNumberContains != nil {
			qrVatIdentificationNumberContains = *o.VatIdentificationNumberContains
		}
		qVatIdentificationNumberContains := qrVatIdentificationNumberContains
		if qVatIdentificationNumberContains != "" {
			if err := r.SetQueryParam("vat_identification_number__contains", qVatIdentificationNumberContains); err != nil {
				return err
			}
		}

	}

	if o.VatIdentificationNumberIcontains != nil {

		// query param vat_identification_number__icontains
		var qrVatIdentificationNumberIcontains string
		if o.VatIdentificationNumberIcontains != nil {
			qrVatIdentificationNumberIcontains = *o.VatIdentificationNumberIcontains
		}
		qVatIdentificationNumberIcontains := qrVatIdentificationNumberIcontains
		if qVatIdentificationNumberIcontains != "" {
			if err := r.SetQueryParam("vat_identification_number__icontains", qVatIdentificationNumberIcontains); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}