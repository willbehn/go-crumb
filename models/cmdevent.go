package models

type CmdEvent struct {
	Id     int     `json:"id"`
	Cmd    string  `json:"cmd"`
	Shell  string  `json:"shell"`
	Dir    string  `json:"dir"`
	Repo   *string `json:"repo"`
	Branch *string `json:"branch"`
	TS     int64   `json:"ts"`
	Exit   int     `json:"exit"`
	Dur    int64   `json:"dur"`
}
