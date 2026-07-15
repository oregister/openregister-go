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
	"github.com/oregister/openregister-go/v2/packages/respjson"
)

// InsolvencyService contains methods and other services that help with interacting
// with the openregister API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInsolvencyService] method instead.
type InsolvencyService struct {
	Options []option.RequestOption
}

// NewInsolvencyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInsolvencyService(opts ...option.RequestOption) (r InsolvencyService) {
	r = InsolvencyService{}
	r.Options = opts
	return
}

// Get detailed insolvency proceeding information
func (r *InsolvencyService) Get(ctx context.Context, insolvencyID string, opts ...option.RequestOption) (res *InsolvencyGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if insolvencyID == "" {
		err = errors.New("missing required insolvency_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/insolvency/%s", insolvencyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Kind of administration ordered for the proceeding.
type InsolvencyAdministrationKind string

const (
	InsolvencyAdministrationKindExternalAdministration InsolvencyAdministrationKind = "external_administration"
	InsolvencyAdministrationKindSelfAdministration     InsolvencyAdministrationKind = "self_administration"
	InsolvencyAdministrationKindProtectiveShield       InsolvencyAdministrationKind = "protective_shield"
)

// Kind of debtor the proceeding concerns.
//
// - legal_person: legal entities (companies, associations, etc.)
// - natural_person: private individuals
type InsolvencyDebtorKind string

const (
	InsolvencyDebtorKindLegalPerson   InsolvencyDebtorKind = "legal_person"
	InsolvencyDebtorKindNaturalPerson InsolvencyDebtorKind = "natural_person"
)

// Kind of insolvency proceeding.
type InsolvencyProceedingKind string

const (
	InsolvencyProceedingKindRegularInsolvency  InsolvencyProceedingKind = "regular_insolvency"
	InsolvencyProceedingKindConsumerInsolvency InsolvencyProceedingKind = "consumer_insolvency"
)

// Current status of the insolvency proceeding.
type InsolvencyStatus string

const (
	InsolvencyStatusPreliminary      InsolvencyStatus = "preliminary"
	InsolvencyStatusOpened           InsolvencyStatus = "opened"
	InsolvencyStatusRejectedNoAssets InsolvencyStatus = "rejected_no_assets"
	InsolvencyStatusMassInsufficient InsolvencyStatus = "mass_insufficient"
	InsolvencyStatusPlanSupervised   InsolvencyStatus = "plan_supervised"
	InsolvencyStatusLifted           InsolvencyStatus = "lifted"
	InsolvencyStatusDiscontinued     InsolvencyStatus = "discontinued"
	InsolvencyStatusDischargePending InsolvencyStatus = "discharge_pending"
	InsolvencyStatusDischargeGranted InsolvencyStatus = "discharge_granted"
	InsolvencyStatusDischargeDenied  InsolvencyStatus = "discharge_denied"
	InsolvencyStatusDischargeRevoked InsolvencyStatus = "discharge_revoked"
	InsolvencyStatusUnknown          InsolvencyStatus = "unknown"
)

// An insolvency proceeding with all of its published events.
type InsolvencyGetResponse struct {
	// Unique identifier of the insolvency proceeding.
	ID string `json:"id" api:"required" format:"uuid"`
	// Case number of the proceeding at the court. Example: "36d IN 3382/25"
	CaseNumber string `json:"case_number" api:"required"`
	// Unique company identifier of the debtor, if the debtor is a registered company.
	// Example: DE-HRB-F1103-267645
	CompanyID string `json:"company_id" api:"required"`
	// Insolvency court handling the proceeding.
	Court string `json:"court" api:"required"`
	// Current status of the insolvency proceeding.
	//
	// Any of "preliminary", "opened", "rejected_no_assets", "mass_insufficient",
	// "plan_supervised", "lifted", "discontinued", "discharge_pending",
	// "discharge_granted", "discharge_denied", "discharge_revoked", "unknown".
	CurrentStatus InsolvencyStatus `json:"current_status" api:"required"`
	// Name of the debtor as published by the court.
	DebtorName string `json:"debtor_name" api:"required"`
	// All published events of the proceeding, ordered by date.
	Events []InsolvencyGetResponseEvent `json:"events" api:"required"`
	// Grounds for the insolvency (e.g. Zahlungsunfähigkeit, Überschuldung).
	InsolvencyGrounds []string `json:"insolvency_grounds" api:"required"`
	// Kind of administration ordered for the proceeding.
	//
	// Any of "external_administration", "self_administration", "protective_shield".
	AdministrationKind InsolvencyAdministrationKind `json:"administration_kind" api:"nullable"`
	// Address of the insolvency administrator.
	AdministratorAddress string `json:"administrator_address" api:"nullable"`
	// Name of the insolvency administrator.
	AdministratorName string `json:"administrator_name" api:"nullable"`
	// Deadline for creditors to file their claims.
	ClaimsFilingDeadline time.Time `json:"claims_filing_deadline" api:"nullable" format:"date-time"`
	// Date the proceeding was closed.
	ClosedAt time.Time `json:"closed_at" api:"nullable" format:"date-time"`
	// Kind of debtor the proceeding concerns.
	//
	// - legal_person: legal entities (companies, associations, etc.)
	// - natural_person: private individuals
	//
	// Any of "legal_person", "natural_person".
	DebtorKind InsolvencyDebtorKind `json:"debtor_kind" api:"nullable"`
	// Legal form of the debtor as published by the court.
	DebtorLegalForm string `json:"debtor_legal_form" api:"nullable"`
	// Amount available for distribution, in euros.
	DistributionAvailable float64 `json:"distribution_available" api:"nullable"`
	// Total registered claims in the distribution, in euros.
	DistributionClaimsTotal float64 `json:"distribution_claims_total" api:"nullable"`
	// Publication date of the first known event of the proceeding.
	FirstEventAt time.Time `json:"first_event_at" api:"nullable" format:"date-time"`
	// Publication date of the most recent known event of the proceeding.
	LastEventAt time.Time `json:"last_event_at" api:"nullable" format:"date-time"`
	// Date the proceeding was opened.
	OpenedAt time.Time `json:"opened_at" api:"nullable" format:"date-time"`
	// Kind of insolvency proceeding.
	//
	// Any of "regular_insolvency", "consumer_insolvency".
	ProceedingKind InsolvencyProceedingKind `json:"proceeding_kind" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		CaseNumber              respjson.Field
		CompanyID               respjson.Field
		Court                   respjson.Field
		CurrentStatus           respjson.Field
		DebtorName              respjson.Field
		Events                  respjson.Field
		InsolvencyGrounds       respjson.Field
		AdministrationKind      respjson.Field
		AdministratorAddress    respjson.Field
		AdministratorName       respjson.Field
		ClaimsFilingDeadline    respjson.Field
		ClosedAt                respjson.Field
		DebtorKind              respjson.Field
		DebtorLegalForm         respjson.Field
		DistributionAvailable   respjson.Field
		DistributionClaimsTotal respjson.Field
		FirstEventAt            respjson.Field
		LastEventAt             respjson.Field
		OpenedAt                respjson.Field
		ProceedingKind          respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsolvencyGetResponse) RawJSON() string { return r.JSON.raw }
func (r *InsolvencyGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single published event of an insolvency proceeding, derived from an official
// court publication.
type InsolvencyGetResponseEvent struct {
	// Unique identifier of the event.
	ID string `json:"id" api:"required" format:"uuid"`
	// Structured details extracted from the publication, where available.
	Details InsolvencyGetResponseEventDetails `json:"details" api:"required"`
	// Lifecycle event type of the insolvency proceeding.
	//
	// Any of "preliminary_measures", "preliminary_measures_lifted",
	// "proceedings_opened", "rejected_insufficient_assets", "in_proceeding_decision",
	// "mass_insufficiency_notified", "distribution_announced",
	// "distribution_record_filed", "final_distribution_announced",
	// "proceedings_lifted", "proceedings_discontinued", "proceedings_terminated",
	// "post_termination_decision", "subsequent_distribution_ordered",
	// "discharge_pending", "discharge_granted", "discharge_denied",
	// "discharge_revoked", "plan_confirmed", "plan_supervision_ordered",
	// "plan_supervision_terminated", "other".
	EventType string `json:"event_type" api:"required"`
	// Date the event was published by the court.
	PublishedAt time.Time `json:"published_at" api:"required" format:"date-time"`
	// Category of the official publication the event was derived from.
	//
	// Any of "security_measures", "rejection_for_insufficiency_of_assets", "openings",
	// "decisions_in_proceedings", "misc", "decisions_after_termination",
	// "distribution_lists", "decisions_in_discharge_proceedings",
	// "supervised_insolvency_plans".
	ReportType string `json:"report_type" api:"required"`
	// Short summary of the publication.
	Summary string `json:"summary" api:"required"`
	// Date of the court decision, if published.
	DecisionDate time.Time `json:"decision_date" api:"nullable" format:"date-time"`
	// Date the decision takes effect, if published.
	EffectiveAt time.Time `json:"effective_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Details      respjson.Field
		EventType    respjson.Field
		PublishedAt  respjson.Field
		ReportType   respjson.Field
		Summary      respjson.Field
		DecisionDate respjson.Field
		EffectiveAt  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsolvencyGetResponseEvent) RawJSON() string { return r.JSON.raw }
func (r *InsolvencyGetResponseEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Structured details extracted from the publication, where available.
type InsolvencyGetResponseEventDetails struct {
	InsolvencyGrounds       []string                                   `json:"insolvency_grounds" api:"required"`
	Meetings                []InsolvencyGetResponseEventDetailsMeeting `json:"meetings" api:"required"`
	AdministratorAddress    string                                     `json:"administrator_address" api:"nullable"`
	AdministratorName       string                                     `json:"administrator_name" api:"nullable"`
	ClaimsFilingDeadline    string                                     `json:"claims_filing_deadline" api:"nullable"`
	ClosedAt                string                                     `json:"closed_at" api:"nullable"`
	DischargeGranted        bool                                       `json:"discharge_granted" api:"nullable"`
	DistributionAvailable   string                                     `json:"distribution_available" api:"nullable"`
	DistributionClaimsTotal string                                     `json:"distribution_claims_total" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InsolvencyGrounds       respjson.Field
		Meetings                respjson.Field
		AdministratorAddress    respjson.Field
		AdministratorName       respjson.Field
		ClaimsFilingDeadline    respjson.Field
		ClosedAt                respjson.Field
		DischargeGranted        respjson.Field
		DistributionAvailable   respjson.Field
		DistributionClaimsTotal respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsolvencyGetResponseEventDetails) RawJSON() string { return r.JSON.raw }
func (r *InsolvencyGetResponseEventDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A creditor meeting or hearing announced in the publication.
type InsolvencyGetResponseEventDetailsMeeting struct {
	Kind     string `json:"kind" api:"required"`
	At       string `json:"at" api:"nullable"`
	Location string `json:"location" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Kind        respjson.Field
		At          respjson.Field
		Location    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InsolvencyGetResponseEventDetailsMeeting) RawJSON() string { return r.JSON.raw }
func (r *InsolvencyGetResponseEventDetailsMeeting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
