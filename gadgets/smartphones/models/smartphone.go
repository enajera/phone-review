package models

//Smartphone model structure for smartphones
type Smartphone struct {
	Id            int64 //long en java
	Name          string
	Price         int
	CountryOrigin string
	OS            string
}

//CreateSmartphoneCMD
type CreateSmartphoneCMD struct {
	Name string `json:"name"`
	Price int `json:"price"`
	CountryOrigin string `json:"country_origin"`
	Os string `json:"os"`
}
