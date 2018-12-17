package message

import "encoding/json"

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

var (
	InternalError           = Error{"service temporary unavaible", 1}
	MissingInformationError = Error{"invalid parameters", 2}
	MalformatedDataError    = Error{"data malformated", 3}
	NoDataFoundError        = Error{"no data found", 4}
)

// JSON return a json string of the error struct
func (err Error) JSON() string {
	json, _ := json.Marshal(err)
	return string(json)
}
