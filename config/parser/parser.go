package parser

import (
	"github.com/dollarkillerx/chronos/config/chronos"
	"io/ioutil"
	"log"
	"strings"

	"github.com/dollarkillerx/chronos/config/lexer"
	"github.com/dollarkillerx/chronos/config/lexertoken"
)

func Parse(filename string) (chronos.Chronos, error) {
	conf, err := ParseConf(filename)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(conf["matchers"])
}

func ParseConf(filename string) (Sections, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return Sections{}, err
	}

	var output SectionItem

	var token lexertoken.Token
	var tokenValue string

	section := ConfSection{}
	key := ""

	l := lexer.BeginLexing(string(file))
	for {
		token = l.NextToken()
		tokenValue = strings.TrimSpace(token.Value)

		if lexertoken.IsEOF(token) {
			if tokenValue != "" && key != "" {
				section.KeyValuePairs = append(section.KeyValuePairs, ConfKeyValue{Key: key, Value: tokenValue})
			}
			output = append(output, section)
			break
		}

		switch token.Type {
		case lexertoken.TOKEN_SECTION:
			if len(section.KeyValuePairs) > 0 {
				output = append(output, section)
			}

			key = ""
			section.Name = tokenValue
			section.KeyValuePairs = make([]ConfKeyValue, 0)
		case lexertoken.TOKEN_KEY:
			key = tokenValue
		case lexertoken.TOKEN_VALUE:
			section.KeyValuePairs = append(section.KeyValuePairs, ConfKeyValue{Key: key, Value: tokenValue})
			key = ""
		}
	}

	result := Sections{}
	for _, v := range output {
		result[v.Name] = v.KeyValuePairs
	}
	return result, nil
}

type Sections map[string][]ConfKeyValue

type SectionItem []ConfSection

type ConfSection struct {
	Name          string         `json:"name"`
	KeyValuePairs []ConfKeyValue `json:"keyValuePairs"`
}

type ConfKeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
