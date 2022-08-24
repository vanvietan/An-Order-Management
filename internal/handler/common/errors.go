package common

// CommonErrorResponse
type CommonErrorResponse struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

var (
	// InternalCommonErrorResponse
	InternalCommonErrorResponse = CommonErrorResponse{
		Code:        "internal_server_error",
		Description: "Something went wrong please try again later!",
	}
)
