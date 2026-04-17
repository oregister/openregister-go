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

// TransparenzregisterExtractService contains methods and other services that help
// with interacting with the openregister API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransparenzregisterExtractService] method instead.
type TransparenzregisterExtractService struct {
	Options []option.RequestOption
}

// NewTransparenzregisterExtractService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewTransparenzregisterExtractService(opts ...option.RequestOption) (r TransparenzregisterExtractService) {
	r = TransparenzregisterExtractService{}
	r.Options = opts
	return
}

// Submit a Transparenzregister extract request and return an extract resource with
// processing status.
//
// Sandbox integration testing (recommended for all non-production testing):
//
//   - Send `X-Credential-Name: sandbox`.
//   - Do not send `company_id` (an empty body `{}` is valid).
//   - OpenRegister uses the Transparenzregister test environment and built-in test
//     authentication.
//   - The request is submitted with the fixed test EKRN `DE727032388716`.
//   - The response has `company_id: null`.
//
// Production usage:
//
//   - Omit `X-Credential-Name` or use `default` / another stored credential name.
//   - `company_id` is required and must resolve to exactly one Transparenzregister
//     legal entity.
func (r *TransparenzregisterExtractService) NewV1(ctx context.Context, params TransparenzregisterExtractNewV1Params, opts ...option.RequestOption) (res *TransparenzregisterExtractNewV1Response, err error) {
	if !param.IsOmitted(params.XCredentialName) {
		opts = append(opts, option.WithHeader("X-Credential-Name", fmt.Sprintf("%v", params.XCredentialName.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/transparenzregister/extracts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get the results of a Transparenzregister extract request. This endpoint handles
// all internal complexity including polling request status, selecting all
// available documents, creating Transparenzregister baskets, and returning
// download URLs when ready. If the request is still processing, it will return a
// pending status. Polling reuses the credential mode stored on the extract at
// create time. Sandbox extracts keep using the Transparenzregister test client
// automatically; no credential header is accepted on this endpoint.
func (r *TransparenzregisterExtractService) GetV1(ctx context.Context, extractID string, opts ...option.RequestOption) (res *TransparenzregisterExtractGetV1Response, err error) {
	opts = slices.Concat(r.Options, opts)
	if extractID == "" {
		err = errors.New("missing required extract_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/transparenzregister/extracts/%s", extractID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Transparenzregister extract resource including processing state, parsed report,
// and downloadable documents.
type TransparenzregisterExtractNewV1Response struct {
	// Stable extract identifier. Example: "tre_12345678"
	ID string `json:"id" api:"required"`
	// Status of the Transparenzregister extract.
	//
	// Any of "completed", "processing", "failed".
	Status TransparenzregisterExtractNewV1ResponseStatus `json:"status" api:"required"`
	// Company identifier associated with this extract request. May be null when using
	// sandbox credentials.
	CompanyID string `json:"company_id" api:"nullable"`
	// Timestamp when extract processing completed.
	CompletedAt time.Time `json:"completed_at" api:"nullable" format:"date-time"`
	// URLs for downloading available extract documents.
	Documents []TransparenzregisterExtractNewV1ResponseDocument `json:"documents"`
	// EKRN used to request this extract.
	Ekrn string `json:"ekrn" api:"nullable"`
	// Transparenzregister reference number from the extract.
	ReferenceNumber string `json:"reference_number" api:"nullable"`
	// Parsed Transparenzregister extract report limited to UBO-relevant fields.
	Report TransparenzregisterExtractNewV1ResponseReport `json:"report" api:"nullable"`
	// Timestamp when extract submission started.
	SubmittedAt time.Time `json:"submitted_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Status          respjson.Field
		CompanyID       respjson.Field
		CompletedAt     respjson.Field
		Documents       respjson.Field
		Ekrn            respjson.Field
		ReferenceNumber respjson.Field
		Report          respjson.Field
		SubmittedAt     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransparenzregisterExtractNewV1Response) RawJSON() string { return r.JSON.raw }
func (r *TransparenzregisterExtractNewV1Response) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the Transparenzregister extract.
type TransparenzregisterExtractNewV1ResponseStatus string

const (
	TransparenzregisterExtractNewV1ResponseStatusCompleted  TransparenzregisterExtractNewV1ResponseStatus = "completed"
	TransparenzregisterExtractNewV1ResponseStatusProcessing TransparenzregisterExtractNewV1ResponseStatus = "processing"
	TransparenzregisterExtractNewV1ResponseStatusFailed     TransparenzregisterExtractNewV1ResponseStatus = "failed"
)

// Download URL for a document with format information.
type TransparenzregisterExtractNewV1ResponseDocument struct {
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
func (r TransparenzregisterExtractNewV1ResponseDocument) RawJSON() string { return r.JSON.raw }
func (r *TransparenzregisterExtractNewV1ResponseDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parsed Transparenzregister extract report limited to UBO-relevant fields.
type TransparenzregisterExtractNewV1ResponseReport struct {
	// Extract creation date.
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date"`
	// Reason indicating no natural person UBO could be determined.
	FictionalUboReason string                                               `json:"fictional_ubo_reason" api:"nullable"`
	Groups             []TransparenzregisterExtractNewV1ResponseReportGroup `json:"groups"`
	// Type of Transparenzregister notice.
	NoticeType  string                                                   `json:"notice_type" api:"nullable"`
	StatusFlags TransparenzregisterExtractNewV1ResponseReportStatusFlags `json:"status_flags" api:"nullable"`
	Ubos        []TransparenzregisterExtractNewV1ResponseReportUbo       `json:"ubos"`
	Validity    TransparenzregisterExtractNewV1ResponseReportValidity    `json:"validity" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt          respjson.Field
		FictionalUboReason respjson.Field
		Groups             respjson.Field
		NoticeType         respjson.Field
		StatusFlags        respjson.Field
		Ubos               respjson.Field
		Validity           respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransparenzregisterExtractNewV1ResponseReport) RawJSON() string { return r.JSON.raw }
func (r *TransparenzregisterExtractNewV1ResponseReport) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransparenzregisterExtractNewV1ResponseReportGroup struct {
	Description  string `json:"description" api:"nullable"`
	InterestType string `json:"interest_type" api:"nullable"`
	Position     int64  `json:"position"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description  respjson.Field
		InterestType respjson.Field
		Position     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransparenzregisterExtractNewV1ResponseReportGroup) RawJSON() string { return r.JSON.raw }
func (r *TransparenzregisterExtractNewV1ResponseReportGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransparenzregisterExtractNewV1ResponseReportStatusFlags struct {
	CorrectedByReference string    `json:"corrected_by_reference" api:"nullable"`
	CorrectedReferences  []string  `json:"corrected_references"`
	Deleted              bool      `json:"deleted"`
	DeletionDate         time.Time `json:"deletion_date" api:"nullable" format:"date"`
	DiscrepancyNote      string    `json:"discrepancy_note" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CorrectedByReference respjson.Field
		CorrectedReferences  respjson.Field
		Deleted              respjson.Field
		DeletionDate         respjson.Field
		DiscrepancyNote      respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransparenzregisterExtractNewV1ResponseReportStatusFlags) RawJSON() string { return r.JSON.raw }
func (r *TransparenzregisterExtractNewV1ResponseReportStatusFlags) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransparenzregisterExtractNewV1ResponseReportUbo struct {
	Interest      TransparenzregisterExtractNewV1ResponseReportUboInterest      `json:"interest" api:"nullable"`
	NaturalPerson TransparenzregisterExtractNewV1ResponseReportUboNaturalPerson `json:"natural_person" api:"nullable"`
	Position      int64                                                         `json:"position"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Interest      respjson.Field
		NaturalPerson respjson.Field
		Position      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransparenzregisterExtractNewV1ResponseReportUbo) RawJSON() string { return r.JSON.raw }
func (r *TransparenzregisterExtractNewV1ResponseReportUbo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransparenzregisterExtractNewV1ResponseReportUboInterest struct {
	Percentage float64 `json:"percentage" api:"nullable"`
	Scope      string  `json:"scope" api:"nullable"`
	Type       string  `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Percentage  respjson.Field
		Scope       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransparenzregisterExtractNewV1ResponseReportUboInterest) RawJSON() string { return r.JSON.raw }
func (r *TransparenzregisterExtractNewV1ResponseReportUboInterest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransparenzregisterExtractNewV1ResponseReportUboNaturalPerson struct {
	City        string    `json:"city" api:"nullable"`
	Country     string    `json:"country" api:"nullable"`
	DateOfBirth time.Time `json:"date_of_birth" api:"nullable" format:"date"`
	FirstName   string    `json:"first_name" api:"nullable"`
	FullName    string    `json:"full_name" api:"nullable"`
	LastName    string    `json:"last_name" api:"nullable"`
	// ISO 3166-1 alpha-2 nationality codes where available.
	Nationalities []string `json:"nationalities"`
	Title         string   `json:"title" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City          respjson.Field
		Country       respjson.Field
		DateOfBirth   respjson.Field
		FirstName     respjson.Field
		FullName      respjson.Field
		LastName      respjson.Field
		Nationalities respjson.Field
		Title         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransparenzregisterExtractNewV1ResponseReportUboNaturalPerson) RawJSON() string {
	return r.JSON.raw
}
func (r *TransparenzregisterExtractNewV1ResponseReportUboNaturalPerson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransparenzregisterExtractNewV1ResponseReportValidity struct {
	From  TransparenzregisterExtractNewV1ResponseReportValidityFrom  `json:"from" api:"nullable"`
	Until TransparenzregisterExtractNewV1ResponseReportValidityUntil `json:"until" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From        respjson.Field
		Until       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransparenzregisterExtractNewV1ResponseReportValidity) RawJSON() string { return r.JSON.raw }
func (r *TransparenzregisterExtractNewV1ResponseReportValidity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransparenzregisterExtractNewV1ResponseReportValidityFrom struct {
	Date time.Time `json:"date" api:"nullable" format:"date"`
	Note string    `json:"note" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date        respjson.Field
		Note        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransparenzregisterExtractNewV1ResponseReportValidityFrom) RawJSON() string {
	return r.JSON.raw
}
func (r *TransparenzregisterExtractNewV1ResponseReportValidityFrom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransparenzregisterExtractNewV1ResponseReportValidityUntil struct {
	Date time.Time `json:"date" api:"nullable" format:"date"`
	Note string    `json:"note" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date        respjson.Field
		Note        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransparenzregisterExtractNewV1ResponseReportValidityUntil) RawJSON() string {
	return r.JSON.raw
}
func (r *TransparenzregisterExtractNewV1ResponseReportValidityUntil) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Transparenzregister extract resource including processing state, parsed report,
// and downloadable documents.
type TransparenzregisterExtractGetV1Response struct {
	// Stable extract identifier. Example: "tre_12345678"
	ID string `json:"id" api:"required"`
	// Status of the Transparenzregister extract.
	//
	// Any of "completed", "processing", "failed".
	Status TransparenzregisterExtractGetV1ResponseStatus `json:"status" api:"required"`
	// Company identifier associated with this extract request. May be null when using
	// sandbox credentials.
	CompanyID string `json:"company_id" api:"nullable"`
	// Timestamp when extract processing completed.
	CompletedAt time.Time `json:"completed_at" api:"nullable" format:"date-time"`
	// URLs for downloading available extract documents.
	Documents []TransparenzregisterExtractGetV1ResponseDocument `json:"documents"`
	// EKRN used to request this extract.
	Ekrn string `json:"ekrn" api:"nullable"`
	// Transparenzregister reference number from the extract.
	ReferenceNumber string `json:"reference_number" api:"nullable"`
	// Parsed Transparenzregister extract report limited to UBO-relevant fields.
	Report TransparenzregisterExtractGetV1ResponseReport `json:"report" api:"nullable"`
	// Timestamp when extract submission started.
	SubmittedAt time.Time `json:"submitted_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Status          respjson.Field
		CompanyID       respjson.Field
		CompletedAt     respjson.Field
		Documents       respjson.Field
		Ekrn            respjson.Field
		ReferenceNumber respjson.Field
		Report          respjson.Field
		SubmittedAt     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransparenzregisterExtractGetV1Response) RawJSON() string { return r.JSON.raw }
func (r *TransparenzregisterExtractGetV1Response) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the Transparenzregister extract.
type TransparenzregisterExtractGetV1ResponseStatus string

const (
	TransparenzregisterExtractGetV1ResponseStatusCompleted  TransparenzregisterExtractGetV1ResponseStatus = "completed"
	TransparenzregisterExtractGetV1ResponseStatusProcessing TransparenzregisterExtractGetV1ResponseStatus = "processing"
	TransparenzregisterExtractGetV1ResponseStatusFailed     TransparenzregisterExtractGetV1ResponseStatus = "failed"
)

// Download URL for a document with format information.
type TransparenzregisterExtractGetV1ResponseDocument struct {
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
func (r TransparenzregisterExtractGetV1ResponseDocument) RawJSON() string { return r.JSON.raw }
func (r *TransparenzregisterExtractGetV1ResponseDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parsed Transparenzregister extract report limited to UBO-relevant fields.
type TransparenzregisterExtractGetV1ResponseReport struct {
	// Extract creation date.
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date"`
	// Reason indicating no natural person UBO could be determined.
	FictionalUboReason string                                               `json:"fictional_ubo_reason" api:"nullable"`
	Groups             []TransparenzregisterExtractGetV1ResponseReportGroup `json:"groups"`
	// Type of Transparenzregister notice.
	NoticeType  string                                                   `json:"notice_type" api:"nullable"`
	StatusFlags TransparenzregisterExtractGetV1ResponseReportStatusFlags `json:"status_flags" api:"nullable"`
	Ubos        []TransparenzregisterExtractGetV1ResponseReportUbo       `json:"ubos"`
	Validity    TransparenzregisterExtractGetV1ResponseReportValidity    `json:"validity" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt          respjson.Field
		FictionalUboReason respjson.Field
		Groups             respjson.Field
		NoticeType         respjson.Field
		StatusFlags        respjson.Field
		Ubos               respjson.Field
		Validity           respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransparenzregisterExtractGetV1ResponseReport) RawJSON() string { return r.JSON.raw }
func (r *TransparenzregisterExtractGetV1ResponseReport) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransparenzregisterExtractGetV1ResponseReportGroup struct {
	Description  string `json:"description" api:"nullable"`
	InterestType string `json:"interest_type" api:"nullable"`
	Position     int64  `json:"position"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description  respjson.Field
		InterestType respjson.Field
		Position     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransparenzregisterExtractGetV1ResponseReportGroup) RawJSON() string { return r.JSON.raw }
func (r *TransparenzregisterExtractGetV1ResponseReportGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransparenzregisterExtractGetV1ResponseReportStatusFlags struct {
	CorrectedByReference string    `json:"corrected_by_reference" api:"nullable"`
	CorrectedReferences  []string  `json:"corrected_references"`
	Deleted              bool      `json:"deleted"`
	DeletionDate         time.Time `json:"deletion_date" api:"nullable" format:"date"`
	DiscrepancyNote      string    `json:"discrepancy_note" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CorrectedByReference respjson.Field
		CorrectedReferences  respjson.Field
		Deleted              respjson.Field
		DeletionDate         respjson.Field
		DiscrepancyNote      respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransparenzregisterExtractGetV1ResponseReportStatusFlags) RawJSON() string { return r.JSON.raw }
func (r *TransparenzregisterExtractGetV1ResponseReportStatusFlags) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransparenzregisterExtractGetV1ResponseReportUbo struct {
	Interest      TransparenzregisterExtractGetV1ResponseReportUboInterest      `json:"interest" api:"nullable"`
	NaturalPerson TransparenzregisterExtractGetV1ResponseReportUboNaturalPerson `json:"natural_person" api:"nullable"`
	Position      int64                                                         `json:"position"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Interest      respjson.Field
		NaturalPerson respjson.Field
		Position      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransparenzregisterExtractGetV1ResponseReportUbo) RawJSON() string { return r.JSON.raw }
func (r *TransparenzregisterExtractGetV1ResponseReportUbo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransparenzregisterExtractGetV1ResponseReportUboInterest struct {
	Percentage float64 `json:"percentage" api:"nullable"`
	Scope      string  `json:"scope" api:"nullable"`
	Type       string  `json:"type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Percentage  respjson.Field
		Scope       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransparenzregisterExtractGetV1ResponseReportUboInterest) RawJSON() string { return r.JSON.raw }
func (r *TransparenzregisterExtractGetV1ResponseReportUboInterest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransparenzregisterExtractGetV1ResponseReportUboNaturalPerson struct {
	City        string    `json:"city" api:"nullable"`
	Country     string    `json:"country" api:"nullable"`
	DateOfBirth time.Time `json:"date_of_birth" api:"nullable" format:"date"`
	FirstName   string    `json:"first_name" api:"nullable"`
	FullName    string    `json:"full_name" api:"nullable"`
	LastName    string    `json:"last_name" api:"nullable"`
	// ISO 3166-1 alpha-2 nationality codes where available.
	Nationalities []string `json:"nationalities"`
	Title         string   `json:"title" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City          respjson.Field
		Country       respjson.Field
		DateOfBirth   respjson.Field
		FirstName     respjson.Field
		FullName      respjson.Field
		LastName      respjson.Field
		Nationalities respjson.Field
		Title         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransparenzregisterExtractGetV1ResponseReportUboNaturalPerson) RawJSON() string {
	return r.JSON.raw
}
func (r *TransparenzregisterExtractGetV1ResponseReportUboNaturalPerson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransparenzregisterExtractGetV1ResponseReportValidity struct {
	From  TransparenzregisterExtractGetV1ResponseReportValidityFrom  `json:"from" api:"nullable"`
	Until TransparenzregisterExtractGetV1ResponseReportValidityUntil `json:"until" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From        respjson.Field
		Until       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransparenzregisterExtractGetV1ResponseReportValidity) RawJSON() string { return r.JSON.raw }
func (r *TransparenzregisterExtractGetV1ResponseReportValidity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransparenzregisterExtractGetV1ResponseReportValidityFrom struct {
	Date time.Time `json:"date" api:"nullable" format:"date"`
	Note string    `json:"note" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date        respjson.Field
		Note        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransparenzregisterExtractGetV1ResponseReportValidityFrom) RawJSON() string {
	return r.JSON.raw
}
func (r *TransparenzregisterExtractGetV1ResponseReportValidityFrom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransparenzregisterExtractGetV1ResponseReportValidityUntil struct {
	Date time.Time `json:"date" api:"nullable" format:"date"`
	Note string    `json:"note" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date        respjson.Field
		Note        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransparenzregisterExtractGetV1ResponseReportValidityUntil) RawJSON() string {
	return r.JSON.raw
}
func (r *TransparenzregisterExtractGetV1ResponseReportValidityUntil) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransparenzregisterExtractNewV1Params struct {
	// Unique company identifier. Required unless `X-Credential-Name` is set to
	// `sandbox`. In sandbox mode this field should be omitted. Example:
	// DE-HRB-F1103-267645
	CompanyID       param.Opt[string] `json:"company_id,omitzero"`
	XCredentialName param.Opt[string] `header:"X-Credential-Name,omitzero" json:"-"`
	paramObj
}

func (r TransparenzregisterExtractNewV1Params) MarshalJSON() (data []byte, err error) {
	type shadow TransparenzregisterExtractNewV1Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransparenzregisterExtractNewV1Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
