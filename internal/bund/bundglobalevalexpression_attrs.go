package bund

import (
	tc "github.com/vulogov/ThreadComputation"
	"github.com/gammazero/deque"
)

func AttrsToQueue(attrs []string) *deque.Deque {
	q := deque.New()
	for _, v := range(attrs) {
		res := tc.GetSimpleData(v)
		if res == nil {
			q.PushBack(v)
		} else {
			q.PushBack(res)
		}
	}
	return q
}
