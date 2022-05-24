package rest

type GenTreeError struct {
	Result  string `json:"result" xml:"result" yaml:"result"`
	Message string `json:"message" xml:"message" yaml:"message"`
}

var (
	ErrInternal            = GenTreeError{Result: "error", Message: "Internal server error. Try again later or contact the support team."}
	ErrFailedToDeserialize = GenTreeError{Result: "error", Message: "Failed to deserialize. Check request payload."}
	ErrResourceNotFound    = GenTreeError{Result: "error", Message: "Resource not found in the database. Check query values."}
)

// ErrInternal exists to avoid forwarding DB errors directly, which could lead to security breach.
