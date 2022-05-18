package errorcode

type ErrorCode string

const (
	Unknown        ErrorCode = "unknown_error"
	Validation               = "validation_error"
	NotFound                 = "notfound_error"
	NotFoundMaster           = "notfound_master_error"
	Database                 = "database_error"
	Redis                    = "redis_error"
	Internal                 = "internal_error"
	BadParams                = "bad_params_error"
)
