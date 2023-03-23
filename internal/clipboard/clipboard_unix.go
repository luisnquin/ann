package clipboard

import (
	"bytes"
	"os/exec"
)

func get() (string, error) {
	cmd := exec.Command("xclip", "-o")
	b := new(bytes.Buffer)
	cmd.Stdout = b

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return b.String(), nil
}

func set(text string) error {
	cmd := exec.Command("xclip", "-selection", "clipboard")
	cmd.Stdin = bytes.NewReader([]byte(text))

	return cmd.Run()
}
