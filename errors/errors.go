package errors

const (
	/********* DB ERRORS *********/

	// ErrDBEngineInvalid is returned when db's engine implementation does not exist
	ErrDBEngineInvalid modelError = "No implementation for the engine: %s"

	/********* FILE RELATED ERRORS *********/

	// ErrOpenFile is returned when a problem occurs while opening a file
	ErrOpenFile modelError = "Error while opening %s => %s"

	/********* PARSING ERRORS *********/

	// ErrJSONDecoding is returned when an error occurs while parsing a JSON file
	ErrJSONDecoding modelError = "Error while decoding %s : %s"
)

type modelError string

func (e modelError) Error() string {
	return string(e)
}
