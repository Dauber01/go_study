package error

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
)

var OriginError error = errors.New("源的error")

type QueryErrors interface {
	Error() string
}

type QueryError struct {
	Query string
	Err   error
}

func New() interface{} {
	return &QueryError{"哈哈", OriginError}
}

func (qe *QueryError) Error() string {
	return qe.Err.Error() + qe.Query
}

// 把类型和中间的包装进行分离
func TestQuery(t *testing.T) {
	qe := New()
	// 可以看到有track信息的error为github.com/pkg/errors包
	e := errors.Wrap(OriginError, "哈哈哈哈")
	fmt.Printf("strack: \n%+v\n", e)
	// 对应的类型可以进行判断
	if e, ok := qe.(*QueryError); ok && e.Err == OriginError {
		fmt.Printf("the error is: \n%+v\n", qe)
	}
	// 被Wrap包装过的可以通过is方法进行比较
	if errors.Is(e, OriginError) {
		t.Log("is可以判断出原始类型的信息")
	}
	ei := fmt.Errorf("这个也不对吧:%v", e)
	// Errorf包装之后,整体的类型发生变化,只留存文字的信息
	if errors.Is(ei, OriginError) {
		t.Log("ei也可以进行对比")
	}
	t.Logf("新建错误: %v", ei)
	t.Logf("新建错误: %T", ei)
}
