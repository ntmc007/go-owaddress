package test

import (
	"github.com/star001007/go-owaddress"
	"testing"
)

func Test_luna_AddressVerify_Valid(t *testing.T) {

	coin := "luna"
	expect := true

	Address := "terra1sarnv732ew8yj2fxwqn373stxpapx7ulu73vel"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify luna valid address")
	}
}

func Test_luna_AddressVerify_InValid(t *testing.T) {

	coin := "luna"
	expect := false

	Address := "terraa13lshmsctpuyp40gk7n6pc7r4y8wj6uv2x38wet"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify luna invalid address")
	}
}
