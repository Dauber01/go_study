package current_test

type Conn struct {
}

func (c *Conn) DoQuery(query string) Result {
	return Result{data: "lilei"}
}

type Result struct {
	data string
}

// 用来在可并行的任务获取第一个有效返回值的demo
func Query(conns []Conn, query string) Result {

	// 建立一个阻塞的通道
	ch := make(chan Result)
	for _, conn := range conns {
		go func(c Conn) {
			select {
			// 在select中进行阻塞,当有结果值返回的时候写入通道
			case ch <- c.DoQuery(query):
			default:
			}
		}(conn)
	}
	// revice先于send发生,当有一个执行成功的时候,会进行返回,未执行完的会直接走default结束
	return <-ch
}
