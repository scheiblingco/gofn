package webtools

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"
	"strings"
)

type MultipartField struct {
	Key         string
	value       io.Reader
	filename    *string
	contentType *string
}

func NewMultipartField(key string) *MultipartField {
	return &MultipartField{
		Key: key,
	}
}

func (m *MultipartField) WithStringValue(value string) *MultipartField {
	m.value = strings.NewReader(value)
	return m
}

func (m *MultipartField) WithBytesValue(value []byte) *MultipartField {
	m.value = strings.NewReader(string(value))
	return m
}

func (m *MultipartField) WithReaderValue(value io.Reader) *MultipartField {
	m.value = value
	return m
}

func (m *MultipartField) WithFilename(filename string) *MultipartField {
	m.filename = &filename
	return m
}

func (m *MultipartField) WithContentType(contentType string) *MultipartField {
	m.contentType = &contentType
	return m
}

func (m *MultipartField) WithPipe(pipe *io.PipeReader) *MultipartField {
	m.value = pipe
	return m
}

func (m *MultipartField) AddToWriter(w *multipart.Writer) error {
	if x, ok := m.value.(io.Closer); ok {
		defer x.Close()
	}

	var fw io.Writer
	var err error

	if m.filename != nil || m.contentType != nil {
		partHeader := textproto.MIMEHeader{}

		if m.filename != nil {
			disp := fmt.Sprintf(`form-data; name="%s"; filename="%s"`, m.Key, *m.filename)
			partHeader.Add("Content-Disposition", disp)
		}

		if m.contentType != nil {
			partHeader.Add("Content-Type", *m.contentType)
		}

		fw, err = w.CreatePart(partHeader)
	} else {
		fw, err = w.CreateFormField(m.Key)
	}

	if err != nil {
		return err
	}

	if _, err := io.Copy(fw, m.value); err != nil {
		return err
	}

	return nil
}
