package service

import "github.com/project-inari/core-business-member/dto"

func (s *service) DoDBTest() (*[]dto.TestModel, error) {
	entity, err := s.databaseRepository.QueryTest()
	if err != nil {
		return nil, err
	}

	res := []dto.TestModel{}
	for _, e := range *entity {
		res = append(res, dto.TestModel(e))
	}

	return &res, nil
}
