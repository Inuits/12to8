// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/inuits/12to8/client/attachments"
	"github.com/inuits/12to8/client/companies"
	"github.com/inuits/12to8/client/contract_groups"
	"github.com/inuits/12to8/client/contract_roles"
	"github.com/inuits/12to8/client/contract_users"
	"github.com/inuits/12to8/client/contracts"
	"github.com/inuits/12to8/client/employment_contract_types"
	"github.com/inuits/12to8/client/employment_contracts"
	"github.com/inuits/12to8/client/groups"
	"github.com/inuits/12to8/client/holidays"
	"github.com/inuits/12to8/client/leave_dates"
	"github.com/inuits/12to8/client/leave_types"
	"github.com/inuits/12to8/client/leaves"
	"github.com/inuits/12to8/client/my_contracts"
	"github.com/inuits/12to8/client/my_leave_dates"
	"github.com/inuits/12to8/client/my_leaves"
	"github.com/inuits/12to8/client/my_performances"
	"github.com/inuits/12to8/client/my_timesheets"
	"github.com/inuits/12to8/client/my_workschedules"
	"github.com/inuits/12to8/client/performance_types"
	"github.com/inuits/12to8/client/performances"
	"github.com/inuits/12to8/client/project_estimates"
	"github.com/inuits/12to8/client/services"
	"github.com/inuits/12to8/client/timesheets"
	"github.com/inuits/12to8/client/user_infos"
	"github.com/inuits/12to8/client/user_relatives"
	"github.com/inuits/12to8/client/users"
	"github.com/inuits/12to8/client/whereabouts"
	"github.com/inuits/12to8/client/work_schedules"
)

// Default ninetofiver HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new ninetofiver HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Ninetofiver {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new ninetofiver HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Ninetofiver {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new ninetofiver client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Ninetofiver {
	cli := new(Ninetofiver)
	cli.Transport = transport

	cli.Attachments = attachments.New(transport, formats)

	cli.Companies = companies.New(transport, formats)

	cli.ContractGroups = contract_groups.New(transport, formats)

	cli.ContractRoles = contract_roles.New(transport, formats)

	cli.ContractUsers = contract_users.New(transport, formats)

	cli.Contracts = contracts.New(transport, formats)

	cli.EmploymentContractTypes = employment_contract_types.New(transport, formats)

	cli.EmploymentContracts = employment_contracts.New(transport, formats)

	cli.Groups = groups.New(transport, formats)

	cli.Holidays = holidays.New(transport, formats)

	cli.LeaveDates = leave_dates.New(transport, formats)

	cli.LeaveTypes = leave_types.New(transport, formats)

	cli.Leaves = leaves.New(transport, formats)

	cli.MyContracts = my_contracts.New(transport, formats)

	cli.MyLeaveDates = my_leave_dates.New(transport, formats)

	cli.MyLeaves = my_leaves.New(transport, formats)

	cli.MyPerformances = my_performances.New(transport, formats)

	cli.MyTimesheets = my_timesheets.New(transport, formats)

	cli.MyWorkschedules = my_workschedules.New(transport, formats)

	cli.PerformanceTypes = performance_types.New(transport, formats)

	cli.Performances = performances.New(transport, formats)

	cli.ProjectEstimates = project_estimates.New(transport, formats)

	cli.Services = services.New(transport, formats)

	cli.Timesheets = timesheets.New(transport, formats)

	cli.UserInfos = user_infos.New(transport, formats)

	cli.UserRelatives = user_relatives.New(transport, formats)

	cli.Users = users.New(transport, formats)

	cli.Whereabouts = whereabouts.New(transport, formats)

	cli.WorkSchedules = work_schedules.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Ninetofiver is a client for ninetofiver
type Ninetofiver struct {
	Attachments *attachments.Client

	Companies *companies.Client

	ContractGroups *contract_groups.Client

	ContractRoles *contract_roles.Client

	ContractUsers *contract_users.Client

	Contracts *contracts.Client

	EmploymentContractTypes *employment_contract_types.Client

	EmploymentContracts *employment_contracts.Client

	Groups *groups.Client

	Holidays *holidays.Client

	LeaveDates *leave_dates.Client

	LeaveTypes *leave_types.Client

	Leaves *leaves.Client

	MyContracts *my_contracts.Client

	MyLeaveDates *my_leave_dates.Client

	MyLeaves *my_leaves.Client

	MyPerformances *my_performances.Client

	MyTimesheets *my_timesheets.Client

	MyWorkschedules *my_workschedules.Client

	PerformanceTypes *performance_types.Client

	Performances *performances.Client

	ProjectEstimates *project_estimates.Client

	Services *services.Client

	Timesheets *timesheets.Client

	UserInfos *user_infos.Client

	UserRelatives *user_relatives.Client

	Users *users.Client

	Whereabouts *whereabouts.Client

	WorkSchedules *work_schedules.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Ninetofiver) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Attachments.SetTransport(transport)

	c.Companies.SetTransport(transport)

	c.ContractGroups.SetTransport(transport)

	c.ContractRoles.SetTransport(transport)

	c.ContractUsers.SetTransport(transport)

	c.Contracts.SetTransport(transport)

	c.EmploymentContractTypes.SetTransport(transport)

	c.EmploymentContracts.SetTransport(transport)

	c.Groups.SetTransport(transport)

	c.Holidays.SetTransport(transport)

	c.LeaveDates.SetTransport(transport)

	c.LeaveTypes.SetTransport(transport)

	c.Leaves.SetTransport(transport)

	c.MyContracts.SetTransport(transport)

	c.MyLeaveDates.SetTransport(transport)

	c.MyLeaves.SetTransport(transport)

	c.MyPerformances.SetTransport(transport)

	c.MyTimesheets.SetTransport(transport)

	c.MyWorkschedules.SetTransport(transport)

	c.PerformanceTypes.SetTransport(transport)

	c.Performances.SetTransport(transport)

	c.ProjectEstimates.SetTransport(transport)

	c.Services.SetTransport(transport)

	c.Timesheets.SetTransport(transport)

	c.UserInfos.SetTransport(transport)

	c.UserRelatives.SetTransport(transport)

	c.Users.SetTransport(transport)

	c.Whereabouts.SetTransport(transport)

	c.WorkSchedules.SetTransport(transport)

}
