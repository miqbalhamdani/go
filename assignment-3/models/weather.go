package models

type Weather struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type Data struct {
	Status Weather `json:"status"`
}
