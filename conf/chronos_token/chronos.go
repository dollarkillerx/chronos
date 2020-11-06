package chronos_token

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
	LINKER_AND LinkerType = "and"
	LINKER_OR  LinkerType = "or"
	LINKER_END LinkerType = "end"
)

type ChronosType string

const (
	CHRONOS_EVAL  ChronosType = "eval"
	CHRONOS_EQ    ChronosType = "eq"
	CHRONOS_COMBO ChronosType = "combo"
)
