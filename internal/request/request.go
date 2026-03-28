package request

import (
	"bytes"
	"fmt"
	"io"
)

type parseState string

const (
	stateInit  parseState = "init"
	stateDone  parseState = "done"
	stateError parseState = "error"
)

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

type Request struct {
	RequestLine RequestLine
	state       parseState
}

func newRequest() *Request {
	return &Request{
		state: stateInit,
	}
}

func (r *Request) done() bool {
	return r.state == stateDone || r.state == stateError
}

func (r *Request) error() bool {
	return r.state == stateError
}

func (r *Request) parse(data []byte) (int, error) {
	read := 0
outer:
	for {
		switch r.state {
		case stateError:
			return 0, ERROR_REQUEST_IN_ERROR
		case stateInit:
			rl, n, err := parseRequestLine(data[read:])
			if err != nil {
				r.state = stateError
				return 0, err
			}
			if n == 0 {
				return 0, err
			}
			r.RequestLine = *rl
			read += n

			r.state = stateDone
		case stateDone:
			break outer
		}
	}
	return read, nil
}

var ERROR_BAD_START_LINE = fmt.Errorf("bad request-line")
var ERROR_REQUEST_IN_ERROR = fmt.Errorf("request in error state")
var SEPARATOR = []byte("\r\n")

func parseRequestLine(b []byte) (*RequestLine, int, error) {
	idx := bytes.Index(b, SEPARATOR)

	if idx == -1 {
		return nil, 0, nil
	}

	//example: "GET / HTTP/1.1"
	startLine := b[:idx]

	// example: "Host: localhost:42069\r\nUser-Agent: curl/7.81.0\r\nAccept: */*\r\n\r\n""
	read := idx + len(SEPARATOR)
	parts := bytes.Split(startLine, []byte(" "))
	if len(parts) != 3 {
		return nil, 0, ERROR_BAD_START_LINE
	}
	// get http method RequestTarget, http version
	httpParts := bytes.Split(parts[2], []byte("/"))
	if len(httpParts) != 2 {
		return nil, 0, ERROR_BAD_START_LINE
	}
	return &RequestLine{
		Method:        string(parts[0]),
		RequestTarget: string(parts[1]),
		HttpVersion:   string(httpParts[1]),
	}, read, nil
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	request := newRequest()

	// TODO: maybe buf will overflow, we don't handle this case now
	buf := make([]byte, 1024)
	bufIdx := 0
	for !request.done() {
		n, err := reader.Read(buf[bufIdx:])
		// TODO
		if err != nil {
			return nil, err
		}
		bufIdx += n
		readN, err := request.parse(buf[:bufIdx])
		if err != nil {
			return nil, err
		}
		copy(buf, buf[readN:bufIdx])
		bufIdx -= readN
	}
	return request, nil
}
