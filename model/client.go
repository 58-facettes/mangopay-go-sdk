package model

import "encoding/base64"

// Client represents a client.
type Client struct {
	// Name it the pretty name for the client.
	Name string `json:"name"`
	// RegisteredName it the registered name of your company.
	RegisteredName string `json:"RegisteredName"`
	// ClientID is an ID for the client (i.e. url friendly, lowercase etc - sort of namespace identifier).
	ClientID string `json:"ClientId"`
	// PrimaryThemeColour is the primary branding colour to use for your merchant.
	PrimaryThemeColour string `json:"PrimaryThemeColour"`
	// PrimaryButtonColour is the primary branding colour to use for buttons for your merchant.
	PrimaryButtonColour string `json:"PrimaryButtonColour"`
	// Logo is the URL of the logo of your client.
	Logo string `json:"Logo"`
	// TechEmails is a list of email addresses to use when contacting you for technical issues/communications.
	TechEmails []string `json:"TechEmails"`
	// AdminEmails is a list of email addresses to use when contacting you for admin/commercial issues/communications.
	AdminEmails []string `json:"AdminEmails"`
	// FraudEmails is a list of email addresses to use when contacting you for fraud/compliance issues/communications.
	FraudEmails []string `json:"FraudEmails"`
	// BillingEmails is a list of email addresses to use when contacting you for billing issues/communications.
	BillingEmails []string `json:"BillingEmails"`
	// PlatformCategorization is the Categorization of your platform, in terms of Business Type and Sector.
	PlatformCategorization PlatformCategorization `json:"PlatformCategorization"`
	// PlatformDescription is a description of what your platform does.
	PlatformDescription string `json:"PlatformDescription"`
	// PlatformURL is the URL for your website.
	PlatformURL string `json:"PlatformURL"`
	// HeadquartersAddress is the address of the company’s headquarters. This field is mandatory to accept payout (More info here).
	HeadquartersAddress Address `json:"HeadquartersAddress"`
	// TaxNumber is the tax (or VAT) number for your company.
	TaxNumber string `json:"TaxNumber"`
	// CompanyReference is your unique MANGOPAY reference which you should use when contacting us.
	CompanyReference string `json:"CompanyReference"`
}

// ClientUpdate is the payload used for updating a client.
type ClientUpdate struct {
	// AdminEmails is a list of email addresses to use when contacting you for admin/commercial issues/communications.
	AdminEmails []string `json:"AdminEmails,omitempty"`
	// TechEmails is a list of email addresses to use when contacting you for technical issues/communications.
	TechEmails []string `json:"TechEmails,omitempty"`
	// BillingEmails is a list of email addresses to use when contacting you for billing issues/communications.
	BillingEmails []string `json:"BillingEmails,omitempty"`
	// FraudEmails is a list of email addresses to use when contacting you for fraud/compliance issues/communications.
	FraudEmails []string `json:"FraudEmails,omitempty"`
	// HeadquartersAddress is the address of the company’s headquarters. This field is mandatory to accept payout (More info here).
	HeadquartersAddress Address `json:"HeadquartersAddress,omitempty"`
	// TaxNumber is the tax (or VAT) number for your company.
	TaxNumber string `json:"TaxNumber,omitempty"`
	// PlatformCategorization is the Categorization of your platform, in terms of Business Type and Sector.
	PlatformCategorization PlatformCategorization `json:"PlatformCategorization,omitempty"`
	// PlatformDescription is a description of what your platform does.
	PlatformDescription string `json:"PlatformDescription,omitempty"`
	// PlatformURL is the URL for your website.
	PlatformURL string `json:"PlatformURL,omitempty"`
	// PrimaryThemeColour is the primary branding colour to use for your merchant.
	PrimaryThemeColour string `json:"PrimaryThemeColour,omitempty"`
	// PrimaryButtonColour is the primary branding colour to use for buttons for your merchant.
	PrimaryButtonColour string `json:"PrimaryButtonColour,omitempty"`
}

// ClientLogo is the payload used for updating the Logo image of the client.
// the value in file must be encoded in a base64 format.
type ClientLogo struct {
	File string `json:"File"`
}

// NewClientLogo is giving a new ClientLogo from a data bytes file,
// it is encoding into base64 format the given bytes.
func NewClientLogo(data []byte) *ClientLogo {
	return &ClientLogo{
		File: base64.StdEncoding.EncodeToString(data),
	}
}

// Address represent an address.
type Address struct {
	// Line1 is the first line of the address.
	Line1 string `json:"AddressLine1"`
	// Line2 is the second line of the address.
	Line2 string `json:"AddressLine2,omitempty"`
	// City is the city of the address.
	City string `json:"City"`
	// Region is the region of the address - this is optional except if the Country is US, CA or MX.
	Region string `json:"Region"`
	// PostalCode is the postal code of the address - can be alphanumeric, dashes or spaces.
	PostalCode string `json:"PostalCode"`
	// Country is the Country of the Address.
	Country string `json:"Countrys"`
}

// Business for the business type.
type Business string

const (
	// BusinessMarketplace is for a bunisess in marketplace.
	BusinessMarketplace Business = "MARKETPLACE"
	// BusinessCrowdfunding is for a bunisess in crowdfunding.
	BusinessCrowdfunding Business = "CROWDFUNDING"
	// BusinessFranchise is for a bunisess in franchise.
	BusinessFranchise Business = "FRANCHISE"
	// BusinessOther is for a bunisess in other.
	BusinessOther Business = "OTHER"
)

// PlatformCategorization is used to categorize a platform.
type PlatformCategorization struct {
	// BusinessType is the business type of your Platform.
	BusinessType Business // OPTIONAL
	// Sector is the sector of your platform activity.
	Sector Sector // OPTIONAL
}

// Sector is the sector of the business.
type Sector string

const (
	// SectorRentals is for a rentals sector.
	SectorRentals Sector = "RENTALS"
	// SectorStoreFashionAccessoriesObjects is for a stores fashion accessorties objects sector.
	SectorStoreFashionAccessoriesObjects Sector = "STORES_FASHION_ACCESSORIES_OBJECTS"
	// SectorBeautyCosmeticsHealth is for a beauty cosmetics health sector.
	SectorBeautyCosmeticsHealth Sector = "BEAUTY_COSMETICS_HEALTH"
	// SectorFoodWineRestaurants is for a food wine restaurants sector.
	SectorFoodWineRestaurants Sector = "FOOD_WINE_RESTAURANTS"
	// SectorHospitalityTravelCording is for a hospitality travel coriding sector.
	SectorHospitalityTravelCording Sector = "HOSPITALITY_TRAVEL_CORIDING"
	// SectorArtMusicEntretainement is for an art music entrertainment sector.
	SectorArtMusicEntretainement Sector = "ART_MUSIC_ENTERTAINMENT"
	// SectorFurnitureGarden is for a furniture garden sector.
	SectorFurnitureGarden Sector = "FURNITURE_GARDEN"
	// SectorServicesJobbingEducation is for a service jobbing education sector.
	SectorServicesJobbingEducation Sector = "SERVICES_JOBBING_EDUCATION"
	// SectorSportRecreationActivities is for a sport recreation activities sector.
	SectorSportRecreationActivities Sector = "SPORT_RECREATION_ACTIVITIES"
	// SectorTicketing is for a ticketing sector.
	SectorTicketing Sector = "TICKETING"
	// SectorLoan is for a loan sector.
	SectorLoan Sector = "LOAN"
	// SectorEquity is for a eauity sector.
	SectorEquity Sector = "EQUITY"
	// SectorPropertyEquity is for a property equity sector.
	SectorPropertyEquity Sector = "PROPERTY_EQUITY"
	// SectorRewardsCharity is for a rewards charity sector.
	SectorRewardsCharity Sector = "REWARDS_CHARITY"
	// SectorPoolGroupPayment is for a pool group payment sector.
	SectorPoolGroupPayment Sector = "POOL_GROUP_PAYMENT"
	// SectorFranchise is for a franchise sector.
	SectorFranchise Sector = "FRANCHISE"
	// SectorOther is for a other sector.
	SectorOther Sector = "OTHER"
)
