package tecutils

import (
	"log"
	"reflect"
	"runtime"
)

type ErpError struct {
	ErrorText string
	Function  string
}

func (e *ErpError) Error() string {
	return e.ErrorText
}

func FromError(caller interface{}, err error) error {
	if err != nil {
		errorText, callerName := err.Error(), ""
		if caller != nil {
			callerName = getFunctionName(caller)
		}
		e := &ErpError{ErrorText: errorText, Function: callerName}
		log.Printf("Error logged at function: %s\n%s\n", e.Function, errorText)
		return e
	} else {
		return nil
	}
}

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
