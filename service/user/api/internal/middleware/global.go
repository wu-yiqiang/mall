package middleware

import (
	"bytes"
	"fmt"
	"net/http"
)

type bodyCopy struct {
	http.ResponseWriter
	body *bytes.Buffer
}

func NewBodyCopy(w http.ResponseWriter) *bodyCopy {
	return &bodyCopy{w, bytes.NewBuffer([]byte{})}
}

func (bc bodyCopy) Write(b []byte) (int, error) {
	bc.body.Write(b)
	return bc.ResponseWriter.Write(b)
}

// 2。再往HTTP响应里写响应内容
func CopyResp(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bc := NewBodyCopy(w)
		next(bc, r)
		fmt.Println(r.URL.RawPath, bc.body.String())
	}
}
