package stdlib

import (
  "github.com/pieterclaerhout/go-log"
  "github.com/gammazero/deque"
  tc "github.com/vulogov/ThreadComputation"
)


func BUNDApplyFunction(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  val := q.PopFront()
  if q.Len() == 0 {
    log.Debugf("apply[] received no parameters")
    return val, nil
  }
  return nil, nil
}


func init() {
  tc.SetFunction("apply", BUNDApplyFunction)
}
