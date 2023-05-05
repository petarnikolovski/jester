package api

import "strconv"

func StringToUint(text string) (*uint, error) {
	// This can be a bug on 32 bit platforms - uint is platform specific
	parsed_text, err := strconv.ParseUint(text, 10, 64)
	if err != nil {
		return nil, err
	}

	converted := uint(parsed_text)
	return &converted, nil
}
