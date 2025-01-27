package dto

// AcceptInviteReq represents the request for accepting an invite to join a business
type AcceptInviteReq struct {
	InviteeUsername string `json:"inviteeUsername" validate:"required"`
	BusinessName    string `json:"businessName" validate:"required"`
}

// AcceptInviteRes represents the response for accepting an invite to join a business
type AcceptInviteRes struct {
	InviteeUsername string `json:"inviteeUsername"`
	BusinessName    string `json:"businessName"`
	Status          string `json:"status"`
}
