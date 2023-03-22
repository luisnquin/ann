package clipboard

import "github.com/d-tsuji/clipboard"

func get() (string, error) {
	return clipboard.Get()
}

func set(text string) error {
	return clipboard.Set(text)
}
