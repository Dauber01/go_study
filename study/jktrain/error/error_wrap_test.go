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

func TestParse(t *testing.T) {
	if 3 > 2 {
		// 可以使用 errors.Errorf 和 errors.New 来返回错误
		// 如果调用其它包的函数, 可以直接返回error
		fmt.Printf("%T", errors.Errorf("the demo error"))
	}
}

var err error = errors.New("new error")

func TestWrapf(t *testing.T) {
	if 3 > 2 {
		// 在包装类型除了可以使用 wrap 之外也可以使用 wrapf 来对包装的信息进行格式化
		fmt.Printf("%v", errors.Wrapf(err, "the error is %q", "hah"))
	}
}

func wrapError() error {
	if 3 > 2 {
		return errors.Wrap(err, "the message is wrong")
	}
	return nil
}

func TestCause(t *testing.T) {
	erro := wrapError()
	if errors.Cause(erro) == err {
		fmt.Println("可以通过cause方法对不出")
	}
}
