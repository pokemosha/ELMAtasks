package model

type TaskData struct {
	Username string `json:"user_name"`
	Task     string `json:"task"`
	Result   `json:"results"`
}

type Result struct {
	Payload [][]interface{} `json:"payload"`
	Results []interface{}   `json:"results"`
}

type Solution struct {
	Percent int            `json:"percent"`
	Fails   []SolutionFail `json:"fails"`
}
type SolutionFail struct {
	OriginalResult interface{}
	ExternalResult interface{}
}
