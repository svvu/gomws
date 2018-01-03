package reports

// Reference http://docs.developer.amazonservices.com/en_US/reports/Reports_Overview.html

import (
	"github.com/svvu/gomws/mws"
)

// Reports is the client for the api
type Reports struct {
	*mws.Client
}

// NewClient generate a new product client
func NewClient(config mws.Config) (*Reports, error) {
	report := new(Reports)
	base, err := mws.NewClient(config, report.Version(), report.Name())
	if err != nil {
		return nil, err
	}
	report.Client = base
	return report, nil
}

// Version return the current version of api
func (r Reports) Version() string {
	return "2009-01-01"
}

// Name return the name of the api
func (r Reports) Name() string {
	return "Reports"
}

// RequestReport Creates a report request and submits the request to Amazon MWS.
// Optional Parameters:
// 	StartDate - string. The start of a date range used for selecting the data to report. Values in ISO 8601 date time format.
//  EndDate - string. The end of a date range used for selecting the data to report. Values in ISO 8601 date time format.
//  ReportOptions - string. Additional information to pass to the report.
//  MarketplaceIdList - []string. A list of one or more marketplace IDs for the marketplaces you are registered to sell in.
// http://docs.developer.amazonservices.com/en_US/reports/Reports_RequestReport.html
func (r Reports) RequestReport(reportType string, optional ...mws.Parameters) (*mws.Response, error) {
	op := mws.OptionalParams([]string{
		"StartDate", "EndDate", "ReportOptions", "MarketplaceIdList",
	}, optional)
	params := mws.Parameters{
		"Action":     "RequestReport",
		"ReportType": reportType,
	}.Merge(op)

	structuredParams := params.StructureKeys("MarketplaceIdList", "Id")

	return r.SendRequest(structuredParams)
}

// GetReportRequestList Returns a list of report requests that you can use to get the ReportRequestId for a report.
// Optional Parameters:
// 	ReportRequestIdList - []string. A structured list of ReportRequestId values. If you pass in ReportRequestId values, other query conditions are ignored.
//  ReportTypeList - []string. A structured list of ReportType enumeration values.
//  ReportProcessingStatusList - []string. A structured list of report processing statuses by which to filter report requests.
//      Values: _SUBMITTED_, _IN_PROGRESS_, _CANCELLED_, _DONE_, _DONE_NO_DATA_.
// 		Default: All.
//  MaxCount - int. A non-negative integer that represents the maximum number of report requests to return. Max: 100.
//  RequestedFromDate: string. The start of the date range used for selecting the data to report, in ISO 8601 date time format .
//  RequestedToDate: string. The end of the date range used for selecting the data to report, in ISO 8601 date time format .
// http://docs.developer.amazonservices.com/en_US/reports/Reports_GetReportRequestList.html
func (r Reports) GetReportRequestList(optional ...mws.Parameters) (*mws.Response, error) {
	op := mws.OptionalParams([]string{
		"ReportRequestIdList", "ReportTypeList", "ReportProcessingStatusList",
		"MaxCount", "RequestedFromDate", "RequestedToDate",
	}, optional)
	params := mws.Parameters{"Action": "GetReportRequestList"}.Merge(op)
	structuredParams := params.StructureKeys("ReportRequestIdList", "Id").
		StructureKeys("ReportTypeList", "Type").
		StructureKeys("ReportProcessingStatusList", "Status")

	return r.SendRequest(structuredParams)
}

// GetReportRequestListByNextToken Returns a list of report requests using the NextToken, which was supplied by a previous request to either GetReportRequestListByNextToken or GetReportRequestList, where the value of HasNext was true in that previous request.
// http://docs.developer.amazonservices.com/en_US/reports/Reports_GetReportRequestListByNextToken.html
func (r Reports) GetReportRequestListByNextToken(nextToken string) (*mws.Response, error) {
	params := mws.Parameters{
		"Action":    "GetReportRequestListByNextToken",
		"NextToken": nextToken,
	}

	return r.SendRequest(params)
}

// GetReportRequestCount Returns a count of report requests that have been submitted to Amazon MWS for processing.
// Optional Parameters:
//  ReportTypeList - []string. A structured list of ReportType enumeration values.
//  ReportProcessingStatusList - []string. A structured list of report processing statuses by which to filter report requests.
//      Values: _SUBMITTED_, _IN_PROGRESS_, _CANCELLED_, _DONE_, _DONE_NO_DATA_.
// 		Default: All.
//  RequestedFromDate: string. The start of the date range used for selecting the data to report, in ISO 8601 date time format .
//  RequestedToDate: string. The end of the date range used for selecting the data to report, in ISO 8601 date time format .
// http://docs.developer.amazonservices.com/en_US/reports/Reports_GetReportRequestCount.html
func (r Reports) GetReportRequestCount(optional ...mws.Parameters) (*mws.Response, error) {
	op := mws.OptionalParams([]string{
		"RequestedFromDate", "RequestedToDate", "ReportTypeList", "ReportProcessingStatusList",
	}, optional)
	params := mws.Parameters{"Action": "GetReportRequestCount"}.Merge(op)
	structuredParams := params.StructureKeys("ReportTypeList", "Type").
		StructureKeys("ReportProcessingStatusList", "Status")

	return r.SendRequest(structuredParams)
}

// CancelReportRequests Cancels one or more report requests.
// Optional Parameters:
// 	ReportRequestIdList - []string. A structured list of ReportRequestId values. If you pass in ReportRequestId values, other query conditions are ignored.
//  ReportTypeList - []string. A structured list of ReportType enumeration values.
//  ReportProcessingStatusList - []string. A structured list of report processing statuses by which to filter report requests.
//      Values: _SUBMITTED_, _IN_PROGRESS_, _CANCELLED_, _DONE_, _DONE_NO_DATA_.
// 		Default: All.
//  RequestedFromDate: string. The start of the date range used for selecting the data to report, in ISO 8601 date time format .
//  RequestedToDate: string. The end of the date range used for selecting the data to report, in ISO 8601 date time format .
// http://docs.developer.amazonservices.com/en_US/reports/Reports_CancelReportRequests.html
func (r Reports) CancelReportRequests(optional ...mws.Parameters) (*mws.Response, error) {
	op := mws.OptionalParams([]string{
		"ReportRequestIdList", "ReportTypeList", "ReportProcessingStatusList",
		"RequestedFromDate", "RequestedToDate",
	}, optional)
	params := mws.Parameters{"Action": "CancelReportRequests"}.Merge(op)
	structuredParams := params.StructureKeys("ReportRequestIdList", "Id").
		StructureKeys("ReportTypeList", "Type").
		StructureKeys("ReportProcessingStatusList", "Status")

	return r.SendRequest(structuredParams)
}

// GetReportList Returns a list of reports that were created in the previous 90 days.
// Optional Parameters:
//  MaxCount - int. A non-negative integer that represents the maximum number of report requests to return. Max: 100.
//  ReportTypeList - []string. A structured list of ReportType enumeration values.
// 	Acknowledged - bool. A Boolean value that indicates if an order report has been acknowledged by a prior call to UpdateReportAcknowledgements.
//  AvailableFromDate: string. The start of the date range used for selecting the data to report, in ISO 8601 date time format .
//  AvailableToDate: string. The end of the date range used for selecting the data to report, in ISO 8601 date time format .
// 	ReportRequestIdList - []string. A structured list of ReportRequestId values. If you pass in ReportRequestId values, other query conditions are ignored.
// http://docs.developer.amazonservices.com/en_US/reports/Reports_GetReportList.html
func (r Reports) GetReportList(optional ...mws.Parameters) (*mws.Response, error) {
	op := mws.OptionalParams([]string{
		"MaxCount", "ReportTypeList", "Acknowledged",
		"AvailableFromDate", "AvailableToDate", "ReportRequestIdList",
	}, optional)
	params := mws.Parameters{"Action": "GetReportList"}.Merge(op)
	structuredParams := params.StructureKeys("ReportRequestIdList", "Id").
		StructureKeys("ReportTypeList", "Type")

	return r.SendRequest(structuredParams)
}

// GetReportListByNextToken Returns a list of reports using the NextToken, which was supplied by a previous request to either GetReportListByNextToken or GetReportList, where the value of HasNext was true in the previous call.
// http://docs.developer.amazonservices.com/en_US/reports/Reports_GetReportListByNextToken.html
func (r Reports) GetReportListByNextToken(nextToken string) (*mws.Response, error) {
	params := mws.Parameters{
		"Action":    "GetReportListByNextToken",
		"NextToken": nextToken,
	}

	return r.SendRequest(params)
}

// GetReportCount Returns a count of the reports, created in the previous 90 days, with a status of _DONE_ and that are available for download.
// Optional Parameters:
//  ReportTypeList - []string. A structured list of ReportType enumeration values.
// 	Acknowledged - bool. A Boolean value that indicates if an order report has been acknowledged by a prior call to UpdateReportAcknowledgements.
//  AvailableFromDate: string. The start of the date range used for selecting the data to report, in ISO 8601 date time format .
//  AvailableToDate: string. The end of the date range used for selecting the data to report, in ISO 8601 date time format .
// http://docs.developer.amazonservices.com/en_US/reports/Reports_GetReportCount.html
func (r Reports) GetReportCount(optional ...mws.Parameters) (*mws.Response, error) {
	op := mws.OptionalParams([]string{
		"ReportTypeList", "Acknowledged",
		"AvailableFromDate", "AvailableToDate",
	}, optional)
	params := mws.Parameters{"Action": "GetReportCount"}.Merge(op)
	structuredParams := params.StructureKeys("ReportTypeList", "Type")

	return r.SendRequest(structuredParams)
}

// GetReport Returns the contents of a report and the Content-MD5 header for the returned report body.
// http://docs.developer.amazonservices.com/en_US/reports/Reports_GetReport.html
func (r Reports) GetReport(reportID string) (*mws.Response, error) {
	params := mws.Parameters{
		"Action":   "GetReport",
		"ReportId": reportID,
	}

	return r.SendRequest(params)
}

// ManageReportSchedule Creates, updates, or deletes a report request schedule for a specified report type.
// http://docs.developer.amazonservices.com/en_US/reports/Reports_ManageReportSchedule.html
// schedule - string. A value of the Schedule that indicates how often a report request should be created.
// 	Detail: http://docs.developer.amazonservices.com/en_US/reports/Reports_Schedule.html
// 	Values: _15_MINUTES_, _30_MINUTES_, _1_HOUR_, _2_HOURS_, _4_HOURS_, _8_HOURS_,
// 		_12_HOURS_, _1_DAY_, _2_DAYS_, _72_HOURS_, _1_WEEK_, _14_DAYS_,
// 		_15_DAYS_, _30_DAYS_, _NEVER_
// Optional Parameters:
//  ScheduleDate - string. The date when the next report request is scheduled to be submitted.
// 		Value can be no more than 366 days in the future. In ISO 8601 date time format .
func (r Reports) ManageReportSchedule(reportType string, schedule string, optional ...mws.Parameters) (*mws.Response, error) {
	op := mws.OptionalParams([]string{"ScheduleDate"}, optional)
	params := mws.Parameters{
		"Action":     "ManageReportSchedule",
		"ReportType": reportType,
		"Schedule":   schedule,
	}.Merge(op)

	return r.SendRequest(params)
}

// GetReportScheduleList Returns a list of order report requests that are scheduled to be submitted to Amazon MWS for processing.
// http://docs.developer.amazonservices.com/en_US/reports/Reports_GetReportScheduleList.html
func (r Reports) GetReportScheduleList(optional ...mws.Parameters) (*mws.Response, error) {
	op := mws.OptionalParams([]string{"ReportTypeList"}, optional)
	params := mws.Parameters{"Action": "GetReportScheduleList"}.Merge(op)
	structuredParams := params.StructureKeys("ReportTypeList", "Type")

	return r.SendRequest(structuredParams)
}

// GetReportScheduleCount Returns a count of order report requests that are scheduled to be submitted to Amazon MWS.
// http://docs.developer.amazonservices.com/en_US/reports/Reports_GetReportScheduleCount.html
func (r Reports) GetReportScheduleCount(optional ...mws.Parameters) (*mws.Response, error) {
	op := mws.OptionalParams([]string{"ReportTypeList"}, optional)
	params := mws.Parameters{"Action": "GetReportScheduleCount"}.Merge(op)
	structuredParams := params.StructureKeys("ReportTypeList", "Type")

	return r.SendRequest(structuredParams)
}

// UpdateReportAcknowledgements Updates the acknowledged status of one or more reports.
// http://docs.developer.amazonservices.com/en_US/reports/Reports_UpdateReportAcknowledgements.html
// Optional Parameters:
//  Acknowledged - string. A Boolean value that indicates that you have received and stored a report.
func (r Reports) UpdateReportAcknowledgements(ids []string, optional ...mws.Parameters) (*mws.Response, error) {
	op := mws.OptionalParams([]string{"Acknowledged"}, optional)
	params := mws.Parameters{
		"Action":       "UpdateReportAcknowledgements",
		"ReportIdList": ids,
	}.Merge(op)
	structuredParams := params.StructureKeys("ReportIdList", "Id")

	return r.SendRequest(structuredParams)
}
