package stdlib

import (
  "fmt"
  "github.com/pieterclaerhout/go-log"
  tc "github.com/vulogov/ThreadComputation"
)


type TCContextGenericFunction func(*tc.TCExecListener, interface{}) interface{}


func GetContextFunctionCallback(name string, x interface{}) TCContextGenericFunction {
  x_type := tc.TCType(x)
  fname := fmt.Sprintf("fun.%v.%v", name, x_type)
  log.Debugf("Looking for function for %v(%T)", name, x)
  if fun, ok := Callbacks.Load(fname); ok {
    log.Debugf("Got: %v", fname)
    return fun.(TCContextGenericFunction)
  }
  fname = fmt.Sprintf("fun.%v.%v", name, tc.Any)
  log.Debugf("Looking for function for %v(Any)", name)
  if fun, ok := Callbacks.Load(fname); ok {
    log.Debugf("Got: %v", fname)
    return fun.(TCContextGenericFunction)
  }
  log.Debugf("No function for: %v(%T)", name, x)
  return nil
}

func RegisterContextFunctionCallback(name string, x_type int, fun TCContextGenericFunction) {
  fname := fmt.Sprintf("fun.%v.%v", name, x_type)
  Callbacks.Delete(fname)
  Callbacks.Store(fname, fun)
}
