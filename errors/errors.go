package errors

import (
	"mortgage-project/enums"
	"mortgage-project/globals"
	"strings"
)

var (
	/********* DB ERRORS *********/
	// ErrDBQueryExec returns an error when something went wrong with a db query
	ErrDBQueryExec ModelError = "Oups something went wrong: Query problem"

	// ErrDBRowFetching for errors while fetching the db rows
	ErrDBRowFetching ModelError = "Oups something went wrong: Row fetching problem"

	/********* FILES ERRORS *********/
	ErrorFileOpening ModelError = "Oups can't open this file: Opening file problem"

	ErrorFileWriting ModelError = "Oups can't write to the file: Writing file problem"
)

type ModelError string

func (e ModelError) Error() string {
	msgs := strings.SplitAfter(string(e), ":")

	// public error message
	if globals.Config.GetMode() != enums.Dev.String() {
		// return the first part of the error message without ":"
		return msgs[0][:len(msgs[0])-1]
	}

	// private error message
	return msgs[1]
}
