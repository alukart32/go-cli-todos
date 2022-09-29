package errorx

import "errors"

func GetErr(msg string) error {
	return errors.New(msg)
}
