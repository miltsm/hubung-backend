package repository

import (
	"github.com/miltsm/hubung-service/internal/model"
	"gorm.io/gorm"
)

type Repository interface {
	GetProfile(id string) model.Profile
}

type repository struct {
	db gorm.DB
}

func New() Repository {
	db := establishDB()
	return &repository{ db: *db }
}

func (r *repository) GetProfile(id string) (model.Profile) {
	hubungan := []model.Hubung{ 
		{
			Name: "linkedin",
			Link: "https://www.linkedin.com/in/izzat-syamil-b86302186/",
		},
	}

	categories := []model.Category{
		{
			Name: "professional",
			Hubungan: hubungan,
		},
	}

	return model.Profile{
		Name: "Izzat Syamil",
		ImageUrl: "",
		Categories: categories,
	}
}