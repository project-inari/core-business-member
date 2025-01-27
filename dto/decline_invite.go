package dto

// DeclineInviteReq represents the request for declining an invite to join a business
type DeclineInviteReq struct {
	InviteeUsername string `json:"inviteeUsername" validate:"required"`
	BusinessName    string `json:"businessName" validate:"required"`
}

// DeclineInviteRes represents the response for declining an invite to join a business
type DeclineInviteRes struct {
	InviteeUsername string `json:"inviteeUsername"`
	BusinessName    string `json:"businessName"`
	Status          string `json:"status"`
}
