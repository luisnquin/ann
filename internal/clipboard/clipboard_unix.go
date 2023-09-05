package clipboard

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func get() (string, error) {
	cmd := exec.Command("xclip", "-selection", "c", "-o")
	var b, bErr bytes.Buffer

	cmd.Stdout = &b
	cmd.Stderr = &bErr

	if err := cmd.Run(); err != nil {
		msg := strings.TrimPrefix(bErr.String(), "Error: ")

		return "", fmt.Errorf("xclip: %w", errors.New(msg))
	}

	return b.String(), nil
}

func set(text string) error {
	cmd := exec.Command("xclip", "-selection", "clipboard")
	cmd.Stdin = bytes.NewReader([]byte(text))

	return cmd.Run()
}
