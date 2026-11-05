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
	Company CompanyGetDetailsV1Response `json:"company" api:"required"`
	// The fetched document of a `document` job.
	Document JobResultsDocument `json:"document" api:"required"`
	// The refreshed owners of an `owners` job.
	Owners CompanyGetOwnersV1Response `json:"owners" api:"required"`
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
