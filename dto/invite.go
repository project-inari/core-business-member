package dto

// InviteReq represents the request for inviting a user to join a business
type InviteReq struct {
	InviterUsername string `json:"inviterUsername" validate:"required"`
	InviteeUsername string `json:"inviteeUsername" validate:"required"`
	BusinessName    string `json:"businessName" validate:"required"`
}

// InviteRes represents the response for inviting a user to join a business
type InviteRes struct {
	InviterUsername string `json:"inviterUsername"`
	InviteeUsername string `json:"inviteeUsername"`
	BusinessName    string `json:"businessName"`
	Status          string `json:"status"`
}
