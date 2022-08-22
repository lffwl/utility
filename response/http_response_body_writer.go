package response

import (
	"bytes"
	"net/http"
)

func NewResponseBodyWriter(w http.ResponseWriter) *ResponseBodyWriter {
	return &ResponseBodyWriter{
		ResponseWriter: w,
		body:           bytes.NewBuffer(nil),
	}
}

type ResponseBodyWriter struct {
	http.ResponseWriter
	body *bytes.Buffer
}

func (w *ResponseBodyWriter) Write(data []byte) (int, error) {
	w.body.Write(data)
	return len(data), nil
}

func (w *ResponseBodyWriter) GetBodyBytes() []byte {
	return w.body.Bytes()
}

func (w *ResponseBodyWriter) GetBodyBytesAndReset() []byte {
	defer w.body.Reset()
	return w.body.Bytes()
}

func (w *ResponseBodyWriter) BodyReset() {
	w.body.Reset()
}

func (w *ResponseBodyWriter) OutPut() {
	if w.body.Len() > 0 {
		w.ResponseWriter.Write(w.body.Bytes())
		w.BodyReset()
	}
}
