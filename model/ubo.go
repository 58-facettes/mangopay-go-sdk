package model

import (
	"github.com/58-facettes/mangopay-go-sdk/model/country"
)

// UBO is a Declaration is an electronic version of the previous KYC document "Shareholder Declaration",
// in order to declare all the Ultimate Beneficial Owners of a BUSINESS-typed
// legal User (ie the shareholders with >25% of capital or voting rights).
type UBO struct {
	// The item's ID
	ID string `json:"Id"`
	// When the item was created
	CreationDate int64 `json:"CreationDate"`
	// The date when the document was processed by MANGOPAY
	ProcessedDate int64 `json:"ProcessedDate"`
	// Status of a UBO Declaration
	Status UBODeclarationStatus `json:"Status"`
	// Refused or incomplete Reason for a UBO Declaration
	Reason UBOReason `json:"Reason"`
	// Refused or incomplete Message for a UBO Declaration
	Message string `json:"Message"`
	// An array of UBOs declared as Ultimate Beneficial Owners of a BUSINESS Legal User.
	// ?? this field is not describe in the documentation ??
	// Ubos DeclaredUbos
	// ?? end ??
}

// UBODeclarationStatus is for UBOs declaration type.
type UBODeclarationStatus string

const (
	// UBODeclarationStatusCreated is for an UBO status created.
	UBODeclarationStatusCreated UBODeclarationStatus = "CREATED"
	// UBODeclarationStatusValidationAsked is for an UBO status validation asked.
	UBODeclarationStatusValidationAsked UBODeclarationStatus = "VALIDATION_ASKED"
	// UBODeclarationStatusIncomplete is for an UBO status incomplete.
	UBODeclarationStatusIncomplete UBODeclarationStatus = "INCOMPLETE"
	// UBODeclarationStatusValidated is for an UBO status validated.
	UBODeclarationStatusValidated UBODeclarationStatus = "VALIDATED"
	// UBODeclarationStatusRefused is for an UBO status refused.
	UBODeclarationStatusRefused UBODeclarationStatus = "REFUSED"
)

// UBOReason is a reason type for UBO declaration
type UBOReason string

const (
	// UBOReasonMissing stands for MISSING_UBO.
	UBOReasonMissing UBOReason = "MISSING_UBO"
	// UBOReasonWrongInformation stands for WRONG_UBO_INFORMATION.
	UBOReasonWrongInformation UBOReason = "WRONG_UBO_INFORMATION"
	// UBOReasonIdentityNeeded stands for UBO_IDENTITY_NEEDED.
	UBOReasonIdentityNeeded UBOReason = "UBO_IDENTITY_NEEDED"
	// UBOReasonShareHolderDeclarationNeeded stands forSHAREHOLDERS_DECLARATION_NEEDED.
	UBOReasonShareHolderDeclarationNeeded UBOReason = "SHAREHOLDERS_DECLARATION_NEEDED"
	// UBOReasonOrganizationChartNeeded stands for ORGANIZATION_CHART_NEEDED.
	UBOReasonOrganizationChartNeeded UBOReason = "ORGANIZATION_CHART_NEEDED"
	// UBOReasonOrganizationDocumentNeeded stands for DOCUMENTS_NEEDED.
	UBOReasonOrganizationDocumentNeeded UBOReason = "DOCUMENTS_NEEDED"
	// UBOReasonDeclarationNotMatch stands for DECLARATION_DO_NOT_MATCH_UBO_INFORMATION.
	UBOReasonDeclarationNotMatch UBOReason = "DECLARATION_DO_NOT_MATCH_UBO_INFORMATION"
	// UBOReasonSpecificCase stands for SPECIFIC_CASE.
	UBOReasonSpecificCase UBOReason = "SPECIFIC_CASE"
)

// UBOCreate is for creating a new UBO.
type UBOCreate struct {
	// The name of the UBO REQUIRED.
	FirstName string `json:"FirstName"`
	// The last name of the UBO REQUIRED.
	LastName string `json:"LastName"`
	// The address REQUIRED.
	Address Address `json:"Address"`
	// The UBO's nationality. ISO 3166-1 alpha-2 format is expected REQUIRED.
	Nationality country.ISO2 `json:"Nationality"`
	// The date of birth of the UBO - be careful to set the right timezone (should be UTC) REQUIRED.
	// to avoid 00h becoming 23h (and hence interpreted as the day before)
	Birthday int64 `json:"Birthday"`
	// The UBO's birthplace REQUIRED.
	Birthplace Birthplace `json:"Birthplace"`
}

// Birthplace is use to describe the place of birth.
type Birthplace struct {
	// The city of the address REQUIRED.
	City string `json:"City"`
	// The Country of the Address REQUIRED.
	Country country.ISO2 `json:"Country"`
}

// UBOUpdate is used to update an existing UBO.
type UBOUpdate struct {
	// The name of the UBO OPTIONAL.
	FirstName string `json:"FirstName,omitempty"`
	// The last name of the UBO OPTIONAL.
	LastName string `json:"LastName,omitempty"`
	// The address OPTIONAL.
	Address Address `json:"Address,omitempty"`
	// The UBO's nationality. ISO 3166-1 alpha-2 format is expected OPTIONAL.
	Nationality country.ISO2 `json:"Nationality,omitempty"`
	// The date of birth of the UBO - be careful to set the right timezone (should be UTC)
	// to avoid 00h becoming 23h (and hence interpreted as the day before) OPTIONAL.
	Birthday int64 `json:"Birthday,omitempty"`
	// The UBO's birthplace OPTIONAL.
	Birthplace Birthplace `json:"Birthplace,omitempty"`
	// To remove a UBO from your UBO declaration only put isActive to "false". You cannot activate back a UBO OPTIONAL.
	IsActive bool `json:"isActive,omitempty"`
}
