package status

var (
	ErrInternalService = appsta{Code: 10000, Message: "internal server error"}
	ErrInvalidParam    = appsta{Code: 10001, Message: "invalid param"}
	ErrDbOpt           = appsta{Code: 10002, Message: "database operation failed"}
)
