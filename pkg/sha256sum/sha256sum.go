package sha256sum

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"

	"goyave.dev/goyave/v5/util/errors"
)

// CalcuLateSHA256Sum 计算文件SHA256 Hash值
func CalcuLateSHA256Sum(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", errors.New(err)
	}
	defer file.Close()

	hasher := sha256.New()
	// 4MB 缓冲
	buf := make([]byte, 4*1024*1024)

	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			return "", errors.New(err)
		}
		if n == 0 {
			break
		}

		hasher.Write(buf[n:])
	}

	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}
