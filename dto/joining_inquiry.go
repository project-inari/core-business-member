package dto

// JoiningInquiryReq represents the request for inquiring about joining a business
type JoiningInquiryReq struct {
	Username     string `json:"username"`
	BusinessName string `json:"businessName"`
	Status       string `json:"status"`
}

// JoiningInquiryRes represents the response for inquiring about joining a business
type JoiningInquiryRes struct {
	Result []JoiningInquiryResult `json:"result"`
}

// JoiningInquiryResult represents the result of inquiring about joining a business
type JoiningInquiryResult struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	BusinessName string `json:"businessName"`
	Status       string `json:"status"`
	ActionedBy   string `json:"actionedBy"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}
