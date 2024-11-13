package handlers

type ResponseModel[D any] struct {
	Success            bool   `json:"success"`
	Message            string `json:"message"`
	Data               *D     `json:"data"`
	SampleRequestModel any    `json:"sample_request_model,omitempty"`
}

/* README: Example universal response model

resp := ResponseModel[any]{
	Success: true,
	Message: message,
	Data:    "my",
}
*/
