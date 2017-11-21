package main

import "testing"

func TestShouldDraw88(t *testing.T) {
	result := ToLCDMultiple(8, 8, 0, 8, 1)
	expect := ` _  _  _  _ ` + `
|_||_|| ||_|
|_||_||_||_|`
	if result != expect {
		t.Fatalf("expected\n%s\nbut got\n%s", expect, result)
	}
	t.Logf("\n%s", result)
}
