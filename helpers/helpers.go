package helpers

import (
	"fmt"
	"log"
	"mortgage-project/customerrors"
	"mortgage-project/enums"
	"mortgage-project/globals"
	"runtime"
	"strings"
)

func CheckErr(err error, msg customerrors.ModelError, args ...string) {
	if err == nil {
		return
	}
	if globals.Config.GetMode() != enums.Dev.String() {
		log.Fatalln(msg.Error())
	} else {
		m := fmt.Sprint(msg.Error(), ": ", strings.Join(args, " | "))
		_, filename, line, _ := runtime.Caller(1)
		src := fmt.Sprintf("[%s:%d]", filename, line)

		log.Fatalln(m, "\n\t", err, "\n\t", src)
	}
}
