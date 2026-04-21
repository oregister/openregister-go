// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openregister

import (
	"context"
	"net/http"
	"slices"

	"github.com/oregister/openregister-go/v2/internal/apijson"
	"github.com/oregister/openregister-go/v2/internal/requestconfig"
	"github.com/oregister/openregister-go/v2/option"
	"github.com/oregister/openregister-go/v2/packages/param"
)

// TransparenzregisterService contains methods and other services that help with
// interacting with the openregister API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransparenzregisterService] method instead.
type TransparenzregisterService struct {
	Options []option.RequestOption
	Extract TransparenzregisterExtractService
}

// NewTransparenzregisterService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTransparenzregisterService(opts ...option.RequestOption) (r TransparenzregisterService) {
	r = TransparenzregisterService{}
	r.Options = opts
	r.Extract = NewTransparenzregisterExtractService(opts...)
	return
}

// Store username and password credentials for accessing the Transparenzregister
// API. These credentials will be used for subsequent requests to retrieve company
// documents. Credential names are user-scoped; the reserved name `sandbox` cannot
// be used. Credentials are validated against Transparenzregister before they are
// persisted.
func (r *TransparenzregisterService) SetCredentialsV1(ctx context.Context, body TransparenzregisterSetCredentialsV1Params, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v1/transparenzregister/credentials"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

type TransparenzregisterSetCredentialsV1Params struct {
	// Password for Transparenzregister API access.
	Password string `json:"password" api:"required" format:"password"`
	// Username for Transparenzregister API access. Example: "compliance@example.com"
	Username string `json:"username" api:"required"`
	// Name to identify this set of credentials. Allows storing multiple
	// Transparenzregister credentials per user (e.g., for different accounts or
	// clients). Defaults to 'default' if not provided. Cannot be `sandbox` because
	// that name is reserved for test-mode extracts. Example: "client_a"
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r TransparenzregisterSetCredentialsV1Params) MarshalJSON() (data []byte, err error) {
	type shadow TransparenzregisterSetCredentialsV1Params
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransparenzregisterSetCredentialsV1Params) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
