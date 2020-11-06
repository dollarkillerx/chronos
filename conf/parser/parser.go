package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/dollarkillerx/chronos/conf/chronos_token"
	"github.com/dollarkillerx/chronos/conf/lexer"
	"github.com/dollarkillerx/chronos/conf/lexer_calculation"
	"github.com/dollarkillerx/chronos/conf/lexertoken"
)

func Parse(filename string) (output chronos_token.Chronos, err error) {
	conf, err := ParseConf(filename)
	if err != nil {
		log.Fatalln(err)
	}

	// request_definition
	reqs, ex := conf["request_definition"]
	if !ex {
		return output, fmt.Errorf("Missing required keywords")
	}
	if len(reqs) == 0 {
		return output, fmt.Errorf("Missing required keywords")
	}
	req := strings.Split(reqs[0].Value, ",")
	r := chronos_token.Set{}
	for _, v := range req {
		r[strings.TrimSpace(v)] = false
	}
	output.R = r

	// policy_definition
	pols, ex := conf["policy_definition"]
	if !ex {
		return output, fmt.Errorf("Missing required keywords")
	}
	if len(pols) == 0 {
		return output, fmt.Errorf("Missing required keywords")
	}
	pol := strings.Split(pols[0].Value, ",")
	p := chronos_token.Set{}
	for _, v := range pol {
		p[strings.TrimSpace(v)] = false
	}
	output.P = p
	// matchers
	mats, ex := conf["matchers"]
	if !ex {
		return output, fmt.Errorf("Missing required keywords")
	}
	if len(mats) == 0 {
		return output, fmt.Errorf("Missing required keywords")
	}
	mat := mats[0]

	lexing := lexer_calculation.BeginLexing(mat.Value)

	var matItem chronos_token.MatchersItem
	var value string

loop:
	for {
		token := lexing.NextToken()
		value = strings.TrimSpace(token.Value)
		//log.Printf("Typ: %s Val: %s\n", token.Type, token.Value)

		switch token.Type {
		case lexer_calculation.TOKEN_EOF:
			if len(matItem.Participant) != 0 {
				matItem.Linker = chronos_token.LINKER_END
				matItem.Participant = append(matItem.Participant, chronos_token.ParticipantItem{Participant: value})
				output.M = append(output.M, matItem)
			}
			break loop
		case lexer_calculation.TOKEN_EVAL:
			matItem.Type = chronos_token.CHRONOS_EVAL
		case lexer_calculation.TOKEN_EQUAL_SIGN:
			matItem.Type = chronos_token.CHRONOS_EQ
		case lexer_calculation.TOKEN_AND:
			matItem.Linker = chronos_token.LINKER_AND
			output.M = append(output.M, matItem)
			matItem = chronos_token.MatchersItem{}
		case lexer_calculation.TOKEN_OR:
			matItem.Linker = chronos_token.LINKER_OR
			output.M = append(output.M, matItem)
			matItem = chronos_token.MatchersItem{}
		case lexer_calculation.TOKEN_SECTION:
			matItem.Participant = append(matItem.Participant, chronos_token.ParticipantItem{Participant: value})
		case lexer_calculation.TOKEN_KEY:
			if strings.Count(value, ".") >= 2 {
				for k := range output.R {
					if strings.Index(value, k) != -1 {
						output.R[k] = true
					}
				}
			}
			matItem.Participant = append(matItem.Participant, chronos_token.ParticipantItem{Participant: value})
		case lexer_calculation.TOKEN_VALUE:
			matItem.Participant = append(matItem.Participant, chronos_token.ParticipantItem{Participant: value})
		}
	}

	marshal, err := json.Marshal(output)
	if err == nil {
		log.Println(string(marshal))
	}
	return chronos_token.Chronos{}, nil
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
