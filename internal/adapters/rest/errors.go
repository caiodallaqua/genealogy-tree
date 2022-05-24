package rest

type GenTreeError struct {
	Error string `json:"error" xml:"error" yaml:"error" example:"Failed to deserialize. Check request."`
}

var (
	ErrInternal            = GenTreeError{Error: "Internal server error. Try again later or contact the support team."}
	ErrFailedToDeserialize = GenTreeError{Error: "Failed to deserialize. Check request payload."}
	ErrResourceNotFound    = GenTreeError{Error: "Resource not found in the database. Check query values."}
)

// ErrInternal exists to avoid forwarding DB errors directly, which could lead to security breach.
