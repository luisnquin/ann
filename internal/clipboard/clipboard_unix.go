package clipboard

import (
	"bytes"
	"os/exec"

	"github.com/d-tsuji/clipboard"
)

func get() (string, error) {
	return clipboard.Get()
}

func set(text string) error {
	cmd := exec.Command("xclip")
	cmd.Stdin = bytes.NewReader([]byte(text))

	return cmd.Run()
}
