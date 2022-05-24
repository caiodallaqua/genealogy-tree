package rest

type GenTreeError struct {
	Error string `json:"error" xml:"error" yaml:"error" example:"Failed to deserialize. Check request."`
}

var (
	// To avoid forwarding DB errors directly, which could lead to security breach.
	ErrInternal = GenTreeError{Error: "Internal server error. Try again later or contact the support team."}

	// Data contract issues.
	ErrFailedToDeserialize = GenTreeError{Error: "Failed to deserialize. Check request payload."}

	// Usually get requests of resources that do not exist in the DB at that moment.
	ErrResourceNotFound = GenTreeError{Error: "Resource not found in the database. Check query values."}
)
