package clipboard

// Returns the current text data of the clipboard.
func Get() (string, error) {
	return get()
}

// Sets the current text data of the clipboard.
func Set(text string) error {
	return set(text)
}
