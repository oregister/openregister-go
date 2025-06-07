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
func (r *DocumentService) Get(ctx context.Context, documentID string, opts ...option.RequestOption) (res *DocumentGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if documentID == "" {
		err = errors.New("missing required document_id parameter")
		return
	}
	path := fmt.Sprintf("v0/document/%s", documentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type DocumentGetResponse struct {
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
	Type DocumentGetResponseType `json:"type,required"`
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
func (r DocumentGetResponse) RawJSON() string { return r.JSON.raw }
func (r *DocumentGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of document.
type DocumentGetResponseType string

const (
	DocumentGetResponseTypeArticlesOfAssociation DocumentGetResponseType = "articles_of_association"
	DocumentGetResponseTypeSampleProtocol        DocumentGetResponseType = "sample_protocol"
	DocumentGetResponseTypeShareholderList       DocumentGetResponseType = "shareholder_list"
)
