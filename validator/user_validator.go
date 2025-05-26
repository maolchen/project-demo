package validator

import (
	"fmt"
	"github.com/maolchen/project_demo/constants"
	"regexp"

	"github.com/go-playground/validator/v10"
)

// ValidatePassword 校验密码复杂度（至少8位，包含大小写、数字、特殊字符）
func ValidatePassword(password string) error {
	v := validator.New()
	v.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		pwd := fl.Field().String()

		if len(pwd) < 8 {
			return false
		}

		hasUpper := regexp.MustCompile(`[A-Z]`).MatchString
		hasLower := regexp.MustCompile(`[a-z]`).MatchString
		hasDigit := regexp.MustCompile(`\d`).MatchString
		hasSpecial := regexp.MustCompile(`[^a-zA-Z\d\s:]`).MatchString

		return hasUpper(pwd) && hasLower(pwd) && hasDigit(pwd) && hasSpecial(pwd)
	})

	err := v.Var(password, "password")
	if err != nil {
		return fmt.Errorf(constants.UserPassValidatorFail)
	}
	return nil
}

// ValidateUsername 校验用户名只能是字母或字母+数字（不能有下划线、符号等）
func ValidateUsername(username string) error {
	v := validator.New()
	v.RegisterValidation("username", func(fl validator.FieldLevel) bool {
		matched, _ := regexp.MatchString(`^[A-Za-z][A-Za-z0-9]*$`, fl.Field().String())
		return matched
	})

	err := v.Var(username, "username")
	if err != nil {
		return fmt.Errorf(constants.UsernameValidatorFail)
	}
	return nil
}
