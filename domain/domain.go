package domain

type returnStatement struct {
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

type function struct {
	Name       string          `json:"name"`
	Return     returnStatement `json:"returnStatement"`
	Parameters []Parameter     `json:"parameters"`
}

type testCase struct {
	Input  []interface{} `json:"input"`
	Output interface{}   `json:"output"`
}

type Problem struct {
	Id          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	TimeLimit   int64      `json:"timeLimit"`
	Function    function   `json:"func"`
	TestCases   []testCase `json:"testCases"`
	Level       int32      `json:"level"`
}
