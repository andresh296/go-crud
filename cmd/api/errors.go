package api

type ValidationError struct {
	Status int `json:"status"`
	Message string `json:"message"`
}

func (v ValidationError) Error() string {
	return v.Message
}