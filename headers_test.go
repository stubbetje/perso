package main

import (
	"testing"
)

func TestAddrBraces(t *testing.T) {
	h := ciHeader(make(map[string][]string))
	addr, err := h.parseNonstandardAddress("<something@something.com> (Name Surname)")
	if err != nil {
		t.Error("Unexpected error")
	}
	if addr.Name != "Name Surname" {
		t.Error("Unexpected name ", addr.Name)
	}
	if addr.Address != "something@something.com" {
		t.Error("Unexpected email ", addr.Address)
	}
}
