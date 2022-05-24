package stdlib

import (
  "github.com/gammazero/deque"
  "github.com/elastic/go-sysinfo"
  tc "github.com/vulogov/ThreadComputation"
)

func BUNDruntimeMAC(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  host,_ := sysinfo.Host()
  hi := host.Info()
  res := new(tc.TCList)
  for _, i := range hi.MACs {
    res.Add(i)
  }
  return res, nil
}


func init() {
  tc.SetCommand("runtime.MAC", BUNDruntimeMAC)
}
