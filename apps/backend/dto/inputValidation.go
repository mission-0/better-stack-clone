package dto

import (
	"github.com/google/uuid"
	"github.com/mission-0/better-stack-backend/models"
)

type WebsiteInput struct {
	Url     string            `json:"url" validate:"required,url"`
	Regions models.RegionList `json:"regions" validate:"required,validregion"`
	UserId  uuid.UUID         `json:"userId" validate:"required"`
}

type UserInput struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Fullname string `json:"fullname" validate:"omitempty"`
}
