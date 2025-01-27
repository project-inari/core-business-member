package dto

// BusinessEntity represents the business entity in the database for table tbl_businesses
type BusinessEntity struct {
	ID               int    `sql:"id"`
	Name             string `sql:"name"`
	IndustryType     string `sql:"industry_type"`
	BusinessType     string `sql:"business_type"`
	Description      string `sql:"description"`
	PhoneNo          string `sql:"phone_no"`
	OperatingHours   string `sql:"operating_hours"`
	Address          string `sql:"address"`
	BusinessImageURL string `sql:"business_image_url"`
	CreatedAt        string `sql:"created_at"`
	UpdatedAt        string `sql:"updated_at"`
}

// BusinessModel represents the business model
type BusinessModel struct {
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

// BusinessMemberEntity represents the business member entity in the database for table tbl_business_members
type BusinessMemberEntity struct {
	ID           int    `sql:"id"`
	BusinessName string `sql:"business_name"`
	Username     string `sql:"username"`
	Role         string `sql:"role"`
	CreatedAt    string `sql:"created_at"`
	UpdatedAt    string `sql:"updated_at"`
}

// BusinessJoiningEntity represents the business joining entity in the database for table tbl_business_joining
type BusinessJoiningEntity struct {
	ID           int    `sql:"id"`
	BusinessName string `sql:"business_name"`
	Username     string `sql:"username"`
	Status       string `sql:"status"`
	ActionedBy   string `sql:"actioned_by"`
	CreatedAt    string `sql:"created_at"`
	UpdatedAt    string `sql:"updated_at"`
}

// BusinessJoiningQueryFilter represents the query filter for business joining
type BusinessJoiningQueryFilter struct {
	BusinessName string `sql:"business_name"`
	Username     string `sql:"username"`
	Status       string `sql:"status"`
}

// UserBusinessesEntity represents the user businesses entity in the database for table tbl_businesses with the user role
type UserBusinessesEntity struct {
	ID               int    `sql:"id"`
	Name             string `sql:"name"`
	IndustryType     string `sql:"industry_type"`
	BusinessType     string `sql:"business_type"`
	Description      string `sql:"description"`
	PhoneNo          string `sql:"phone_no"`
	OperatingHours   string `sql:"operating_hours"`
	Address          string `sql:"address"`
	BusinessImageURL string `sql:"business_image_url"`
	CreatedAt        string `sql:"created_at"`
	UpdatedAt        string `sql:"updated_at"`
	UserRole         string `sql:"user_role"`
}
