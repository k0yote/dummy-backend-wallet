package issuer

type ErrorResult struct {
	Message   string `json:"message"`
	RequestID string `json:"requestID,omitempty"`
	Code      int    `json:"code,omitempty"`
	Error     string `json:"error,omitempty"`
}
