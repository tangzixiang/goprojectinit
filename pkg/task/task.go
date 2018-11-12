package task

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

// 任务
type Handle func() error

// 任务下载 Hub
type T struct {
	m   map[string]Handle
	err error
}

// 任务列表
var (
	once  sync.Once
	Tasks T
)

// HandleFuc 注册任务
func (t *T) HandleFuc(name string, handle Handle) *T {
	once.Do(func() {
		t.m = make(map[string]Handle)
	})

	if name = strings.TrimSpace(name); name == "" {
		t.err = errors.New("handle func failed : need argument name")
		return t
	}

	if handle == nil {
		t.err = errors.New(
			fmt.Sprintf("handle func for %v failed : need argument handle", name))
		return t
	}

	t.m[name] = handle

	return t
}

// Call 调用任务
func (t *T) Call(name string) *T {

	if t.err != nil {
		return t
	}

	if name = strings.TrimSpace(name); name == "" {
		t.err = errors.New("call func failed : need argument name")
		return t
	}

	handle, exists := t.m[name]
	if !exists || handle == nil {
		t.err = errors.New(
			fmt.Sprintf("call func for %v failed : no such handle", name))
		return t
	}

	t.err = handle()
	return t
}

// 获取错误
func (t *T) Err() error {
	return t.err
}
