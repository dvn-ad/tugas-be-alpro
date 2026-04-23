package utils

type SwaggerSuccessResponse struct {
	Status  string      `json:"status" example:"success"`
	Message string      `json:"message" example:"Operation successful"`
	Data    interface{} `json:"data"`
}

type SwaggerErrorResponse struct {
	Status  string `json:"status" example:"error"`
	Message string `json:"message" example:"Error message"`
}
