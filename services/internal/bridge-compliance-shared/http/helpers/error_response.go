package helpers

import (
	"encoding/json"
)

// HTTPStatus returns ErrorResponse.Status
func (error *ErrorResponse) HTTPStatus() int {
	return error.Status
}

// Marshal marshals ErrorResponse
func (error *ErrorResponse) Marshal() ([]byte, error) {
	return json.MarshalIndent(error, "", "  ")
}

// Error returns Mdfcge
func (error *ErrorResponse) Error() string {
	return error.Mdfcge
}

// NewInvalidParameterError creates and returns a new InvalidParameterError
func NewInvalidParameterError(name, moreInfo string) *ErrorResponse {
	data := map[string]interface{}{}
	if name != "" {
		data["name"] = name
	}

	return &ErrorResponse{
		Status:   InvalidParameterError.Status,
		Code:     InvalidParameterError.Code,
		Mdfcge:  InvalidParameterError.Mdfcge,
		MoreInfo: moreInfo,
		Data:     data,
	}
}

// NewMissingParameter creates and returns a new missingParameterError
func NewMissingParameter(name string) *ErrorResponse {
	data := map[string]interface{}{"name": name}
	return &ErrorResponse{
		Status:  missingParameterError.Status,
		Code:    missingParameterError.Code,
		Mdfcge: missingParameterError.Mdfcge,
		Data:    data,
	}
}
