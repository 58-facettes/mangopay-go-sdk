package model

import (
	"encoding/base64"
	"io/ioutil"
)

type KYCDocument struct {
	// ?? some fields may missing ??
	// ID of the current card's registration.
	ID string `json:"Id"` // ?? new ??
	// Tag of the card.
	Tag string `json:"Tag"` // ?? new ??
	// CreationDate
	CreationDate int64 `json:"CreationDate"` // ?? new ??
	// ?? end ??

	// The object owner's UserId
	UserID string `json:"UserId"`
	// The status of this KYC/Dispute document
	Status DocumentStatus `json:"Status"`
	// The message accompanying a refusal
	RefusedReasonMessage string `json:"RefusedReasonMessage"`
	// The type of reason for refusal
	RefusedReason RefusedReason `json:"RefusedReason"`
	// The date when the document was processed by MANGOPAY
	ProcessedDate int64 `json:"ProcessedDate"`
}

// DocumentStatus is the Status of the Document
type DocumentStatus string

const (
	DocumentStatusCreated         DocumentStatus = "CREATED"
	DocumentStatusValidationAsked DocumentStatus = "VALIDATION_ASKED"
	DocumentStatusValidated       DocumentStatus = "VALIDATED"
	DocumentStatusRefused         DocumentStatus = "REFUSED"
)

// RefusedReason The type of refusal reason.
type RefusedReason string

const (
	RefusedReasonDocUnreadable          RefusedReason = "DOCUMENT_UNREADABLE"
	RefusedReasonDocNotAccepted         RefusedReason = "DOCUMENT_NOT_ACCEPTED"
	RefusedReasonDocHasExpired          RefusedReason = "DOCUMENT_HAS_EXPIRED"
	RefusedReasonDocIncomplete          RefusedReason = "DOCUMENT_INCOMPLETE"
	RefusedReasonDocMissing             RefusedReason = "DOCUMENT_MISSING"
	RefusedReasonDocNotMatchWithUser    RefusedReason = "DOCUMENT_DO_NOT_MATCH_USER_DATA"
	RefusedReasonDocNotMatchWithAccount RefusedReason = "DOCUMENT_DO_NOT_MATCH_ACCOUNT_DATA"
	RefusedReasonDocFalsified           RefusedReason = "DOCUMENT_FALSIFIED"
	RefusedReasonSecificCase            RefusedReason = "SPECIFIC_CASE"
	RefusedReasonUnderAgePerson         RefusedReason = "UNDERAGE_PERSON"
)

type KYCDocumentCreate struct {
}

type DocumentType string

const (
	DocumentTypeIdentity             DocumentType = "IDENTITY_PROOF"
	DocumentTypeRegistration         DocumentType = "REGISTRATION_PROOF"
	DocumentTypeArticleOfAssociation DocumentType = "ARTICLES_OF_ASSOCIATION"
	DocumentTypeShareHolder          DocumentType = "SHAREHOLDER_DECLARATION"
	DocumentTypeAddress              DocumentType = "ADDRESS_PROOF"
)

type KYCPage struct {
	File string `json:"File"`
}

// NewKYCPage is reading from the given filepath the file and insert it into the KYCPage struct encoded in base64.
// The maximum size per page is about 7MB.
// The following formats are accepted : .pdf, .jpeg, .jpg, .gif and .png.
// The minimum size is 1Kb.
func NewKYCPage(filepath string) (*KYCPage, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return &KYCPage{
		File: base64.StdEncoding.EncodeToString(data),
	}, nil
}
