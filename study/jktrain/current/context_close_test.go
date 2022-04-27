package current_test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

const shortDuration = 1 * time.Millisecond

// 不知道为啥,虽然是done方法返回,但一直有上下文超过deadline的错误
func TestContextClose(t *testing.T) {
	d := time.Now().Add(shortDuration)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// for i := 0; i < 10; i++ {
	// 	t.Log("哈哈哈哈")
	// }

	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		fmt.Println("over")
	}
}
