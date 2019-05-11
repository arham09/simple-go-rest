package models

type ResponseBooks struct {
	Status int     `json:"status"`
	Data   []Books `json:"data"`
}

type ResponseBook struct {
	Status int   `json:"status"`
	Data   Books `json:"data"`
}
