package utils

import (
	"os"
	"path/filepath"
	"testing"
	"archive/zip"
)

func TestUnzip(t *testing.T) {
	// 1. Create a dummy zip file
	tmpDir := t.TempDir()
	zipPath := filepath.Join(tmpDir, "test.zip")
	
	zipFile, err := os.Create(zipPath)
	if err != nil {
		t.Fatal(err)
	}
	
	w := zip.NewWriter(zipFile)
	f, err := w.Create("hello.txt")
	if err != nil {
		t.Fatal(err)
	}
	_, err = f.Write([]byte("Hello World"))
	if err != nil {
		t.Fatal(err)
	}
	w.Close()
	zipFile.Close()

	// 2. Unzip it
	outDir := filepath.Join(tmpDir, "output")
	err = Unzip(zipPath, outDir)
	if err != nil {
		t.Fatalf("Unzip failed: %v", err)
	}

	// 3. Verify content
	content, err := os.ReadFile(filepath.Join(outDir, "hello.txt"))
	if err != nil {
		t.Fatal(err)
	}

	if string(content) != "Hello World" {
		t.Errorf("Expected 'Hello World', got '%s'", string(content))
	}
}
