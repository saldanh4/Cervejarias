package controllers

type Brewery struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Brewery_type   string `json:"brewery_type"`
	Address_1      string `json:"address_1"`
	Address_2      string `json:"address_2"`
	Address_3      string `json:"address_3"`
	City           string `json:"city"`
	State_province string `json:"state_province"`
	Postal_code    string `json:"postal_code"`
	Country        string `json:"country"`
	Longitude      string `json:"longitude"`
	Latitude       string `json:"latitude"`
	Phone          string `json:"phone"`
	Website_url    string `json:"website_url"`
	State          string `json:"state"`
	Street         string `json:"street"`
}
