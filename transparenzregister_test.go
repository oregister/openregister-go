// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openregister_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/oregister/openregister-go/v2"
	"github.com/oregister/openregister-go/v2/internal/testutil"
	"github.com/oregister/openregister-go/v2/option"
)

func TestTransparenzregisterSetCredentialsV1WithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := openregister.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Transparenzregister.SetCredentialsV1(context.TODO(), openregister.TransparenzregisterSetCredentialsV1Params{
		Password: "password",
		Username: "username",
		Name:     openregister.String("name"),
	})
	if err != nil {
		var apierr *openregister.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
