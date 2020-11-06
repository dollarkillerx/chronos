package chronos

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dollarkillerx/chronos/adapter"
	"github.com/dollarkillerx/chronos/conf/chronos_token"
	"github.com/dollarkillerx/chronos/conf/parser"
	"github.com/dollarkillerx/chronos/utils"
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

	if len(parse.R) == 0 || len(parse.M) == 0 || len(parse.P) == 0 {
		return nil, fmt.Errorf("The rule is null.")
	}
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

type Item struct {
	Idx    int
	Struct string
}
type EvalItem struct {
	Count       int
	Calculation CalculationType
	Name        string
}

type CalculationType string

const (
	CalculationGreater CalculationType = ">"
	CalculationLess    CalculationType = "<"
)

func (c *Chronos) Enforce(params ...interface{}) (bool, error) {
	r := c.token.R
	if len(params) != len(r) {
		return false, fmt.Errorf("params number!!! ")
	}

	var eval bool

	var keySlice []string
	mKey := map[int]Item{}
	for idx, k := range c.token.P {
		for _, v := range c.token.M {
			if v.Type == chronos_token.CHRONOS_EVAL {
				eval = true
			}

			i := 0
			for _, v2 := range v.Participant {
				if strings.Index(v2.Participant, k) == -1 {
					i++
					continue
				}
			}
			if len(v.Participant) == i {
				continue
			}

			if v.Type == chronos_token.CHRONOS_EQ {
				var rr string
				// 找到第一个Key的M
				if len(v.Participant) == 2 {
					if strings.Index(v.Participant[0].Participant, "r.") != -1 {
						rr = v.Participant[0].Participant
					} else {
						rr = v.Participant[1].Participant
					}
				}
				if strings.Count(rr, ".") == 2 {
					for k1, v1 := range c.token.R {
						if strings.Index(rr, v1.Data) != -1 {
							mKey[idx] = Item{
								Idx:    k1,
								Struct: rr[strings.LastIndex(rr, ".")+1:],
							}
						}
					}
				} else {
					for k1, v1 := range c.token.R {
						if strings.Index(rr, v1.Data) != -1 {
							mKey[idx] = Item{
								Idx: k1,
							}
						}
					}
				}
			}
		}
	}

	for k := range params {
		var data string
		if mKey[k].Struct != "" {
			val, err := utils.GetStructVal(params[k], mKey[k].Struct)
			if err != nil {
				return false, err
			}
			data = val[mKey[k].Struct].Data.(string)
		} else {
			data = params[mKey[k].Idx].(string)
		}
		keySlice = append(keySlice, data)
	}

	if !eval {
		keySlice = append(keySlice[:len(keySlice)-1])
	}

	rule, err := c.adapter.GetRule(keySlice...)
	if err != nil {
		return false, err
	}

	if eval {
		for ic, v := range c.token.R {
			idx := strings.Index(rule, v.Data)
			if idx == -1 {
				continue
			}

			rules := rule[idx+len(v.Data)+1:]
			if i := strings.Index(rules, string(CalculationGreater)); i != -1 {
				key := rules[:i]
				num, err := strconv.Atoi(rules[i+1:])
				if err != nil {
					return false, err
				}
				val, err := utils.GetStructVal(params[ic], key)
				if err != nil {
					return false, err
				}
				if val[key].Data.(int64) > int64(num) {
					return true, nil
				}
				return false, nil
			} else if i := strings.Index(rules, string(CalculationLess)); i != -1 {
				key := rules[:i]
				num, err := strconv.Atoi(rules[i+1:])
				if err != nil {
					return false, err
				}
				val, err := utils.GetStructVal(params[ic], key)
				if err != nil {
					return false, err
				}
				if val[key].Data.(int64) < int64(num) {
					return true, nil
				}
				return false, nil
			}
		}
	} else {
		if rule == params[len(params)-1].(string) {
			return true, nil
		}
	}

	return false, nil
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
