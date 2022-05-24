package stdlib

import (
  "github.com/elastic/go-sysinfo"
  tc "github.com/vulogov/ThreadComputation"
)



func init() {
  host,_ := sysinfo.Host()
  hi := host.Info()
  tc.SetVariable("time.Timezone", hi.Timezone)
}
