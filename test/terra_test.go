package test

import (
	"github.com/star001007/go-owaddress"
	"testing"
)

func Test_terra_AddressVerify_Valid(t *testing.T) {

	coin := "terra"
	expect := true

	Address := "terra1sarnv732ew8yj2fxwqn373stxpapx7ulu73vel"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify terra valid address")
	}
}

func Test_terra_AddressVerify_InValid(t *testing.T) {

	coin := "terra"
	expect := false

	Address := "terraa13lshmsctpuyp40gk7n6pc7r4y8wj6uv2x38wet"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify terra invalid address")
	}
}
