package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();unique"`
	Email    string    `json:"email" gorm:"unique;notNull" validate:"required,email"`
	Password string    `json:"password" validate:"required,min=8,max=16"`
	Fullname string    `json:"fullname" validate:"required"`
	Website  []Website `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"website"`
}

type Website struct {
	ID      uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	URL     string     `json:"url" validate:"required,url"`
	Regions RegionList `json:"regions" validate:"required,validregion"`
	UserID  uuid.UUID  `gorm:"type:uuid;notNull;index" json:"userId" validate:"-"`
	User    User       `gorm:"foreignKey:UserID;references:ID" json:"user" validate:"-"`
	Logs    []Logs     `gorm:"foreignKey:WebsiteID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"logs"`
}

type Logs struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Latency   string    `json:"latency"`
	WebsiteID uuid.UUID `gorm:"type:uuid;notNull;index" json:"websiteID" validate:"-"`
	Website   Website   `gorm:"foreignKey:WebsiteID;references:ID" json:"website" validate:"-"`
	Logs      string    `json:"logs" validate:"required,logs"`
	Time      time.Time
}

type RegionList string

const (
	Asia         RegionList = "Asia"
	Europe       RegionList = "Europe"
	NorthAmerica RegionList = "North America"
	MiddleEast   RegionList = "Middle East"
)

var allowedRegions = map[RegionList]bool{
	Asia:         true,
	Europe:       true,
	NorthAmerica: true,
	MiddleEast:   true,
}

func IsValidRegion(r RegionList) bool {
	_, ok := allowedRegions[r]
	return ok
}
