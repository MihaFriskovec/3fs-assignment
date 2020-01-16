package customerrors

type CustomError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (e *CustomError) Error() CustomError {
	err := CustomError{}

	err.Type = e.Type
	err.Message = e.Message
	err.Status = e.Status

	return err
}
