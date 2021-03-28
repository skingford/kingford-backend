package collection

import (
	"kingford-backend/dto"
	"kingford-backend/global"
	"kingford-backend/repository"
)

type GetService struct {
}

func (s *GetService) Get(id string) *dto.Response {
	repo := repository.CollectionRepository{DB: global.DB}
	item, err := repo.Get(id)

	if err != nil {
		return &dto.Response{
			Status: 400,
			Data:   item,
			Msg:    err.Error(),
		}
	}
	return &dto.Response{
		Status: 200,
		Data:   item,
		Msg:    "success",
	}
}