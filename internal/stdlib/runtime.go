package stdlib

import (
  "github.com/elastic/go-sysinfo"
  tc "github.com/vulogov/ThreadComputation"
)

func init() {
  rt := sysinfo.Go()
  host,_ := sysinfo.Host()
  hi := host.Info()
  tc.SetVariable("runtime.OS", rt.OS)
  tc.SetVariable("runtime.Arch", rt.Arch)
  tc.SetVariable("runtime.MaxProcs", rt.MaxProcs)
  tc.SetVariable("runtime.Version", rt.Version)
  tc.SetVariable("runtime.HostArchitecture", hi.Architecture)
  tc.SetVariable("system.Hostname", hi.Hostname)
  cnt := hi.Containerized
  if cnt != nil {
    tc.SetVariable("runtime.Containerized", cnt)
  } else {
    tc.SetVariable("runtime.Containerized", false)
  }
  tc.SetVariable("os.Name", hi.OS.Name)
  tc.SetVariable("os.Type", hi.OS.Type)
  tc.SetVariable("os.Family", hi.OS.Family)
  tc.SetVariable("os.Platform", hi.OS.Platform)
  tc.SetVariable("os.Version", hi.OS.Version)
  tc.SetVariable("os.Build", hi.OS.Build)
  tc.SetVariable("os.Codename", hi.OS.Codename)
}
