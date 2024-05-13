package mp3

import (
	"fmt"
	"github.com/MamushevArup/ds-voice/internal/usecase/ogg"
	"os"
	"os/exec"
	"path/filepath"
)

func Download(url, uploadID string) (string, error) {
	// Download the file
	oggFilePath, err := ogg.Download(url, uploadID)
	if err != nil {
		return "", fmt.Errorf("failed to download OGG file: %v", err)
	}

	// Define the output file path
	mp3FilePath := oggFilePath[:len(oggFilePath)-len(".ogg")] + ".mp3"

	oggFilePath = filepath.Clean(oggFilePath)
	mp3FilePath = filepath.Clean(mp3FilePath)

	// Create the ffmpeg command
	cmd := exec.Command("ffmpeg", "-i", oggFilePath, mp3FilePath)

	// Run the command
	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to convert OGG to MP3: %v", err)
	}

	// Delete the original OGG file
	err = os.Remove(oggFilePath)
	if err != nil {
		return "", fmt.Errorf("failed to delete OGG file: %v", err)
	}

	// Return the path to the MP3 file
	return mp3FilePath, nil
}
