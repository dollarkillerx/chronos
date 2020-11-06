package adapter

type Adapter interface {
	AddRule(rules ...string) error
	RemoveRule(rules ...string) error
	GetRule(rules ...string) (rule string, err error)
}
