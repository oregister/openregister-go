// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openregister

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/oregister/openregister-go/v2/internal/apijson"
	"github.com/oregister/openregister-go/v2/internal/apiquery"
	"github.com/oregister/openregister-go/v2/internal/requestconfig"
	"github.com/oregister/openregister-go/v2/option"
	"github.com/oregister/openregister-go/v2/packages/param"
	"github.com/oregister/openregister-go/v2/packages/respjson"
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
	opts = slices.Concat(r.Options, opts)
	if companyID == "" {
		err = errors.New("missing required company_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v0/company/%s/contact", companyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get detailed company information
func (r *CompanyService) GetDetailsV1(ctx context.Context, companyID string, query CompanyGetDetailsV1Params, opts ...option.RequestOption) (res *CompanyGetDetailsV1Response, err error) {
	opts = slices.Concat(r.Options, opts)
	if companyID == "" {
		err = errors.New("missing required company_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/company/%s", companyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get financial reports
func (r *CompanyService) GetFinancialsV1(ctx context.Context, companyID string, opts ...option.RequestOption) (res *CompanyGetFinancialsV1Response, err error) {
	opts = slices.Concat(r.Options, opts)
	if companyID == "" {
		err = errors.New("missing required company_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/company/%s/financials", companyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get historical owner changes
func (r *CompanyService) GetHistoricalOwnersV1(ctx context.Context, companyID string, opts ...option.RequestOption) (res *CompanyGetHistoricalOwnersV1Response, err error) {
	opts = slices.Concat(r.Options, opts)
	if companyID == "" {
		err = errors.New("missing required company_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/company/%s/owners/historical", companyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get company holdings
func (r *CompanyService) GetHoldingsV1(ctx context.Context, companyID string, opts ...option.RequestOption) (res *CompanyGetHoldingsV1Response, err error) {
	opts = slices.Concat(r.Options, opts)
	if companyID == "" {
		err = errors.New("missing required company_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/company/%s/holdings", companyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get company owners
func (r *CompanyService) GetOwnersV1(ctx context.Context, companyID string, query CompanyGetOwnersV1Params, opts ...option.RequestOption) (res *CompanyGetOwnersV1Response, err error) {
	opts = slices.Concat(r.Options, opts)
	if companyID == "" {
		err = errors.New("missing required company_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/company/%s/owners", companyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get company end owners
func (r *CompanyService) GetUbosV1(ctx context.Context, companyID string, opts ...option.RequestOption) (res *CompanyGetUbosV1Response, err error) {
	opts = slices.Concat(r.Options, opts)
	if companyID == "" {
		err = errors.New("missing required company_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/company/%s/ubo", companyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type CompanyAddress struct {
	// City or locality name. Example: "Berlin"
	City string `json:"city" api:"required"`
	// Country name. Example: "Germany"
	Country string `json:"country" api:"required"`
	// Complete address formatted as a single string. Example: "Musterstraße 1, 10117
	// Berlin, Germany"
	FormattedValue string `json:"formatted_value" api:"required"`
	// Date when this address became effective. Format: ISO 8601 (YYYY-MM-DD) Example:
	// "2022-01-01"
	StartDate string `json:"start_date" api:"required"`
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
	Amount float64 `json:"amount" api:"required"`
	// Currency code for the capital amount. Example: "EUR" for Euro
	//
	// Any of "EUR", "DEM", "USD".
	Currency CompanyCapitalCurrency `json:"currency" api:"required"`
	// Date when this capital amount became effective. Format: ISO 8601 (YYYY-MM-DD)
	// Example: "2023-01-01"
	StartDate string `json:"start_date" api:"required"`
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

type CompanyDocument struct {
	// Unique identifier for the document. Example:
	// "f47ac10b-58cc-4372-a567-0e02b2c3d479"
	ID string `json:"id" api:"required"`
	// Document publication or filing date. Format: ISO 8601 (YYYY-MM-DD) Example:
	// "2022-01-01"
	Date string `json:"date" api:"required"`
	// Whether this is the latest version of the document_type.
	Latest bool `json:"latest" api:"required"`
	// Categorization of the document:
	//
	// - articles_of_association: Company statutes/bylaws
	// - sample_protocol: Standard founding protocol
	// - shareholder_list: List of company shareholders
	//
	// Any of "articles_of_association", "sample_protocol", "shareholder_list".
	Type CompanyDocumentType `json:"type" api:"required"`
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
func (r CompanyDocument) RawJSON() string { return r.JSON.raw }
func (r *CompanyDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Categorization of the document:
//
// - articles_of_association: Company statutes/bylaws
// - sample_protocol: Standard founding protocol
// - shareholder_list: List of company shareholders
type CompanyDocumentType string

const (
	CompanyDocumentTypeArticlesOfAssociation CompanyDocumentType = "articles_of_association"
	CompanyDocumentTypeSampleProtocol        CompanyDocumentType = "sample_protocol"
	CompanyDocumentTypeShareholderList       CompanyDocumentType = "shareholder_list"
)

type CompanyName struct {
	// Legal form of the company at this point in time. Example: "gmbh" for
	// Gesellschaft mit beschränkter Haftung
	//
	// Any of "ag", "eg", "ek", "ev", "ewiv", "foreign", "gbr", "ggmbh", "gmbh", "kg",
	// "kgaa", "unknown", "llp", "municipal", "ohg", "se", "ug".
	LegalForm CompanyLegalForm `json:"legal_form" api:"required"`
	// Official company name including any legal form designations. Example: "Descartes
	// Technologies UG (haftungsbeschränkt)"
	Name string `json:"name" api:"required"`
	// Date when this name became effective. Format: ISO 8601 (YYYY-MM-DD) Example:
	// "2022-01-01"
	StartDate string `json:"start_date" api:"required"`
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

type CompanyOwnerLegalPerson struct {
	City string `json:"city" api:"required"`
	// Country where the owner is located, in ISO 3166-1 alpha-2 format. Example: "DE"
	// for Germany
	Country string `json:"country" api:"required"`
	Name    string `json:"name" api:"required"`
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
func (r CompanyOwnerLegalPerson) RawJSON() string { return r.JSON.raw }
func (r *CompanyOwnerLegalPerson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyOwnerNaturalPerson struct {
	City        string `json:"city" api:"required"`
	Country     string `json:"country" api:"required"`
	DateOfBirth string `json:"date_of_birth" api:"required"`
	FirstName   string `json:"first_name" api:"required"`
	FullName    string `json:"full_name" api:"required"`
	LastName    string `json:"last_name" api:"required"`
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
func (r CompanyOwnerNaturalPerson) RawJSON() string { return r.JSON.raw }
func (r *CompanyOwnerNaturalPerson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyPurpose struct {
	// Official description of the company's business activities and objectives. This
	// is the registered purpose as stated in official documents.
	Purpose string `json:"purpose" api:"required"`
	// Date when this purpose became effective. Format: ISO 8601 (YYYY-MM-DD) Example:
	// "2022-01-01"
	StartDate string `json:"start_date" api:"required"`
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
	RegisterCourt string `json:"register_court" api:"required"`
	// Registration number in the company register. Example: "230633"
	RegisterNumber string `json:"register_number" api:"required"`
	// Type of register where the company is recorded. Example: "HRB" (Commercial
	// Register B)
	//
	// Any of "HRB", "HRA", "PR", "GnR", "VR".
	RegisterType CompanyRegisterType `json:"register_type" api:"required"`
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

// Report row with values from multiple report periods
type MergedReportRow struct {
	Children      []MergedReportRow `json:"children" api:"required"`
	FormattedName string            `json:"formatted_name" api:"required"`
	Name          string            `json:"name" api:"required"`
	// Report end date to value mapping (ISO date string as key)
	Values map[string]int64 `json:"values" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Children      respjson.Field
		FormattedName respjson.Field
		Name          respjson.Field
		Values        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MergedReportRow) RawJSON() string { return r.JSON.raw }
func (r *MergedReportRow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Report table with data merged across multiple report periods
type MergedReportTable struct {
	Rows []MergedReportRow `json:"rows" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Rows        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MergedReportTable) RawJSON() string { return r.JSON.raw }
func (r *MergedReportTable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReportRow struct {
	Children      []ReportRow `json:"children" api:"required"`
	CurrentValue  int64       `json:"current_value" api:"required"`
	FormattedName string      `json:"formatted_name" api:"required"`
	Name          string      `json:"name" api:"required"`
	PreviousValue int64       `json:"previous_value" api:"required"`
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

type ReportTable struct {
	Rows []ReportRow `json:"rows" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Rows        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReportTable) RawJSON() string { return r.JSON.raw }
func (r *ReportTable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RepresentationRole string

const (
	RepresentationRoleDirector               RepresentationRole = "DIRECTOR"
	RepresentationRoleProkura                RepresentationRole = "PROKURA"
	RepresentationRoleShareholder            RepresentationRole = "SHAREHOLDER"
	RepresentationRoleOwner                  RepresentationRole = "OWNER"
	RepresentationRolePartner                RepresentationRole = "PARTNER"
	RepresentationRolePersonalLiableDirector RepresentationRole = "PERSONAL_LIABLE_DIRECTOR"
	RepresentationRoleLiquidator             RepresentationRole = "LIQUIDATOR"
	RepresentationRoleOther                  RepresentationRole = "OTHER"
)

type Source struct {
	// Url of the source document. In the form of a presigned url accessible for 30
	// minutes.
	DocumentURL string `json:"document_url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DocumentURL respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Source) RawJSON() string { return r.JSON.raw }
func (r *Source) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyGetContactV0Response struct {
	// Where the contact information was found. Example: "https://openregister.de"
	SourceURL string `json:"source_url" api:"required" format:"uri"`
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
	ID string `json:"id" api:"required"`
	// Companies that were merged into this company (Verschmelzung durch Aufnahme, as
	// the acquiring entity).
	Acquisitions []CompanyGetDetailsV1ResponseAcquisition `json:"acquisitions" api:"required"`
	// Current registered address of the company.
	Address CompanyAddress `json:"address" api:"required"`
	// Historical addresses. Shows how the company address changed over time.
	Addresses []CompanyAddress `json:"addresses" api:"required"`
	// Spin-offs (Ausgliederung, § 123 Abs. 3 UmwG) in which this company transferred
	// assets to another company as the transferring entity.
	AssetSpinOffs []CompanyGetDetailsV1ResponseAssetSpinOff `json:"asset_spin_offs" api:"required"`
	// Current registered capital of the company.
	Capital CompanyCapital `json:"capital" api:"required"`
	// Historical capital changes. Shows how the company capital changed over time.
	Capitals []CompanyCapital `json:"capitals" api:"required"`
	// Contact information of the company.
	Contact CompanyGetDetailsV1ResponseContact `json:"contact" api:"required"`
	// Available official documents related to the company.
	Documents []CompanyDocument `json:"documents" api:"required"`
	// Date when the company was officially registered. Format: ISO 8601 (YYYY-MM-DD)
	// Example: "2022-01-01"
	IncorporatedAt string `json:"incorporated_at" api:"required"`
	// Key company indicators like net income, employee count, revenue, etc..
	Indicators []CompanyGetDetailsV1ResponseIndicator `json:"indicators" api:"required"`
	// Industry codes of the company.
	IndustryCodes CompanyGetDetailsV1ResponseIndustryCodes `json:"industry_codes" api:"required"`
	// Legal form of the company. Example: "gmbh" for Gesellschaft mit beschränkter
	// Haftung
	//
	// Any of "ag", "eg", "ek", "ev", "ewiv", "foreign", "gbr", "ggmbh", "gmbh", "kg",
	// "kgaa", "unknown", "llp", "municipal", "ohg", "se", "ug".
	LegalForm CompanyLegalForm `json:"legal_form" api:"required"`
	// If the company ceased to exist through a merger (Verschmelzung), the company it
	// was merged into.
	MergedInto CompanyGetDetailsV1ResponseMergedInto `json:"merged_into" api:"required"`
	// Current official name of the company.
	Name CompanyName `json:"name" api:"required"`
	// Historical company names. Shows how the company name changed over time.
	Names []CompanyName `json:"names" api:"required"`
	// Date of the notarized company agreement (Gesellschaftsvertrag or Satzung).
	// Format: ISO 8601 (YYYY-MM-DD) Example: "2021-12-21"
	NotarizedAt string `json:"notarized_at" api:"required"`
	// The company's current profit and loss transfer agreement
	// (Gewinnabführungsvertrag), if one exists. The referenced company is the parent
	// receiving this company's profit (Organträger). Null if the company has no active
	// agreement.
	ProfitTransferAgreement CompanyGetDetailsV1ResponseProfitTransferAgreement `json:"profit_transfer_agreement" api:"required"`
	// Current official business purpose of the company.
	Purpose CompanyPurpose `json:"purpose" api:"required"`
	// Historical business purposes. Shows how the company purpose changed over time.
	Purposes []CompanyPurpose `json:"purposes" api:"required"`
	// Current registration information of the company.
	Register CompanyRegister `json:"register" api:"required"`
	// Historical registration changes. Shows how registration details changed over
	// time.
	Registers []CompanyRegister `json:"registers" api:"required"`
	// List of individuals or entities authorized to represent the company. Includes
	// directors, officers, and authorized signatories.
	Representation []CompanyGetDetailsV1ResponseRepresentation `json:"representation" api:"required"`
	// The company's current general representation rule (allgemeine
	// Vertretungsregelung), as published in the register. Example: "Ist nur ein
	// Geschäftsführer bestellt, so vertritt er die Gesellschaft allein. Sind mehrere
	// Geschäftsführer bestellt, so wird die Gesellschaft durch zwei Geschäftsführer
	// oder durch einen Geschäftsführer gemeinsam mit einem Prokuristen vertreten."
	RepresentationRule string `json:"representation_rule" api:"required"`
	// Sources of the company data.
	Sources []Source `json:"sources" api:"required"`
	// Current status of the company:
	//
	// - active: Operating normally
	// - inactive: No longer operating
	// - liquidation: In the process of being dissolved
	//
	// Any of "active", "inactive", "liquidation".
	Status CompanyGetDetailsV1ResponseStatus `json:"status" api:"required"`
	// Date when the company was officially terminated (if applicable). Format: ISO
	// 8601 (YYYY-MM-DD) Example: "2024-01-01"
	TerminatedAt string `json:"terminated_at" api:"required"`
	// Insolvency proceedings of the company, if any. Contains basic information per
	// proceeding; use the insolvency endpoint to retrieve all events of a proceeding.
	Insolvencies []CompanyGetDetailsV1ResponseInsolvency `json:"insolvencies"`
	// Legal Entity Identifier (LEI), if available.
	Lei string `json:"lei"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		Acquisitions            respjson.Field
		Address                 respjson.Field
		Addresses               respjson.Field
		AssetSpinOffs           respjson.Field
		Capital                 respjson.Field
		Capitals                respjson.Field
		Contact                 respjson.Field
		Documents               respjson.Field
		IncorporatedAt          respjson.Field
		Indicators              respjson.Field
		IndustryCodes           respjson.Field
		LegalForm               respjson.Field
		MergedInto              respjson.Field
		Name                    respjson.Field
		Names                   respjson.Field
		NotarizedAt             respjson.Field
		ProfitTransferAgreement respjson.Field
		Purpose                 respjson.Field
		Purposes                respjson.Field
		Register                respjson.Field
		Registers               respjson.Field
		Representation          respjson.Field
		RepresentationRule      respjson.Field
		Sources                 respjson.Field
		Status                  respjson.Field
		TerminatedAt            respjson.Field
		Insolvencies            respjson.Field
		Lei                     respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetDetailsV1Response) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetDetailsV1Response) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyGetDetailsV1ResponseAcquisition struct {
	// Date the underlying contract (Verschmelzungsvertrag) was concluded, as cited in
	// the register entry. Null when the register text does not cite a contract date.
	// Entries sharing an agreement_date belong to the same transaction. Format: ISO
	// 8601 (YYYY-MM-DD)
	AgreementDate string `json:"agreement_date" api:"required"`
	// Unique company identifier of the company that was merged into this company.
	// Example: DE-HRB-F1103-267645
	CompanyID string `json:"company_id" api:"required"`
	// Current name of the company that was merged into this company.
	Name string `json:"name" api:"required"`
	// Date the merger was registered. Format: ISO 8601 (YYYY-MM-DD)
	RegistrationDate string `json:"registration_date" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgreementDate    respjson.Field
		CompanyID        respjson.Field
		Name             respjson.Field
		RegistrationDate respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetDetailsV1ResponseAcquisition) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetDetailsV1ResponseAcquisition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyGetDetailsV1ResponseAssetSpinOff struct {
	// Date the underlying contract (Ausgliederungsvertrag) was concluded, as cited in
	// the register entry. Null when the register text does not cite a contract date.
	// Entries sharing an agreement_date belong to the same transaction. Format: ISO
	// 8601 (YYYY-MM-DD)
	AgreementDate string `json:"agreement_date" api:"required"`
	// Unique company identifier of the company that received the assets. Example:
	// DE-HRB-F1103-267645
	CompanyID string `json:"company_id" api:"required"`
	// Current name of the company that received the assets.
	Name string `json:"name" api:"required"`
	// Date the spin-off was registered. Format: ISO 8601 (YYYY-MM-DD)
	RegistrationDate string `json:"registration_date" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgreementDate    respjson.Field
		CompanyID        respjson.Field
		Name             respjson.Field
		RegistrationDate respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetDetailsV1ResponseAssetSpinOff) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetDetailsV1ResponseAssetSpinOff) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contact information of the company.
type CompanyGetDetailsV1ResponseContact struct {
	SocialMedia CompanyGetDetailsV1ResponseContactSocialMedia `json:"social_media" api:"required"`
	WebsiteURL  string                                        `json:"website_url" api:"required"`
	Email       string                                        `json:"email"`
	Phone       string                                        `json:"phone"`
	VatID       string                                        `json:"vat_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SocialMedia respjson.Field
		WebsiteURL  respjson.Field
		Email       respjson.Field
		Phone       respjson.Field
		VatID       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetDetailsV1ResponseContact) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetDetailsV1ResponseContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyGetDetailsV1ResponseContactSocialMedia struct {
	Facebook  string `json:"facebook"`
	GitHub    string `json:"github"`
	Instagram string `json:"instagram"`
	Linkedin  string `json:"linkedin"`
	Tiktok    string `json:"tiktok"`
	Twitter   string `json:"twitter"`
	Xing      string `json:"xing"`
	Youtube   string `json:"youtube"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Facebook    respjson.Field
		GitHub      respjson.Field
		Instagram   respjson.Field
		Linkedin    respjson.Field
		Tiktok      respjson.Field
		Twitter     respjson.Field
		Xing        respjson.Field
		Youtube     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetDetailsV1ResponseContactSocialMedia) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetDetailsV1ResponseContactSocialMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A focused subset of the key company indicators for a given year. Values of the
// indicator are given in the smallest currency unit (cents). Example: 2099
// represents €20.99 for monetary values. For non-monetary values (e.g.,
// employees), the actual number.
type CompanyGetDetailsV1ResponseIndicator struct {
	// The balance sheet total of that year (in cents).
	BalanceSheetTotal int64 `json:"balance_sheet_total" api:"required"`
	// The capital reserves of that year (in cents).
	CapitalReserves int64 `json:"capital_reserves" api:"required"`
	// The cash of that year (in cents).
	Cash int64 `json:"cash" api:"required"`
	// Date to which this financial indicators apply. Format: ISO 8601 (YYYY-MM-DD)
	// Example: "2022-01-01"
	Date string `json:"date" api:"required"`
	// The number of employees of that year.
	Employees int64 `json:"employees" api:"required"`
	// The equity of that year (in cents).
	Equity int64 `json:"equity" api:"required"`
	// The liabilities of that year (in cents).
	Liabilities int64 `json:"liabilities" api:"required"`
	// The materials of that year (in cents).
	Materials int64 `json:"materials" api:"required"`
	// The net income of that year (in cents).
	NetIncome int64 `json:"net_income" api:"required"`
	// The pension provisions of that year (in cents).
	PensionProvisions int64 `json:"pension_provisions" api:"required"`
	// The real estate of that year (in cents).
	RealEstate int64 `json:"real_estate" api:"required"`
	// The report id (source) of the indicators.
	ReportID string `json:"report_id" api:"required" format:"uuid"`
	// The revenue of that year (in cents).
	Revenue int64 `json:"revenue" api:"required"`
	// The salaries of that year (in cents).
	Salaries int64 `json:"salaries" api:"required"`
	// The taxes of that year (in cents).
	Taxes int64 `json:"taxes" api:"required"`
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
	Wz2025 []CompanyGetDetailsV1ResponseIndustryCodesWz2025 `json:"WZ2025" api:"required"`
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
	Code string `json:"code" api:"required"`
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

// If the company ceased to exist through a merger (Verschmelzung), the company it
// was merged into.
type CompanyGetDetailsV1ResponseMergedInto struct {
	// Date the underlying contract (Verschmelzungsvertrag) was concluded, as cited in
	// the register entry. Null when the register text does not cite a contract date.
	// Entries sharing an agreement_date belong to the same transaction. Format: ISO
	// 8601 (YYYY-MM-DD)
	AgreementDate string `json:"agreement_date" api:"required"`
	// Unique company identifier of the company this company was merged into. Example:
	// DE-HRB-F1103-267645
	CompanyID string `json:"company_id" api:"required"`
	// Current name of the company this company was merged into.
	Name string `json:"name" api:"required"`
	// Date the merger was registered. Format: ISO 8601 (YYYY-MM-DD)
	RegistrationDate string `json:"registration_date" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgreementDate    respjson.Field
		CompanyID        respjson.Field
		Name             respjson.Field
		RegistrationDate respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetDetailsV1ResponseMergedInto) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetDetailsV1ResponseMergedInto) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The company's current profit and loss transfer agreement
// (Gewinnabführungsvertrag), if one exists. The referenced company is the parent
// receiving this company's profit (Organträger). Null if the company has no active
// agreement.
type CompanyGetDetailsV1ResponseProfitTransferAgreement struct {
	// Date the underlying contract (Gewinnabführungsvertrag) was concluded, as cited
	// in the register entry. Null when the register text does not cite a contract
	// date. Entries sharing an agreement_date belong to the same transaction. Format:
	// ISO 8601 (YYYY-MM-DD)
	AgreementDate string `json:"agreement_date" api:"required"`
	// Unique company identifier of the parent company receiving this company's profit
	// (Organträger). Example: DE-HRB-F1103-267645
	CompanyID string `json:"company_id" api:"required"`
	// Current name of the parent company.
	Name string `json:"name" api:"required"`
	// Date the agreement was registered. Format: ISO 8601 (YYYY-MM-DD)
	RegistrationDate string `json:"registration_date" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgreementDate    respjson.Field
		CompanyID        respjson.Field
		Name             respjson.Field
		RegistrationDate respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetDetailsV1ResponseProfitTransferAgreement) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetDetailsV1ResponseProfitTransferAgreement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyGetDetailsV1ResponseRepresentation struct {
	// Unique identifier for the representative. For companies: Format matches
	// company_id pattern For individuals: UUID Example: "DE-HRB-F1103-267645" or UUID
	// May be null for certain representatives.
	ID string `json:"id" api:"required"`
	// The representative's current individual representation authority (individuelle
	// Vertretungsbefugnis), as published in the register. Null if no special authority
	// is recorded. Example: "einzelvertretungsberechtigt mit der Befugnis, im Namen
	// der Gesellschaft mit sich im eigenen Namen Rechtsgeschäfte abzuschließen"
	Authority string `json:"authority" api:"required"`
	// Date when this representative role ended (if applicable). Format: ISO 8601
	// (YYYY-MM-DD) Example: "2022-01-01"
	EndDate string `json:"end_date" api:"required"`
	// The name of the representative. E.g. "Max Mustermann" or "Max Mustermann GmbH"
	Name string `json:"name" api:"required"`
	// The role of the representation. E.g. "DIRECTOR"
	//
	// Any of "DIRECTOR", "PROKURA", "SHAREHOLDER", "OWNER", "PARTNER",
	// "PERSONAL_LIABLE_DIRECTOR", "LIQUIDATOR", "OTHER".
	Role RepresentationRole `json:"role" api:"required"`
	// Date when this representative role became effective. Format: ISO 8601
	// (YYYY-MM-DD) Example: "2022-01-01"
	StartDate string `json:"start_date" api:"required"`
	// Whether the representation is a natural person or a legal entity.
	//
	// Any of "natural_person", "legal_person".
	Type          EntityType                                             `json:"type" api:"required"`
	LegalPerson   CompanyGetDetailsV1ResponseRepresentationLegalPerson   `json:"legal_person" api:"nullable"`
	NaturalPerson CompanyGetDetailsV1ResponseRepresentationNaturalPerson `json:"natural_person" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Authority     respjson.Field
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
	City string `json:"city" api:"required"`
	// Country where the representative is located, in ISO 3166-1 alpha-2 format.
	// Example: "DE" for Germany
	Country string `json:"country" api:"required"`
	Name    string `json:"name" api:"required"`
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
	City string `json:"city" api:"required"`
	// Date of birth of the representative. May still be null for natural persons if it
	// is not available. Format: ISO 8601 (YYYY-MM-DD) Example: "1990-01-01"
	DateOfBirth string `json:"date_of_birth" api:"required"`
	// First name of the representative. Example: "Max"
	FirstName string `json:"first_name" api:"required"`
	// Last name of the representative. Example: "Mustermann"
	LastName string `json:"last_name" api:"required"`
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

// Basic information about an insolvency proceeding of the company. Use the
// insolvency endpoint to retrieve all events of the proceeding.
type CompanyGetDetailsV1ResponseInsolvency struct {
	// Unique identifier of the insolvency proceeding.
	ID string `json:"id" api:"required" format:"uuid"`
	// Case number of the proceeding at the court. Example: "36d IN 3382/25"
	CaseNumber string `json:"case_number" api:"required"`
	// Insolvency court handling the proceeding.
	Court string `json:"court" api:"required"`
	// Current status of the insolvency proceeding.
	//
	// Any of "preliminary", "opened", "rejected_no_assets", "mass_insufficient",
	// "plan_supervised", "lifted", "discontinued", "discharge_pending",
	// "discharge_granted", "discharge_denied", "discharge_revoked", "unknown".
	CurrentStatus InsolvencyStatus `json:"current_status" api:"required"`
	// Kind of administration ordered for the proceeding.
	//
	// Any of "external_administration", "self_administration", "protective_shield".
	AdministrationKind InsolvencyAdministrationKind `json:"administration_kind" api:"nullable"`
	// Date the proceeding was closed.
	ClosedAt time.Time `json:"closed_at" api:"nullable" format:"date-time"`
	// Date the proceeding was opened.
	OpenedAt time.Time `json:"opened_at" api:"nullable" format:"date-time"`
	// Kind of insolvency proceeding.
	//
	// Any of "regular_insolvency", "consumer_insolvency".
	ProceedingKind InsolvencyProceedingKind `json:"proceeding_kind" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CaseNumber         respjson.Field
		Court              respjson.Field
		CurrentStatus      respjson.Field
		AdministrationKind respjson.Field
		ClosedAt           respjson.Field
		OpenedAt           respjson.Field
		ProceedingKind     respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetDetailsV1ResponseInsolvency) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetDetailsV1ResponseInsolvency) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyGetFinancialsV1Response struct {
	// Key financial indicators per fiscal year, sorted by date (latest first).
	Indicators []CompanyGetFinancialsV1ResponseIndicator `json:"indicators" api:"required"`
	// Financial data merged across all available report periods
	Merged  CompanyGetFinancialsV1ResponseMerged   `json:"merged" api:"required"`
	Reports []CompanyGetFinancialsV1ResponseReport `json:"reports" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Indicators  respjson.Field
		Merged      respjson.Field
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

// The indicators of the company for a given year. Values of the indicator are
// given in the smallest currency unit (cents). Example: 2099 represents €20.99 for
// monetary values For non-monetary values (e.g., employees), the actual number.
type CompanyGetFinancialsV1ResponseIndicator struct {
	// The active accruals of that year (in cents).
	ActiveAccruals int64 `json:"active_accruals" api:"required"`
	// The affiliated liabilities of that year (in cents).
	AffiliatedLiabilities int64 `json:"affiliated_liabilities" api:"required"`
	// The balance sheet total of that year (in cents).
	BalanceSheetTotal int64 `json:"balance_sheet_total" api:"required"`
	// The bank debt of that year (in cents).
	BankDebt int64 `json:"bank_debt" api:"required"`
	// The capital reserves of that year (in cents).
	CapitalReserves int64 `json:"capital_reserves" api:"required"`
	// The cash of that year (in cents).
	Cash int64 `json:"cash" api:"required"`
	// Commission expense (Provisionsaufwendungen) of that year (in cents).
	CommissionExpense int64 `json:"commission_expense" api:"required"`
	// Commission income (Provisionserträge) of that year (in cents).
	CommissionIncome int64 `json:"commission_income" api:"required"`
	// The current assets of that year (in cents).
	CurrentAssets int64 `json:"current_assets" api:"required"`
	// Date to which this financial indicators apply. Format: ISO 8601 (YYYY-MM-DD)
	// Example: "2022-01-01"
	Date string `json:"date" api:"required"`
	// The earnings before interest and taxes of that year (in cents).
	Ebit int64 `json:"ebit" api:"required"`
	// The earnings before interest, taxes, depreciation, and amortization of that year
	// (in cents).
	Ebitda int64 `json:"ebitda" api:"required"`
	// The number of employees of that year.
	Employees int64 `json:"employees" api:"required"`
	// The equity of that year (in cents).
	Equity int64 `json:"equity" api:"required"`
	// The financial assets of that year (in cents).
	FinancialAssets int64 `json:"financial_assets" api:"required"`
	// The financial debt of that year (in cents).
	FinancialDebt int64 `json:"financial_debt" api:"required"`
	// The signed financial asset depreciation, write-down, or reversal of that year
	// (in cents).
	FinancialDepreciation int64 `json:"financial_depreciation" api:"required"`
	// The fixed assets of that year (in cents).
	FixedAssets int64 `json:"fixed_assets" api:"required"`
	// The gross profit (Rohergebnis) of that year (in cents).
	GrossProfit int64 `json:"gross_profit" api:"required"`
	// The income after income taxes of that year (in cents).
	IncomeAfterTax int64 `json:"income_after_tax" api:"required"`
	// The income before income taxes of that year (in cents).
	IncomeBeforeTax int64 `json:"income_before_tax" api:"required"`
	// The intangible assets of that year (in cents).
	IntangibleAssets int64 `json:"intangible_assets" api:"required"`
	// The interest expense of that year (in cents).
	InterestExpense int64 `json:"interest_expense" api:"required"`
	// The interest income of that year (in cents).
	InterestIncome int64 `json:"interest_income" api:"required"`
	// The inventory of that year (in cents).
	Inventory int64 `json:"inventory" api:"required"`
	// The liabilities of that year (in cents).
	Liabilities int64 `json:"liabilities" api:"required"`
	// The materials of that year (in cents).
	Materials int64 `json:"materials" api:"required"`
	// The net income of that year (in cents).
	NetIncome int64 `json:"net_income" api:"required"`
	// The operating depreciation and amortization of that year (in cents).
	OperatingDepreciation int64 `json:"operating_depreciation" api:"required"`
	// The other liabilities of that year (in cents).
	OtherLiabilities int64 `json:"other_liabilities" api:"required"`
	// The other operating expenses of that year (in cents).
	OtherOperatingExpenses int64 `json:"other_operating_expenses" api:"required"`
	// The other operating income of that year (in cents).
	OtherOperatingIncome int64 `json:"other_operating_income" api:"required"`
	// The other provisions of that year (in cents).
	OtherProvisions int64 `json:"other_provisions" api:"required"`
	// Other taxes (Sonstige Steuern) of that year (in cents).
	OtherTaxes int64 `json:"other_taxes" api:"required"`
	// The parent-attributed net income of that year (in cents).
	ParentNetIncome int64 `json:"parent_net_income" api:"required"`
	// The passive accruals of that year (in cents).
	PassiveAccruals int64 `json:"passive_accruals" api:"required"`
	// The pension provisions of that year (in cents).
	PensionProvisions int64 `json:"pension_provisions" api:"required"`
	// The profit carryforward of that year (in cents).
	ProfitCarryforward int64 `json:"profit_carryforward" api:"required"`
	// The provisions of that year (in cents).
	Provisions int64 `json:"provisions" api:"required"`
	// The real estate of that year (in cents).
	RealEstate int64 `json:"real_estate" api:"required"`
	// The receivables of that year (in cents).
	Receivables int64 `json:"receivables" api:"required"`
	// The report id (source) of the indicators.
	ReportID string `json:"report_id" api:"required" format:"uuid"`
	// The retained earnings of that year (in cents).
	RetainedEarnings int64 `json:"retained_earnings" api:"required"`
	// The revenue of that year (in cents).
	Revenue int64 `json:"revenue" api:"required"`
	// The salaries of that year (in cents).
	Salaries int64 `json:"salaries" api:"required"`
	// The shareholder liabilities of that year (in cents).
	ShareholderLiabilities int64 `json:"shareholder_liabilities" api:"required"`
	// The tangible assets of that year (in cents).
	TangibleAssets int64 `json:"tangible_assets" api:"required"`
	// The taxes of that year (in cents).
	Taxes int64 `json:"taxes" api:"required"`
	// The trade payables of that year (in cents).
	TradePayables int64 `json:"trade_payables" api:"required"`
	// The trade receivables of that year (in cents).
	TradeReceivables int64 `json:"trade_receivables" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActiveAccruals         respjson.Field
		AffiliatedLiabilities  respjson.Field
		BalanceSheetTotal      respjson.Field
		BankDebt               respjson.Field
		CapitalReserves        respjson.Field
		Cash                   respjson.Field
		CommissionExpense      respjson.Field
		CommissionIncome       respjson.Field
		CurrentAssets          respjson.Field
		Date                   respjson.Field
		Ebit                   respjson.Field
		Ebitda                 respjson.Field
		Employees              respjson.Field
		Equity                 respjson.Field
		FinancialAssets        respjson.Field
		FinancialDebt          respjson.Field
		FinancialDepreciation  respjson.Field
		FixedAssets            respjson.Field
		GrossProfit            respjson.Field
		IncomeAfterTax         respjson.Field
		IncomeBeforeTax        respjson.Field
		IntangibleAssets       respjson.Field
		InterestExpense        respjson.Field
		InterestIncome         respjson.Field
		Inventory              respjson.Field
		Liabilities            respjson.Field
		Materials              respjson.Field
		NetIncome              respjson.Field
		OperatingDepreciation  respjson.Field
		OtherLiabilities       respjson.Field
		OtherOperatingExpenses respjson.Field
		OtherOperatingIncome   respjson.Field
		OtherProvisions        respjson.Field
		OtherTaxes             respjson.Field
		ParentNetIncome        respjson.Field
		PassiveAccruals        respjson.Field
		PensionProvisions      respjson.Field
		ProfitCarryforward     respjson.Field
		Provisions             respjson.Field
		RealEstate             respjson.Field
		Receivables            respjson.Field
		ReportID               respjson.Field
		RetainedEarnings       respjson.Field
		Revenue                respjson.Field
		Salaries               respjson.Field
		ShareholderLiabilities respjson.Field
		TangibleAssets         respjson.Field
		Taxes                  respjson.Field
		TradePayables          respjson.Field
		TradeReceivables       respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetFinancialsV1ResponseIndicator) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetFinancialsV1ResponseIndicator) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Financial data merged across all available report periods
type CompanyGetFinancialsV1ResponseMerged struct {
	// Report table with data merged across multiple report periods
	Aktiva MergedReportTable `json:"aktiva" api:"required"`
	// Report table with data merged across multiple report periods
	Passiva MergedReportTable `json:"passiva" api:"required"`
	// Report table with data merged across multiple report periods
	Guv MergedReportTable `json:"guv"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Aktiva      respjson.Field
		Passiva     respjson.Field
		Guv         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetFinancialsV1ResponseMerged) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetFinancialsV1ResponseMerged) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyGetFinancialsV1ResponseReport struct {
	Aktiva ReportTable `json:"aktiva" api:"required"`
	// Whether the report is a consolidated report or not.
	Consolidated  bool        `json:"consolidated" api:"required"`
	Passiva       ReportTable `json:"passiva" api:"required"`
	ReportEndDate time.Time   `json:"report_end_date" api:"required" format:"date"`
	// Unique identifier for the financial report. Example:
	// f47ac10b-58cc-4372-a567-0e02b2c3d479
	ReportID        string    `json:"report_id" api:"required"`
	ReportStartDate time.Time `json:"report_start_date" api:"required" format:"date"`
	// Sources of the report data. Presigned URLs accessible for 30 minutes.
	Sources []CompanyGetFinancialsV1ResponseReportSource `json:"sources" api:"required"`
	Guv     ReportTable                                  `json:"guv" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Aktiva          respjson.Field
		Consolidated    respjson.Field
		Passiva         respjson.Field
		ReportEndDate   respjson.Field
		ReportID        respjson.Field
		ReportStartDate respjson.Field
		Sources         respjson.Field
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

type CompanyGetFinancialsV1ResponseReportSource struct {
	// Url of the rendered HTML report. In the form of a presigned url accessible for
	// 30 minutes.
	HTMLURL string `json:"html_url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HTMLURL     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetFinancialsV1ResponseReportSource) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetFinancialsV1ResponseReportSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyGetHistoricalOwnersV1Response struct {
	Owners []CompanyGetHistoricalOwnersV1ResponseOwner `json:"owners" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Owners      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetHistoricalOwnersV1Response) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetHistoricalOwnersV1Response) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyGetHistoricalOwnersV1ResponseOwner struct {
	// Unique identifier for the owner. For companies, this is the company register ID
	// (e.g. DE-HRB-F1103-267645) which can be used to look up the company. For natural
	// persons, this is the entity UUID. For other entity types (foreign companies,
	// foundations, etc.), this is empty.
	ID string `json:"id" api:"required"`
	// Type of the owner entity
	//
	// Any of "natural_person", "german_company", "foreign_company",
	// "german_government_entity", "german_foundation", "german_multiple_shareholder".
	EntityType string `json:"entity_type" api:"required"`
	// Date when this owner first appeared
	FirstAppearance time.Time `json:"first_appearance" api:"required" format:"date-time"`
	// Name of the owner
	Name string `json:"name" api:"required"`
	// Historical ownership data across all documents
	OwnershipHistory []CompanyGetHistoricalOwnersV1ResponseOwnerOwnershipHistory `json:"ownership_history" api:"required"`
	// Current status of the owner
	//
	// Any of "active", "removed".
	Status string `json:"status" api:"required"`
	// Country of the owner
	Country string `json:"country"`
	// Date when this owner last appeared (null if still active)
	LastAppearance time.Time `json:"last_appearance" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		EntityType       respjson.Field
		FirstAppearance  respjson.Field
		Name             respjson.Field
		OwnershipHistory respjson.Field
		Status           respjson.Field
		Country          respjson.Field
		LastAppearance   respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetHistoricalOwnersV1ResponseOwner) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetHistoricalOwnersV1ResponseOwner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyGetHistoricalOwnersV1ResponseOwnerOwnershipHistory struct {
	// Date of the document
	DocumentDate time.Time `json:"document_date" api:"required" format:"date-time"`
	// Document where this ownership data was found
	DocumentID string `json:"document_id" api:"required" format:"uuid"`
	// Nominal value of shares in this document
	NominalShares int64 `json:"nominal_shares" api:"required"`
	// Percentage ownership in this document
	PercentageShares float64 `json:"percentage_shares" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DocumentDate     respjson.Field
		DocumentID       respjson.Field
		NominalShares    respjson.Field
		PercentageShares respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetHistoricalOwnersV1ResponseOwnerOwnershipHistory) RawJSON() string {
	return r.JSON.raw
}
func (r *CompanyGetHistoricalOwnersV1ResponseOwnerOwnershipHistory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Companies this entity owns or has invested in.
type CompanyGetHoldingsV1Response struct {
	// Unique company identifier. Example: DE-HRB-F1103-267645
	CompanyID string                                `json:"company_id" api:"required"`
	Holdings  []CompanyGetHoldingsV1ResponseHolding `json:"holdings" api:"required"`
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
	CompanyID string `json:"company_id" api:"required"`
	// Date when the ownership ended. Format: ISO 8601 (YYYY-MM-DD) Example:
	// "2022-01-01"
	End string `json:"end" api:"required"`
	// Name of the company.
	Name string `json:"name" api:"required"`
	// Amount of shares or capital in the company. Example: 100
	NominalShare float64 `json:"nominal_share" api:"required"`
	// Share of the company. Example: 0.5 represents 50% ownership
	PercentageShare float64 `json:"percentage_share" api:"required"`
	// Type of relationship between the entity and the company.
	//
	// Any of "shareholder", "stockholder", "limited_partner", "general_partner".
	RelationType CompanyRelationType `json:"relation_type" api:"required"`
	// Date when the ownership started. Format: ISO 8601 (YYYY-MM-DD) Example:
	// "2022-01-01"
	Start string `json:"start" api:"required"`
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
	// When true, the returned owner data is the best available but may not reflect the
	// most current ownership state. This applies to AG and SE companies where
	// ownership data is sourced from Handelsregister decision and articles of
	// association documents, which are not filed on every ownership change.
	BestAvailable bool `json:"best_available" api:"required"`
	// Unique company identifier. Example: DE-HRB-F1103-267645
	CompanyID string                            `json:"company_id" api:"required"`
	Owners    []CompanyGetOwnersV1ResponseOwner `json:"owners" api:"required"`
	// Sources of the company owners data.
	Sources []Source `json:"sources" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BestAvailable respjson.Field
		CompanyID     respjson.Field
		Owners        respjson.Field
		Sources       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
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
	ID string `json:"id" api:"required"`
	// Details about the legal person.
	LegalPerson CompanyOwnerLegalPerson `json:"legal_person" api:"required"`
	// The name of the shareholder. E.g. "Max Mustermann" or "Max Mustermann GmbH"
	Name string `json:"name" api:"required"`
	// Details about the natural person.
	NaturalPerson CompanyOwnerNaturalPerson `json:"natural_person" api:"required"`
	// Nominal value of shares in Euro. Example: 100
	NominalShare float64 `json:"nominal_share" api:"required"`
	// Percentage of company ownership. Example: 5.36 represents 5.36% ownership
	PercentageShare float64 `json:"percentage_share" api:"required"`
	// Type of relationship between the entity and the company.
	//
	// Any of "shareholder", "stockholder", "limited_partner", "general_partner".
	RelationType CompanyRelationType `json:"relation_type" api:"required"`
	// Date when the relation started. Only available for some types of owners. Format:
	// ISO 8601 (YYYY-MM-DD) Example: "2022-01-01"
	Start string `json:"start" api:"required"`
	// The type of shareholder.
	//
	// Any of "natural_person", "legal_person".
	Type EntityType `json:"type" api:"required"`
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

type CompanyGetUbosV1Response struct {
	// Unique company identifier. Example: DE-HRB-F1103-267645
	CompanyID string                        `json:"company_id" api:"required"`
	Ubos      []CompanyGetUbosV1ResponseUbo `json:"ubos" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompanyID   respjson.Field
		Ubos        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetUbosV1Response) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetUbosV1Response) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyGetUbosV1ResponseUbo struct {
	// Unique identifier for the shareholder. For individuals: UUID For companies:
	// Format matches company_id pattern Example: "DE-HRB-F1103-267645" or UUID May be
	// null for certain shareholders.
	ID          string                  `json:"id" api:"required"`
	LegalPerson CompanyOwnerLegalPerson `json:"legal_person" api:"required"`
	// Maximum percentage of company ownership. Example: 5.36 represents maximum of
	// 5.36% ownership There is no exact percentage share for owners that hold a stake
	// as or through a limited partner. For these owners, we can only show the maximum
	// percentage share they could have based on their deposit as a limited partner. Is
	// null for all owners that have an exact percentage share or owners that hold a
	// stake as or through a personal liable director.
	MaxPercentageShare float64 `json:"max_percentage_share" api:"required"`
	// The name of the shareholder. E.g. "Max Mustermann"
	Name          string                    `json:"name" api:"required"`
	NaturalPerson CompanyOwnerNaturalPerson `json:"natural_person" api:"required"`
	// Percentage of company ownership. Example: 5.36 represents 5.36% ownership Is
	// null for all owners that hold a stake as or through a personal liable directors
	// or limited partner.
	PercentageShare float64 `json:"percentage_share" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		LegalPerson        respjson.Field
		MaxPercentageShare respjson.Field
		Name               respjson.Field
		NaturalPerson      respjson.Field
		PercentageShare    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetUbosV1ResponseUbo) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetUbosV1ResponseUbo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyGetDetailsV1Params struct {
	// Setting this to true will return the company without sources.
	Export param.Opt[bool] `query:"export,omitzero" json:"-"`
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
	// When set to true, returns the best available owner data for AG and SE companies.
	// This data is extracted from Handelsregister documents and may not reflect the
	// most current ownership state, as these document types are not filed on every
	// ownership change. Requests for AG/SE companies without this flag return 404.
	// Note: realtime and best_available cannot be used together at the moment.
	BestAvailable param.Opt[bool] `query:"best_available,omitzero" json:"-"`
	// Setting this to true will return the owners of the company if they exist but
	// will skip processing the documents in case they weren't processed yet.
	Export param.Opt[bool] `query:"export,omitzero" json:"-"`
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
