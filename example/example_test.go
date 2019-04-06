package main_test

import (
	"os/exec"
	"testing"
)

func TestExamples(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Log("Test #", i+1)
		cmd := exec.Command("go", "run", "example.go")
		err := cmd.Run()
		if err != nil {
			t.Fatal("...FAILED\n")
		}
		t.Log("...SUCCEEDED\n")
	}
}
