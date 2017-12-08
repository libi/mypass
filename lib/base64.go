package lib

import (
	"strings"
	"encoding/base64"
)

func DecodeSegment(seg string) ([]byte, error) {
	if l := len(seg) % 4; l > 0 {
		seg += strings.Repeat("=", 4-l)
	}

	return base64.URLEncoding.DecodeString(seg)
}
func EncodeSegment(s []byte)(string){
	return strings.Replace(base64.URLEncoding.EncodeToString(s),"=","",-1)
}
