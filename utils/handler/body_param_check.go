package handler

import (
	"errors"
	"fmt"
	"strings"

	"github.com/beego/beego/v2/core/validation"
)

var valid = validation.Validation{}

type CustomErrorMessage struct {
	validation.Validation
}

func (c CustomErrorMessage) String() string {
	var keys string = "[ "
	for _, err := range c.Errors {
		keys += strings.Split(err.Key, ".")[0] + " "
	}
	return keys + "] can not be empty"
}

func BodyParamCheck(p interface{}) error {
	b, err := valid.Valid(p)
	if err != nil {
		return err
	}
	if !b {
		var message CustomErrorMessage
		message.Errors = valid.Errors
		valid.Clear()

		return errors.New(fmt.Sprintf("%v", message))
	}
	return nil
}
