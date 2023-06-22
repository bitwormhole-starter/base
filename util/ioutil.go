package util

import (
	"io"

	"github.com/starter-go/vlog"
)

// Close 关闭指定的 closer
func Close(c io.Closer) error {
	if c == nil {
		return nil
	}
	err := c.Close()
	if err != nil {
		vlog.Warn(err.Error())
	}
	return err
}

// PumpStream 把数据从 src 读出来，并写入到 dst. 返回成功写入的字节数.
func PumpStream(src io.Reader, dst io.Writer, buffer []byte) (int64, error) {
	if buffer == nil {
		// 默认缓冲区大小为 4k
		buffer = make([]byte, 1024*4)
	}
	return io.CopyBuffer(dst, src, buffer)
}

// IsEOF 检查并判断 err 是否为 EOF
func IsEOF(err error) bool {
	return err == io.EOF
}
