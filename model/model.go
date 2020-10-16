package model

type Area struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type ConfigElasticSearch struct {
	Host  string `json:"host"`
	Index string `json:"index"`
}
