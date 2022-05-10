package errors

import (
	"mortgage-project/enums"
	"mortgage-project/globals"
	"strings"
)

var (
	/*************************************************************************************************/
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
	/*************************************************************************************************/

	/********* DB ERRORS *********/
	ErrDBQueryExec ModelError = "Oups something went wrong: Query problem"

	ErrDBRowFetching ModelError = "Oups something went wrong: Row fetching problem"

	/********* FILES ERRORS *********/
	ErrorFileOpening ModelError = "Oups can't open this file: Opening file problem"
	ErrorFileWriting ModelError = "Oups can't write to the file: Writing file problem"
)

type ModelError string

func (e ModelError) Error() string {
	msgs := strings.SplitAfter(string(e), ":")

	if globals.Config.GetMode() != enums.Dev.String() {
		// return the first part of the error message without ":"
		return msgs[0][:len(msgs[0])-1]
	}
	return msgs[1]
}
