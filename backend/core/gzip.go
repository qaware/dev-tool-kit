package core

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"errors"
	"io/ioutil"
)

func GzipEncode(text string) (string, error) {
	if text == "" {
		return "", nil
	}

	var buf bytes.Buffer
	writer := gzip.NewWriter(&buf)

	_, err := writer.Write([]byte(text))
	if err != nil {
		DebugError(err)
		return "", errors.New("Error compressing input")
	}

	err = writer.Close()
	if err != nil {
		DebugError(err)
		return "", errors.New("Error compressing input")
	}

	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

func GzipDecode(text string) (string, error) {
	if text == "" {
		return "", nil
	}

	decoded, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		DebugError(err)
		return "", errors.New("Invalid base64 encoding")
	}

	decompressed, err := gzipDecompress(decoded)
	if err != nil {
		DebugError(err)
		return "", errors.New("Invalid gzip compression")
	}
	return string(decompressed), nil
}

func gzipDecompress(input []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(input))
	if err != nil {
		return []byte{}, err
	}

	decompressed, err := ioutil.ReadAll(reader)
	if err != nil {
		return []byte{}, err
	}

	return decompressed, nil
}
