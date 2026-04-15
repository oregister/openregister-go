// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openregister

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/oregister/openregister-go/v2/internal/apijson"
	"github.com/oregister/openregister-go/v2/internal/apiquery"
	"github.com/oregister/openregister-go/v2/internal/requestconfig"
	"github.com/oregister/openregister-go/v2/option"
	"github.com/oregister/openregister-go/v2/packages/param"
	"github.com/oregister/openregister-go/v2/packages/respjson"
)

// TransparenzregisterRequestService contains methods and other services that help
// with interacting with the openregister API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransparenzregisterRequestService] method instead.
type TransparenzregisterRequestService struct {
	Options []option.RequestOption
}

// NewTransparenzregisterRequestService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewTransparenzregisterRequestService(opts ...option.RequestOption) (r TransparenzregisterRequestService) {
	r = TransparenzregisterRequestService{}
	r.Options = opts
	return
}

// Submit a Transparenzregister request for a company using its company ID. This
// endpoint will initiate the Transparenzregister request process and return a
// request ID for tracking.
func (r *TransparenzregisterRequestService) NewV1(ctx context.Context, params TransparenzregisterRequestNewV1Params, opts ...option.RequestOption) (res *TransparenzregisterRequestNewV1Response, err error) {
	if !param.IsOmitted(params.XCredentialLabel) {
		opts = append(opts, option.WithHeader("X-Credential-Label", fmt.Sprintf("%v", params.XCredentialLabel.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/transparenzregister/request"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get the results of a Transparenzregister request. This endpoint handles all
// internal complexity including polling request status, selecting all available
// documents, creating Transparenzregister baskets, and returning download URLs
// when ready. If the request is still processing, it will return a pending status.
func (r *TransparenzregisterRequestService) GetV1(ctx context.Context, requestID string, opts ...option.RequestOption) (res *TransparenzregisterRequestGetV1Response, err error) {
	opts = slices.Concat(r.Options, opts)
	if requestID == "" {
		err = errors.New("missing required request_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/transparenzregister/request/%s", requestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Response from submitting a document request.
type TransparenzregisterRequestNewV1Response struct {
	// Request ID for tracking the document request. Example: "req_12345678"
	RequestID string `json:"request_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RequestID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransparenzregisterRequestNewV1Response) RawJSON() string { return r.JSON.raw }
func (r *TransparenzregisterRequestNewV1Response) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response containing document request results and download URLs. All internal
// complexity (document IDs, baskets, polling) is handled automatically.
type TransparenzregisterRequestGetV1Response struct {
	// Request ID that was used to submit the request. Example: "req_12345678"
	RequestID string `json:"request_id" api:"required"`
	// Status of the Transparenzregister request.
	//
	// Any of "completed", "processing", "failed".
	Status TransparenzregisterRequestGetV1ResponseStatus `json:"status" api:"required"`
	// URLs for downloading all available documents.
	DownloadURLs []TransparenzregisterRequestGetV1ResponseDownloadURL `json:"download_urls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RequestID    respjson.Field
		Status       respjson.Field
		DownloadURLs respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransparenzregisterRequestGetV1Response) RawJSON() string { return r.JSON.raw }
func (r *TransparenzregisterRequestGetV1Response) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the Transparenzregister request.
type TransparenzregisterRequestGetV1ResponseStatus string

const (
	TransparenzregisterRequestGetV1ResponseStatusCompleted  TransparenzregisterRequestGetV1ResponseStatus = "completed"
	TransparenzregisterRequestGetV1ResponseStatusProcessing TransparenzregisterRequestGetV1ResponseStatus = "processing"
	TransparenzregisterRequestGetV1ResponseStatusFailed     TransparenzregisterRequestGetV1ResponseStatus = "failed"
)

// Download URL for a document with format information.
type TransparenzregisterRequestGetV1ResponseDownloadURL struct {
	// Stable UUID for this document.
	DocumentID string `json:"document_id" api:"required" format:"uuid"`
	// Suggested filename for the download. Example: "registerauszug_company_12345.pdf"
	Filename string `json:"filename" api:"required"`
	// Format of the downloadable document. Example: "xml", "pdf", "json"
	Format string `json:"format" api:"required"`
	// Download URL for the document. Example:
	// "https://api.example.com/download/abc123"
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DocumentID  respjson.Field
		Filename    respjson.Field
		Format      respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransparenzregisterRequestGetV1ResponseDownloadURL) RawJSON() string { return r.JSON.raw }
func (r *TransparenzregisterRequestGetV1ResponseDownloadURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransparenzregisterRequestNewV1Params struct {
	// Unique company identifier. Required unless using X-Credential-Label=test.
	// Example: DE-HRB-F1103-267645
	CompanyID        param.Opt[string] `query:"company_id,omitzero" json:"-"`
	XCredentialLabel param.Opt[string] `header:"X-Credential-Label,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TransparenzregisterRequestNewV1Params]'s query parameters
// as `url.Values`.
func (r TransparenzregisterRequestNewV1Params) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
