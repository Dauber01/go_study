package error

import (
	"fmt"
	"io"
)

type Header struct {
	Key, Value string
}

type Status struct {
	Code   int
	Reason string
}

type errorWriter struct {
	io.Writer
	err error
}

// 通过把error包装起来, 减少在业务代码中对 error 判断的处理
func (e *errorWriter) Write(buf []byte) (int, error) {
	if e.err != nil {
		return 0, e.err
	}
	var n int
	n, e.err = e.Writer.Write(buf)
	return n, e.err
}

// 处理完之后的业务流程代码
func WriteResponse(w io.Writer, st Status, headers []Header, body io.Reader) error {
	ew := &errorWriter{Writer: w}
	fmt.Fprintf(ew, "HTTP?1.1 %d %s \r\n", st.Code, st.Reason)

	for _, h := range headers {
		fmt.Fprintf(ew, "%s: %s\r\n", h.Key, h.Value)
	}

	fmt.Fprintf(ew, "\r\n")
	io.Copy(ew, body)
	return ew.err
}
