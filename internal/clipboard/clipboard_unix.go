package clipboard

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func get() (string, error) {
	var cmd *exec.Cmd

	if isUsingWayland() {
		cmd = exec.Command("wl-paste")
	} else {
		cmd = exec.Command("xclip", "-selection", "c", "-o")
	}

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
	var cmd *exec.Cmd

	if isUsingWayland() {
		cmd = exec.Command("wl-copy")
	} else {
		cmd = exec.Command("xclip", "-selection", "clipboard")
	}

	cmd.Stdin = bytes.NewReader([]byte(text))

	return cmd.Run()
}

func isUsingWayland() bool {
	sessionType, waylandDisplay := os.Getenv("XDG_SESSION_TYPE"), os.Getenv("WAYLAND_DISPLAY")

	return sessionType == "wayland" || strings.HasPrefix(waylandDisplay, "wayland")
}
