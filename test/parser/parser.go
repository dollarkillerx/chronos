package parser

import (
	"github.com/dollarkillerx/chronos/test/ini"
	"github.com/dollarkillerx/chronos/test/lexer_factory"
	"github.com/dollarkillerx/chronos/test/lexertoken"
	"log"
	"strings"
)

func isEOF(token lexertoken.Token) bool {
	return token.Type == lexertoken.TOKEN_EOF
}

func Parse(fileName, input string) ini.IniFile {
	output := ini.IniFile{
		FileName: fileName,
		Sections: make([]ini.IniSection, 0),
	}

	var token lexertoken.Token
	var tokenValue string

	/* State variables */
	section := ini.IniSection{}
	key := ""

	log.Println("Starting lexer and parser for file", fileName, "...")

	l := lexer_factory.BeginLexing(fileName, input)

	for {
		token = l.NextToken()

		if token.Type != lexertoken.TOKEN_VALUE {
			tokenValue = strings.TrimSpace(token.Value)
		} else {
			tokenValue = token.Value
		}


		//log.Println(len(output.Sections))
		if isEOF(token) {
			output.Sections = append(output.Sections, section)
			break
		}
		//
		//log.Println(tokenValue)
		//log.Println(len(output.Sections))
		//break


		switch token.Type {
		case lexertoken.TOKEN_SECTION:
			/*
			 * Reset tracking variables
			 */
			if len(section.KeyValuePairs) > 0 {
				output.Sections = append(output.Sections, section)
			}

			key = ""

			section.Name = tokenValue
			section.KeyValuePairs = make([]ini.IniKeyValue, 0)

		case lexertoken.TOKEN_KEY:
			key = tokenValue

		case lexertoken.TOKEN_VALUE:
			section.KeyValuePairs = append(section.KeyValuePairs, ini.IniKeyValue{Key: key, Value: tokenValue})
			key = ""
		}
	}

	log.Println("Parser has been shutdown")
	return output
}
