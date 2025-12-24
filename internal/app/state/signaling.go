package state

type SignalingError struct {
	// Summary must contain a user friendly message.
	Summary string
	// Description must contain detailed information that is useful for debugging by a developer.
	Description string
	Err         JSONSerializableError
}

func NewSignalingError(summary string, description string, err error) *SignalingError {
	return &SignalingError{
		Summary:     summary,
		Description: description,
		Err:         &JSONSerializableErr{Err: err},
	}
}

// Error implements the error interface for SignalingError.
func (se *SignalingError) Error() string {
	return se.Err.Error()
}
