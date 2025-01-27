package dto

// UserVerifiedAccountCache represents the cache for a verified user account
type UserVerifiedAccountCache struct {
	Username   string               `json:"username"`
	UID        string               `json:"uid"`
	FirstName  string               `json:"firstName"`
	LastName   string               `json:"lastName"`
	PhoneNo    string               `json:"phoneNo"`
	Email      string               `json:"email"`
	Businesses []BusinessCacheModel `json:"businesses"`
}

// BusinessCacheModel represents the business model
type BusinessCacheModel struct {
	ID               int            `json:"id"`
	Name             string         `json:"name"`
	IndustryType     string         `json:"industryType"`
	BusinessType     string         `json:"businessType"`
	Description      string         `json:"description"`
	PhoneNo          string         `json:"phoneNo"`
	OperatingHours   OperatingHours `json:"operatingHours"`
	Address          string         `json:"address"`
	BusinessImageURL string         `json:"businessImageUrl"`
	CreatedAt        string         `json:"createdAt"`
	UpdatedAt        string         `json:"updatedAt"`
	UserRole         string         `json:"userRole"`
}

// OperatingHours represents the operating hours of a business for the week
type OperatingHours struct {
	Monday    OpenTime `json:"monday"`
	Tuesday   OpenTime `json:"tuesday"`
	Wednesday OpenTime `json:"wednesday"`
	Thursday  OpenTime `json:"thursday"`
	Friday    OpenTime `json:"friday"`
	Saturday  OpenTime `json:"saturday"`
	Sunday    OpenTime `json:"sunday"`
}

// OpenTime represents the open and close time of a business
type OpenTime struct {
	Open      bool   `json:"open"`
	OpenTime  string `json:"openTime,omitempty"`
	CloseTime string `json:"closeTime,omitempty"`
}
