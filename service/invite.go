package service

import (
	"context"

	"github.com/project-inari/core-business-member/dto"
)

const (
	statusInvitePending = "INVITE_PENDING"
)

// Invite invites a user to join a business
func (s *service) Invite(ctx context.Context, req dto.InviteReq) (*dto.InviteRes, error) {
	_ = ctx

	err := s.databaseRepository.InviteToJoinBusiness(req.InviterUsername, req.InviteeUsername, req.BusinessName, statusInvitePending)
	if err != nil {
		return nil, err
	}

	return &dto.InviteRes{
		InviterUsername: req.InviterUsername,
		InviteeUsername: req.InviteeUsername,
		BusinessName:    req.BusinessName,
		Status:          statusInvitePending,
	}, nil
}
