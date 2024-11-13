package handlers

type ResponseModel[D any] struct {
	Success bool
	Message string
	Data    *D
}

/* README: Example universal response model

resp := ResponseModel[any]{
	Success: true,
	Message: message,
	Data:    "my",
}
*/
