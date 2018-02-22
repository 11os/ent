package utils

import (
	"fmt"

	"golang.org/x/text/encoding/simplifiedchinese"
)

func decodeToGBK(text string) (string, error) {
	dst := make([]byte, len(text)*2)
	tr := simplifiedchinese.GB18030.NewDecoder()
	nDst, _, err := tr.Transform(dst, []byte(text), true)
	if err != nil {
		return text, err
	}

	return string(dst[:nDst]), nil
}

func Print(msg string) {
	result, _ := decodeToGBK(msg)
	fmt.Println(result)
}
