package dto

// UserVerifiedAccountCache represents the cache for a verified user account
type UserVerifiedAccountCache struct {
	Username   string          `json:"username"`
	UID        string          `json:"uid"`
	FirstName  string          `json:"firstName"`
	LastName   string          `json:"lastName"`
	PhoneNo    string          `json:"phoneNo"`
	Email      string          `json:"email"`
	Businesses []BusinessModel `json:"businesses"`
}
