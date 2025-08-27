// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openregister

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/oregister/openregister-go/internal/apijson"
	"github.com/oregister/openregister-go/internal/requestconfig"
	"github.com/oregister/openregister-go/option"
	"github.com/oregister/openregister-go/packages/respjson"
)

// PersonService contains methods and other services that help with interacting
// with the openregister API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPersonService] method instead.
type PersonService struct {
	Options []option.RequestOption
}

// NewPersonService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPersonService(opts ...option.RequestOption) (r PersonService) {
	r = PersonService{}
	r.Options = opts
	return
}

// Get detailed person information
func (r *PersonService) GetDetailsV1(ctx context.Context, personID string, opts ...option.RequestOption) (res *PersonGetDetailsV1Response, err error) {
	opts = append(r.Options[:], opts...)
	if personID == "" {
		err = errors.New("missing required person_id parameter")
		return
	}
	path := fmt.Sprintf("v1/person/%s", personID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get person holdings
func (r *PersonService) GetHoldingsV1(ctx context.Context, personID string, opts ...option.RequestOption) (res *PersonGetHoldingsV1Response, err error) {
	opts = append(r.Options[:], opts...)
	if personID == "" {
		err = errors.New("missing required person_id parameter")
		return
	}
	path := fmt.Sprintf("v1/person/%s/holdings", personID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PersonGetDetailsV1Response struct {
	// Unique person identifier. Example: cc78ab54-d958-49b8-bae7-2f6c0c308837
	ID string `json:"id,required"`
	// Age of the person.
	Age int64 `json:"age,required"`
	// City of the person.
	City string `json:"city,required"`
	// Date of birth of the person. Format: ISO 8601 (YYYY-MM-DD) Example: "1990-01-01"
	DateOfBirth time.Time `json:"date_of_birth,required" format:"date-time"`
	// First name of the person.
	FirstName string `json:"first_name,required"`
	// Last name of the person.
	LastName string `json:"last_name,required"`
	// Management positions of the person.
	ManagementPositions []PersonGetDetailsV1ResponseManagementPosition `json:"management_positions,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Age                 respjson.Field
		City                respjson.Field
		DateOfBirth         respjson.Field
		FirstName           respjson.Field
		LastName            respjson.Field
		ManagementPositions respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonGetDetailsV1Response) RawJSON() string { return r.JSON.raw }
func (r *PersonGetDetailsV1Response) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// All current and past management positions of the person.
type PersonGetDetailsV1ResponseManagementPosition struct {
	// Name of the company. Example: "Descartes Technologies GmbH"
	CompanyName string `json:"company_name,required"`
	// Register ID of the company. Example: DE-HRB-F1103-267645
	RegisterID string `json:"register_id,required"`
	// Role of the person in the company. Example: "DIRECTOR"
	Role string `json:"role,required"`
	// Date when the person ended the management position. Format: ISO 8601
	// (YYYY-MM-DD) Example: "2023-01-01"
	EndDate time.Time `json:"end_date" format:"date-time"`
	// Date when the person started the management position. Format: ISO 8601
	// (YYYY-MM-DD) Example: "2022-01-01"
	StartDate time.Time `json:"start_date" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompanyName respjson.Field
		RegisterID  respjson.Field
		Role        respjson.Field
		EndDate     respjson.Field
		StartDate   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonGetDetailsV1ResponseManagementPosition) RawJSON() string { return r.JSON.raw }
func (r *PersonGetDetailsV1ResponseManagementPosition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Companies this entity owns or has invested in.
type PersonGetHoldingsV1Response struct {
	// Shareholder and limited partner positions of the person.
	Holdings []PersonGetHoldingsV1ResponseHolding `json:"holdings,required"`
	// Unique person identifier. Example: cc78ab54-d958-49b8-bae7-2f6c0c308837
	PersonID string `json:"person_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Holdings    respjson.Field
		PersonID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PersonGetHoldingsV1Response) RawJSON() string { return r.JSON.raw }
func (r *PersonGetHoldingsV1Response) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonGetHoldingsV1ResponseHolding struct {
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
func (r PersonGetHoldingsV1ResponseHolding) RawJSON() string { return r.JSON.raw }
func (r *PersonGetHoldingsV1ResponseHolding) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
