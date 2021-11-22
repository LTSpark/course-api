package utils

import "errors"

var errNotFound = errors.New("value not found")

func Contains(s []string, str string) error {
	for _, v := range s {
		if v == str {
			return nil
		}
	}

	return errNotFound
}
