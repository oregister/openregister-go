// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openregister

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/oregister/openregister-go/v2/internal/apijson"
	"github.com/oregister/openregister-go/v2/internal/requestconfig"
	"github.com/oregister/openregister-go/v2/option"
	"github.com/oregister/openregister-go/v2/packages/param"
	"github.com/oregister/openregister-go/v2/packages/respjson"
)

// MonitorService contains methods and other services that help with interacting
// with the openregister API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMonitorService] method instead.
type MonitorService struct {
	Options []option.RequestOption
}

// NewMonitorService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMonitorService(opts ...option.RequestOption) (r MonitorService) {
	r = MonitorService{}
	r.Options = opts
	return
}

// Create webhook monitor item
func (r *MonitorService) New(ctx context.Context, body MonitorNewParams, opts ...option.RequestOption) (res *MonitorNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/monitor"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List webhook monitor items
func (r *MonitorService) List(ctx context.Context, opts ...option.RequestOption) (res *MonitorListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/monitor"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete webhook monitor item
func (r *MonitorService) Delete(ctx context.Context, entityID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return err
	}
	path := fmt.Sprintf("v1/monitor/%s", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type MonitorNewResponse struct {
	// Whether the monitor item is disabled.
	Disabled bool `json:"disabled" api:"required"`
	// For `company` this is the register ID (e.g. `DE-HRB-F1103-267645`). For `person`
	// this is the person UUID.
	EntityID string `json:"entity_id" api:"required"`
	// Type of the entity to monitor.
	//
	// Any of "company", "person".
	EntityType MonitorNewResponseEntityType `json:"entity_type" api:"required"`
	// Preferences for the entity to monitor. Use `WebhookMonitorCompanyPreference`
	// values when `entity_type` is `company`, and `WebhookMonitorPersonPreference`
	// values when `entity_type` is `person`.
	//
	// Any of "basic", "representation", "financials", "documents", "ownership",
	// "holdings", "management_positions".
	Preferences []string `json:"preferences" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Disabled    respjson.Field
		EntityID    respjson.Field
		EntityType  respjson.Field
		Preferences respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorNewResponse) RawJSON() string { return r.JSON.raw }
func (r *MonitorNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the entity to monitor.
type MonitorNewResponseEntityType string

const (
	MonitorNewResponseEntityTypeCompany MonitorNewResponseEntityType = "company"
	MonitorNewResponseEntityTypePerson  MonitorNewResponseEntityType = "person"
)

type MonitorListResponse struct {
	// List of webhook monitor items.
	Items []MonitorListResponseItem `json:"items" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListResponse) RawJSON() string { return r.JSON.raw }
func (r *MonitorListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorListResponseItem struct {
	// Whether the monitor item is disabled.
	Disabled bool `json:"disabled" api:"required"`
	// For `company` this is the register ID (e.g. `DE-HRB-F1103-267645`). For `person`
	// this is the person UUID.
	EntityID string `json:"entity_id" api:"required"`
	// Type of the entity to monitor.
	//
	// Any of "company", "person".
	EntityType string `json:"entity_type" api:"required"`
	// Preferences for the entity to monitor. Use `WebhookMonitorCompanyPreference`
	// values when `entity_type` is `company`, and `WebhookMonitorPersonPreference`
	// values when `entity_type` is `person`.
	//
	// Any of "basic", "representation", "financials", "documents", "ownership",
	// "holdings", "management_positions".
	Preferences []string `json:"preferences" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Disabled    respjson.Field
		EntityID    respjson.Field
		EntityType  respjson.Field
		Preferences respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MonitorListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *MonitorListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MonitorNewParams struct {
	// For `company` this is the register ID (e.g. `DE-HRB-F1103-267645`). For `person`
	// this is the person UUID.
	EntityID string `json:"entity_id" api:"required"`
	// Type of the entity to monitor.
	//
	// Any of "company", "person".
	EntityType MonitorNewParamsEntityType `json:"entity_type,omitzero" api:"required"`
	// Preferences for the entity to monitor. Use `WebhookMonitorCompanyPreference`
	// values when `entity_type` is `company`, and `WebhookMonitorPersonPreference`
	// values when `entity_type` is `person`.
	//
	// Any of "basic", "representation", "financials", "documents", "ownership",
	// "holdings", "management_positions".
	Preferences []string `json:"preferences,omitzero" api:"required"`
	paramObj
}

func (r MonitorNewParams) MarshalJSON() (data []byte, err error) {
	type shadow MonitorNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MonitorNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the entity to monitor.
type MonitorNewParamsEntityType string

const (
	MonitorNewParamsEntityTypeCompany MonitorNewParamsEntityType = "company"
	MonitorNewParamsEntityTypePerson  MonitorNewParamsEntityType = "person"
)
