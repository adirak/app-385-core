package zzztest

import (
	"testing"

	"github.com/adirak/app-385-core/apis/regis"
	"github.com/adirak/app-385-core/data"
)

// TestSaveRegisterLog step of regis
func TestSaveRegisterLog(t *testing.T) {

	reqt := data.ReqtData{}
	inData := make(map[string]interface{})
	inData["userId"] = ""
	inData["eventId"] = 1

	reqt.InData = inData

	// Call Test
	resp := regis.SaveRegisterLog(reqt)

	// Print result
	t.Log(resp)

}

func TestAdd(t *testing.T) {
	result := add(2, 4)
	expected := 6
	if result != expected {
		t.Errorf("add() test returned an unexpected result: got %v want %v", result, expected)
	}
}

func add(num1 int, num2 int) int {
	return num1 + num2
}
