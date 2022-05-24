package stdlib

import (
  "fmt"
  "github.com/pieterclaerhout/go-log"
  "github.com/gammazero/deque"
  tc "github.com/vulogov/ThreadComputation"
)


func ExecuteSingleArgumentFunction(l *tc.TCExecListener, name string, q *deque.Deque) *tc.TCError {
  var fun TCContextGenericFunction

  for q.Len() > 0 {
    e := q.PopFront()
    switch e.(type) {
    case *tc.TCValue:
      fun = GetContextFunctionCallback(name, e.(*tc.TCValue).Value)
    default:
      fun = GetContextFunctionCallback(name, e)
    }
    if fun == nil {
      return l.TC.MakeError(fmt.Sprintf("callback for %v(%T) not found", name, e))
    }
    res := fun(l, e)
    if res == nil {
      //
      // Yes, we do not care if execute function returned nothing
      //
      // return l.TC.MakeError(fmt.Sprintf("callback for %v(%T) returned nil", name, e))
      return nil
    }
    log.Debugf("function %v(%T) = %v", name, e, res)
    switch e.(type) {
    case *tc.TCValue:
      tc.ReturnFromFunction(l, name, tc.MakeValue(res, e.(*tc.TCValue).P, e.(*tc.TCValue).TTL))
    default:
      tc.ReturnFromFunction(l, name, res)
    }
  }
  return nil
}
