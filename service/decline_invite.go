package service

import (
	"context"

	"github.com/project-inari/core-business-member/dto"
)

const (
	statusInviteDeclined = "INVITE_DECLINED"
)

// DeclineInvite declines a business invitation
func (s *service) DeclineInvite(ctx context.Context, req dto.DeclineInviteReq) (*dto.DeclineInviteRes, error) {
	if err := s.databaseRepository.DeclineBusinessInvitation(ctx, req.InviteeUsername, req.BusinessName); err != nil {
		return nil, err
	}

	return &dto.DeclineInviteRes{
		InviteeUsername: req.InviteeUsername,
		BusinessName:    req.BusinessName,
		Status:          statusInviteDeclined,
	}, nil
}
