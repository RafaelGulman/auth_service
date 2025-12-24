package auth_service

import (
	"errors"
	mathV2 "math/rand/v2"
	"strings"
)

var accounts []Account

// Function is generate random ASCII symbol from range which defined in RangeType
func generateFromRange(rt ...RangeType) string {
	tempRange := rt[mathV2.IntN(len(rt))]
	if tempRange.Max-tempRange.Min == 0 {
		return string(tempRange.Min)
	}
	return string(mathV2.IntN(tempRange.Max-tempRange.Min) + tempRange.Min)
}

// Generate random password
func GenerateRandomPasswordV2() string {
	var completePassword string

	completePassword += generateFromRange(SpecSymbRange...)
	completePassword += generateFromRange(DigitCaseRange...)
	completePassword += generateFromRange(UpperCaseRange...)
	completePassword += generateFromRange(DownerCaseRange...)

	itterationCount := mathV2.IntN(8) + 4

	for i := 0; i < itterationCount; i++ {
		switch mathV2.IntN(4) {
		case 0:
			completePassword += generateFromRange(SpecSymbRange...)
		case 1:
			completePassword += generateFromRange(DigitCaseRange...)
		case 2:
			completePassword += generateFromRange(UpperCaseRange...)
		case 3:
			completePassword += generateFromRange(DownerCaseRange...)
		}
	}
	newStrArr := strings.Fields(completePassword)
	mathV2.Shuffle(len(newStrArr), func(i, j int) { newStrArr[i], newStrArr[j] = newStrArr[j], newStrArr[i] })
	completePassword = strings.Join(newStrArr, "")
	return completePassword
}

// Check string on compliance propertis of passwords ()
func PasswordValidation(pass string) []error {
	var allErrors []error

	if len(pass) < 8 {
		allErrors = append(allErrors, errors.New("PASSWORD_IS_SHORT"))
	}
	if !checkWithAscii(pass, UpperCaseRange...) {
		allErrors = append(allErrors, errors.New("PASSWORD_HAVE_NOT_UPPER_CASE"))
	}

	if !checkWithAscii(pass, DownerCaseRange...) {
		allErrors = append(allErrors, errors.New("PASSWORD_HAVE_NOT_DOWNER_CASE"))
	}

	if !checkWithAscii(pass, SpecSymbRange...) {
		allErrors = append(allErrors, errors.New("PASSWORD_HAVE_NOT_SPECIAL_SYMBOL"))
	}
	if !checkWithAscii(pass, DigitCaseRange...) {
		allErrors = append(allErrors, errors.New("PASSWORD_HAVE_NOT_DIGIT"))
	}

	return allErrors
}

func checkWithAscii(ascii string, rt ...RangeType) bool {
	for _, v := range rt {
		for _, strVal := range ascii {
			if uint(strVal) >= uint(v.Min) && uint(strVal) <= uint(v.Max) {
				return true
			}

		}
	}
	return false
}

func CheckLogin(acc Account) bool {
	for _, v := range accounts {
		if v.Login == acc.Login && v.Password == acc.Password {
			return true
		}
	}
	return false
}
