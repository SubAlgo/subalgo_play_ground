package response

type ResponseGenderMessageAndError struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}
