package dto

type SuccessResponse struct {
	Success bool  `json:"success"`
	Error   error `json:"error"`
}

type IndexResponse struct {
	Results interface{} `json:"results"`
	Count   int         `json:"count"`
}
