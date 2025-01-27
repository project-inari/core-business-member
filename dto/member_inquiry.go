package dto

// MemberInquiryRes represents the response for inquiring about a member
type MemberInquiryRes struct {
	BusinessName string                `json:"businessName"`
	Result       []MemberInquiryResult `json:"result"`
}

// MemberInquiryResult represents the result of inquiring about a member
type MemberInquiryResult struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	CreatedAt string `json:"createdAt"`
}
