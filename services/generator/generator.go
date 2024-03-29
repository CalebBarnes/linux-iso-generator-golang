package generator

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func GenerateIso(userDataConfig string) error {
	tempDir, err := os.MkdirTemp("", "iso-gen")
	if err != nil {
		log.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir) // Clean up tmp directory when done.

	userDataPath := filepath.Join(tempDir, "user-data")
	if err := os.WriteFile(userDataPath, []byte(userDataConfig), 0644); err != nil {
		log.Fatalf("Failed to write user-data: %v", err)
	}

	isoOutputPath := "./tmp/output.iso"
	cmd := exec.Command("xorriso", "-as", "mkisofs", "-o", isoOutputPath, tempDir)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to generate ISO: %v", err)
	}

	return nil
}
