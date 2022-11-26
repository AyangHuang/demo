package dio

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDioBuf(t *testing.T) {
	_, err := NewDioBuf(4096)
	if err != nil {
		_ = fmt.Errorf("%s", err)
	}
}

// 必须在 linux 环境下测试
func TestDIO(t *testing.T) {
	buf, _ := NewDioBuf(512)
	// syscall.O_DIRECT = 0x4000
	// 第三个参数是权限模式 permission mode
	fp, _ := os.OpenFile("./test.txt", 0x4000|os.O_RDWR, 666)
	defer fp.Close()
	fp.Read(buf)
	fmt.Printf("%s", buf)
	buf[0] = 'O'
	buf[1] = 'I'
	fp.WriteAt(buf, 0)
	fp.ReadAt(buf, 0)
	fmt.Print(" \n 更改后 \n")
	fmt.Printf("%s", buf)
	fp.Read(buf)
	fmt.Print(" \n 下一个 512B \n")
	fmt.Printf("%s", buf)
}
