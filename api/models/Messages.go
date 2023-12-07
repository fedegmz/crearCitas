package models

type OnlyMessage struct {
	Success bool `json:"success"`
    Message string `json:"message"`
}

type Response struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Citas Citas `json:"citas"`
}


