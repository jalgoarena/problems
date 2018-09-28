package domain

type ReturnStatement struct {
	Type    string `json:"type"`
	Comment string `json:"comment"`
	Generic string `json:"generic"`
}

type Parameter struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Comment string `json:"comment"`
	Generic string `json:"generic"`
}

type Function struct {
	Name       string          `json:"name"`
	Return     ReturnStatement `json:"returnStatement"`
	Parameters []Parameter     `json:"parameters"`
}

type TestCase struct {
	Input  []interface{} `json:"input"`
	Output interface{}   `json:"output"`
}

type Problem struct {
	Id          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	TimeLimit   int64      `json:"timeLimit"`
	Function    Function   `json:"func"`
	TestCases   []TestCase `json:"testCases"`
	Level       int32      `json:"level"`
}
