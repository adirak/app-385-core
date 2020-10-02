package main

import (
	"fmt"

	"github.com/adirak/app-385-core/apis/regis"
	"github.com/adirak/app-385-core/data"
)

// TestSaveRegisterLog step of regis
func main() {

	reqt := data.ReqtData{}
	inData := make(map[string]interface{})
	inData["userId"] = ""
	inData["eventId"] = 1

	reqt.InData = inData

	// Call Test
	resp := regis.SaveRegisterLog(reqt)

	// Print result
	fmt.Println(resp)

}
