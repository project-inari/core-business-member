package service

import (
	"context"
	"errors"

	"github.com/project-inari/core-business-member/dto"
	"github.com/project-inari/core-business-member/pkg/utils"
)

const (
	roleMember           = "MEMBER"
	statusInviteAccepted = "INVITE_ACCEPTED"
)

// AcceptInvite accepts a business invitation
func (s *service) AcceptInvite(ctx context.Context, req dto.AcceptInviteReq) (*dto.AcceptInviteRes, error) {
	inviteStatus, err := s.databaseRepository.GetBusinessJoiningStatus(ctx, dto.BusinessJoiningQueryFilter{
		BusinessName: req.BusinessName,
		Username:     req.InviteeUsername,
		Status:       statusInvitePending,
	})
	if err != nil {
		return nil, err
	}

	if len(inviteStatus) == 0 {
		return nil, errors.New("no pending invitation found")
	}

	if err := s.databaseRepository.AcceptBusinessInvitation(ctx, req.InviteeUsername, req.BusinessName); err != nil {
		return nil, err
	}

	business, err := s.databaseRepository.GetBusiness(ctx, req.BusinessName)
	if err != nil {
		return nil, err
	}

	if err := s.cacheRepository.UpdateUserCacheNewBusinessJoined(ctx, req.InviteeUsername, *constructBusinessCacheModel(*business)); err != nil {
		return nil, err
	}

	return &dto.AcceptInviteRes{
		InviteeUsername: req.InviteeUsername,
		BusinessName:    req.BusinessName,
		Status:          statusInviteAccepted,
	}, nil
}

func constructBusinessCacheModel(business dto.BusinessEntity) *dto.BusinessModel {
	return &dto.BusinessModel{
		ID:               business.ID,
		Name:             business.Name,
		IndustryType:     business.IndustryType,
		BusinessType:     business.BusinessType,
		Description:      business.Description,
		PhoneNo:          business.PhoneNo,
		OperatingHours:   *utils.DecodeJSONfromString[dto.OperatingHours](business.OperatingHours),
		Address:          business.Address,
		BusinessImageURL: business.BusinessImageURL,
		CreatedAt:        business.CreatedAt,
		UpdatedAt:        business.UpdatedAt,
		UserRole:         roleMember,
	}
}
