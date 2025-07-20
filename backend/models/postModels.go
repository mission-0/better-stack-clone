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
	Website  []Website `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"website"`
}

type Website struct {
	ID      uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	URL     string     `json:"url" validate:"required,url"`
	Regions RegionList `json:"regions" validate:"required,validregion"`
	UserID  uuid.UUID  `gorm:"type:uuid;notNull;index" json:"userId" validate:"-"`
	User    User       `gorm:"foreignKey:UserId;references:Id" json:"user" validate:"-"`
	Logs    []Logs     `gorm:"foreignKey:LogsId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"logs"`
}

type Logs struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	LogsID uuid.UUID `gorm:"type:uuid;notNull;index" json:"logsId" validate:"-"`
	Logs   string    `json:"logs" validate:"required,logs"`
	time   time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
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
