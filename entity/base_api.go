package entity


type BaseApiResponse struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}