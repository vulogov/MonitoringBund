package stdlib

import (
  "github.com/gammazero/deque"
  tc "github.com/vulogov/ThreadComputation"
)

func BUNDnnNormalize(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  err := l.ExecuteSingleArgumentFunction("mlnormalize", q)
  if err != nil {
    return nil, err
  }
  return nil, nil
}


func init() {
  tc.SetFunction("ml.Normalize", BUNDnnNormalize)
}
