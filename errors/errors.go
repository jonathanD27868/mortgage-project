package errors

import (
	"mortgage-project/enums"
	"mortgage-project/globals"
)

var (
// Those are initialization errors
// WE CAN'T USE THEM AT THIS POINT SINCE globals.Config DOES NOT EXIST YET

// /********* DB ERRORS *********/

// // ErrDBEngineInvalid is returned when db's engine implementation does not exist
// ErrDBEngineInvalid modelError = modelError{"No implementation for the engine: %s", ""}

// /********* FILE RELATED ERRORS *********/

// // ErrOpenFile is returned when a problem occurs while opening a file
// ErrOpenFile modelError = modelError{
// 	"Error while opening a file",
// 	"Error while opening %s => %s",
// }

// /********* PARSING ERRORS *********/

// // ErrJSONDecoding is returned when an error occurs while parsing a JSON file
// ErrJSONDecoding modelError = modelError{
// 	"Error while reading a File",
// 	"Error while decoding %s : %s",
// }
)

type modelError struct {
	public  string
	private string
}

func (e modelError) Error() string {
	if globals.Config.GetMode() == enums.Dev.String() && len(e.private) != 0 {
		return e.private
	}
	return e.public
}
