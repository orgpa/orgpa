package message

import "encoding/json"

// Error will be used to send error message over JSON.
// The Error contain a message ex: "no internet connection"
// and a code refering to the message ex: 1001.
//
// Error contains a JSON method that return the JSON format
// of the error so we can send it over HTTP.
type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

var (
	// InternalError when the problem comes from inside our server or code
	InternalError = Error{"service temporary unavaible", 1000}

	// MissingInformationError when an information or parameter is missing
	MissingInformationError = Error{"invalid parameters", 2000}

	// MalformatedDataError when the given data is malformated or invalid
	MalformatedDataError = Error{"data malformated", 2001}

	// NoDataFoundError when 0 row or data have been found
	NoDataFoundError = Error{"no data found", 3000}
)

// JSON method will return a JSON version of the given Error.
// No error will be returned as the given Error should be
// correct and entered by the developper.
func (err Error) JSON() string {
	json, _ := json.Marshal(err)
	return string(json)
}

// IsNoDataErr check if the given error is a NoDataFoundError.
// If it is a NoDataFoundError "true" value will be returned,
// otherwise "false" will be returned.
func IsNoDataErr(errUser error) bool {
	if NoDataFoundError.Message == errUser.Error() {
		return true
	}
	return false
}
