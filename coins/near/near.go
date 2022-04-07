package near

import (
	"regexp"
)

// for register
var (
	DefaultStruct = &Verifier{}
	CoinName      = "near"

	reg = regexp.MustCompile("^(([a-z\\d]+[\\-_])*[a-z\\d]+\\.)*([a-z\\d]+[\\-_])*[a-z\\d]+$")
)

type Verifier struct{}

func (b Verifier) IsValid(address string) bool {
	return reg.MatchString(address)
}
