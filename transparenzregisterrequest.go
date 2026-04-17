// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openregister

import (
	"github.com/oregister/openregister-go/v2/option"
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
