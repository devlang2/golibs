package compress

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

func compressWithGzip(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	if _, err := zw.Write(data); err != nil {
		return nil, err
	}

	if err := zw.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func decompressWithGzip(s []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(s))
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	if err := reader.Close(); err != nil {
		return nil, err
	}
	return data, nil
}
