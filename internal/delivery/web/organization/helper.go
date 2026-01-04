package organization

type ErrorResponse struct {
	Error string `json:"error"`
	Meta  string `json:"meta,omitempty"`
}

func DefaultErr(meta string) *ErrorResponse {
	return &ErrorResponse{
		Error: "مشکلی پیش آمده",
		Meta:  meta,
	}
}
