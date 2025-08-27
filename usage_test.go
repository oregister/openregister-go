// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openregister_test

import (
	"context"
	"os"
	"testing"

	"github.com/oregister/openregister-go"
	"github.com/oregister/openregister-go/internal/testutil"
	"github.com/oregister/openregister-go/option"
)

func TestUsage(t *testing.T) {
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
	response, err := client.Search.AutocompleteCompaniesV1(context.TODO(), openregister.SearchAutocompleteCompaniesV1Params{
		Query: "query",
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", response.Results)
}
