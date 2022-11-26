package dio

import (
	"errors"
	"unsafe"
)

var (
	sectorSize = 512
)

func AlterSectorSize(size int) {
	sectorSize = size
}

func align(buf []byte) int {
	//return int(uintptr(unsafe.Pointer(&buf[0])) % uintptr(sectorSize))
	// 上下结果相同，位运算会快一点
	return int(uintptr(unsafe.Pointer(&buf[0])) & uintptr(sectorSize-1))
}

// NewDioBuf bufSize 必须是 sectorSize 的整数倍
func NewDioBuf(bufSize int) ([]byte, error) {
	if bufSize%sectorSize != 0 {
		panic("缓冲区的大小必须和扇区大小对齐")
	}
	buf := make([]byte, bufSize+sectorSize)
	offset := sectorSize - align(buf)
	if offset != 0 {
		buf = buf[offset : offset+bufSize]
	} else {
		buf = buf[:bufSize]
	}
	// 最终判断
	if judgeAlign(buf) {
		return buf, nil
	}
	return buf, errors.New("err")
}

func judgeAlign(buf []byte) bool {
	if len(buf)%sectorSize == 0 {
		if align(buf) == 0 {
			return true
		}
	}
	return false
}
