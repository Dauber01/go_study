package answerques_test

import (
	"sync"
)

type safeResource struct {
	// 当有公共资源需要进行暴露的时候,要用行为去限制使用者,而不是用注释提醒,
	// 指望使用者去看到或者遵守
	resource interface{}
	lock     sync.Mutex
}

func (s *safeResource) DoSomethingTpResource() {
	s.lock.Lock()
	defer s.lock.Unlock()
	// 处理对应的业务逻辑
}

// 这个在处理的时候没有采用任何同步的操作,但是对于和用户之间有默契的情况下
// 比如所有的调用都在初始化的过程中由单个的gorotine来处理,就不需要考虑线程安全的问题
// 或者该结构在一个时间点前,只写,而在一个时间点后只读,也算半个安全
type Registry struct {
	resource map[string]interface{}
}

func (r *Registry) Register(name string, resource interface{}) {
	r.resource[name] = resource
}

func (r *Registry) Get(name string) (interface{}, error) {
	return nil, nil
}
