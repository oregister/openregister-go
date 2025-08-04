// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openregister

import (
	"context"
	"net/http"
	"net/url"

	"github.com/oregister/openregister-go/internal/apijson"
	"github.com/oregister/openregister-go/internal/apiquery"
	"github.com/oregister/openregister-go/internal/requestconfig"
	"github.com/oregister/openregister-go/option"
	"github.com/oregister/openregister-go/packages/param"
	"github.com/oregister/openregister-go/packages/respjson"
)

// SearchService contains methods and other services that help with interacting
// with the openregister API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSearchService] method instead.
type SearchService struct {
	Options []option.RequestOption
}

// NewSearchService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSearchService(opts ...option.RequestOption) (r SearchService) {
	r = SearchService{}
	r.Options = opts
	return
}

// Autocomplete company search
func (r *SearchService) AutocompleteCompaniesV1(ctx context.Context, query SearchAutocompleteCompaniesV1Params, opts ...option.RequestOption) (res *SearchAutocompleteCompaniesV1Response, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/autocomplete/company"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Search for companies
func (r *SearchService) FindCompaniesV0(ctx context.Context, query SearchFindCompaniesV0Params, opts ...option.RequestOption) (res *CompanySearch, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/search/company"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Search for companies
func (r *SearchService) FindCompaniesV1(ctx context.Context, body SearchFindCompaniesV1Params, opts ...option.RequestOption) (res *CompanySearch, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/search/company"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Search for people
func (r *SearchService) FindPerson(ctx context.Context, body SearchFindPersonParams, opts ...option.RequestOption) (res *SearchFindPersonResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/search/person"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Find company by website URL
func (r *SearchService) LookupCompanyByURL(ctx context.Context, query SearchLookupCompanyByURLParams, opts ...option.RequestOption) (res *SearchLookupCompanyByURLResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/search/lookup"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Legal form of the company. Common German legal forms:
//
// - gmbh: Gesellschaft mit beschränkter Haftung (Limited Liability Company)
// - ag: Aktiengesellschaft (Stock Corporation)
// - ug: Unternehmergesellschaft (Entrepreneurial Company with limited liability)
// - ohg: Offene Handelsgesellschaft (General Partnership)
// - kg: Kommanditgesellschaft (Limited Partnership)
// - ev: Eingetragener Verein (Registered Association)
type CompanyLegalForm string

const (
	CompanyLegalFormAg        CompanyLegalForm = "ag"
	CompanyLegalFormEg        CompanyLegalForm = "eg"
	CompanyLegalFormEk        CompanyLegalForm = "ek"
	CompanyLegalFormEv        CompanyLegalForm = "ev"
	CompanyLegalFormEwiv      CompanyLegalForm = "ewiv"
	CompanyLegalFormForeign   CompanyLegalForm = "foreign"
	CompanyLegalFormGbr       CompanyLegalForm = "gbr"
	CompanyLegalFormGgmbh     CompanyLegalForm = "ggmbh"
	CompanyLegalFormGmbh      CompanyLegalForm = "gmbh"
	CompanyLegalFormKg        CompanyLegalForm = "kg"
	CompanyLegalFormKgaa      CompanyLegalForm = "kgaa"
	CompanyLegalFormUnknown   CompanyLegalForm = "unknown"
	CompanyLegalFormLlp       CompanyLegalForm = "llp"
	CompanyLegalFormMunicipal CompanyLegalForm = "municipal"
	CompanyLegalFormOhg       CompanyLegalForm = "ohg"
	CompanyLegalFormSe        CompanyLegalForm = "se"
	CompanyLegalFormUg        CompanyLegalForm = "ug"
)

// Type of company register where the entity is recorded. Common types:
//
// - HRB: Commercial Register B (limited liability companies, stock corporations)
// - HRA: Commercial Register A (partnerships, sole proprietorships)
// - PR: Partnership Register
// - GnR: Cooperative Register
// - VR: Association Register
type CompanyRegisterType string

const (
	CompanyRegisterTypeHrb CompanyRegisterType = "HRB"
	CompanyRegisterTypeHra CompanyRegisterType = "HRA"
	CompanyRegisterTypePr  CompanyRegisterType = "PR"
	CompanyRegisterTypeGnR CompanyRegisterType = "GnR"
	CompanyRegisterTypeVr  CompanyRegisterType = "VR"
)

type CompanySearch struct {
	Pagination CompanySearchPagination `json:"pagination,required"`
	// List of companies matching the search criteria.
	Results []CompanySearchResult `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pagination  respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanySearch) RawJSON() string { return r.JSON.raw }
func (r *CompanySearch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanySearchPagination struct {
	// Current page number.
	Page int64 `json:"page,required"`
	// Number of results per page.
	PerPage int64 `json:"per_page,required"`
	// Total number of pages.
	TotalPages int64 `json:"total_pages,required"`
	// Total number of results.
	TotalResults int64 `json:"total_results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Page         respjson.Field
		PerPage      respjson.Field
		TotalPages   respjson.Field
		TotalResults respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanySearchPagination) RawJSON() string { return r.JSON.raw }
func (r *CompanySearchPagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanySearchResult struct {
	// Company status - true if active, false if inactive.
	Active bool `json:"active,required"`
	// Unique company identifier. Example: DE-HRB-F1103-267645
	CompanyID string `json:"company_id,required"`
	// Legal form of the company. Example: "gmbh" for Gesellschaft mit beschränkter
	// Haftung
	//
	// Any of "ag", "eg", "ek", "ev", "ewiv", "foreign", "gbr", "ggmbh", "gmbh", "kg",
	// "kgaa", "unknown", "llp", "municipal", "ohg", "se", "ug".
	LegalForm CompanyLegalForm `json:"legal_form,required"`
	// Official registered company name. Example: "Max Mustermann GmbH"
	Name string `json:"name,required"`
	// Court where the company is registered. Example: "Berlin (Charlottenburg)"
	RegisterCourt string `json:"register_court,required"`
	// Registration number in the company register. Example: "230633"
	RegisterNumber string `json:"register_number,required"`
	// Type of company register. Example: "HRB" for Commercial Register B
	//
	// Any of "HRB", "HRA", "PR", "GnR", "VR".
	RegisterType CompanyRegisterType `json:"register_type,required"`
	// Country where the company is registered using ISO 3166-1 alpha-2 code. Example:
	// "DE" for Germany
	Country string `json:"country"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active         respjson.Field
		CompanyID      respjson.Field
		LegalForm      respjson.Field
		Name           respjson.Field
		RegisterCourt  respjson.Field
		RegisterNumber respjson.Field
		RegisterType   respjson.Field
		Country        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanySearchResult) RawJSON() string { return r.JSON.raw }
func (r *CompanySearchResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchAutocompleteCompaniesV1Response struct {
	// List of companies matching the search criteria.
	Results []SearchAutocompleteCompaniesV1ResponseResult `json:"results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchAutocompleteCompaniesV1Response) RawJSON() string { return r.JSON.raw }
func (r *SearchAutocompleteCompaniesV1Response) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchAutocompleteCompaniesV1ResponseResult struct {
	// Company status - true if active, false if inactive.
	Active bool `json:"active,required"`
	// Unique company identifier. Example: DE-HRB-F1103-267645
	CompanyID string `json:"company_id,required"`
	// Legal form of the company. Example: "gmbh" for Gesellschaft mit beschränkter
	// Haftung
	//
	// Any of "ag", "eg", "ek", "ev", "ewiv", "foreign", "gbr", "ggmbh", "gmbh", "kg",
	// "kgaa", "unknown", "llp", "municipal", "ohg", "se", "ug".
	LegalForm CompanyLegalForm `json:"legal_form,required"`
	// Official registered company name. Example: "Max Mustermann GmbH"
	Name string `json:"name,required"`
	// Court where the company is registered. Example: "Berlin (Charlottenburg)"
	RegisterCourt string `json:"register_court,required"`
	// Registration number in the company register. Example: "230633"
	RegisterNumber string `json:"register_number,required"`
	// Type of company register. Example: "HRB" for Commercial Register B
	//
	// Any of "HRB", "HRA", "PR", "GnR", "VR".
	RegisterType CompanyRegisterType `json:"register_type,required"`
	// Country where the company is registered using ISO 3166-1 alpha-2 code. Example:
	// "DE" for Germany
	Country string `json:"country"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active         respjson.Field
		CompanyID      respjson.Field
		LegalForm      respjson.Field
		Name           respjson.Field
		RegisterCourt  respjson.Field
		RegisterNumber respjson.Field
		RegisterType   respjson.Field
		Country        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchAutocompleteCompaniesV1ResponseResult) RawJSON() string { return r.JSON.raw }
func (r *SearchAutocompleteCompaniesV1ResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchFindPersonResponse struct {
	Pagination SearchFindPersonResponsePagination `json:"pagination,required"`
	// List of people matching the search criteria.
	Results []SearchFindPersonResponseResult `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pagination  respjson.Field
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchFindPersonResponse) RawJSON() string { return r.JSON.raw }
func (r *SearchFindPersonResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchFindPersonResponsePagination struct {
	// Current page number.
	Page int64 `json:"page,required"`
	// Number of results per page.
	PerPage int64 `json:"per_page,required"`
	// Total number of pages.
	TotalPages int64 `json:"total_pages,required"`
	// Total number of results.
	TotalResults int64 `json:"total_results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Page         respjson.Field
		PerPage      respjson.Field
		TotalPages   respjson.Field
		TotalResults respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchFindPersonResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *SearchFindPersonResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchFindPersonResponseResult struct {
	// Unique person identifier. Example: 1234-5678-9012-345678901234
	ID string `json:"id,required"`
	// Person status - true if active, false if inactive.
	Active bool `json:"active,required"`
	// Date of birth of the person. Format: ISO 8601 (YYYY-MM-DD) Example: "1990-01-01"
	DateOfBirth string `json:"date_of_birth,required"`
	// Name of the person. Example: "Max Mustermann"
	Name string `json:"name,required"`
	// City of the person. Example: "Berlin"
	City string `json:"city"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Active      respjson.Field
		DateOfBirth respjson.Field
		Name        respjson.Field
		City        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchFindPersonResponseResult) RawJSON() string { return r.JSON.raw }
func (r *SearchFindPersonResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchLookupCompanyByURLResponse struct {
	// Unique company identifier. Example: DE-HRB-F1103-267645
	CompanyID string `json:"company_id,required"`
	// Email address of the company. Example: "info@maxmustermann.de"
	Email string `json:"email"`
	// Phone number of the company. Example: "+49 123 456 789"
	Phone string `json:"phone"`
	// Value Added Tax identification number. Example: "DE123456789"
	VatID string `json:"vat_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompanyID   respjson.Field
		Email       respjson.Field
		Phone       respjson.Field
		VatID       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchLookupCompanyByURLResponse) RawJSON() string { return r.JSON.raw }
func (r *SearchLookupCompanyByURLResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchAutocompleteCompaniesV1Params struct {
	// Text search query to find companies by name. Example: "Descartes Technologies
	// UG"
	Query string `query:"query,required" json:"-"`
	paramObj
}

// URLQuery serializes [SearchAutocompleteCompaniesV1Params]'s query parameters as
// `url.Values`.
func (r SearchAutocompleteCompaniesV1Params) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SearchFindCompaniesV0Params struct {
	// Filter for active or inactive companies. Set to true for active companies only,
	// false for inactive only.
	Active param.Opt[bool] `query:"active,omitzero" json:"-"`
	// Date of incorporation of the company. Format: ISO 8601 (YYYY-MM-DD) Example:
	// "2022-01-01"
	IncorporationDate param.Opt[string] `query:"incorporation_date,omitzero" json:"-"`
	// Page number for pagination.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Number of results per page (max 50).
	PerPage param.Opt[int64] `query:"per_page,omitzero" json:"-"`
	// Text search query to find companies by name. Example: "Descartes Technologies
	// UG"
	Query param.Opt[string] `query:"query,omitzero" json:"-"`
	// Court where the company is registered. Example: "Berlin (Charlottenburg)"
	RegisterCourt param.Opt[string] `query:"register_court,omitzero" json:"-"`
	// Company register number for exact matching. Example: "230633"
	RegisterNumber param.Opt[string] `query:"register_number,omitzero" json:"-"`
	// Legal form of the company. Example: "gmbh" for "Gesellschaft mit beschränkter
	// Haftung"
	//
	// Any of "ag", "eg", "ek", "ev", "ewiv", "foreign", "gbr", "ggmbh", "gmbh", "kg",
	// "kgaa", "unknown", "llp", "municipal", "ohg", "se", "ug".
	LegalForm CompanyLegalForm `query:"legal_form,omitzero" json:"-"`
	// Type of register to filter results. Example: "HRB" (Commercial Register B)
	//
	// Any of "HRB", "HRA", "PR", "GnR", "VR".
	RegisterType CompanyRegisterType `query:"register_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SearchFindCompaniesV0Params]'s query parameters as
// `url.Values`.
func (r SearchFindCompaniesV0Params) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SearchFindCompaniesV1Params struct {
	// Filters to filter companies.
	Filters []SearchFindCompaniesV1ParamsFilter `json:"filters,omitzero"`
	// Location to filter companies.
	Location SearchFindCompaniesV1ParamsLocation `json:"location,omitzero"`
	// Pagination parameters.
	Pagination SearchFindCompaniesV1ParamsPagination `json:"pagination,omitzero"`
	// Search query to filter companies.
	Query SearchFindCompaniesV1ParamsQuery `json:"query,omitzero"`
	paramObj
}

func (r SearchFindCompaniesV1Params) MarshalJSON() (data []byte, err error) {
	type shadow SearchFindCompaniesV1Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SearchFindCompaniesV1Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter by field. The properties values, value, keywords and min/max are mutually
// exclusive. Dates must be in the format YYYY-MM-DD.
type SearchFindCompaniesV1ParamsFilter struct {
	// Maximum value to filter on.
	Max param.Opt[string] `json:"max,omitzero"`
	// Minimum value to filter on.
	Min param.Opt[string] `json:"min,omitzero"`
	// Value to filter on.
	Value param.Opt[string] `json:"value,omitzero"`
	// Any of "date_of_birth", "city", "active", "status", "legal_form",
	// "register_number", "register_court", "register_type", "incorporated_at", "zip",
	// "address", "balance_sheet_total", "revenue", "cash", "employees", "equity",
	// "real_estate", "materials", "pension_provisions", "salaries", "taxes",
	// "liabilities", "capital_reserves", "net_income", "industry_codes",
	// "capital_amount", "capital_currency".
	Field SearchFindCompaniesV1ParamsFilterField `json:"field,omitzero"`
	// Keywords to filter on.
	Keywords []string `json:"keywords,omitzero"`
	// Values to filter on.
	Values []string `json:"values,omitzero"`
	paramObj
}

func (r SearchFindCompaniesV1ParamsFilter) MarshalJSON() (data []byte, err error) {
	type shadow SearchFindCompaniesV1ParamsFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SearchFindCompaniesV1ParamsFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchFindCompaniesV1ParamsFilterField string

const (
	SearchFindCompaniesV1ParamsFilterFieldDateOfBirth       SearchFindCompaniesV1ParamsFilterField = "date_of_birth"
	SearchFindCompaniesV1ParamsFilterFieldCity              SearchFindCompaniesV1ParamsFilterField = "city"
	SearchFindCompaniesV1ParamsFilterFieldActive            SearchFindCompaniesV1ParamsFilterField = "active"
	SearchFindCompaniesV1ParamsFilterFieldStatus            SearchFindCompaniesV1ParamsFilterField = "status"
	SearchFindCompaniesV1ParamsFilterFieldLegalForm         SearchFindCompaniesV1ParamsFilterField = "legal_form"
	SearchFindCompaniesV1ParamsFilterFieldRegisterNumber    SearchFindCompaniesV1ParamsFilterField = "register_number"
	SearchFindCompaniesV1ParamsFilterFieldRegisterCourt     SearchFindCompaniesV1ParamsFilterField = "register_court"
	SearchFindCompaniesV1ParamsFilterFieldRegisterType      SearchFindCompaniesV1ParamsFilterField = "register_type"
	SearchFindCompaniesV1ParamsFilterFieldIncorporatedAt    SearchFindCompaniesV1ParamsFilterField = "incorporated_at"
	SearchFindCompaniesV1ParamsFilterFieldZip               SearchFindCompaniesV1ParamsFilterField = "zip"
	SearchFindCompaniesV1ParamsFilterFieldAddress           SearchFindCompaniesV1ParamsFilterField = "address"
	SearchFindCompaniesV1ParamsFilterFieldBalanceSheetTotal SearchFindCompaniesV1ParamsFilterField = "balance_sheet_total"
	SearchFindCompaniesV1ParamsFilterFieldRevenue           SearchFindCompaniesV1ParamsFilterField = "revenue"
	SearchFindCompaniesV1ParamsFilterFieldCash              SearchFindCompaniesV1ParamsFilterField = "cash"
	SearchFindCompaniesV1ParamsFilterFieldEmployees         SearchFindCompaniesV1ParamsFilterField = "employees"
	SearchFindCompaniesV1ParamsFilterFieldEquity            SearchFindCompaniesV1ParamsFilterField = "equity"
	SearchFindCompaniesV1ParamsFilterFieldRealEstate        SearchFindCompaniesV1ParamsFilterField = "real_estate"
	SearchFindCompaniesV1ParamsFilterFieldMaterials         SearchFindCompaniesV1ParamsFilterField = "materials"
	SearchFindCompaniesV1ParamsFilterFieldPensionProvisions SearchFindCompaniesV1ParamsFilterField = "pension_provisions"
	SearchFindCompaniesV1ParamsFilterFieldSalaries          SearchFindCompaniesV1ParamsFilterField = "salaries"
	SearchFindCompaniesV1ParamsFilterFieldTaxes             SearchFindCompaniesV1ParamsFilterField = "taxes"
	SearchFindCompaniesV1ParamsFilterFieldLiabilities       SearchFindCompaniesV1ParamsFilterField = "liabilities"
	SearchFindCompaniesV1ParamsFilterFieldCapitalReserves   SearchFindCompaniesV1ParamsFilterField = "capital_reserves"
	SearchFindCompaniesV1ParamsFilterFieldNetIncome         SearchFindCompaniesV1ParamsFilterField = "net_income"
	SearchFindCompaniesV1ParamsFilterFieldIndustryCodes     SearchFindCompaniesV1ParamsFilterField = "industry_codes"
	SearchFindCompaniesV1ParamsFilterFieldCapitalAmount     SearchFindCompaniesV1ParamsFilterField = "capital_amount"
	SearchFindCompaniesV1ParamsFilterFieldCapitalCurrency   SearchFindCompaniesV1ParamsFilterField = "capital_currency"
)

// Location to filter companies.
//
// The properties Latitude, Longitude are required.
type SearchFindCompaniesV1ParamsLocation struct {
	// Latitude to filter on.
	Latitude float64 `json:"latitude,required"`
	// Longitude to filter on.
	Longitude float64 `json:"longitude,required"`
	// Radius in kilometers to filter on. Example: 10
	Radius param.Opt[float64] `json:"radius,omitzero"`
	paramObj
}

func (r SearchFindCompaniesV1ParamsLocation) MarshalJSON() (data []byte, err error) {
	type shadow SearchFindCompaniesV1ParamsLocation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SearchFindCompaniesV1ParamsLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pagination parameters.
type SearchFindCompaniesV1ParamsPagination struct {
	// Page number to return.
	Page param.Opt[int64] `json:"page,omitzero"`
	// Number of results per page.
	PerPage param.Opt[int64] `json:"per_page,omitzero"`
	paramObj
}

func (r SearchFindCompaniesV1ParamsPagination) MarshalJSON() (data []byte, err error) {
	type shadow SearchFindCompaniesV1ParamsPagination
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SearchFindCompaniesV1ParamsPagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Search query to filter companies.
//
// The property Value is required.
type SearchFindCompaniesV1ParamsQuery struct {
	// Search query to filter companies.
	Value string `json:"value,required"`
	paramObj
}

func (r SearchFindCompaniesV1ParamsQuery) MarshalJSON() (data []byte, err error) {
	type shadow SearchFindCompaniesV1ParamsQuery
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SearchFindCompaniesV1ParamsQuery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchFindPersonParams struct {
	// Filters to filter people.
	Filters []SearchFindPersonParamsFilter `json:"filters,omitzero"`
	// Pagination parameters.
	Pagination SearchFindPersonParamsPagination `json:"pagination,omitzero"`
	// Search query to filter people.
	Query SearchFindPersonParamsQuery `json:"query,omitzero"`
	paramObj
}

func (r SearchFindPersonParams) MarshalJSON() (data []byte, err error) {
	type shadow SearchFindPersonParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SearchFindPersonParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Filter by field. The properties values, value, keywords and min/max are mutually
// exclusive. Dates must be in the format YYYY-MM-DD.
type SearchFindPersonParamsFilter struct {
	// Maximum value to filter on.
	Max param.Opt[string] `json:"max,omitzero"`
	// Minimum value to filter on.
	Min param.Opt[string] `json:"min,omitzero"`
	// Value to filter on.
	Value param.Opt[string] `json:"value,omitzero"`
	// Any of "date_of_birth", "city", "active", "status", "legal_form",
	// "register_number", "register_court", "register_type", "incorporated_at", "zip",
	// "address", "balance_sheet_total", "revenue", "cash", "employees", "equity",
	// "real_estate", "materials", "pension_provisions", "salaries", "taxes",
	// "liabilities", "capital_reserves", "net_income", "industry_codes",
	// "capital_amount", "capital_currency".
	Field SearchFindPersonParamsFilterField `json:"field,omitzero"`
	// Keywords to filter on.
	Keywords []string `json:"keywords,omitzero"`
	// Values to filter on.
	Values []string `json:"values,omitzero"`
	paramObj
}

func (r SearchFindPersonParamsFilter) MarshalJSON() (data []byte, err error) {
	type shadow SearchFindPersonParamsFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SearchFindPersonParamsFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchFindPersonParamsFilterField string

const (
	SearchFindPersonParamsFilterFieldDateOfBirth       SearchFindPersonParamsFilterField = "date_of_birth"
	SearchFindPersonParamsFilterFieldCity              SearchFindPersonParamsFilterField = "city"
	SearchFindPersonParamsFilterFieldActive            SearchFindPersonParamsFilterField = "active"
	SearchFindPersonParamsFilterFieldStatus            SearchFindPersonParamsFilterField = "status"
	SearchFindPersonParamsFilterFieldLegalForm         SearchFindPersonParamsFilterField = "legal_form"
	SearchFindPersonParamsFilterFieldRegisterNumber    SearchFindPersonParamsFilterField = "register_number"
	SearchFindPersonParamsFilterFieldRegisterCourt     SearchFindPersonParamsFilterField = "register_court"
	SearchFindPersonParamsFilterFieldRegisterType      SearchFindPersonParamsFilterField = "register_type"
	SearchFindPersonParamsFilterFieldIncorporatedAt    SearchFindPersonParamsFilterField = "incorporated_at"
	SearchFindPersonParamsFilterFieldZip               SearchFindPersonParamsFilterField = "zip"
	SearchFindPersonParamsFilterFieldAddress           SearchFindPersonParamsFilterField = "address"
	SearchFindPersonParamsFilterFieldBalanceSheetTotal SearchFindPersonParamsFilterField = "balance_sheet_total"
	SearchFindPersonParamsFilterFieldRevenue           SearchFindPersonParamsFilterField = "revenue"
	SearchFindPersonParamsFilterFieldCash              SearchFindPersonParamsFilterField = "cash"
	SearchFindPersonParamsFilterFieldEmployees         SearchFindPersonParamsFilterField = "employees"
	SearchFindPersonParamsFilterFieldEquity            SearchFindPersonParamsFilterField = "equity"
	SearchFindPersonParamsFilterFieldRealEstate        SearchFindPersonParamsFilterField = "real_estate"
	SearchFindPersonParamsFilterFieldMaterials         SearchFindPersonParamsFilterField = "materials"
	SearchFindPersonParamsFilterFieldPensionProvisions SearchFindPersonParamsFilterField = "pension_provisions"
	SearchFindPersonParamsFilterFieldSalaries          SearchFindPersonParamsFilterField = "salaries"
	SearchFindPersonParamsFilterFieldTaxes             SearchFindPersonParamsFilterField = "taxes"
	SearchFindPersonParamsFilterFieldLiabilities       SearchFindPersonParamsFilterField = "liabilities"
	SearchFindPersonParamsFilterFieldCapitalReserves   SearchFindPersonParamsFilterField = "capital_reserves"
	SearchFindPersonParamsFilterFieldNetIncome         SearchFindPersonParamsFilterField = "net_income"
	SearchFindPersonParamsFilterFieldIndustryCodes     SearchFindPersonParamsFilterField = "industry_codes"
	SearchFindPersonParamsFilterFieldCapitalAmount     SearchFindPersonParamsFilterField = "capital_amount"
	SearchFindPersonParamsFilterFieldCapitalCurrency   SearchFindPersonParamsFilterField = "capital_currency"
)

// Pagination parameters.
type SearchFindPersonParamsPagination struct {
	// Page number to return.
	Page param.Opt[int64] `json:"page,omitzero"`
	// Number of results per page.
	PerPage param.Opt[int64] `json:"per_page,omitzero"`
	paramObj
}

func (r SearchFindPersonParamsPagination) MarshalJSON() (data []byte, err error) {
	type shadow SearchFindPersonParamsPagination
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SearchFindPersonParamsPagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Search query to filter people.
//
// The property Value is required.
type SearchFindPersonParamsQuery struct {
	// Search query to filter people.
	Value string `json:"value,required"`
	paramObj
}

func (r SearchFindPersonParamsQuery) MarshalJSON() (data []byte, err error) {
	type shadow SearchFindPersonParamsQuery
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SearchFindPersonParamsQuery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchLookupCompanyByURLParams struct {
	// Website URL to search for. Example: "https://openregister.de"
	URL string `query:"url,required" format:"uri" json:"-"`
	paramObj
}

// URLQuery serializes [SearchLookupCompanyByURLParams]'s query parameters as
// `url.Values`.
func (r SearchLookupCompanyByURLParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
