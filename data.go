package ipgeo


type Language struct {
	Code string
	Name string
	Native string
}

type Location struct {
	Geoname_id int
	Capital	string
	Languages []Language
	Country_flag	string
	Country_flag_emoji string
	Country_flag_emoji_unicode string
	Calling_code string
	Is_eu bool
}


type IPStack struct {
	Ip string 
	Ty string `json:"type"`
	Continent_code string
	Continent_name string
	Country_code string
	Country_name string
	Region_code string
	Region_name string
	City string
	Zip string
	Latitude float64
	Longitude float64
	Loc Location `json:"location"`
}