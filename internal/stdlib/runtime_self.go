package stdlib

import (
  "github.com/gammazero/deque"
  "github.com/elastic/go-sysinfo"
  tc "github.com/vulogov/ThreadComputation"
)

func BUNDruntimeSelfMemory(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  self, _ := sysinfo.Self()
  mi, _ := self.Memory()
  res := tc.MakeDict()
  res.D.Set("resident", mi.Resident)
  res.D.Set("virtual", mi.Virtual)
  for k, v := range mi.Metrics {
    res.D.Set(k, v)
  }
  return res, nil
}

func BUNDruntimeSelfUser(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  self, _ := sysinfo.Self()
  user, _ := self.User()
  res := tc.MakeDict()
  res.D.Set("uid", user.UID)
  res.D.Set("euid", user.EUID)
  res.D.Set("suid", user.SUID)
  res.D.Set("gid", user.GID)
  return res, nil
}

func init() {
  self, _ := sysinfo.Self()
  info, _ := self.Info()
  tc.SetVariable("runtime.Self.Name", info.Name)
  tc.SetVariable("runtime.Self.PID", info.PID)
  tc.SetVariable("runtime.Self.PPID", info.PPID)
  tc.SetVariable("runtime.Self.Exe", info.Exe)
  tc.SetCommand("runtime.Self.Memory", BUNDruntimeSelfMemory)
  tc.SetCommand("runtime.Self.User", BUNDruntimeSelfUser)
}
