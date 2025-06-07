// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openregister

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/oregister/openregister-go/internal/apijson"
	"github.com/oregister/openregister-go/internal/apiquery"
	"github.com/oregister/openregister-go/internal/requestconfig"
	"github.com/oregister/openregister-go/option"
	"github.com/oregister/openregister-go/packages/param"
	"github.com/oregister/openregister-go/packages/respjson"
)

// CompanyService contains methods and other services that help with interacting
// with the openregister API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCompanyService] method instead.
type CompanyService struct {
	Options []option.RequestOption
}

// NewCompanyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCompanyService(opts ...option.RequestOption) (r CompanyService) {
	r = CompanyService{}
	r.Options = opts
	return
}

// Get detailed company information
func (r *CompanyService) Get(ctx context.Context, companyID string, query CompanyGetParams, opts ...option.RequestOption) (res *CompanyGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if companyID == "" {
		err = errors.New("missing required company_id parameter")
		return
	}
	path := fmt.Sprintf("v0/company/%s", companyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get company shareholders
func (r *CompanyService) ListShareholders(ctx context.Context, companyID string, opts ...option.RequestOption) (res *CompanyListShareholdersResponse, err error) {
	opts = append(r.Options[:], opts...)
	if companyID == "" {
		err = errors.New("missing required company_id parameter")
		return
	}
	path := fmt.Sprintf("v0/company/%s/shareholders", companyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get company contact information
func (r *CompanyService) GetContact(ctx context.Context, companyID string, opts ...option.RequestOption) (res *CompanyGetContactResponse, err error) {
	opts = append(r.Options[:], opts...)
	if companyID == "" {
		err = errors.New("missing required company_id parameter")
		return
	}
	path := fmt.Sprintf("v0/company/%s/contact", companyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CompanyAddress struct {
	// City or locality name. Example: "Berlin"
	City string `json:"city,required"`
	// Country name. Example: "Germany"
	Country string `json:"country,required"`
	// Complete address formatted as a single string. Example: "Musterstraße 1, 10117
	// Berlin, Germany"
	FormattedValue string `json:"formatted_value,required"`
	// Date when this address became effective. Format: ISO 8601 (YYYY-MM-DD) Example:
	// "2022-01-01"
	StartDate string `json:"start_date,required"`
	// Additional address information such as c/o or attention line. Example: "c/o Max
	// Mustermann"
	Extra string `json:"extra"`
	// Postal or ZIP code. Example: "10117"
	PostalCode string `json:"postal_code"`
	// Street name and number. Example: "Musterstraße 1"
	Street string `json:"street"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City           respjson.Field
		Country        respjson.Field
		FormattedValue respjson.Field
		StartDate      respjson.Field
		Extra          respjson.Field
		PostalCode     respjson.Field
		Street         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyAddress) RawJSON() string { return r.JSON.raw }
func (r *CompanyAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyCapital struct {
	// Capital amount as a decimal number. Example: 100000.00 represents 100,000.00
	// monetary units
	Amount float64 `json:"amount,required"`
	// Currency code for the capital amount. Example: "EUR" for Euro
	//
	// Any of "EUR", "DEM", "USD".
	Currency CompanyCapitalCurrency `json:"currency,required"`
	// Date when this capital amount became effective. Format: ISO 8601 (YYYY-MM-DD)
	// Example: "2023-01-01"
	StartDate string `json:"start_date,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Currency    respjson.Field
		StartDate   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyCapital) RawJSON() string { return r.JSON.raw }
func (r *CompanyCapital) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Currency code for the capital amount. Example: "EUR" for Euro
type CompanyCapitalCurrency string

const (
	CompanyCapitalCurrencyEur CompanyCapitalCurrency = "EUR"
	CompanyCapitalCurrencyDem CompanyCapitalCurrency = "DEM"
	CompanyCapitalCurrencyUsd CompanyCapitalCurrency = "USD"
)

type CompanyName struct {
	// Legal form of the company at this point in time. Example: "gmbh" for
	// Gesellschaft mit beschränkter Haftung
	//
	// Any of "ag", "eg", "ek", "ev", "ewiv", "foreign", "gbr", "ggmbh", "gmbh", "kg",
	// "kgaa", "unknown", "llp", "municipal", "ohg", "se", "ug".
	LegalForm CompanyLegalForm `json:"legal_form,required"`
	// Official company name including any legal form designations. Example: "Descartes
	// Technologies UG (haftungsbeschränkt)"
	Name string `json:"name,required"`
	// Date when this name became effective. Format: ISO 8601 (YYYY-MM-DD) Example:
	// "2022-01-01"
	StartDate string `json:"start_date,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LegalForm   respjson.Field
		Name        respjson.Field
		StartDate   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyName) RawJSON() string { return r.JSON.raw }
func (r *CompanyName) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyPurpose struct {
	// Official description of the company's business activities and objectives. This
	// is the registered purpose as stated in official documents.
	Purpose string `json:"purpose,required"`
	// Date when this purpose became effective. Format: ISO 8601 (YYYY-MM-DD) Example:
	// "2022-01-01"
	StartDate string `json:"start_date,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Purpose     respjson.Field
		StartDate   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyPurpose) RawJSON() string { return r.JSON.raw }
func (r *CompanyPurpose) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyRegister struct {
	// Court where the company is registered. Example: "Berlin (Charlottenburg)"
	RegisterCourt string `json:"register_court,required"`
	// Registration number in the company register. Example: "230633"
	RegisterNumber string `json:"register_number,required"`
	// Type of register where the company is recorded. Example: "HRB" (Commercial
	// Register B)
	//
	// Any of "HRB", "HRA", "PR", "GnR", "VR".
	RegisterType CompanyRegisterType `json:"register_type,required"`
	// Unique company identifier. Example: DE-HRB-F1103-267645
	CompanyID string `json:"company_id"`
	// Date when this registration information became effective. Format: ISO 8601
	// (YYYY-MM-DD) Example: "2022-01-01"
	StartDate string `json:"start_date"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RegisterCourt  respjson.Field
		RegisterNumber respjson.Field
		RegisterType   respjson.Field
		CompanyID      respjson.Field
		StartDate      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyRegister) RawJSON() string { return r.JSON.raw }
func (r *CompanyRegister) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EntityType string

const (
	EntityTypeNaturalPerson EntityType = "natural_person"
	EntityTypeLegalPerson   EntityType = "legal_person"
)

type CompanyGetResponse struct {
	// Unique company identifier. Example: DE-HRB-F1103-267645
	ID string `json:"id,required"`
	// Current registered address of the company.
	Address CompanyAddress `json:"address,required"`
	// Date when the company was officially registered. Format: ISO 8601 (YYYY-MM-DD)
	// Example: "2022-01-01"
	IncorporatedAt string `json:"incorporated_at,required"`
	// Legal form of the company. Example: "gmbh" for Gesellschaft mit beschränkter
	// Haftung
	//
	// Any of "ag", "eg", "ek", "ev", "ewiv", "foreign", "gbr", "ggmbh", "gmbh", "kg",
	// "kgaa", "unknown", "llp", "municipal", "ohg", "se", "ug".
	LegalForm CompanyLegalForm `json:"legal_form,required"`
	// Current official name of the company.
	Name CompanyName `json:"name,required"`
	// Current registration information of the company.
	Register CompanyRegister `json:"register,required"`
	// List of individuals or entities authorized to represent the company. Includes
	// directors, officers, and authorized signatories.
	Representation []CompanyGetResponseRepresentation `json:"representation,required"`
	// Current status of the company:
	//
	// - active: Operating normally
	// - inactive: No longer operating
	// - liquidation: In the process of being dissolved
	//
	// Any of "active", "inactive", "liquidation".
	Status CompanyGetResponseStatus `json:"status,required"`
	// Historical addresses, only included when history=true. Shows how the company
	// address changed over time.
	Addresses []CompanyAddress `json:"addresses"`
	// Current registered capital of the company.
	Capital CompanyCapital `json:"capital"`
	// Historical capital changes, only included when history=true. Shows how the
	// company capital changed over time.
	Capitals []CompanyCapital `json:"capitals"`
	// Available official documents related to the company, only included when
	// documents=true.
	Documents []CompanyGetResponseDocument `json:"documents"`
	// Financial reports and key financial indicators, only included when
	// financials=true.
	Financials CompanyGetResponseFinancials `json:"financials"`
	// Historical company names, only included when history=true. Shows how the company
	// name changed over time.
	Names []CompanyName `json:"names"`
	// Current official business purpose of the company.
	Purpose CompanyPurpose `json:"purpose"`
	// Historical business purposes, only included when history=true. Shows how the
	// company purpose changed over time.
	Purposes []CompanyPurpose `json:"purposes"`
	// Historical registration changes, only included when history=true. Shows how
	// registration details changed over time.
	Registers []CompanyRegister `json:"registers"`
	// Date when the company was officially terminated (if applicable). Format: ISO
	// 8601 (YYYY-MM-DD) Example: "2022-01-01"
	TerminatedAt string `json:"terminated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Address        respjson.Field
		IncorporatedAt respjson.Field
		LegalForm      respjson.Field
		Name           respjson.Field
		Register       respjson.Field
		Representation respjson.Field
		Status         respjson.Field
		Addresses      respjson.Field
		Capital        respjson.Field
		Capitals       respjson.Field
		Documents      respjson.Field
		Financials     respjson.Field
		Names          respjson.Field
		Purpose        respjson.Field
		Purposes       respjson.Field
		Registers      respjson.Field
		TerminatedAt   respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetResponse) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyGetResponseRepresentation struct {
	// City where the representative is located. Example: "Berlin"
	City string `json:"city,required"`
	// Country where the representative is located, in ISO 3166-1 alpha-2 format.
	// Example: "DE" for Germany
	Country string `json:"country,required"`
	// The name of the representative. E.g. "Max Mustermann" or "Max Mustermann GmbH"
	Name string `json:"name,required"`
	// The role of the representation. E.g. "DIRECTOR"
	//
	// Any of "DIRECTOR", "PROKURA", "SHAREHOLDER", "OWNER", "PARTNER",
	// "PERSONAL_LIABLE_DIRECTOR", "LIQUIDATOR", "OTHER".
	Role string `json:"role,required"`
	// Date when this representative role became effective. Format: ISO 8601
	// (YYYY-MM-DD) Example: "2022-01-01"
	StartDate string `json:"start_date,required"`
	// Whether the representation is a natural person or a legal entity.
	//
	// Any of "natural_person", "legal_person".
	Type EntityType `json:"type,required"`
	// Unique identifier for the representative. For companies: Format matches
	// company_id pattern For individuals: UUID Example: "DE-HRB-F1103-267645" or UUID
	// May be null for certain representatives.
	ID string `json:"id"`
	// Date when this representative role ended (if applicable). Format: ISO 8601
	// (YYYY-MM-DD) Example: "2022-01-01"
	EndDate string `json:"end_date"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Country     respjson.Field
		Name        respjson.Field
		Role        respjson.Field
		StartDate   respjson.Field
		Type        respjson.Field
		ID          respjson.Field
		EndDate     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetResponseRepresentation) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetResponseRepresentation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of the company:
//
// - active: Operating normally
// - inactive: No longer operating
// - liquidation: In the process of being dissolved
type CompanyGetResponseStatus string

const (
	CompanyGetResponseStatusActive      CompanyGetResponseStatus = "active"
	CompanyGetResponseStatusInactive    CompanyGetResponseStatus = "inactive"
	CompanyGetResponseStatusLiquidation CompanyGetResponseStatus = "liquidation"
)

type CompanyGetResponseDocument struct {
	// Unique identifier for the document. Example:
	// "f47ac10b-58cc-4372-a567-0e02b2c3d479"
	ID string `json:"id,required"`
	// Document publication or filing date. Format: ISO 8601 (YYYY-MM-DD) Example:
	// "2022-01-01"
	Date string `json:"date,required"`
	// Whether this is the latest version of the document_type.
	Latest bool `json:"latest,required"`
	// Categorization of the document:
	//
	// - articles_of_association: Company statutes/bylaws
	// - sample_protocol: Standard founding protocol
	// - shareholder_list: List of company shareholders
	//
	// Any of "articles_of_association", "sample_protocol", "shareholder_list".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Date        respjson.Field
		Latest      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetResponseDocument) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetResponseDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Financial reports and key financial indicators, only included when
// financials=true.
type CompanyGetResponseFinancials struct {
	// Key financial metrics extracted from the reports. Includes balance sheet totals,
	// revenue, and other important figures.
	Indicators []CompanyGetResponseFinancialsIndicator `json:"indicators,required"`
	// The financial reports of the company.
	Reports []CompanyGetResponseFinancialsReport `json:"reports,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Indicators  respjson.Field
		Reports     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetResponseFinancials) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetResponseFinancials) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyGetResponseFinancialsIndicator struct {
	// Date to which this financial indicator applies. Format: ISO 8601 (YYYY-MM-DD)
	// Example: "2022-01-01"
	Date string `json:"date,required"`
	// The identifier for the financial report this indicator originates from. E.g.
	// "f47ac10b-58cc-4372-a567-0e02b2c3d479"
	ReportID string `json:"report_id,required"`
	// The type of indicator.
	//
	// Any of "balance_sheet_total", "net_income", "revenue", "cash", "employees",
	// "equity", "real_estate", "materials", "pension_provisions", "salaries", "taxes",
	// "liabilities", "capital_reserves".
	Type string `json:"type,required"`
	// Value of the indicator in the smallest currency unit (cents). Example: 2099
	// represents €20.99 for monetary values For non-monetary values (e.g., employees),
	// the actual number.
	Value int64 `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date        respjson.Field
		ReportID    respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetResponseFinancialsIndicator) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetResponseFinancialsIndicator) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyGetResponseFinancialsReport struct {
	// The unique identifier for the financial report. E.g.
	// "f47ac10b-58cc-4372-a567-0e02b2c3d479"
	ID string `json:"id,required"`
	// The name of the financial report. E.g. "Jahresabschluss 2022"
	Name string `json:"name,required"`
	// The date when the financial report was published. E.g. "2022-01-01"
	PublishedAt string `json:"published_at,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		PublishedAt respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetResponseFinancialsReport) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetResponseFinancialsReport) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyListShareholdersResponse struct {
	// Date when this shareholder information became effective. Format: ISO 8601
	// (YYYY-MM-DD) Example: "2022-01-01"
	Date string `json:"date,required"`
	// Unique identifier for the document this was taken from. Example:
	// "f47ac10b-58cc-4372-a567-0e02b2c3d479"
	DocumentID   string                                       `json:"document_id,required"`
	Shareholders []CompanyListShareholdersResponseShareholder `json:"shareholders,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date         respjson.Field
		DocumentID   respjson.Field
		Shareholders respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyListShareholdersResponse) RawJSON() string { return r.JSON.raw }
func (r *CompanyListShareholdersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyListShareholdersResponseShareholder struct {
	// Country where the shareholder is located, in ISO 3166-1 alpha-2 format. Example:
	// "DE" for Germany
	Country string `json:"country,required"`
	// The name of the shareholder. E.g. "Max Mustermann" or "Max Mustermann GmbH"
	Name string `json:"name,required"`
	// Nominal value of shares in Euro. Example: 100
	NominalShare int64 `json:"nominal_share,required"`
	// Percentage of company ownership. Example: 5.36 represents 5.36% ownership
	PercentageShare float64 `json:"percentage_share,required"`
	// The type of shareholder.
	//
	// Any of "natural_person", "legal_person".
	Type EntityType `json:"type,required"`
	// Unique identifier for the shareholder. For companies: Format matches company_id
	// pattern For individuals: UUID Example: "DE-HRB-F1103-267645" or UUID May be null
	// for certain shareholders.
	ID string `json:"id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Country         respjson.Field
		Name            respjson.Field
		NominalShare    respjson.Field
		PercentageShare respjson.Field
		Type            respjson.Field
		ID              respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyListShareholdersResponseShareholder) RawJSON() string { return r.JSON.raw }
func (r *CompanyListShareholdersResponseShareholder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyGetContactResponse struct {
	// Where the contact information was found. Example: "https://openregister.de"
	SourceURL string `json:"source_url,required" format:"uri"`
	// Company contact email address. Example: "founders@openregister.de"
	Email string `json:"email"`
	// Company phone number. Example: "+49 030 12345678"
	Phone string `json:"phone"`
	// Value Added Tax identification number. (Umsatzsteuer-Identifikationsnummer)
	// Example: "DE370146530"
	VatID string `json:"vat_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SourceURL   respjson.Field
		Email       respjson.Field
		Phone       respjson.Field
		VatID       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetContactResponse) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetContactResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyGetParams struct {
	// Include document metadata when set to true. Lists available official documents
	// related to the company.
	Documents param.Opt[bool] `query:"documents,omitzero" json:"-"`
	// Include financial data when set to true. Provides access to financial reports
	// and key financial indicators.
	Financials param.Opt[bool] `query:"financials,omitzero" json:"-"`
	// Include historical company data when set to true. This returns past names,
	// addresses, and other changed information.
	History param.Opt[bool] `query:"history,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CompanyGetParams]'s query parameters as `url.Values`.
func (r CompanyGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
