package stdlib

import (
  "fmt"
  "github.com/vulogov/monitoringbund/internal/conf"
  tc "github.com/vulogov/ThreadComputation"
  "github.com/pieterclaerhout/go-log"
)

func toChar(i int) rune {
  return rune('A' + i)
}

func StoreArgs() {
  if len(conf.Argv) > 0 {
    log.Debugf("[ MBUND ] storing passed arguments: %v", conf.Argv)
    for n1, x := range conf.Argv {
      for n2, y := range x {
        name := fmt.Sprintf("Arg%c%c", toChar(n1), toChar(n2))
        log.Debugf("%v = %v", name, y)
        tc.SetVariable(name, y)
      }
    }
  }
}
