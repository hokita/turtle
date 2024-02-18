package main

import "testing"

func TestTalk(t *testing.T) {
	if talk() != "I'm walking slowly." {
		t.Error("Test failed")
	}
}
