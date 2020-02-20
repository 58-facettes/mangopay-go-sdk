package model

// PermissionGroup sets a specific level of permissions.
// Each SSO is assigned to a permission group that defines which API elements the user is allowed to view, edit and create.
// There are 3 default permission groups:
//
// 		1. The "Admin" permission group ( "Id": "ADMIN" ) that allows to:
//
// 		View and edit client settings
// 		View data such as users details
// 		Edit data such as create a payin
//
// 		2. The "Read & Write" permission group ( "Id": "WRITE" ) that allows to:
//
//		View data such as users details
// 		Edit data such as create a payin
//
// 		3. The "Read Only" permission group ( "Id": "READ" ) that allows to:
//
// View data such as users details
type PermissionGroup struct {
	// The name of the permission group
	Name string `json:"Name"`
	// The type of permission group
	Type GroupType `json:"Type,omitempty"`
	// The Scopes of the permissions
	Scopes Scopes `json:"Scopes,omitempty"`
}

// GroupType of the permission group.
type GroupType string

const (
	// GroupDefault is the default GroupType.
	GroupDefault GroupType = "DEFAULT"
	// GroupCustom is the custom GroupType.
	GroupCustom GroupType = "CUSTOM"
)

// Scopes show all the scopes of permissions possible on the API.
type Scopes struct {
	// API endpoints linked to client details
	ClientDetails Permissions `json:"ClientDetails,omitempty"`
	// API endpoints linked to client logo
	ClientLogo Permissions `json:"ClientLogo,omitempty"`
	// API endpoints linked to client wallets
	ClientWallets Permissions `json:"ClientWallets,omitempty"`
	// API endpoints linked to client bank accounts
	ClientBankAccounts Permissions `json:"ClientBankAccounts,omitempty"`
	// API endpoints linked to client payins
	ClientPayins Permissions `json:"ClientPayins,omitempty"`
	// API endpoints linked to client payouts
	ClientPayouts Permissions `json:"ClientPayouts,omitempty"`
	// API endpoints linked to client transactions
	ClientTransactions Permissions `json:"ClientTransactions,omitempty"`
	// API endpoints linked to SSOs
	SSOs Permissions `json:"SSOs,omitempty"`
	// API endpoints linked to permission groups
	PermissionGroups Permissions `json:"PermissionGroups,omitempty"`
	// API endpoints linked to users
	Users Permissions `json:"Users,omitempty"`
	// API endpoints linked to wallets
	Wallets Permissions `json:"Wallets,omitempty"`
	// API endpoints linked to banking aliases
	BankingAliases Permissions `json:"BankingAliases,omitempty"`
	// API endpoints linked to cards
	Cards Permissions `json:"Cards,omitempty"`
	// API endpoints linked to bank accounts
	BankAccounts Permissions `json:"BankAccounts,omitempty"`
	// API endpoints linked to preauthorizations
	PreAuthorizations Permissions `json:"PreAuthorizations,omitempty"`
	// API endpoints linked to payins
	Payins Permissions `json:"Payins,omitempty"`
	// API endpoints linked to transfers
	Transfers Permissions `json:"Transfers,omitempty"`
	// API endpoints linked to payouts
	Payouts Permissions `json:"Payouts,omitempty"`
	// API endpoints linked to refunds
	Refunds Permissions `json:"Refunds,omitempty"`
	// API endpoints linked to transactions
	Transactions Permissions `json:"Transactions,omitempty"`
	// API endpoints linked to KYC documents
	KYCDocuments Permissions `json:"KYCDocuments,omitempty"`
	// API endpoints linked to disputes
	Disputes Permissions `json:"Disputes,omitempty"`
	// API endpoints linked to repudiations
	Repudiations Permissions `json:"Repudiations,omitempty"`
	// API endpoints linked to mandates
	Mandates Permissions `json:"Mandates,omitempty"`
	// API endpoints linked to reporting
	Reporting Permissions `json:"Reporting,omitempty"`
	// API endpoints linked to responses
	Responses Permissions `json:"Responses,omitempty"`
	// API endpoints linked to events
	Events Permissions `json:"Events,omitempty"`
	// API endpoints linked to hooks
	Hooks Permissions `json:"Hooks,omitempty"`
	// API endpoints linked to UBO declaration
	UboDeclarations Permissions `json:"UboDeclarations,omitempty"`
}

// Permissions describe all permission.
type Permissions struct {
	// allows GET requests on the related endpoints.
	Read bool `json:"Read,omitempty"`
	// allows PUT requests on the related endpoints.
	Edit bool `json:"Edit,omitempty"`
	// allows POST requests on the related endpoints.
	Create bool `json:"Create,omitempty"`
}

// PermissionGroupCreate is used to create some PermissionGroup.
type PermissionGroupCreate struct {
	// The name of the permission group
	Name string `json:"Name"`
	// The type of permission group
	Type GroupType `json:"Type,omitempty"`
	// The Scopes of the permissions
	Scopes Scopes `json:"Scopes,omitempty"`
}

// PermissionGroupUpdate is used to update som PermissionGroup.
type PermissionGroupUpdate struct {
	// The name of the permission group
	Name string `json:"Name,omitempty"`
	// The Scopes of the permissions
	Scopes Scopes `json:"Scopes,omitempty"`
}
