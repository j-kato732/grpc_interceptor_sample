package errors

type ErrorBody struct {
	GrpcCode int32         `json:"grpcCode"`
	Message  string        `json:"message"`
	Details  []ErrorDetail `json:"details"`
}

type ErrorDetail struct {
	Code    int32  `json:"code"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

const ErrorDetailKey = "error-detail" // converted to lowdercase in setTrailer

var (
	Period              = ErrorDetail{Name: "period", Code: 100, Message: "period length should be 6"}
	InvalidUserId       = ErrorDetail{Name: "userId", Code: 200, Message: "userId is invalid"}
	InvalidPeriod       = ErrorDetail{Name: "period", Code: 201, Message: "period is invalid"}
	InvalidDepartmentId = ErrorDetail{Name: "departmentId", Code: 202, Message: "departmentId is invalid"}
	InvalidAimId        = ErrorDetail{Name: "aimId", Code: 203, Message: "aimId is invalid"}
	InvalidAimNumber    = ErrorDetail{Name: "aimNumber", Code: 204, Message: "aimNumber is invalid"}
)
