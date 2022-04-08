package error_test

import (
	"errors"
	"fmt"
	"testing"
)

var Wuhuiexce error = errors.New("哈哈哈")

func TestError(t *testing.T) {
	err := origin()
	if err != nil {
		// 可以看到仍然无法打出错误的堆栈信息
		fmt.Printf("%+v", err)
	}
	ans := errors.Is(Wuhuiexce, err)
	fmt.Printf("%v", ans)
	if errors.Is(err, Wuhuiexce) {
		fmt.Println("可以的")
	}
	//err = Wuhuiexce
	if err == Wuhuiexce {
		fmt.Println("基本可以")
	}
}

func origin() error {
	return fmt.Errorf("这是误会, %+v", caseDemo())
}

func caseDemo() error {

	return fmt.Errorf("别开枪, %+v", Wuhuiexce)
}
