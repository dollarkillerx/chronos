package lexertoken

type Set map[string]bool

type Chronos struct {
	R Set
	P Set
	M []Node
}

// 计算单位定义
type Node struct {
	Type        ChronosType
	Linker      LinkerType
	Participant []ParticipantItem
}

type ParticipantItem struct {
	Participant string
	Object      string
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
