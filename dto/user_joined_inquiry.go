package dto

// UserJoinedInquiryRes represents the response for inquiring if a user has joined a business
type UserJoinedInquiryRes struct {
	Username   string          `json:"username"`
	Businesses []BusinessModel `json:"businesses"`
}
