package greet

import "testing"

func TestGreetName(t *testing.T) {
	given := "Skooldio"
	expect := "Hi, Skooldio"

	actual := greet(given)

	if expect != actual {
		t.Errorf("given %q, expect %q, actual %q\n", given, expect, actual)
	}
}
