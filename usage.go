// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openregister

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/oregister/openregister-go/v2/internal/apijson"
	"github.com/oregister/openregister-go/v2/internal/requestconfig"
	"github.com/oregister/openregister-go/v2/option"
	"github.com/oregister/openregister-go/v2/packages/respjson"
)

// UsageService contains methods and other services that help with interacting with
// the openregister API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUsageService] method instead.
type UsageService struct {
	Options []option.RequestOption
}

// NewUsageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUsageService(opts ...option.RequestOption) (r UsageService) {
	r = UsageService{}
	r.Options = opts
	return
}

// Retrieve public API credit usage
func (r *UsageService) GetCreditsV1(ctx context.Context, opts ...option.RequestOption) (res *UsageGetCreditsV1Response, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/credits"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type UsageGetCreditsV1Response struct {
	IncludedCredits int64 `json:"included_credits" api:"required"`
	// Credits above the included allowance.
	OverageCredits int64                           `json:"overage_credits" api:"required"`
	Paid           bool                            `json:"paid" api:"required"`
	Period         UsageGetCreditsV1ResponsePeriod `json:"period" api:"required"`
	// Never negative; zero once usage exceeds included credits.
	RemainingCredits int64 `json:"remaining_credits" api:"required"`
	UsedCredits      int64 `json:"used_credits" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IncludedCredits  respjson.Field
		OverageCredits   respjson.Field
		Paid             respjson.Field
		Period           respjson.Field
		RemainingCredits respjson.Field
		UsedCredits      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageGetCreditsV1Response) RawJSON() string { return r.JSON.raw }
func (r *UsageGetCreditsV1Response) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageGetCreditsV1ResponsePeriod struct {
	ResetAt time.Time `json:"reset_at" api:"required" format:"date-time"`
	// billing_cycle for paid plans; rolling_30_days for the free plan, where the
	// window starts with the first request.
	//
	// Any of "billing_cycle", "rolling_30_days".
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ResetAt     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageGetCreditsV1ResponsePeriod) RawJSON() string { return r.JSON.raw }
func (r *UsageGetCreditsV1ResponsePeriod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
