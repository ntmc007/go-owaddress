package test

import (
	"github.com/star001007/go-owaddress"
	"github.com/star001007/go-owaddress/coins/near"
	"testing"
)

func Test_near_AddressVerify_Valid(t *testing.T) {
	coin := near.CoinName
	expect := true

	Address := "1f00888f06c20ce63f57a9cbf8856d29c5e0fd59400ca33608230f2b7909431d"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify valid address")
	}
}

func Test_near_AddressVerify_InValid(t *testing.T) {
	coin := near.CoinName
	expect := false
	Address := "019Z7E42k46kxnBjAh8YGXDw3rRGwwxQUBYM7Ccrmwg6ZP0"
	valid, err := owaddress.Verify(coin, Address)
	t.Logf("valid:%v, err:%v", valid, err)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify invalid address")
	}
}
