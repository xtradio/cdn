package main

import (
	"os"
	"testing"
)

func TestRandomString(t *testing.T) {
	getData := randomString()

	if getData == "" {
		t.Error("Was expecitng a string, got empty.")
	}
}

func TestDoesExist(t *testing.T) {
	testFile := "test.txt"
	getResult := doesExist(testFile)

	if getResult != false {
		t.Errorf("Expected false, got %t", getResult)
	}

	testEmptyFile, err := os.Create(testFile)
	if err != nil {
		t.Fatalf("Error creating file for testing: %s", err)
	}
	testEmptyFile.Close()

	getResult = doesExist(testFile)

	if getResult != true {
		t.Errorf("Expected true, got %t", getResult)
	}

	err = os.Remove(testFile)
	if err != nil {
		t.Fatalf("Error removing %s for cleanup", testFile)
	}
}

func TestGenerateFilename(t *testing.T) {
	getFilename, getFilepath := generateFilename("foo")

	if getFilename == "" {
		t.Errorf("Expecting a value, got empty.")
	}

	if getFilepath == "" {
		t.Errorf("Expecting a value, got empty.")
	}
}
