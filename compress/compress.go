package compress

const GZIP = 1

func Compress(data []byte, with int) ([]byte, error) {
	if with == GZIP {
		return compressWithGzip(data)
	}
	return compressWithGzip(data)
}
