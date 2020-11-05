package test

import (
	"encoding/json"
	"github.com/dollarkillerx/chronos/test/parser"
	"log"
	"strings"
	"testing"
)

func TestMu(t *testing.T) {

	log.SetFlags(log.LstdFlags | log.Llongfile)
	sampleInput := `
		key=abcdefg

		[User]
		userName=adampresley
		keyFile=~/path/to/keyfile

		[Servers]
		server1=localhost:8080
	`

	//parsedINIFile := parser.Parse("sample.ini", sampleInput)
	//prettyJSON, err := json.MarshalIndent(parsedINIFile, "", "   ")
	//
	//if err != nil {
	//	log.Println("Error marshalling JSON:", err.Error())
	//	return
	//}
	//
	//log.Println(string(prettyJSON))

	parsedINIFile := parser.Parse("sample.ini", sampleInput)
	_, err := json.MarshalIndent(parsedINIFile, "", "   ")

	if err != nil {
		log.Println("Error marshalling JSON:", err.Error())
		return
	}

	//log.Println(string(prettyJSON))
}

func TestB(t *testing.T) {
	a := "sadasd=sdad"
	log.Println(strings.HasPrefix(a, "="))
}
