package models

import (
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();unique"`
	Email    string    `json:"email" gorm:"unique;notNull" validate:"required,email"`
	Password string    `json:"password" validate:"required,min=8,max=16"`
	Fullname string    `json:"fullname" validate:"required"`
	Website  []Website `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"website"`
}

type Website struct {
	Id      uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Url     string     `json:"url" validate:"required,url"`
	Regions RegionList `json:"regions" validate:"required,validregion"`
	UserId  uuid.UUID  `gorm:"type:uuid;notNull;index" json:"userId" validate:"-"`
	User    User       `gorm:"foreignKey:UserId;references:Id" json:"user" validate:"-"`
}

type RegionList string

const (
	Asia          RegionList = "Asia"
	Europe        RegionList = "Europe"
	North_America RegionList = "North America"
	Middle_East   RegionList = "Middle East"
)

var allowedRegions = map[RegionList]bool{
	Asia:          true,
	Europe:        true,
	North_America: true,
	Middle_East:   true,
}

func IsValidRegion(r RegionList) bool {
	_, ok := allowedRegions[r]
	return ok
}
