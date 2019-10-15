package core

import (
	"bytes"
	"strconv"
)

func DecToBinaryForInt(str string) string {
	i, err := strconv.Atoi(str)
	if err != nil || i == 0 {
		return str
	}
	bys := bytes.Buffer{}
	for ; i > 0; i /= 2 {
		bys.WriteString(strconv.Itoa(i%2))
	}
	return bys.String()
}
