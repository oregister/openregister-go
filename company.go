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

// Get company contact information
func (r *CompanyService) GetContactV0(ctx context.Context, companyID string, opts ...option.RequestOption) (res *CompanyGetContactV0Response, err error) {
	opts = append(r.Options[:], opts...)
	if companyID == "" {
		err = errors.New("missing required company_id parameter")
		return
	}
	path := fmt.Sprintf("v0/company/%s/contact", companyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get detailed company information
func (r *CompanyService) GetDetailsV1(ctx context.Context, companyID string, query CompanyGetDetailsV1Params, opts ...option.RequestOption) (res *CompanyGetDetailsV1Response, err error) {
	opts = append(r.Options[:], opts...)
	if companyID == "" {
		err = errors.New("missing required company_id parameter")
		return
	}
	path := fmt.Sprintf("v1/company/%s", companyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get financial reports
func (r *CompanyService) GetFinancialsV1(ctx context.Context, companyID string, opts ...option.RequestOption) (res *CompanyGetFinancialsV1Response, err error) {
	opts = append(r.Options[:], opts...)
	if companyID == "" {
		err = errors.New("missing required company_id parameter")
		return
	}
	path := fmt.Sprintf("v1/company/%s/financials", companyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get company holdings
func (r *CompanyService) GetHoldingsV1(ctx context.Context, companyID string, opts ...option.RequestOption) (res *CompanyGetHoldingsV1Response, err error) {
	opts = append(r.Options[:], opts...)
	if companyID == "" {
		err = errors.New("missing required company_id parameter")
		return
	}
	path := fmt.Sprintf("v1/company/%s/holdings", companyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get company owners
func (r *CompanyService) GetOwnersV1(ctx context.Context, companyID string, query CompanyGetOwnersV1Params, opts ...option.RequestOption) (res *CompanyGetOwnersV1Response, err error) {
	opts = append(r.Options[:], opts...)
	if companyID == "" {
		err = errors.New("missing required company_id parameter")
		return
	}
	path := fmt.Sprintf("v1/company/%s/owners", companyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
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

type CompanyRelationType string

const (
	CompanyRelationTypeShareholder    CompanyRelationType = "shareholder"
	CompanyRelationTypeStockholder    CompanyRelationType = "stockholder"
	CompanyRelationTypeLimitedPartner CompanyRelationType = "limited_partner"
	CompanyRelationTypeGeneralPartner CompanyRelationType = "general_partner"
)

type EntityType string

const (
	EntityTypeNaturalPerson EntityType = "natural_person"
	EntityTypeLegalPerson   EntityType = "legal_person"
)

type ReportRow struct {
	Children      []ReportRow `json:"children,required"`
	CurrentValue  int64       `json:"current_value,required"`
	FormattedName string      `json:"formatted_name,required"`
	Name          string      `json:"name,required"`
	PreviousValue int64       `json:"previous_value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Children      respjson.Field
		CurrentValue  respjson.Field
		FormattedName respjson.Field
		Name          respjson.Field
		PreviousValue respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReportRow) RawJSON() string { return r.JSON.raw }
func (r *ReportRow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyGetContactV0Response struct {
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
func (r CompanyGetContactV0Response) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetContactV0Response) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyGetDetailsV1Response struct {
	// Unique company identifier. Example: DE-HRB-F1103-267645
	ID string `json:"id,required"`
	// Current registered address of the company.
	Address CompanyAddress `json:"address,required"`
	// Historical addresses. Shows how the company address changed over time.
	Addresses []CompanyAddress `json:"addresses,required"`
	// Current registered capital of the company.
	Capital CompanyCapital `json:"capital,required"`
	// Historical capital changes. Shows how the company capital changed over time.
	Capitals []CompanyCapital `json:"capitals,required"`
	// Available official documents related to the company.
	Documents []CompanyGetDetailsV1ResponseDocument `json:"documents,required"`
	// Date when the company was officially registered. Format: ISO 8601 (YYYY-MM-DD)
	// Example: "2022-01-01"
	IncorporatedAt string `json:"incorporated_at,required"`
	// Key company indicators like net income, employee count, revenue, etc..
	Indicators []CompanyGetDetailsV1ResponseIndicator `json:"indicators,required"`
	// Industry codes of the company.
	IndustryCodes CompanyGetDetailsV1ResponseIndustryCodes `json:"industry_codes,required"`
	// Legal form of the company. Example: "gmbh" for Gesellschaft mit beschränkter
	// Haftung
	//
	// Any of "ag", "eg", "ek", "ev", "ewiv", "foreign", "gbr", "ggmbh", "gmbh", "kg",
	// "kgaa", "unknown", "llp", "municipal", "ohg", "se", "ug".
	LegalForm CompanyLegalForm `json:"legal_form,required"`
	// Current official name of the company.
	Name CompanyName `json:"name,required"`
	// Historical company names. Shows how the company name changed over time.
	Names []CompanyName `json:"names,required"`
	// Current official business purpose of the company.
	Purpose CompanyPurpose `json:"purpose,required"`
	// Historical business purposes. Shows how the company purpose changed over time.
	Purposes []CompanyPurpose `json:"purposes,required"`
	// Current registration information of the company.
	Register CompanyRegister `json:"register,required"`
	// Historical registration changes. Shows how registration details changed over
	// time.
	Registers []CompanyRegister `json:"registers,required"`
	// List of individuals or entities authorized to represent the company. Includes
	// directors, officers, and authorized signatories.
	Representation []CompanyGetDetailsV1ResponseRepresentation `json:"representation,required"`
	// Sources of the company data.
	Sources []CompanyGetDetailsV1ResponseSource `json:"sources,required"`
	// Current status of the company:
	//
	// - active: Operating normally
	// - inactive: No longer operating
	// - liquidation: In the process of being dissolved
	//
	// Any of "active", "inactive", "liquidation".
	Status CompanyGetDetailsV1ResponseStatus `json:"status,required"`
	// Date when the company was officially terminated (if applicable). Format: ISO
	// 8601 (YYYY-MM-DD) Example: "2022-01-01"
	TerminatedAt string `json:"terminated_at,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Address        respjson.Field
		Addresses      respjson.Field
		Capital        respjson.Field
		Capitals       respjson.Field
		Documents      respjson.Field
		IncorporatedAt respjson.Field
		Indicators     respjson.Field
		IndustryCodes  respjson.Field
		LegalForm      respjson.Field
		Name           respjson.Field
		Names          respjson.Field
		Purpose        respjson.Field
		Purposes       respjson.Field
		Register       respjson.Field
		Registers      respjson.Field
		Representation respjson.Field
		Sources        respjson.Field
		Status         respjson.Field
		TerminatedAt   respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetDetailsV1Response) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetDetailsV1Response) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyGetDetailsV1ResponseDocument struct {
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
func (r CompanyGetDetailsV1ResponseDocument) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetDetailsV1ResponseDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The indicators of the company for a given year. Values of the indicator are
// given in the smallest currency unit (cents). Example: 2099 represents €20.99 for
// monetary values For non-monetary values (e.g., employees), the actual number.
type CompanyGetDetailsV1ResponseIndicator struct {
	// The balance sheet total of that year (in cents).
	BalanceSheetTotal int64 `json:"balance_sheet_total,required"`
	// The capital reserves of that year (in cents).
	CapitalReserves int64 `json:"capital_reserves,required"`
	// The cash of that year (in cents).
	Cash int64 `json:"cash,required"`
	// Date to which this financial indicators apply. Format: ISO 8601 (YYYY-MM-DD)
	// Example: "2022-01-01"
	Date string `json:"date,required"`
	// The number of employees of that year.
	Employees int64 `json:"employees,required"`
	// The equity of that year (in cents).
	Equity int64 `json:"equity,required"`
	// The liabilities of that year (in cents).
	Liabilities int64 `json:"liabilities,required"`
	// The materials of that year (in cents).
	Materials int64 `json:"materials,required"`
	// The net income of that year (in cents).
	NetIncome int64 `json:"net_income,required"`
	// The pension provisions of that year (in cents).
	PensionProvisions int64 `json:"pension_provisions,required"`
	// The real estate of that year (in cents).
	RealEstate int64 `json:"real_estate,required"`
	// The report id (source) of the indicators.
	ReportID string `json:"report_id,required" format:"uuid"`
	// The revenue of that year (in cents).
	Revenue int64 `json:"revenue,required"`
	// The salaries of that year (in cents).
	Salaries int64 `json:"salaries,required"`
	// The taxes of that year (in cents).
	Taxes int64 `json:"taxes,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BalanceSheetTotal respjson.Field
		CapitalReserves   respjson.Field
		Cash              respjson.Field
		Date              respjson.Field
		Employees         respjson.Field
		Equity            respjson.Field
		Liabilities       respjson.Field
		Materials         respjson.Field
		NetIncome         respjson.Field
		PensionProvisions respjson.Field
		RealEstate        respjson.Field
		ReportID          respjson.Field
		Revenue           respjson.Field
		Salaries          respjson.Field
		Taxes             respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetDetailsV1ResponseIndicator) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetDetailsV1ResponseIndicator) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Industry codes of the company.
type CompanyGetDetailsV1ResponseIndustryCodes struct {
	Wz2025 []CompanyGetDetailsV1ResponseIndustryCodesWz2025 `json:"WZ2025,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Wz2025      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetDetailsV1ResponseIndustryCodes) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetDetailsV1ResponseIndustryCodes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Industry codes from WZ 2025.
type CompanyGetDetailsV1ResponseIndustryCodesWz2025 struct {
	Code string `json:"code,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetDetailsV1ResponseIndustryCodesWz2025) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetDetailsV1ResponseIndustryCodesWz2025) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyGetDetailsV1ResponseRepresentation struct {
	// Unique identifier for the representative. For companies: Format matches
	// company_id pattern For individuals: UUID Example: "DE-HRB-F1103-267645" or UUID
	// May be null for certain representatives.
	ID string `json:"id,required"`
	// Date when this representative role ended (if applicable). Format: ISO 8601
	// (YYYY-MM-DD) Example: "2022-01-01"
	EndDate string `json:"end_date,required"`
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
	Type          EntityType                                             `json:"type,required"`
	LegalPerson   CompanyGetDetailsV1ResponseRepresentationLegalPerson   `json:"legal_person,nullable"`
	NaturalPerson CompanyGetDetailsV1ResponseRepresentationNaturalPerson `json:"natural_person,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		EndDate       respjson.Field
		Name          respjson.Field
		Role          respjson.Field
		StartDate     respjson.Field
		Type          respjson.Field
		LegalPerson   respjson.Field
		NaturalPerson respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetDetailsV1ResponseRepresentation) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetDetailsV1ResponseRepresentation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyGetDetailsV1ResponseRepresentationLegalPerson struct {
	City string `json:"city,required"`
	// Country where the representative is located, in ISO 3166-1 alpha-2 format.
	// Example: "DE" for Germany
	Country string `json:"country,required"`
	Name    string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Country     respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetDetailsV1ResponseRepresentationLegalPerson) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetDetailsV1ResponseRepresentationLegalPerson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyGetDetailsV1ResponseRepresentationNaturalPerson struct {
	// City where the representative is located. Example: "Berlin"
	City string `json:"city,required"`
	// Date of birth of the representative. May still be null for natural persons if it
	// is not available. Format: ISO 8601 (YYYY-MM-DD) Example: "1990-01-01"
	DateOfBirth string `json:"date_of_birth,required"`
	// First name of the representative. Example: "Max"
	FirstName string `json:"first_name,required"`
	// Last name of the representative. Example: "Mustermann"
	LastName string `json:"last_name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		DateOfBirth respjson.Field
		FirstName   respjson.Field
		LastName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetDetailsV1ResponseRepresentationNaturalPerson) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetDetailsV1ResponseRepresentationNaturalPerson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyGetDetailsV1ResponseSource struct {
	// Url of the source document. In the form of a presigned url accessible for 30
	// minutes.
	DocumentURL string `json:"document_url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DocumentURL respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetDetailsV1ResponseSource) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetDetailsV1ResponseSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of the company:
//
// - active: Operating normally
// - inactive: No longer operating
// - liquidation: In the process of being dissolved
type CompanyGetDetailsV1ResponseStatus string

const (
	CompanyGetDetailsV1ResponseStatusActive      CompanyGetDetailsV1ResponseStatus = "active"
	CompanyGetDetailsV1ResponseStatusInactive    CompanyGetDetailsV1ResponseStatus = "inactive"
	CompanyGetDetailsV1ResponseStatusLiquidation CompanyGetDetailsV1ResponseStatus = "liquidation"
)

type CompanyGetFinancialsV1Response struct {
	Reports []CompanyGetFinancialsV1ResponseReport `json:"reports,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Reports     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetFinancialsV1Response) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetFinancialsV1Response) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyGetFinancialsV1ResponseReport struct {
	Aktiva CompanyGetFinancialsV1ResponseReportAktiva `json:"aktiva,required"`
	// Whether the report is a consolidated report or not.
	Consolidated  bool                                        `json:"consolidated,required"`
	Passiva       CompanyGetFinancialsV1ResponseReportPassiva `json:"passiva,required"`
	ReportEndDate string                                      `json:"report_end_date,required"`
	// Unique identifier for the financial report. Example:
	// f47ac10b-58cc-4372-a567-0e02b2c3d479
	ReportID        string                                  `json:"report_id,required"`
	ReportStartDate string                                  `json:"report_start_date,required"`
	Guv             CompanyGetFinancialsV1ResponseReportGuv `json:"guv,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Aktiva          respjson.Field
		Consolidated    respjson.Field
		Passiva         respjson.Field
		ReportEndDate   respjson.Field
		ReportID        respjson.Field
		ReportStartDate respjson.Field
		Guv             respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetFinancialsV1ResponseReport) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetFinancialsV1ResponseReport) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyGetFinancialsV1ResponseReportAktiva struct {
	Rows []ReportRow `json:"rows,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Rows        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetFinancialsV1ResponseReportAktiva) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetFinancialsV1ResponseReportAktiva) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyGetFinancialsV1ResponseReportPassiva struct {
	Rows []ReportRow `json:"rows,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Rows        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetFinancialsV1ResponseReportPassiva) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetFinancialsV1ResponseReportPassiva) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyGetFinancialsV1ResponseReportGuv struct {
	Rows []ReportRow `json:"rows,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Rows        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetFinancialsV1ResponseReportGuv) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetFinancialsV1ResponseReportGuv) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Companies this entity owns or has invested in.
type CompanyGetHoldingsV1Response struct {
	// Unique company identifier. Example: DE-HRB-F1103-267645
	CompanyID string                                `json:"company_id,required"`
	Holdings  []CompanyGetHoldingsV1ResponseHolding `json:"holdings,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompanyID   respjson.Field
		Holdings    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetHoldingsV1Response) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetHoldingsV1Response) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyGetHoldingsV1ResponseHolding struct {
	// Unique company identifier. Example: DE-HRB-F1103-267645
	CompanyID string `json:"company_id,required"`
	// Date when the ownership ended. Format: ISO 8601 (YYYY-MM-DD) Example:
	// "2022-01-01"
	End string `json:"end,required"`
	// Name of the company.
	Name string `json:"name,required"`
	// Amount of shares or capital in the company. Example: 100
	NominalShare float64 `json:"nominal_share,required"`
	// Share of the company. Example: 0.5 represents 50% ownership
	PercentageShare float64 `json:"percentage_share,required"`
	// Type of relationship between the entity and the company.
	//
	// Any of "shareholder", "stockholder", "limited_partner", "general_partner".
	RelationType CompanyRelationType `json:"relation_type,required"`
	// Date when the ownership started. Format: ISO 8601 (YYYY-MM-DD) Example:
	// "2022-01-01"
	Start string `json:"start,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompanyID       respjson.Field
		End             respjson.Field
		Name            respjson.Field
		NominalShare    respjson.Field
		PercentageShare respjson.Field
		RelationType    respjson.Field
		Start           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetHoldingsV1ResponseHolding) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetHoldingsV1ResponseHolding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyGetOwnersV1Response struct {
	// Unique company identifier. Example: DE-HRB-F1103-267645
	CompanyID string                            `json:"company_id,required"`
	Owners    []CompanyGetOwnersV1ResponseOwner `json:"owners,required"`
	// Sources of the company owners data.
	Sources []CompanyGetOwnersV1ResponseSource `json:"sources,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompanyID   respjson.Field
		Owners      respjson.Field
		Sources     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetOwnersV1Response) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetOwnersV1Response) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyGetOwnersV1ResponseOwner struct {
	// Unique identifier for the shareholder. For companies: Format matches company_id
	// pattern For individuals: UUID Example: "DE-HRB-F1103-267645" or UUID May be null
	// for certain shareholders.
	ID string `json:"id,required"`
	// Details about the legal person.
	LegalPerson CompanyGetOwnersV1ResponseOwnerLegalPerson `json:"legal_person,required"`
	// The name of the shareholder. E.g. "Max Mustermann" or "Max Mustermann GmbH"
	Name string `json:"name,required"`
	// Details about the natural person.
	NaturalPerson CompanyGetOwnersV1ResponseOwnerNaturalPerson `json:"natural_person,required"`
	// Nominal value of shares in Euro. Example: 100
	NominalShare float64 `json:"nominal_share,required"`
	// Percentage of company ownership. Example: 5.36 represents 5.36% ownership
	PercentageShare float64 `json:"percentage_share,required"`
	// Type of relationship between the entity and the company.
	//
	// Any of "shareholder", "stockholder", "limited_partner", "general_partner".
	RelationType CompanyRelationType `json:"relation_type,required"`
	// Date when the relation started. Only available for some types of owners. Format:
	// ISO 8601 (YYYY-MM-DD) Example: "2022-01-01"
	Start string `json:"start,required"`
	// The type of shareholder.
	//
	// Any of "natural_person", "legal_person".
	Type EntityType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		LegalPerson     respjson.Field
		Name            respjson.Field
		NaturalPerson   respjson.Field
		NominalShare    respjson.Field
		PercentageShare respjson.Field
		RelationType    respjson.Field
		Start           respjson.Field
		Type            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetOwnersV1ResponseOwner) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetOwnersV1ResponseOwner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details about the legal person.
type CompanyGetOwnersV1ResponseOwnerLegalPerson struct {
	City string `json:"city,required"`
	// Country where the owner is located, in ISO 3166-1 alpha-2 format. Example: "DE"
	// for Germany
	Country string `json:"country,required"`
	Name    string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Country     respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetOwnersV1ResponseOwnerLegalPerson) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetOwnersV1ResponseOwnerLegalPerson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details about the natural person.
type CompanyGetOwnersV1ResponseOwnerNaturalPerson struct {
	City        string `json:"city,required"`
	Country     string `json:"country,required"`
	DateOfBirth string `json:"date_of_birth,required"`
	FirstName   string `json:"first_name,required"`
	FullName    string `json:"full_name,required"`
	LastName    string `json:"last_name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Country     respjson.Field
		DateOfBirth respjson.Field
		FirstName   respjson.Field
		FullName    respjson.Field
		LastName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetOwnersV1ResponseOwnerNaturalPerson) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetOwnersV1ResponseOwnerNaturalPerson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyGetOwnersV1ResponseSource struct {
	// Url of the source document. In the form of a presigned url accessible for 30
	// minutes.
	DocumentURL string `json:"document_url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DocumentURL respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetOwnersV1ResponseSource) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetOwnersV1ResponseSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyGetDetailsV1Params struct {
	// Get the most up-to-date company information directly from the Handelsregister.
	// When set to true, we fetch the latest data in real-time from the official German
	// commercial register, ensuring you receive the most current company details.
	// Note: Real-time requests take longer but guarantee the freshest data available.
	Realtime param.Opt[bool] `query:"realtime,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CompanyGetDetailsV1Params]'s query parameters as
// `url.Values`.
func (r CompanyGetDetailsV1Params) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CompanyGetOwnersV1Params struct {
	// Get the most up-to-date company information directly from the Handelsregister.
	// When set to true, we fetch the latest data in real-time from the official German
	// commercial register, ensuring you receive the most current company details.
	// Note: Real-time requests take longer but guarantee the freshest data available.
	Realtime param.Opt[bool] `query:"realtime,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CompanyGetOwnersV1Params]'s query parameters as
// `url.Values`.
func (r CompanyGetOwnersV1Params) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
