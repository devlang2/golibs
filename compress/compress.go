package compress

const (
	GZIP = "GZIP"
)

func Compress(data []byte, with string) ([]byte, error) {
	if with == GZIP {
		return compressWithGzip(data)
	}
	return compressWithGzip(data)
}

func Decompress(data []byte, with string) ([]byte, error) {
	if with == GZIP {
		return decompressWithGzip(data)
	}
	return decompressWithGzip(data)
}
