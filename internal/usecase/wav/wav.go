package wav

import (
	"fmt"
	"github.com/MamushevArup/ds-voice/internal/usecase/ogg"
	"os"
	"os/exec"
)

func Download(url, uploadID string) (string, error) {
	// Download the file
	oggFilePath, err := ogg.Download(url, uploadID)

	// Define the output file path
	wavFilePath := oggFilePath[:len(oggFilePath)-len(".ogg")] + ".wav"

	// Create the ffmpeg command
	cmd := exec.Command("ffmpeg", "-i", oggFilePath, wavFilePath)

	// Run the command
	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to convert OGG to WAV: %v", err)
	}

	// Delete the original OGG file
	err = os.Remove(oggFilePath)
	if err != nil {
		return "", fmt.Errorf("failed to delete OGG file: %v", err)
	}

	// Return the path to the WAV file
	return wavFilePath, nil
}
