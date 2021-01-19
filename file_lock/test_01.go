package file_lock

import (
	"log"
	"os"
	"syscall"
)

type Test01 struct {
}

func (t *Test01) Run() {
	t.foo()
}

func (t *Test01) foo() {
	// 打开文件锁
	lock, err := os.Create("./lock")
	if err != nil {
		log.Printf("main: 创建文件锁失败: %s", err)
		os.Exit(1)
	}
	defer os.Remove("./lock")
	defer lock.Close()

	// 尝试独占文件锁
	err = syscall.Flock(int(lock.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	if err != nil {
		log.Printf("main: 独占文件锁失败: %s", err)
		os.Exit(1)
	}
	defer syscall.Flock(int(lock.Fd()), syscall.LOCK_UN)
}
