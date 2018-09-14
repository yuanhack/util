package regx // import "fixup.cc/go/util/regx"

import "regexp"

func IsChineseString(s string) bool {
	if m, _ := regexp.MatchString("^[\\x{4e00}-\\x{9fa5}]+$", s); m {
		return true
	}
	return false
}

func IsEnglishString(s string) bool {
	if m, _ := regexp.MatchString("^[a-zA-Z]+$", s); m {
		return true
	}
	return false
}

func IsEmailAddress(s string) bool {
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, s); m {
		return true
	}
	return false
}

func IsMoblieNumber(s string) bool {
	if m, _ := regexp.MatchString(`^(1[3|4|5|8|7][0-9]\d{4,8})$`, s); m {
		return true
	}
	return false
}

func IsArabicNumeral(s string) bool {
	if m, _ := regexp.MatchString(`^[0-9]+$`, s); m {
		return true
	}
	return false
}

// 二进制
func IsBinNumber(s string) bool {
	if m, _ := regexp.MatchString(`^[+-]{0,1}[01]+$`, s); m {
		return true
	}
	return false
}

// 八进制
func IsOctNumber(s string) bool {
	if m, _ := regexp.MatchString(`^[+-]{0,1}0[1-7]{1,1}[0-7]{0,}$`, s); m {
		return true
	}
	return false
}

// 十进制
func IsDecNumber(s string) bool {
	if m, _ := regexp.MatchString(`^[+-]{0,1}[1-9]\d+$`, s); m {
		return true
	}
	return false
}

// 十六进制
func IsHexNumber(s string) bool {
	if m, _ := regexp.MatchString(`^[+-]{0,1}0[Xx][0-9a-eA-E]+$`, s); m {
		return true
	}
	return false
}
