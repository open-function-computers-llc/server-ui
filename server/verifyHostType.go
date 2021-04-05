package server

import "errors"

func verifyHostType(t string) error {
	validServerHostTypes := []string{
		"modx",
		"wordpress",
	}

	for _, v := range validServerHostTypes {
		if v == t {
			return nil
		}
	}
	return errors.New("Not a valid server type")
}
