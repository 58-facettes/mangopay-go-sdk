package model

// SSO (Single Sign-On) allows you to access to the Dashboard.
// You can create multiple SSOs so that several users can access to the Dashboard
// with personal credentials and specific permissions.
type SSO struct {
	// The name of the user
	FirstName string `json:"FirstName"`
	// The last name of the user
	LastName string `json:"LastName"`
	// The user's email address - must be a valid email (not more than 12 consecutive numbers).
	// This email should not be already used for another SSO. (not more than 12 consecutive numbers)
	Email string `json:"Email"`
	// Whether the SSO is active or not
	Active bool `json:"Active"`
	// Status of the invitation sent to the user to confirm the SSO activation
	InvitationStatus InvitationStatus `json:"InvitationStatus"`
	// Date of the latest authentification
	LastLoginDate int64 `json:"LastLoginDate"`
	// PermissionGroup assigned to this SSO
	PermissionGroupID string `json:"PermissionGroupId"`
	// An ID for the client (i.e. url friendly, lowercase etc - sort of namespace identifier)
	ClientID string `json:"ClientId"`
}

// InvitationStatus is the status of the invitation.
type InvitationStatus string

const (
	// InvitationStatusAccepted stands for invistation status ACCEPTED.
	InvitationStatusAccepted InvitationStatus = "ACCEPTED"
	// InvitationStatusSent stands for invistation status SENT.
	InvitationStatusSent InvitationStatus = "SENT"
	// InvitationStatusExpired stands for invistation status EXPIRED.
	InvitationStatusExpired InvitationStatus = "EXPIRED"
)

// SSOCreate is use to create an new SSO.
type SSOCreate struct {
	// The name of the user REQUIRED
	FirstName string `json:"FirstName"`
	// The last name of the user REQUIRED
	LastName string `json:"LastName"`
	// The user's email address - must be a valid email (not more than 12 consecutive numbers). REQUIRED
	// This email should not be already used for another SSO. (not more than 12 consecutive numbers)
	Email string `json:"Email"`
	// PermissionGroup assigned to this SSO REQUIRED
	PermissionGroupID string `json:"PermissionGroupId"`
}

// SSOUpdate is use to update an existing SSO.
type SSOUpdate struct {
	// The name of the user OPTIONAL.
	FirstName string `json:"FirstName,omitempty"`
	// The last name of the user OPTIONAL.
	LastName string `json:"LastName,omitempty"`
	// PermissionGroup assigned to this SSO OPTIONAL.
	PermissionGroupID string `json:"PermissionGroupId,omitempty"`
	// Whether the SSO is active or not OPTIONAL.
	Active bool `json:"Active,omitempty"`
}
