package models

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"primaryKey;unique;notNull"`
	Password string `json:"password"`
	FullName string `json:"fullname"`
	// Website  *[]Website `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// type Website struct {
// 	gorm.Model
// 	Url     string
// 	Regions RegionList
// }

// type RegionList string

// const (
// 	Asia          RegionList = "Asia"
// 	Europe        RegionList = " Europe"
// 	North_America RegionList = " North America"
// 	Middle_East   RegionList = "Middle East"
// )
