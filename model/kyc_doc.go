package model

import (
	"encoding/base64"
	"io/ioutil"
)

// KYCDocument describe a KYC document.
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
	// DocumentStatusCreated stands for a status created.
	DocumentStatusCreated DocumentStatus = "CREATED"
	// DocumentStatusValidationAsked stands for a status validation asked.
	DocumentStatusValidationAsked DocumentStatus = "VALIDATION_ASKED"
	// DocumentStatusValidated stands for a status validated.
	DocumentStatusValidated DocumentStatus = "VALIDATED"
	// DocumentStatusRefused stands for a status refused.
	DocumentStatusRefused DocumentStatus = "REFUSED"
)

// RefusedReason The type of refusal reason.
type RefusedReason string

const (
	// RefusedReasonDocUnreadable stands for a reason document unreadable.
	RefusedReasonDocUnreadable RefusedReason = "DOCUMENT_UNREADABLE"
	// RefusedReasonDocNotAccepted stands for a reason document not accepted.
	RefusedReasonDocNotAccepted RefusedReason = "DOCUMENT_NOT_ACCEPTED"
	// RefusedReasonDocHasExpired stands for a reason document has exired.
	RefusedReasonDocHasExpired RefusedReason = "DOCUMENT_HAS_EXPIRED"
	// RefusedReasonDocIncomplete stands for a reason document incomplete.
	RefusedReasonDocIncomplete RefusedReason = "DOCUMENT_INCOMPLETE"
	// RefusedReasonDocMissing stands for a reason document missing.
	RefusedReasonDocMissing RefusedReason = "DOCUMENT_MISSING"
	// RefusedReasonDocNotMatchWithUser stands for a reason document not match user data.
	RefusedReasonDocNotMatchWithUser RefusedReason = "DOCUMENT_DO_NOT_MATCH_USER_DATA"
	// RefusedReasonDocNotMatchWithAccount stands for a reason document do not match account data.
	RefusedReasonDocNotMatchWithAccount RefusedReason = "DOCUMENT_DO_NOT_MATCH_ACCOUNT_DATA"
	// RefusedReasonDocFalsified stands for a reason document falsifed.
	RefusedReasonDocFalsified RefusedReason = "DOCUMENT_FALSIFIED"
	// RefusedReasonSecificCase stands for a reason specific case.
	RefusedReasonSecificCase RefusedReason = "SPECIFIC_CASE"
	// RefusedReasonUnderAgePerson stands for a reason underage person.
	RefusedReasonUnderAgePerson RefusedReason = "UNDERAGE_PERSON"
)

// KYCDocumentCreate is a payload use for creating a new KYCDocument.
type KYCDocumentCreate struct {
}

// DocumentType is the type of KYC document.
type DocumentType string

const (
	// DocumentTypeIdentity is for a document of iendtity proof.
	DocumentTypeIdentity DocumentType = "IDENTITY_PROOF"
	// DocumentTypeRegistration is for a document of registration proff.
	DocumentTypeRegistration DocumentType = "REGISTRATION_PROOF"
	// DocumentTypeArticleOfAssociation is for a document of article of association.
	DocumentTypeArticleOfAssociation DocumentType = "ARTICLES_OF_ASSOCIATION"
	// DocumentTypeShareHolder is for a document of shareholder declaration.
	DocumentTypeShareHolder DocumentType = "SHAREHOLDER_DECLARATION"
	// DocumentTypeAddress is for a document of address proof.
	DocumentTypeAddress DocumentType = "ADDRESS_PROOF"
)

// KYCPage is used as a container to hold the files encoded in base64.
// if you are creating a new Page please use the function `NewKYCPage`
// that will help you to read and encode the given file.
type KYCPage struct {
	File string `json:"File"`
}

// NewKYCPage is reading from the given filepath the file and insert it into the KYCPage struct encoded in base64.
// The maximum size per page is about 7MB.
// The following formats are accepted : .pdf, .jpeg, .jpg, .gif and .png.
// The minimum size is 1Kb.
func NewKYCPage(filepath string) (*KYCPage, error) {
	var data string
	return &KYCPage{
		File: data,
	}, encodeBase64File(&data, filepath)
}

func encodeBase64File(data *string, filepath string) error {
	read, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	*data = base64.StdEncoding.EncodeToString(read)
	return nil
}
