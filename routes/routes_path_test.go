package routes

import (
	"testing"
)

func TestTrimSuffixPoint(t *testing.T) {
	name := "file.txt"
	name = TrimSuffixPoint(name)
	if name != "file" {
		t.Fatal("Error trim suffix point")
	}
}

func TestGetSuffixPoint(t *testing.T) {
	name := "file.txt"
	name = GetSuffixPoint(name)
	if name != "txt" {
		t.Fatal("Error get suffix point")
	}
}

func TestTrimSuffixSlash(t *testing.T) {
	path := "../test/data/"
	path = TrimSuffixSlash(path)
	if path != "../test/data" {
		t.Fatal("Error trim suffix slash")
	}
}
