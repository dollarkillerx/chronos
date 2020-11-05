package ini

type IniFile struct {
	FileName string       `json:"fileName"`
	Sections []IniSection `json:"sections"`
}

type IniSection struct {
	Name          string        `json:"name"`
	KeyValuePairs []IniKeyValue `json:"keyValuePairs"`
}

type IniKeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

