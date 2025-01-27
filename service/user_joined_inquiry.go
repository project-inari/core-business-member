package service

import (
	"context"

	"github.com/project-inari/core-business-member/dto"
	"github.com/project-inari/core-business-member/pkg/utils"
)

func (s *service) UserJoinedInquiry(ctx context.Context, username string) (*dto.UserJoinedInquiryRes, error) {
	businesses, err := s.databaseRepository.GetUserJoinedBusinesses(ctx, username)
	if err != nil {
		return nil, err
	}

	return constructUserJoinedInquiryRes(username, businesses), nil
}

func constructUserJoinedInquiryRes(username string, businesses []*dto.UserBusinessesEntity) *dto.UserJoinedInquiryRes {
	result := make([]dto.BusinessModel, 0)
	for _, business := range businesses {
		result = append(result, dto.BusinessModel{
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
			UserRole:         business.UserRole,
		})
	}

	return &dto.UserJoinedInquiryRes{
		Username:   username,
		Businesses: result,
	}
}
