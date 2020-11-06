package conf

import (
	"encoding/json"
	"github.com/dollarkillerx/chronos/conf/parser"
	"log"
	"testing"
)

func TestMu(t *testing.T) {

	log.SetFlags(log.LstdFlags | log.Llongfile)

	parsedINIFile, err := parser.ParseConf("../base.conf")
	if err != nil {
		log.Fatalln(err)
	}
	prettyJSON, err := json.MarshalIndent(parsedINIFile, "", "   ")

	if err != nil {
		log.Println("Error marshalling JSON:", err.Error())
		return
	}

	log.Println(string(prettyJSON))
}

func TestMu2(t *testing.T) {

	log.SetFlags(log.LstdFlags | log.Llongfile)

	parser.Parse("../base.conf")
}
