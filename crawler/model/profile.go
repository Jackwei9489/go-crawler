package model

type Profile struct {
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	Age        int    `json:"age"`
	Height     int    `json:"height"`
	Weight     int    `json:"weight"`
	Income     string `json:"income"`
	Marriage   string `json:"marriage"`
	Education  string `json:"education"`
	Occupation string `json:"occupation"`
	Hokou      string `json:"hokou"`
	Xinzuo     string `json:"xinzuo"`
	House      string `json:"house"`
	Car        string `json:"car"`
}
