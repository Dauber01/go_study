package current_test

import (
	"testing"
	//"golang.org/x/sync/errgroup"
)

func TestErrGroup(t *testing.T) {
	/** g, ctx := errgroup.WithContext(context.Background())
	var a, b, c []int
	//调用广告
	g.Go(func() error {
		a = append(a, 1)
		return errors.New("test")
	})
	//调用AI
	g.Go(func() error {
		c = append(c, 1)
		return nil
	})
	//调用运营平台
	g.Go(func() error {
		b = append(b, 1)
		return nil
	})
	err := g.Wait()
	fmt.Println(err)
	fmt.Println(ctx.Err()) */
}
