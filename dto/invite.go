package dto

type InviteReq struct {
	InviterUsername string `json:"inviterUsername" validate:"required"`
	InviteeUsername string `json:"inviteeUsername" validate:"required"`
	BusinessName    string `json:"businessName" validate:"required"`
}

type InviteRes struct {
	InviterUsername string `json:"inviterUsername"`
	InviteeUsername string `json:"inviteeUsername"`
	BusinessName    string `json:"businessName"`
	Status          string `json:"status"`
}
