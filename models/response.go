package models

type ResponseBooks struct {
	Status int     `json:"status"`
	Data   []Books `json:"data"`
}

type ResponseBook struct {
	Status int   `json:"status"`
	Data   Books `json:"data"`
}

type ResponseLogin struct {
	Status      int    `json:"status"`
	Message     string `json:"message"`
	AccessToken string `json:"access_token"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
