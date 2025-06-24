// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openregister

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/oregister/openregister-go/internal/apijson"
	"github.com/oregister/openregister-go/internal/requestconfig"
	"github.com/oregister/openregister-go/option"
	"github.com/oregister/openregister-go/packages/param"
	"github.com/oregister/openregister-go/packages/respjson"
)

// JobDocumentService contains methods and other services that help with
// interacting with the openregister API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewJobDocumentService] method instead.
type JobDocumentService struct {
	Options []option.RequestOption
}

// NewJobDocumentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewJobDocumentService(opts ...option.RequestOption) (r JobDocumentService) {
	r = JobDocumentService{}
	r.Options = opts
	return
}

// Create a document job
func (r *JobDocumentService) New(ctx context.Context, body JobDocumentNewParams, opts ...option.RequestOption) (res *JobDocumentNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/jobs/document"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get document job status
func (r *JobDocumentService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *JobDocumentGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v0/jobs/document/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type JobDocumentNewResponse struct {
	// Unique job identifier. Example: f47ac10b-58cc-4372-a567-0e02b2c3d479
	ID string `json:"id,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobDocumentNewResponse) RawJSON() string { return r.JSON.raw }
func (r *JobDocumentNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobDocumentGetResponse struct {
	// Any of "pending", "completed", "failed".
	Status JobDocumentGetResponseStatus `json:"status,required"`
	// Date when the job was created. Format: ISO 8601 (YYYY-MM-DD) Example:
	// "2022-01-01"
	Date string `json:"date"`
	URL  string `json:"url" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		Date        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobDocumentGetResponse) RawJSON() string { return r.JSON.raw }
func (r *JobDocumentGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobDocumentGetResponseStatus string

const (
	JobDocumentGetResponseStatusPending   JobDocumentGetResponseStatus = "pending"
	JobDocumentGetResponseStatusCompleted JobDocumentGetResponseStatus = "completed"
	JobDocumentGetResponseStatusFailed    JobDocumentGetResponseStatus = "failed"
)

type JobDocumentNewParams struct {
	// Unique company identifier. Example: DE-HRB-F1103-267645
	CompanyID string `json:"company_id,required"`
	// Any of "current_printout", "chronological_printout", "historical_printout",
	// "structured_information", "shareholder_list", "articles_of_association".
	DocumentCategory JobDocumentNewParamsDocumentCategory `json:"document_category,omitzero,required"`
	paramObj
}

func (r JobDocumentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow JobDocumentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JobDocumentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobDocumentNewParamsDocumentCategory string

const (
	JobDocumentNewParamsDocumentCategoryCurrentPrintout       JobDocumentNewParamsDocumentCategory = "current_printout"
	JobDocumentNewParamsDocumentCategoryChronologicalPrintout JobDocumentNewParamsDocumentCategory = "chronological_printout"
	JobDocumentNewParamsDocumentCategoryHistoricalPrintout    JobDocumentNewParamsDocumentCategory = "historical_printout"
	JobDocumentNewParamsDocumentCategoryStructuredInformation JobDocumentNewParamsDocumentCategory = "structured_information"
	JobDocumentNewParamsDocumentCategoryShareholderList       JobDocumentNewParamsDocumentCategory = "shareholder_list"
	JobDocumentNewParamsDocumentCategoryArticlesOfAssociation JobDocumentNewParamsDocumentCategory = "articles_of_association"
)
