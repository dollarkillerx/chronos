package chronos

import (
	"fmt"

	"github.com/dollarkillerx/chronos/adapter"
	"github.com/dollarkillerx/chronos/conf/chronos_token"
	"github.com/dollarkillerx/chronos/conf/parser"
)

type Chronos struct {
	token   chronos_token.Chronos
	adapter adapter.Adapter
}

func NewEnforcer(confPath string, adapter adapter.Adapter) (*Chronos, error) {
	e := &Chronos{
		adapter: adapter,
	}
	parse, err := parser.Parse(confPath)
	if err != nil {
		return nil, err
	}
	e.token = parse

	return e, nil
}

func (c *Chronos) AddPolicy(params ...interface{}) (bool, error) {
	var rules []string
	for _, v := range params {
		s, ex := v.(string)
		if !ex {
			continue
		}
		rules = append(rules, s)
	}

	err := c.adapter.AddRule(rules...)
	if err == nil {
		return true, nil
	}
	return false, err
}

func (c *Chronos) Enforce(params ...interface{}) (bool, error) {
	r := c.token.R
	if len(params) != len(r) {
		return false, fmt.Errorf("params number!!! what fuck?")
	}

}

func (c *Chronos) RemovePolicy(params ...interface{}) (bool, error) {
	var rules []string
	for _, v := range params {
		s, ex := v.(string)
		if !ex {
			continue
		}
		rules = append(rules, s)
	}

	err := c.adapter.RemoveRule(rules...)
	if err == nil {
		return true, nil
	}
	return false, err
}

func (c *Chronos) GetFilteredPolicy(params ...interface{}) (string, error) {
	var rules []string
	for _, v := range params {
		s, ex := v.(string)
		if !ex {
			continue
		}
		rules = append(rules, s)
	}

	return c.adapter.GetRule(rules...)
}
