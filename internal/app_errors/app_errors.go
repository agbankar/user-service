package app_errors

type AppError struct {
	ErrorCode        string `json:"errorCode"`
	ErrorDescription string `json:"errorDescription"`
	MoreInfo         string `json:"moreInfo,omitempty"`
}
