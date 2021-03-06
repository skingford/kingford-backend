package collection

import (
	"kingford-backend/global"
	"kingford-backend/modules/collection/repository"
)

type DeleteService struct {
}

func (s *DeleteService) Delete(id string) *global.Response {
	repo := repository.CollectionRepository{DB: global.DB}
	err := repo.Delete(id)

	if err != nil {
		return &global.Response{
			Status: 400,

			Msg: err.Error(),
		}
	}
	return &global.Response{
		Status: 200,
		Data:   true,
		Msg:    "success",
	}
}
