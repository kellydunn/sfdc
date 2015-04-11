package commands

import("testing")

func TestConvert15CharIdTo18CharId(t * testing.T) {
	expected := "a3XE00000006w0SMAQ"
	actual := Convert15CharIdTo18CharId("a3XE00000006w0S")
	if expected != actual {
		t.Errorf("Unexpected ID conversion result: %s.  Expected: %s.", actual, expected)
	}
}
