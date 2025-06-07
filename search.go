// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openregister

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/openregister-go/internal/apijson"
	"github.com/stainless-sdks/openregister-go/internal/apiquery"
	"github.com/stainless-sdks/openregister-go/internal/requestconfig"
	"github.com/stainless-sdks/openregister-go/option"
	"github.com/stainless-sdks/openregister-go/packages/param"
	"github.com/stainless-sdks/openregister-go/packages/respjson"
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

// Search for companies
func (r *SearchService) FindCompanies(ctx context.Context, query SearchFindCompaniesParams, opts ...option.RequestOption) (res *SearchFindCompaniesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/search/company"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
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

type SearchFindCompaniesResponse struct {
	// List of companies matching the search criteria. Limited to a maximum of 10
	// results per query.
	Results []SearchFindCompaniesResponseResult `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchFindCompaniesResponse) RawJSON() string { return r.JSON.raw }
func (r *SearchFindCompaniesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchFindCompaniesResponseResult struct {
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
func (r SearchFindCompaniesResponseResult) RawJSON() string { return r.JSON.raw }
func (r *SearchFindCompaniesResponseResult) UnmarshalJSON(data []byte) error {
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

type SearchFindCompaniesParams struct {
	// Filter for active or inactive companies. Set to true for active companies only,
	// false for inactive only.
	Active param.Opt[bool] `query:"active,omitzero" json:"-"`
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

// URLQuery serializes [SearchFindCompaniesParams]'s query parameters as
// `url.Values`.
func (r SearchFindCompaniesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
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
