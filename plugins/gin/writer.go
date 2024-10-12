package gin

import (
	"bytes"
	"github.com/gin-gonic/gin"
)

type CustomWriter struct {
	gin.ResponseWriter
	Data  *bytes.Buffer
	Limit uint
}

func (w CustomWriter) Write(b []byte) (int, error) {
	w.Data.Write(Limit(b, w.Limit))
	return w.ResponseWriter.Write(b)
}

func (w CustomWriter) WriteString(s string) (int, error) {
	w.Data.WriteString(string(Limit([]byte(s), w.Limit)))
	return w.ResponseWriter.WriteString(s)
}
