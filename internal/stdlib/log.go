package stdlib

import (
  "github.com/pieterclaerhout/go-log"
  "github.com/gammazero/deque"
  tc "github.com/vulogov/ThreadComputation"
)



func BUNDLogDebug(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  msg := StdlibFormat(q)
  log.Debug(msg)
  return msg, nil
}

func BUNDLogInfo(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  msg := StdlibFormat(q)
  log.Info(msg)
  return msg, nil
}

func BUNDLogWarning(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  msg := StdlibFormat(q)
  log.Warn(msg)
  return msg, nil
}

func BUNDLogError(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  msg := StdlibFormat(q)
  log.Error(msg)
  return msg, nil
}

func BUNDLogFatal(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  msg := StdlibFormat(q)
  log.Fatal(msg)
  return msg, nil
}

func init() {
  tc.SetCommand("log.Debug", BUNDLogDebug)
  tc.SetCommand("log.Info", BUNDLogInfo)
  tc.SetCommand("log.Warning", BUNDLogWarning)
  tc.SetCommand("log.Error", BUNDLogError)
  tc.SetCommand("log.Fatal", BUNDLogFatal)

}
