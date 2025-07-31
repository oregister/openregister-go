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
	"github.com/oregister/openregister-go/packages/respjson"
)

// AutocompleteService contains methods and other services that help with
// interacting with the openregister API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAutocompleteService] method instead.
type AutocompleteService struct {
	Options []option.RequestOption
}

// NewAutocompleteService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAutocompleteService(opts ...option.RequestOption) (r AutocompleteService) {
	r = AutocompleteService{}
	r.Options = opts
	return
}

// Autocomplete company search
func (r *AutocompleteService) AutocompleteCompaniesV1(ctx context.Context, query AutocompleteAutocompleteCompaniesV1Params, opts ...option.RequestOption) (res *AutocompleteAutocompleteCompaniesV1Response, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/autocomplete/company"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AutocompleteAutocompleteCompaniesV1Response struct {
	// List of companies matching the search criteria.
	Results []AutocompleteAutocompleteCompaniesV1ResponseResult `json:"results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AutocompleteAutocompleteCompaniesV1Response) RawJSON() string { return r.JSON.raw }
func (r *AutocompleteAutocompleteCompaniesV1Response) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutocompleteAutocompleteCompaniesV1ResponseResult struct {
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
func (r AutocompleteAutocompleteCompaniesV1ResponseResult) RawJSON() string { return r.JSON.raw }
func (r *AutocompleteAutocompleteCompaniesV1ResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AutocompleteAutocompleteCompaniesV1Params struct {
	// Text search query to find companies by name. Example: "Descartes Technologies
	// UG"
	Query string `query:"query,required" json:"-"`
	paramObj
}

// URLQuery serializes [AutocompleteAutocompleteCompaniesV1Params]'s query
// parameters as `url.Values`.
func (r AutocompleteAutocompleteCompaniesV1Params) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
