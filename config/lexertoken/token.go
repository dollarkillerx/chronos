package lexertoken

type Set map[string]bool // true is Struct

type Chronos struct {
	R Set
	P Set
	M Matchers
}

type Matchers []MatchersItem

// 计算单位定义
type MatchersItem struct {
	Type        ChronosType
	Linker      LinkerType
	Participant []ParticipantItem
}

type ParticipantItem struct {
	Participant string
}

type LinkerType string

const (
	LinkerAnd LinkerType = "and"
	LinkerOr  LinkerType = "or"
	LinkerEnd LinkerType = "end"
)

type ChronosType string

const (
	ChronosEval  ChronosType = "eval"
	ChronosEq    ChronosType = "eq"
	ChronosCombo ChronosType = "combo"
)
