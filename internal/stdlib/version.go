package stdlib

import (
  "github.com/vulogov/monitoringbund/internal/conf"
  "github.com/gammazero/deque"
  tc "github.com/vulogov/ThreadComputation"
)

var VERSION = conf.EVersion

func BUNDVersionCore(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  return tc.VERSION, nil
}

func BUNDVersionBund(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  return VERSION, nil
}

func init() {
  tc.SetCommand("version.Core", BUNDVersionCore)
  tc.SetCommand("version.Bund", BUNDVersionBund)
}
