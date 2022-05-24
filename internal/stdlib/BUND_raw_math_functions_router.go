package stdlib

import (
  "github.com/pieterclaerhout/go-log"
  "github.com/gammazero/deque"
  tc "github.com/vulogov/ThreadComputation"
)

func BUNDRawMathFunctionsRouter(l *tc.TCExecListener, name string, q *deque.Deque) (bool, interface{}, error) {
  q = tc.SingleValueFromStackAndAttr(l, q)
  log.Debugf("raw-math-function-router: %v len=%v", name, q.Len())
  err := ExecuteOperator(l, "rawmath",name, q)
  if err == nil {
    log.Debugf("raw-math function served: %v", name)
    return true, nil, nil
  }
  log.Debug("raw-math error: %v", err)
  return false, nil, nil
}


func init() {
  tc.TCAddRouterFunction(BUNDRawMathFunctionsRouter)
}
