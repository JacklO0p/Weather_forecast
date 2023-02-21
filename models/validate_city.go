package models

type CityValidator struct {
	Res []struct {
		ID  string `json:"id"`
		N   string `json:"n"`
		D   string `json:"d"`
		Iso string `json:"iso"`
	} `json:"res"`
	NearObjCount int `json:"nearObjCount"`
}