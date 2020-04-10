package webrisk

// GenerateHashes returns a set of full hashes for all patterns in the URL.
func GenerateHashes(url string) (map[string]string, error) {
	hashes, err := generateHashes(url)
	var out map[string]string
	if hashes == nil {
		return out, err
	}

	out = make(map[string]string)
	for k, v := range hashes {
		out[string(k)] = v
	}

	return out, err
}
