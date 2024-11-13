package handlers

type ResponseModel[D any] struct {
	IsSuccess          bool   `json:"is_success"`
	IsError            bool   `json:"is_error,omitempty"`
	Message            string `json:"message"`
	Data               *D     `json:"data,omitempty"`
	SampleRequestModel any    `json:"sample_request_model,omitempty"`
}
