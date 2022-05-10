package helpers

import (
	"log"
	"mortgage-project/enums"
	"mortgage-project/errors"
	"mortgage-project/globals"
)

func CheckErr(err error, msg errors.ModelError) {
	if err == nil {
		return
	}
	if globals.Config.GetMode() != enums.Dev.String() {
		log.Fatalln(msg.Error())
	} else {
		log.Fatalln(msg.Error(), "=> ", err)
	}
}
