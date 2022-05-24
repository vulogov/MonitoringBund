package stdlib

import (
  "github.com/lrita/cmap"
  "github.com/pieterclaerhout/go-log"
  tc "github.com/vulogov/ThreadComputation"
)

var Callbacks cmap.Cmap

type BUNDEnv struct {
  TC        *tc.TCstate
}

func InitBUND() *BUNDEnv {
  res := new(BUNDEnv)
  res.TC = tc.Init()
  return res
}

func (env *BUNDEnv) Eval(code interface{}) *BUNDEnv {
  switch toRun := code.(type) {
  case string:
    env.TC = env.TC.Eval(toRun)
  default:
    log.Errorf("Error while evaluating: %T", code)
    env.TC.SetError("Invalid input for Eval()")
  }
  return env
}
