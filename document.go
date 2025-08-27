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
	"github.com/oregister/openregister-go/packages/respjson"
)

// DocumentService contains methods and other services that help with interacting
// with the openregister API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDocumentService] method instead.
type DocumentService struct {
	Options []option.RequestOption
}

// NewDocumentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDocumentService(opts ...option.RequestOption) (r DocumentService) {
	r = DocumentService{}
	r.Options = opts
	return
}

// Get document information
func (r *DocumentService) DocumentCached(ctx context.Context, documentID string, opts ...option.RequestOption) (res *DocumentDocumentCachedResponse, err error) {
	opts = append(r.Options[:], opts...)
	if documentID == "" {
		err = errors.New("missing required document_id parameter")
		return
	}
	path := fmt.Sprintf("v1/document/%s", documentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetch a document in realtime.
func (r *DocumentService) Fetch(ctx context.Context, query DocumentFetchParams, opts ...option.RequestOption) (res *DocumentFetchResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/document"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type DocumentDocumentCachedResponse struct {
	// The unique identifier for the document. E.g.
	// "f47ac10b-58cc-4372-a567-0e02b2c3d479"
	ID string `json:"id,required"`
	// The date of the document. E.g. "2022-01-01"
	Date string `json:"date,required"`
	// The name of the document. E.g. "Musterprotokoll vom 01.01.2022"
	Name string `json:"name,required"`
	// The type of document.
	//
	// Any of "articles_of_association", "sample_protocol", "shareholder_list".
	Type DocumentDocumentCachedResponseType `json:"type,required"`
	// The URL of the document. It can be downloaded from there. Only valid for 15
	// minutes after the request.
	URL string `json:"url,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Date        respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentDocumentCachedResponse) RawJSON() string { return r.JSON.raw }
func (r *DocumentDocumentCachedResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of document.
type DocumentDocumentCachedResponseType string

const (
	DocumentDocumentCachedResponseTypeArticlesOfAssociation DocumentDocumentCachedResponseType = "articles_of_association"
	DocumentDocumentCachedResponseTypeSampleProtocol        DocumentDocumentCachedResponseType = "sample_protocol"
	DocumentDocumentCachedResponseTypeShareholderList       DocumentDocumentCachedResponseType = "shareholder_list"
)

type DocumentFetchResponse struct {
	// Any of "current_printout", "chronological_printout", "historical_printout",
	// "structured_information", "shareholder_list", "articles_of_association".
	Category DocumentFetchResponseCategory `json:"category,required"`
	FileDate string                        `json:"file_date,required" format:"date-only"`
	FileName string                        `json:"file_name,required"`
	URL      string                        `json:"url,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Category    respjson.Field
		FileDate    respjson.Field
		FileName    respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentFetchResponse) RawJSON() string { return r.JSON.raw }
func (r *DocumentFetchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentFetchResponseCategory string

const (
	DocumentFetchResponseCategoryCurrentPrintout       DocumentFetchResponseCategory = "current_printout"
	DocumentFetchResponseCategoryChronologicalPrintout DocumentFetchResponseCategory = "chronological_printout"
	DocumentFetchResponseCategoryHistoricalPrintout    DocumentFetchResponseCategory = "historical_printout"
	DocumentFetchResponseCategoryStructuredInformation DocumentFetchResponseCategory = "structured_information"
	DocumentFetchResponseCategoryShareholderList       DocumentFetchResponseCategory = "shareholder_list"
	DocumentFetchResponseCategoryArticlesOfAssociation DocumentFetchResponseCategory = "articles_of_association"
)

type DocumentFetchParams struct {
	CompanyID string `query:"company_id,required" json:"-"`
	// Any of "current_printout", "chronological_printout", "historical_printout",
	// "structured_information", "shareholder_list", "articles_of_association".
	DocumentCategory DocumentFetchParamsDocumentCategory `query:"document_category,omitzero,required" json:"-"`
	paramObj
}

// URLQuery serializes [DocumentFetchParams]'s query parameters as `url.Values`.
func (r DocumentFetchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DocumentFetchParamsDocumentCategory string

const (
	DocumentFetchParamsDocumentCategoryCurrentPrintout       DocumentFetchParamsDocumentCategory = "current_printout"
	DocumentFetchParamsDocumentCategoryChronologicalPrintout DocumentFetchParamsDocumentCategory = "chronological_printout"
	DocumentFetchParamsDocumentCategoryHistoricalPrintout    DocumentFetchParamsDocumentCategory = "historical_printout"
	DocumentFetchParamsDocumentCategoryStructuredInformation DocumentFetchParamsDocumentCategory = "structured_information"
	DocumentFetchParamsDocumentCategoryShareholderList       DocumentFetchParamsDocumentCategory = "shareholder_list"
	DocumentFetchParamsDocumentCategoryArticlesOfAssociation DocumentFetchParamsDocumentCategory = "articles_of_association"
)
