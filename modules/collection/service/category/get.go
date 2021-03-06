package category

import (
	"kingford-backend/global"
	"kingford-backend/modules/collection/repository"
)

type GetService struct {
}

func (s *GetService) Get(id string) *global.Response {
	repo := repository.CollectionCategoryRepository{DB: global.DB}
	item, err := repo.Get(id)

	if err != nil {
		return &global.Response{
			Status: 400,
			Data:   item,
			Msg:    err.Error(),
		}
	}
	return &global.Response{
		Status: 200,
		Data:   item,
		Msg:    "success",
	}
}
