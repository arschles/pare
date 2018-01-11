package run

import (
	"fmt"

	"github.com/robertkrimen/otto"
)

func checkStringInObject(obj *otto.Object, key, expected string) error {
	str, err := stringFromObject(obj, key)
	if err != nil {
		return err
	}
	if str != expected {
		return fmt.Errorf("expected %s to be %s, but got %s", key, expected, str)
	}
	return err
}

func stringFromObject(obj *otto.Object, key string) (string, error) {
	val, err := obj.Get(key)
	if err != nil {
		return "", err
	}
	if !val.IsString() {
		return "", fmt.Errorf("%s wasn't a string", key)
	}
	ret, err := val.ToString()
	if err != nil {
		return "", err
	}
	return ret, nil
}

func numberFromObject(obj *otto.Object, key string) (int64, error) {
	val, err := obj.Get(key)
	if err != nil {
		return 0, err
	}
	if !val.IsNumber() {
		return 0, fmt.Errorf("%s wasn't a number", key)
	}
	ret, err := val.ToInteger()
	if err != nil {
		return 0, err
	}
	return ret, nil
}
