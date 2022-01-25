package Utilitys

import (
	"net/mail"
	"regexp"
	"unicode"
)

func CheckPassword(pass string) *LogInstance {
	var (
		upp, low, num, sym bool
		tot                uint8
	)
	for _, char := range pass {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
			tot++
		default:
			return Logger("CheckPassword", "Invalid char in Pass", "Password is hidden", false)
		}
	}

	if !upp || !low || !num || !sym || tot < 8 {
		return Logger("CheckPassword", "Invalid Pass Length", "Password is hidden", false)
	}

	return nil
}

func CheckMail(address string) *LogInstance {
	if _, err := mail.ParseAddress(address); err != nil {
		return Logger("CheckMail", "Mail Invalid", address, err)
	}
	return nil
}

func CheckPhoneNumber(CellNo string) *LogInstance {
	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	if re.MatchString(CellNo) {
		return nil
	}
	return Logger("CheckPhoneNumber", "Phone Number", CellNo, false)
}

func CheckName(name string) *LogInstance {
	if name == "" {
		return Logger("CheckName", "Name is invalid", name, "")
	}
	if len(name) < 2 || len(name) > 40 {
		return Logger("CheckName", "Name length is invalid", name, len(name))
	}
	return nil
}
