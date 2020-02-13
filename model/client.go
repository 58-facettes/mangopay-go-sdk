package model

import "encoding/base64"

type Client struct {
	// Name it the pretty name for the client.
	Name string `json:"name"`
	// RegisteredName it the registered name of your company.
	RegisteredName string `json:"RegisteredName"`
	// ClientId is an ID for the client (i.e. url friendly, lowercase etc - sort of namespace identifier).
	ClientId string `json:"ClientId"`
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

type Address struct {
	// AddressLine1 is the first line of the address.
	AddressLine1 string `json:""`
	// AddressLine2 is the second line of the address.
	AddressLine2 string `json:"AddressLine2,omitempty"`
	// City is the city of the address.
	City string `json:"City"`
	// Region is the region of the address - this is optional except if the Country is US, CA or MX.
	Region string `json:"Region"`
	// PostalCode is the postal code of the address - can be alphanumeric, dashes or spaces.
	PostalCode string `json:"PostalCode"`
	// Country is the Country of the Address.
	Country string `json:"Countrys"`
}

type Business string

const (
	BusinessMarketplace  Business = "MARKETPLACE"
	BusinessCrowdfunding Business = "CROWDFUNDING"
	BusinessFranchise    Business = "FRANCHISE"
	BusinessOther        Business = "OTHER"
)

type PlatformCategorization struct {
	// BusinessType is the business type of your Platform.
	BusinessType Business // OPTIONAL
	// Sector is the sector of your platform activity.
	Sector Sector // OPTIONAL
}

type Sector string

const (
	SectorRentals                        Sector = "RENTALS"
	SectorStoreFashionAccessoriesObjects Sector = "STORES_FASHION_ACCESSORIES_OBJECTS"
	SectorBeautyCosmeticsHealth          Sector = "BEAUTY_COSMETICS_HEALTH"
	SectorFoodWineRestaurants            Sector = "FOOD_WINE_RESTAURANTS"
	SectorHospitalityTravelCording       Sector = "HOSPITALITY_TRAVEL_CORIDING"
	SectorArtMusicEntretainement         Sector = "ART_MUSIC_ENTERTAINMENT"
	SectorFurnitureGarden                Sector = "FURNITURE_GARDEN"
	SectorServicesJobbingEducation       Sector = "SERVICES_JOBBING_EDUCATION"
	SectorSportRecreationActivities      Sector = "SPORT_RECREATION_ACTIVITIES"
	SectorTicketing                      Sector = "TICKETING"
	SectorLoan                           Sector = "LOAN"
	SectorEquity                         Sector = "EQUITY"
	SectorPropertyEquity                 Sector = "PROPERTY_EQUITY"
	SectorRewardsCharity                 Sector = "REWARDS_CHARITY"
	SectorPoolGroupPayment               Sector = "POOL_GROUP_PAYMENT"
	SectorFranchise                      Sector = "FRANCHISE"
	SectorOther                          Sector = "OTHER"
)
