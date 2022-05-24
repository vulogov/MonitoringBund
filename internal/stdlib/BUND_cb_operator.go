package stdlib

import (
  "fmt"
  "github.com/pieterclaerhout/go-log"
  tc "github.com/vulogov/ThreadComputation"
)


type TCContextOperatorFunction func(*tc.TCExecListener, interface{}, interface{}) interface{}


func GetContextOperatorCallback(name string, x interface{}, y interface{}) TCContextOperatorFunction {
  x_type := tc.TCType(x)
  y_type := tc.TCType(y)
  fname := fmt.Sprintf("op.%v.%v.%v", name, x_type, y_type)
  log.Debugf("Looking for context operator for %v(%T %T)=%v", name, x, y, fname)
  if fun, ok := Callbacks.Load(fname); ok {
    log.Debugf("Got: %v", fname)
    return fun.(TCContextOperatorFunction)
  }
  fname = fmt.Sprintf("op.%v.%v.%v", name, x_type, tc.Any)
  log.Debugf("Looking for context operator for %v(%T Any)", name, x)
  if fun, ok := Callbacks.Load(fname); ok {
    log.Debugf("Got: %v", fname)
    return fun.(TCContextOperatorFunction)
  }
  log.Debugf("No context operator for: %v(%T %T)", name, x, y)
  return nil
}

func RegisterContextOperatorCallback(name string, x_type int, y_type int, fun TCContextOperatorFunction) {
  fname := fmt.Sprintf("op.%v.%v.%v", name, x_type, y_type)
  Callbacks.Delete(fname)
  Callbacks.Store(fname, fun)
}
