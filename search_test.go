// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openregister_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/oregister/openregister-go"
	"github.com/oregister/openregister-go/internal/testutil"
	"github.com/oregister/openregister-go/option"
)

func TestSearchFindCompaniesV0WithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Search.FindCompaniesV0(context.TODO(), openregister.SearchFindCompaniesV0Params{
		Active:            openregister.Bool(true),
		IncorporationDate: openregister.String("incorporation_date"),
		LegalForm:         openregister.CompanyLegalFormAg,
		Page:              openregister.Int(0),
		PerPage:           openregister.Int(0),
		Query:             openregister.String("query"),
		RegisterCourt:     openregister.String("register_court"),
		RegisterNumber:    openregister.String("register_number"),
		RegisterType:      openregister.CompanyRegisterTypeHrb,
	})
	if err != nil {
		var apierr *openregister.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSearchLookupCompanyByURL(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Search.LookupCompanyByURL(context.TODO(), openregister.SearchLookupCompanyByURLParams{
		URL: "https://example.com",
	})
	if err != nil {
		var apierr *openregister.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
