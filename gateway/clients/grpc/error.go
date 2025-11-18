package grpc

type GrpcError struct {
	Code    int32
	Message string
}

func (e *GrpcError) Error() string {
	return e.Message
}

func NewError(code int32, message string) *GrpcError {
	return &GrpcError{
		Code:    code,
		Message: message,
	}
}
