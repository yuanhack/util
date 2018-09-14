package regx

import (
	"share/logging"
	"testing"
)

func init() {
	logging.SetLogModel(true, true)
}

func TestIsChineseString(t *testing.T) {
	s := "全是中文吗"
	logging.Info("test string: %v", s)
	if IsChineseString(s) {
		logging.Info("pass")
	} else {
		logging.Info("no pass")
	}

	s += " "
	logging.Info("test string: %v", s)
	if IsChineseString(s) {
		logging.Info("pass")
	} else {
		logging.Info("no pass")
	}
}

func TestIsDecNumber(t *testing.T) {
	s := "1234789"
	logging.Info("test string: %v", s)
	if IsDecNumber(s) {
		logging.Info("pass")
	} else {
		logging.Info("no pass")
	}

	s = "1234789"
	logging.Info("test string: %v", s)
	if IsDecNumber(s) {
		logging.Info("pass")
	} else {
		logging.Info("no pass")
	}
}

func TestIsOctNumber(t *testing.T) {
	s := "-0111"
	logging.Info("test string: %v", s)
	if IsOctNumber(s) {
		logging.Info("pass")
	} else {
		logging.Info("no pass")
	}
}

func TestIsHexNumber(t *testing.T) {
	s := "0X0012312ae"
	logging.Info("test string: %v", s)
	if IsHexNumber(s) {
		logging.Info("pass")
	} else {
		logging.Info("no pass")
	}
}

func TestIsBinNumber(t *testing.T) {
	s := "01111"
	logging.Info("test string: %v", s)
	if IsBinNumber(s) {
		logging.Info("pass")
	} else {
		logging.Info("no pass")
	}
}
