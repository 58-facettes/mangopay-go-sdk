package model

// DisputeDocument is the KYC document.
// You need to create document in order to upload pages on this document.
// Note that dispute documents function in just the same way KYC documents.
// ?? Why we don't use the same status of the KYC documents ??
type DisputeDocument struct {
	// The Id of a Dispute.
	DisputeID string `json:"DisputeId"`
	// The status of this KYC/Dispute document.
	Status DocumentStatus `json:"Status"`
	// Gives the type of the dispute document.
	Type DisputeDocumentType `json:"Type"`
	// The message accompanying a refusal.
	RefusedReasonMessage string `json:"RefusedReasonMessage"`
	// The type of reason for refusal.
	RefusedReasonType DisputeDocRefusedReasonType `json:"RefusedReasonType"`
	// The date when the document was processed by MANGOPAY.
	ProcessedDate int64 `json:"ProcessedDate"`
}

// DisputeDocumentType is the type of dispute document.
type DisputeDocumentType string

const (
	// DisputeDocumentDeliveryProof is for dispute document type of DELIVERY_PROOF.
	DisputeDocumentDeliveryProof DisputeDocumentType = "DELIVERY_PROOF"
	// DisputeDocumentInvoice is for dispute document type of INVOICE.
	DisputeDocumentInvoice DisputeDocumentType = "INVOICE"
	// DisputeDocumentRefundProof is for dispute document type of REFUND_PROOF.
	DisputeDocumentRefundProof DisputeDocumentType = "REFUND_PROOF"
	// DisputeDocumentUserCorrespondance is for dispute document type of USER_CORRESPONDANCE.
	DisputeDocumentUserCorrespondance DisputeDocumentType = "USER_CORRESPONDANCE"
	// DisputeDocumentUserAcceptanceProof is for dispute document type of USER_ACCEPTANCE_PROOF.
	DisputeDocumentUserAcceptanceProof DisputeDocumentType = "USER_ACCEPTANCE_PROOF"
	// DisputeDocumentProductReplacementProof is for dispute document type of PRODUCT_REPLACEMENT_PROOF.
	DisputeDocumentProductReplacementProof DisputeDocumentType = "PRODUCT_REPLACEMENT_PROOF"
	// DisputeDocumentOther is for dispute document type of OTHER.
	DisputeDocumentOther DisputeDocumentType = "OTHER"
)

// DisputeDocRefusedReasonType is the type of refusal reason.
type DisputeDocRefusedReasonType string

const (
	// DisputeDocRefusedReasonDocumentUnreadable is for dispute doc reason DOCUMENT_UNREADABLE.
	DisputeDocRefusedReasonDocumentUnreadable DisputeDocRefusedReasonType = "DOCUMENT_UNREADABLE"
	// DisputeDocRefusedReasonDocumentNotAccepted is for dispute doc reason DOCUMENT_NOT_ACCEPTED.
	DisputeDocRefusedReasonDocumentNotAccepted DisputeDocRefusedReasonType = "DOCUMENT_NOT_ACCEPTED"
	// DisputeDocRefusedReasonDocumentHasExpired is for dispute doc reason DOCUMENT_HAS_EXPIRED.
	DisputeDocRefusedReasonDocumentHasExpired DisputeDocRefusedReasonType = "DOCUMENT_HAS_EXPIRED"
	// DisputeDocRefusedReasonDocumentIncomplete is for dispute doc reason DOCUMENT_INCOMPLETE.
	DisputeDocRefusedReasonDocumentIncomplete DisputeDocRefusedReasonType = "DOCUMENT_INCOMPLETE"
	// DisputeDocRefusedReasonDocumentMissing is for dispute doc reason DOCUMENT_MISSING.
	DisputeDocRefusedReasonDocumentMissing DisputeDocRefusedReasonType = "DOCUMENT_MISSING"
	// DisputeDocRefusedReasonSpecificCase is for dispute doc reason SPECIFIC_CASE.
	DisputeDocRefusedReasonSpecificCase DisputeDocRefusedReasonType = "SPECIFIC_CASE"
	// DisputeDocRefusedReasonDocumentFalsified is for dispute doc reason DOCUMENT_FALSIFIED.
	DisputeDocRefusedReasonDocumentFalsified DisputeDocRefusedReasonType = "DOCUMENT_FALSIFIED"
	// DisputeDocRefusedReasonOther is for dispute doc reason OTHER.
	DisputeDocRefusedReasonOther DisputeDocRefusedReasonType = "OTHER"
)

// DisputeDocumentCreate is the payload used to send the data of the document in base64.
type DisputeDocumentCreate struct {
	File string `json:"File"`
}

// NewDisputeDocument is reading from a file an encode the given []bytes into a base64 fromat.
func NewDisputeDocument(filepath string) (*KYCPage, error) {
	var data string
	return &KYCPage{
		File: data,
	}, encodeBase64File(&data, filepath)
}
