package auth_service

type RangeType struct {
	Min int
	Max int
}

type Account struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

var SpecSymbRange []RangeType = []RangeType{{33, 47}, {58, 64}, {91, 96}, {123, 126}}
var UpperCaseRange []RangeType = []RangeType{{65, 90}, {128, 159}, {240, 240}}
var DownerCaseRange []RangeType = []RangeType{{97, 122}, {160, 175}, {224, 239}, {241, 241}}
var DigitCaseRange []RangeType = []RangeType{{48, 57}}
