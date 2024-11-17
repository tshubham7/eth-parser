package utils

import (
	"fmt"
	"strconv"
)

func HexToInt(hexStr string) (int, error) {
	if len(hexStr) < 2 || hexStr[:2] != "0x" {
		return 0, fmt.Errorf("invalid hex string: %s", hexStr)
	}
	value, err := strconv.ParseInt(hexStr[2:], 16, 64)
	if err != nil {
		return 0, err
	}
	return int(value), nil
}

func IntToHex(value int) string {
	return fmt.Sprintf("0x%x", value)
}
