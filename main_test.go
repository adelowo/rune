package main

import (
	"path/filepath"
	"testing"
)

func TestIsGoFile(t *testing.T) {

	table := []struct {
		path  string
		valid bool
	}{
		{"gofile.go", true},
		{"gofile.py", false},
		{"gofile.php", false},
		{"gofile.rb", false},
	}

	for _, v := range table {
		if got := isGoFile(v.path); got != v.valid {
			t.Errorf("Want %v. Got %v", v.valid, got)
		}
	}
}

func TestPathWalker(t *testing.T) {

	filepath.Walk(dir, pathWalker)

	//Fragile test though
	//2 since pathWalker internally acknowledges *.go files only
	if len(files) != 2 {
		t.Errorf("This project currently has %d files. Got %d instead", 2, len(files))
	}
}

func TestIsGetRepo(t *testing.T) {
	if !isGitRepo(dir) {
		t.Error("This is supposed to be a git repository")
	}
}
