package models

import (
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();unique"`
	Name     string    `json:"name"`
	Email    string    `json:"email" gorm:"unique;notNull"`
	Password string    `json:"password"`
	FullName string    `json:"fullname"`
	Website  []Website `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"website"`
}

type Website struct {
	Id      uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Url     string     `json:"url"`
	Regions RegionList `json:"regions" gorm:"type:RegionList"`
	UserId  uuid.UUID  `gorm:"type:uuid;notNull;index" json:"userId"`
	User    User       `gorm:"foreignKey:UserId;references:Id" json:"user"`
}

type RegionList string

const (
	Asia          RegionList = "Asia"
	Europe        RegionList = "Europe"
	North_America RegionList = "North America"
	Middle_East   RegionList = "Middle East"
)
