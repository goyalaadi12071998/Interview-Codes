package errorclass

const (
	BadRequestValidationError = "BAD_REQUEST_VALIDATION_ERROR"
	BadRequestError           = "BAD_REQUEST_ERROR"
	InternalServerError       = "INTERNAL_SERVER_ERROR"
)

var errorList = map[string]*Error{
	BadRequestValidationError: {
		code: BadRequestValidationError,
		name: "BadRequestValidationError",
	},
	InternalServerError: {
		code: InternalServerError,
		name: "InternalServerError",
	},
}
