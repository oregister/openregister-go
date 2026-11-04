// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openregister

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/oregister/openregister-go/v2/internal/apijson"
	"github.com/oregister/openregister-go/v2/internal/requestconfig"
	"github.com/oregister/openregister-go/v2/option"
	"github.com/oregister/openregister-go/v2/packages/param"
	"github.com/oregister/openregister-go/v2/packages/respjson"
)

// JobService contains methods and other services that help with interacting with
// the openregister API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewJobService] method instead.
type JobService struct {
	Options []option.RequestOption
}

// NewJobService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewJobService(opts ...option.RequestOption) (r JobService) {
	r = JobService{}
	r.Options = opts
	return
}

// Start a realtime request without waiting for it. The request returns at once
// with a pending job. Use this instead of `realtime=true` when your client cannot
// keep a request open for the minutes the Handelsregister may take.
//
// Job types:
//
//   - `company`: refreshes the company from the Handelsregister, like
//     `GET /v1/company/{company_id}?realtime=true`.
//   - `owners`: refreshes the company's owners, like
//     `GET /v1/company/{company_id}/owners?realtime=true`.
//   - `document`: fetches one document, like `GET /v1/document`. Requires
//     `document_category`.
//
// Follow a job either way, or both:
//
//   - Poll `GET /v1/jobs/{job_id}` until `status` is `completed` or `failed`. A
//     completed job carries its result in `results`: the document for a `document`
//     job, the company or its owners for a `company` or `owners` job.
//   - Configure a webhook endpoint: when the job finishes, OpenRegister sends a
//     `job.completed` or `job.failed` webhook to it. The job can be polled all the
//     same.
//
// Jobs are available on paid plans. A completed job costs what the same request
// with `realtime=true` costs: 20 credits for `company` and `owners`, 10 for
// `document`. Failed jobs and polling are free.
func (r *JobService) NewV1(ctx context.Context, body JobNewV1Params, opts ...option.RequestOption) (res *Job, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/jobs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get a job's status and, once it completed, its result in `results`. Polling is
// free.
//
//   - `document` jobs return `results.document` with a download URL, presigned again
//     on every read.
//   - `company` jobs return `results.company`, the same as
//     `GET /v1/company/{company_id}`.
//   - `owners` jobs return `results.owners`, the same as
//     `GET /v1/company/{company_id}/owners`.
//
// Results are returned for 12 hours after the job finished (`expires_at`); after
// that, `results` is null and the regular endpoints serve the data.
func (r *JobService) GetV1(ctx context.Context, jobID string, opts ...option.RequestOption) (res *Job, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/jobs/%s", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type Job struct {
	// Unique job identifier. Example: f47ac10b-58cc-4372-a567-0e02b2c3d479
	ID string `json:"id" api:"required" format:"uuid"`
	// Company the job refreshes. For `company` and `owners` jobs this is the company's
	// latest register ID, which differs from the requested one when the company moved
	// to another register. Example: DE-HRB-F1103-267645
	CompanyID string    `json:"company_id" api:"required"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Any of "current_printout", "chronological_printout", "historical_printout",
	// "structured_information", "shareholder_list", "articles_of_association".
	DocumentCategory JobDocumentCategory `json:"document_category" api:"required"`
	// Until when a completed job returns `results`, 12 hours after it finished.
	ExpiresAt  time.Time `json:"expires_at" api:"required" format:"date-time"`
	FinishedAt time.Time `json:"finished_at" api:"required" format:"date-time"`
	// Why a failed job failed:
	//
	// - `not_available`: the register holds no such document for this company.
	// - `timeout`: the Handelsregister did not answer in time; try again later.
	// - `internal_error`: any other failure.
	//
	// Any of "not_available", "timeout", "internal_error".
	Reason JobReason `json:"reason" api:"required"`
	// The result of a completed job. Only the field of the job's type is set.
	Results JobResults `json:"results" api:"required"`
	// Any of "pending", "completed", "failed".
	Status JobStatus `json:"status" api:"required"`
	// Any of "company", "owners", "document".
	Type JobType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CompanyID        respjson.Field
		CreatedAt        respjson.Field
		DocumentCategory respjson.Field
		ExpiresAt        respjson.Field
		FinishedAt       respjson.Field
		Reason           respjson.Field
		Results          respjson.Field
		Status           respjson.Field
		Type             respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Job) RawJSON() string { return r.JSON.raw }
func (r *Job) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Why a failed job failed:
//
// - `not_available`: the register holds no such document for this company.
// - `timeout`: the Handelsregister did not answer in time; try again later.
// - `internal_error`: any other failure.
type JobReason string

const (
	JobReasonNotAvailable  JobReason = "not_available"
	JobReasonTimeout       JobReason = "timeout"
	JobReasonInternalError JobReason = "internal_error"
)

// The result of a completed job. Only the field of the job's type is set.
type JobResults struct {
	// The refreshed company of a `company` job.
	Company JobResultsCompany `json:"company" api:"required"`
	// The fetched document of a `document` job.
	Document JobResultsDocument `json:"document" api:"required"`
	// The refreshed owners of an `owners` job.
	Owners JobResultsOwners `json:"owners" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Company     respjson.Field
		Document    respjson.Field
		Owners      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobResults) RawJSON() string { return r.JSON.raw }
func (r *JobResults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The refreshed company of a `company` job.
type JobResultsCompany struct {
	// Unique company identifier. Example: DE-HRB-F1103-267645
	ID string `json:"id" api:"required"`
	// Companies that were merged into this company (Verschmelzung durch Aufnahme, as
	// the acquiring entity).
	Acquisitions []JobResultsCompanyAcquisition `json:"acquisitions" api:"required"`
	// Current registered address of the company.
	Address CompanyAddress `json:"address" api:"required"`
	// Historical addresses. Shows how the company address changed over time.
	Addresses []CompanyAddress `json:"addresses" api:"required"`
	// Spin-offs (Ausgliederung, § 123 Abs. 3 UmwG) in which this company transferred
	// assets to another company as the transferring entity.
	AssetSpinOffs []JobResultsCompanyAssetSpinOff `json:"asset_spin_offs" api:"required"`
	// Current registered capital of the company.
	Capital CompanyCapital `json:"capital" api:"required"`
	// Historical capital changes. Shows how the company capital changed over time.
	Capitals []CompanyCapital `json:"capitals" api:"required"`
	// Contact information of the company.
	Contact JobResultsCompanyContact `json:"contact" api:"required"`
	// Available official documents related to the company.
	Documents []CompanyDocument `json:"documents" api:"required"`
	// Modelled revenue, EBITDA and employees for companies that do not report them.
	// One estimate per metric, for the latest fiscal year.
	Estimates JobResultsCompanyEstimates `json:"estimates" api:"required"`
	// Date when the company was officially registered. Format: ISO 8601 (YYYY-MM-DD)
	// Example: "2022-01-01"
	IncorporatedAt time.Time `json:"incorporated_at" api:"required" format:"date"`
	// Key company indicators like net income, employee count, revenue, etc..
	Indicators []JobResultsCompanyIndicator `json:"indicators" api:"required"`
	// Industry codes of the company, classified from its current purpose
	// (Unternehmensgegenstand). Empty when the purpose has not been classified or
	// describes no codable activity.
	IndustryCodes JobResultsCompanyIndustryCodes `json:"industry_codes" api:"required"`
	// Legal form of the company. Example: "gmbh" for Gesellschaft mit beschränkter
	// Haftung
	//
	// Any of "ag", "eg", "ek", "ev", "ewiv", "foreign", "foundation", "gbr", "ggmbh",
	// "gmbh", "kg", "kgaa", "unknown", "llp", "municipal", "ohg", "se", "ug".
	LegalForm CompanyLegalForm `json:"legal_form" api:"required"`
	// If the company ceased to exist through a merger (Verschmelzung), the company it
	// was merged into.
	MergedInto JobResultsCompanyMergedInto `json:"merged_into" api:"required"`
	// Current official name of the company.
	Name CompanyName `json:"name" api:"required"`
	// Historical company names. Shows how the company name changed over time.
	Names []CompanyName `json:"names" api:"required"`
	// Date of the notarized company agreement (Gesellschaftsvertrag or Satzung).
	// Format: ISO 8601 (YYYY-MM-DD) Example: "2021-12-21"
	NotarizedAt time.Time `json:"notarized_at" api:"required" format:"date"`
	// The company's current profit and loss transfer agreement
	// (Gewinnabführungsvertrag), if one exists. The referenced company is the parent
	// receiving this company's profit (Organträger). Null if the company has no active
	// agreement.
	ProfitTransferAgreement JobResultsCompanyProfitTransferAgreement `json:"profit_transfer_agreement" api:"required"`
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
	Representation []JobResultsCompanyRepresentation `json:"representation" api:"required"`
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
	Status string `json:"status" api:"required"`
	// Date when the company was officially terminated (if applicable). Format: ISO
	// 8601 (YYYY-MM-DD) Example: "2024-01-01"
	TerminatedAt time.Time `json:"terminated_at" api:"required" format:"date"`
	// Insolvency proceedings of the company, if any. Contains basic information per
	// proceeding; use the insolvency endpoint to retrieve all events of a proceeding.
	Insolvencies []JobResultsCompanyInsolvency `json:"insolvencies"`
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
		Estimates               respjson.Field
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
func (r JobResultsCompany) RawJSON() string { return r.JSON.raw }
func (r *JobResultsCompany) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobResultsCompanyAcquisition struct {
	// Date the underlying contract (Verschmelzungsvertrag) was concluded, as cited in
	// the register entry. Null when the register text does not cite a contract date.
	// Entries sharing an agreement_date belong to the same transaction. Format: ISO
	// 8601 (YYYY-MM-DD)
	AgreementDate time.Time `json:"agreement_date" api:"required" format:"date"`
	// Unique company identifier of the company that was merged into this company.
	// Example: DE-HRB-F1103-267645
	CompanyID string `json:"company_id" api:"required"`
	// Current name of the company that was merged into this company.
	Name string `json:"name" api:"required"`
	// Date the merger was registered. Format: ISO 8601 (YYYY-MM-DD)
	RegistrationDate time.Time `json:"registration_date" api:"required" format:"date"`
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
func (r JobResultsCompanyAcquisition) RawJSON() string { return r.JSON.raw }
func (r *JobResultsCompanyAcquisition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobResultsCompanyAssetSpinOff struct {
	// Date the underlying contract (Ausgliederungsvertrag) was concluded, as cited in
	// the register entry. Null when the register text does not cite a contract date.
	// Entries sharing an agreement_date belong to the same transaction. Format: ISO
	// 8601 (YYYY-MM-DD)
	AgreementDate time.Time `json:"agreement_date" api:"required" format:"date"`
	// Unique company identifier of the company that received the assets. Example:
	// DE-HRB-F1103-267645
	CompanyID string `json:"company_id" api:"required"`
	// Current name of the company that received the assets.
	Name string `json:"name" api:"required"`
	// Date the spin-off was registered. Format: ISO 8601 (YYYY-MM-DD)
	RegistrationDate time.Time `json:"registration_date" api:"required" format:"date"`
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
func (r JobResultsCompanyAssetSpinOff) RawJSON() string { return r.JSON.raw }
func (r *JobResultsCompanyAssetSpinOff) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contact information of the company.
type JobResultsCompanyContact struct {
	SocialMedia JobResultsCompanyContactSocialMedia `json:"social_media" api:"required"`
	WebsiteURL  string                              `json:"website_url" api:"required"`
	Email       string                              `json:"email"`
	Phone       string                              `json:"phone"`
	VatID       string                              `json:"vat_id"`
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
func (r JobResultsCompanyContact) RawJSON() string { return r.JSON.raw }
func (r *JobResultsCompanyContact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobResultsCompanyContactSocialMedia struct {
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
func (r JobResultsCompanyContactSocialMedia) RawJSON() string { return r.JSON.raw }
func (r *JobResultsCompanyContactSocialMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Modelled revenue, EBITDA and employees for companies that do not report them.
// One estimate per metric, for the latest fiscal year.
type JobResultsCompanyEstimates struct {
	// A modelled value for one fiscal year. min and max are not hard limits: the true
	// value lies between them with 80% probability.
	Ebitda JobResultsCompanyEstimatesEbitda `json:"ebitda" api:"required"`
	// A modelled value for one fiscal year. min and max are not hard limits: the true
	// value lies between them with 80% probability.
	Employees JobResultsCompanyEstimatesEmployees `json:"employees" api:"required"`
	// A modelled value for one fiscal year. min and max are not hard limits: the true
	// value lies between them with 80% probability.
	Revenue JobResultsCompanyEstimatesRevenue `json:"revenue" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ebitda      respjson.Field
		Employees   respjson.Field
		Revenue     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobResultsCompanyEstimates) RawJSON() string { return r.JSON.raw }
func (r *JobResultsCompanyEstimates) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A modelled value for one fiscal year. min and max are not hard limits: the true
// value lies between them with 80% probability.
type JobResultsCompanyEstimatesEbitda struct {
	// How narrow the 80% range is, the same tiers the search filters
	// estimated\_\*\_confidence use. Revenue and employees are judged by the ratio of
	// max to min, EBITDA by the range relative to the latest reported balance sheet
	// total (low when there is none).
	//
	// Any of "high", "medium", "low".
	Confidence string `json:"confidence" api:"required"`
	// End of the fiscal year the estimate applies to, matching the company's reporting
	// year end. Format: ISO 8601 (YYYY-MM-DD) Example: "2024-12-31"
	Date time.Time `json:"date" api:"required" format:"date"`
	// Upper end of the 80% range.
	Max int64 `json:"max" api:"required"`
	// Lower end of the 80% range.
	Min int64 `json:"min" api:"required"`
	// The estimated figure, the middle of the modelled range.
	Value int64 `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Confidence  respjson.Field
		Date        respjson.Field
		Max         respjson.Field
		Min         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobResultsCompanyEstimatesEbitda) RawJSON() string { return r.JSON.raw }
func (r *JobResultsCompanyEstimatesEbitda) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A modelled value for one fiscal year. min and max are not hard limits: the true
// value lies between them with 80% probability.
type JobResultsCompanyEstimatesEmployees struct {
	// How narrow the 80% range is, the same tiers the search filters
	// estimated\_\*\_confidence use. Revenue and employees are judged by the ratio of
	// max to min, EBITDA by the range relative to the latest reported balance sheet
	// total (low when there is none).
	//
	// Any of "high", "medium", "low".
	Confidence string `json:"confidence" api:"required"`
	// End of the fiscal year the estimate applies to, matching the company's reporting
	// year end. Format: ISO 8601 (YYYY-MM-DD) Example: "2024-12-31"
	Date time.Time `json:"date" api:"required" format:"date"`
	// Upper end of the 80% range.
	Max int64 `json:"max" api:"required"`
	// Lower end of the 80% range.
	Min int64 `json:"min" api:"required"`
	// The estimated figure, the middle of the modelled range.
	Value int64 `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Confidence  respjson.Field
		Date        respjson.Field
		Max         respjson.Field
		Min         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobResultsCompanyEstimatesEmployees) RawJSON() string { return r.JSON.raw }
func (r *JobResultsCompanyEstimatesEmployees) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A modelled value for one fiscal year. min and max are not hard limits: the true
// value lies between them with 80% probability.
type JobResultsCompanyEstimatesRevenue struct {
	// How narrow the 80% range is, the same tiers the search filters
	// estimated\_\*\_confidence use. Revenue and employees are judged by the ratio of
	// max to min, EBITDA by the range relative to the latest reported balance sheet
	// total (low when there is none).
	//
	// Any of "high", "medium", "low".
	Confidence string `json:"confidence" api:"required"`
	// End of the fiscal year the estimate applies to, matching the company's reporting
	// year end. Format: ISO 8601 (YYYY-MM-DD) Example: "2024-12-31"
	Date time.Time `json:"date" api:"required" format:"date"`
	// Upper end of the 80% range.
	Max int64 `json:"max" api:"required"`
	// Lower end of the 80% range.
	Min int64 `json:"min" api:"required"`
	// The estimated figure, the middle of the modelled range.
	Value int64 `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Confidence  respjson.Field
		Date        respjson.Field
		Max         respjson.Field
		Min         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobResultsCompanyEstimatesRevenue) RawJSON() string { return r.JSON.raw }
func (r *JobResultsCompanyEstimatesRevenue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A focused subset of the key company indicators for a given year. Values of the
// indicator are given in the smallest currency unit (cents). Example: 2099
// represents €20.99 for monetary values. For non-monetary values (e.g.,
// employees), the actual number.
type JobResultsCompanyIndicator struct {
	// The balance sheet total of that year (in cents).
	BalanceSheetTotal int64 `json:"balance_sheet_total" api:"required"`
	// The capital reserves of that year (in cents).
	CapitalReserves int64 `json:"capital_reserves" api:"required"`
	// The cash of that year (in cents).
	Cash int64 `json:"cash" api:"required"`
	// Date to which this financial indicators apply. Format: ISO 8601 (YYYY-MM-DD)
	// Example: "2022-01-01"
	Date time.Time `json:"date" api:"required" format:"date"`
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
	// A short fiscal year (Rumpfgeschäftsjahr), shorter than twelve months
	ShortYear bool `json:"short_year" api:"required"`
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
		ShortYear         respjson.Field
		Taxes             respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobResultsCompanyIndicator) RawJSON() string { return r.JSON.raw }
func (r *JobResultsCompanyIndicator) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Industry codes of the company, classified from its current purpose
// (Unternehmensgegenstand). Empty when the purpose has not been classified or
// describes no codable activity.
type JobResultsCompanyIndustryCodes struct {
	Wz2025 []JobResultsCompanyIndustryCodesWz2025 `json:"WZ2025" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Wz2025      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobResultsCompanyIndustryCodes) RawJSON() string { return r.JSON.raw }
func (r *JobResultsCompanyIndustryCodes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A WZ 2025 class (Klasse) the company is active in, with the subclasses
// (Unterklassen) behind it. The class of the company's main activity comes first.
type JobResultsCompanyIndustryCodesWz2025 struct {
	// WZ 2025 class. Example: "64.19"
	Code string `json:"code" api:"required"`
	// Official German title of the class. Example: "Kreditinstitute, ohne
	// Spezialkreditinstitute"
	LabelDe string `json:"label_de" api:"required"`
	// Whether the company's main activity falls in this class. True for at most one
	// class. The main activity is classified from the purpose, not recorded in the
	// register.
	Primary bool `json:"primary" api:"required"`
	// The WZ 2025 subclasses within this class, the main activity first.
	Subclasses []JobResultsCompanyIndustryCodesWz2025Subclass `json:"subclasses" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		LabelDe     respjson.Field
		Primary     respjson.Field
		Subclasses  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobResultsCompanyIndustryCodesWz2025) RawJSON() string { return r.JSON.raw }
func (r *JobResultsCompanyIndustryCodesWz2025) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A WZ 2025 subclass (Unterklasse).
type JobResultsCompanyIndustryCodesWz2025Subclass struct {
	// WZ 2025 subclass. Example: "64.19.1"
	Code string `json:"code" api:"required"`
	// Official German title of the subclass. Example: "Kreditbanken einschließlich
	// Zweigstellen ausländischer Banken"
	LabelDe string `json:"label_de" api:"required"`
	// Whether this is the company's main activity. True for at most one subclass. The
	// main activity is classified from the purpose, not recorded in the register.
	Primary bool `json:"primary" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		LabelDe     respjson.Field
		Primary     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobResultsCompanyIndustryCodesWz2025Subclass) RawJSON() string { return r.JSON.raw }
func (r *JobResultsCompanyIndustryCodesWz2025Subclass) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// If the company ceased to exist through a merger (Verschmelzung), the company it
// was merged into.
type JobResultsCompanyMergedInto struct {
	// Date the underlying contract (Verschmelzungsvertrag) was concluded, as cited in
	// the register entry. Null when the register text does not cite a contract date.
	// Entries sharing an agreement_date belong to the same transaction. Format: ISO
	// 8601 (YYYY-MM-DD)
	AgreementDate time.Time `json:"agreement_date" api:"required" format:"date"`
	// Unique company identifier of the company this company was merged into. Example:
	// DE-HRB-F1103-267645
	CompanyID string `json:"company_id" api:"required"`
	// Current name of the company this company was merged into.
	Name string `json:"name" api:"required"`
	// Date the merger was registered. Format: ISO 8601 (YYYY-MM-DD)
	RegistrationDate time.Time `json:"registration_date" api:"required" format:"date"`
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
func (r JobResultsCompanyMergedInto) RawJSON() string { return r.JSON.raw }
func (r *JobResultsCompanyMergedInto) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The company's current profit and loss transfer agreement
// (Gewinnabführungsvertrag), if one exists. The referenced company is the parent
// receiving this company's profit (Organträger). Null if the company has no active
// agreement.
type JobResultsCompanyProfitTransferAgreement struct {
	// Date the underlying contract (Gewinnabführungsvertrag) was concluded, as cited
	// in the register entry. Null when the register text does not cite a contract
	// date. Entries sharing an agreement_date belong to the same transaction. Format:
	// ISO 8601 (YYYY-MM-DD)
	AgreementDate time.Time `json:"agreement_date" api:"required" format:"date"`
	// Unique company identifier of the parent company receiving this company's profit
	// (Organträger). Example: DE-HRB-F1103-267645
	CompanyID string `json:"company_id" api:"required"`
	// Current name of the parent company.
	Name string `json:"name" api:"required"`
	// Date the agreement was registered. Format: ISO 8601 (YYYY-MM-DD)
	RegistrationDate time.Time `json:"registration_date" api:"required" format:"date"`
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
func (r JobResultsCompanyProfitTransferAgreement) RawJSON() string { return r.JSON.raw }
func (r *JobResultsCompanyProfitTransferAgreement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobResultsCompanyRepresentation struct {
	// Unique identifier for the representative. For companies: Format matches
	// company_id pattern For individuals: UUID Example: "DE-HRB-F1103-267645" or UUID
	// May be null for certain representatives.
	ID string `json:"id" api:"required"`
	// The representative's current individual representation authority (individuelle
	// Vertretungsbefugnis), as published in the register. Null if no special authority
	// is recorded. Example: "einzelvertretungsberechtigt mit der Befugnis, im Namen
	// der Gesellschaft mit sich im eigenen Namen Rechtsgeschäfte abzuschließen"
	Authority string `json:"authority" api:"required"`
	// Structured facts read from `authority`. Null when none are stated.
	AuthorityDetail JobResultsCompanyRepresentationAuthorityDetail `json:"authority_detail" api:"required"`
	// Date when this representative role ended (if applicable). Format: ISO 8601
	// (YYYY-MM-DD) Example: "2022-01-01"
	EndDate time.Time `json:"end_date" api:"required" format:"date"`
	// The name of the representative. E.g. "Max Mustermann" or "Max Mustermann GmbH"
	Name string `json:"name" api:"required"`
	// Coarse role bucket. Stable set of values, safe to switch on exhaustively. Use
	// `role_detail` for the precise register role. E.g. "DIRECTOR"
	//
	// Any of "DIRECTOR", "PROKURA", "SHAREHOLDER", "OWNER", "PARTNER",
	// "PERSONAL_LIABLE_DIRECTOR", "LIQUIDATOR", "OTHER".
	Role RepresentationRole `json:"role" api:"required"`
	// The precise role as recorded in the register, plus ready-to-display titles. A
	// Vorstand and a Geschäftsführer both have role "DIRECTOR" but differ here.
	RoleDetail JobResultsCompanyRepresentationRoleDetail `json:"role_detail" api:"required"`
	// Date when this representative role became effective. Format: ISO 8601
	// (YYYY-MM-DD) Example: "2022-01-01"
	StartDate time.Time `json:"start_date" api:"required" format:"date"`
	// Whether the representation is a natural person or a legal entity.
	//
	// Any of "natural_person", "legal_person".
	Type          EntityType                                   `json:"type" api:"required"`
	LegalPerson   JobResultsCompanyRepresentationLegalPerson   `json:"legal_person" api:"nullable"`
	NaturalPerson JobResultsCompanyRepresentationNaturalPerson `json:"natural_person" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Authority       respjson.Field
		AuthorityDetail respjson.Field
		EndDate         respjson.Field
		Name            respjson.Field
		Role            respjson.Field
		RoleDetail      respjson.Field
		StartDate       respjson.Field
		Type            respjson.Field
		LegalPerson     respjson.Field
		NaturalPerson   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobResultsCompanyRepresentation) RawJSON() string { return r.JSON.raw }
func (r *JobResultsCompanyRepresentation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Structured facts read from `authority`. Null when none are stated.
type JobResultsCompanyRepresentationAuthorityDetail struct {
	// Fully exempt from self-dealing restrictions (in Germany § 181 BGB). Null when
	// not stated or only partly exempt.
	SelfDealingAllowed bool `json:"self_dealing_allowed" api:"required"`
	// SOLE may represent alone, JOINT only together with others, NONE is excluded from
	// representation, GENERAL follows the general representation rule.
	//
	// Any of "SOLE", "JOINT", "NONE", "GENERAL".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SelfDealingAllowed respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobResultsCompanyRepresentationAuthorityDetail) RawJSON() string { return r.JSON.raw }
func (r *JobResultsCompanyRepresentationAuthorityDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The precise role as recorded in the register, plus ready-to-display titles. A
// Vorstand and a Geschäftsführer both have role "DIRECTOR" but differ here.
type JobResultsCompanyRepresentationRoleDetail struct {
	// The precise role as recorded in the register. Example: "EXECUTIVE" for a
	// Vorstand, "DIRECTOR" for a Geschäftsführer — both have role "DIRECTOR".
	//
	// Any of "DIRECTOR", "EMERGENCY_DIRECTOR", "EXECUTIVE", "CHAIRMAN",
	// "DEPUTY_CHAIRMAN", "BOARD", "PRESIDENT_BOARD", "VICE_PRESIDENT_BOARD",
	// "PRESIDENT", "MEMBER", "NON_EXECUTIVE_DIRECTOR", "BOARD_SECRETARY",
	// "BOARD_TREASURER", "BOARD_SPORTS", "BOARD_OTHER", "REPRESENTATIVE",
	// "SECONDARY_REPRESENTATIVE", "COURT_ORDERED_REPRESENTATIVE", "PROKURA",
	// "PROKURA_SINGLE", "PROKURA_JOINT", "LIQUIDATOR", "INSOLVENCY_DIRECTOR",
	// "PRELIMINARY_INSOLVENCY_DIRECTOR", "OWNER", "SHAREHOLDER", "PARTNER",
	// "PERSONAL_LIABLE_DIRECTOR", "LIMITED_PARTNER", "OTHER".
	Code string `json:"code" api:"required"`
	// German title, gender-neutral, as shown on openregister.de. Example: "Vorstand"
	LabelDe string `json:"label_de" api:"required"`
	// Abbreviated German title, for dense layouts such as tables. Example: "Vst."
	LabelDeShort string `json:"label_de_short" api:"required"`
	// English title. A translation for display, not a legal term — the German title is
	// the one the register uses. Example: "Executive Board Member"
	LabelEn string `json:"label_en" api:"required"`
	// Abbreviated English title, for dense layouts such as tables. Example: "Exec.
	// Board"
	LabelEnShort string `json:"label_en_short" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code         respjson.Field
		LabelDe      respjson.Field
		LabelDeShort respjson.Field
		LabelEn      respjson.Field
		LabelEnShort respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobResultsCompanyRepresentationRoleDetail) RawJSON() string { return r.JSON.raw }
func (r *JobResultsCompanyRepresentationRoleDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobResultsCompanyRepresentationLegalPerson struct {
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
func (r JobResultsCompanyRepresentationLegalPerson) RawJSON() string { return r.JSON.raw }
func (r *JobResultsCompanyRepresentationLegalPerson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobResultsCompanyRepresentationNaturalPerson struct {
	// City where the representative is located. Example: "Berlin"
	City string `json:"city" api:"required"`
	// Date of birth of the representative. May still be null for natural persons if it
	// is not available. Format: ISO 8601 (YYYY-MM-DD) Example: "1990-01-01"
	DateOfBirth time.Time `json:"date_of_birth" api:"required" format:"date"`
	// First name of the representative. Example: "Max"
	FirstName string `json:"first_name" api:"required"`
	// Last name of the representative. Example: "Mustermann"
	LastName string `json:"last_name" api:"required"`
	// Academic or professional title of the representative. Example: "Dr."
	Title string `json:"title" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		DateOfBirth respjson.Field
		FirstName   respjson.Field
		LastName    respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobResultsCompanyRepresentationNaturalPerson) RawJSON() string { return r.JSON.raw }
func (r *JobResultsCompanyRepresentationNaturalPerson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Basic information about an insolvency proceeding of the company. Use the
// insolvency endpoint to retrieve all events of the proceeding.
type JobResultsCompanyInsolvency struct {
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
	// Date the proceeding was closed. Format: ISO 8601 (YYYY-MM-DD)
	ClosedAt time.Time `json:"closed_at" api:"nullable" format:"date"`
	// Date the proceeding was opened. Format: ISO 8601 (YYYY-MM-DD)
	OpenedAt time.Time `json:"opened_at" api:"nullable" format:"date"`
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
func (r JobResultsCompanyInsolvency) RawJSON() string { return r.JSON.raw }
func (r *JobResultsCompanyInsolvency) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The fetched document of a `document` job.
type JobResultsDocument struct {
	FileDate time.Time `json:"file_date" api:"required" format:"date"`
	FileName string    `json:"file_name" api:"required"`
	// Download URL of the document. It expires; read the job again for a new one.
	URL string `json:"url" api:"required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileDate    respjson.Field
		FileName    respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobResultsDocument) RawJSON() string { return r.JSON.raw }
func (r *JobResultsDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The refreshed owners of an `owners` job.
type JobResultsOwners struct {
	// When true, the returned owner data is the best available but may not reflect the
	// most current ownership state. This applies to AG and SE companies where
	// ownership data is sourced from Handelsregister decision and articles of
	// association documents, which are not filed on every ownership change.
	BestAvailable bool `json:"best_available" api:"required"`
	// Unique company identifier. Example: DE-HRB-F1103-267645
	CompanyID string                  `json:"company_id" api:"required"`
	Owners    []JobResultsOwnersOwner `json:"owners" api:"required"`
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
func (r JobResultsOwners) RawJSON() string { return r.JSON.raw }
func (r *JobResultsOwners) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobResultsOwnersOwner struct {
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
	// Type of relationship between the entity and the company. Use `role_detail` for
	// the precise register role.
	//
	// Any of "shareholder", "stockholder", "limited_partner", "general_partner".
	RelationType CompanyRelationType `json:"relation_type" api:"required"`
	// The precise role as recorded in the register, plus ready-to-display titles. A
	// Kommanditist and a Gesellschafter both appear here with their register-level
	// role, while `relation_type` only carries the coarse bucket.
	RoleDetail JobResultsOwnersOwnerRoleDetail `json:"role_detail" api:"required"`
	// Date when the relation started. Only available for some types of owners. Format:
	// ISO 8601 (YYYY-MM-DD) Example: "2022-01-01"
	Start time.Time `json:"start" api:"required" format:"date"`
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
		RoleDetail      respjson.Field
		Start           respjson.Field
		Type            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobResultsOwnersOwner) RawJSON() string { return r.JSON.raw }
func (r *JobResultsOwnersOwner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The precise role as recorded in the register, plus ready-to-display titles. A
// Kommanditist and a Gesellschafter both appear here with their register-level
// role, while `relation_type` only carries the coarse bucket.
type JobResultsOwnersOwnerRoleDetail struct {
	// The precise role as recorded in the register. Example: "EXECUTIVE" for a
	// Vorstand, "DIRECTOR" for a Geschäftsführer — both have role "DIRECTOR".
	//
	// Any of "DIRECTOR", "EMERGENCY_DIRECTOR", "EXECUTIVE", "CHAIRMAN",
	// "DEPUTY_CHAIRMAN", "BOARD", "PRESIDENT_BOARD", "VICE_PRESIDENT_BOARD",
	// "PRESIDENT", "MEMBER", "NON_EXECUTIVE_DIRECTOR", "BOARD_SECRETARY",
	// "BOARD_TREASURER", "BOARD_SPORTS", "BOARD_OTHER", "REPRESENTATIVE",
	// "SECONDARY_REPRESENTATIVE", "COURT_ORDERED_REPRESENTATIVE", "PROKURA",
	// "PROKURA_SINGLE", "PROKURA_JOINT", "LIQUIDATOR", "INSOLVENCY_DIRECTOR",
	// "PRELIMINARY_INSOLVENCY_DIRECTOR", "OWNER", "SHAREHOLDER", "PARTNER",
	// "PERSONAL_LIABLE_DIRECTOR", "LIMITED_PARTNER", "OTHER".
	Code string `json:"code" api:"required"`
	// German title, gender-neutral, as shown on openregister.de. Example: "Vorstand"
	LabelDe string `json:"label_de" api:"required"`
	// Abbreviated German title, for dense layouts such as tables. Example: "Vst."
	LabelDeShort string `json:"label_de_short" api:"required"`
	// English title. A translation for display, not a legal term — the German title is
	// the one the register uses. Example: "Executive Board Member"
	LabelEn string `json:"label_en" api:"required"`
	// Abbreviated English title, for dense layouts such as tables. Example: "Exec.
	// Board"
	LabelEnShort string `json:"label_en_short" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code         respjson.Field
		LabelDe      respjson.Field
		LabelDeShort respjson.Field
		LabelEn      respjson.Field
		LabelEnShort respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobResultsOwnersOwnerRoleDetail) RawJSON() string { return r.JSON.raw }
func (r *JobResultsOwnersOwnerRoleDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobStatus string

const (
	JobStatusPending   JobStatus = "pending"
	JobStatusCompleted JobStatus = "completed"
	JobStatusFailed    JobStatus = "failed"
)

type JobDocumentCategory string

const (
	JobDocumentCategoryCurrentPrintout       JobDocumentCategory = "current_printout"
	JobDocumentCategoryChronologicalPrintout JobDocumentCategory = "chronological_printout"
	JobDocumentCategoryHistoricalPrintout    JobDocumentCategory = "historical_printout"
	JobDocumentCategoryStructuredInformation JobDocumentCategory = "structured_information"
	JobDocumentCategoryShareholderList       JobDocumentCategory = "shareholder_list"
	JobDocumentCategoryArticlesOfAssociation JobDocumentCategory = "articles_of_association"
)

type JobType string

const (
	JobTypeCompany  JobType = "company"
	JobTypeOwners   JobType = "owners"
	JobTypeDocument JobType = "document"
)

type JobNewV1Params struct {
	// Unique company identifier. Example: DE-HRB-F1103-267645
	CompanyID string `json:"company_id" api:"required"`
	// Any of "company", "owners", "document".
	Type JobType `json:"type,omitzero" api:"required"`
	// Any of "current_printout", "chronological_printout", "historical_printout",
	// "structured_information", "shareholder_list", "articles_of_association".
	DocumentCategory JobDocumentCategory `json:"document_category,omitzero"`
	paramObj
}

func (r JobNewV1Params) MarshalJSON() (data []byte, err error) {
	type shadow JobNewV1Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JobNewV1Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
