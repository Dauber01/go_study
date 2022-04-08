package error_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/pkg/errors"
)

func ReadFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		// 包装对应的信息
		return nil, errors.Wrap(err, "open failed")
	}
	defer f.Close()
	return nil, nil
}

func ReadConfig() ([]byte, error) {
	home := os.Getenv("HOME")
	config, err := ReadFile(filepath.Join(home, ".settings.xml"))
	return config, errors.WithMessage(err, "could not read cinfig")
}

func TestReadFile(t *testing.T) {
	_, err := ReadConfig()
	if err != nil {
		// 打印出error的原始类型, 和对应的信息
		fmt.Printf("original error: %T %v\n", errors.Cause(err), errors.Cause(err))
		// 打印出对应的堆栈信息
		fmt.Printf("stack trace: \n%+v\n", err)
		os.Exit(1)
	}
}

func TestOther(t *testing.T) {
	_, err := ReadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
