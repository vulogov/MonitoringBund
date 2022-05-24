package stdlib

import (
  "fmt"
  "github.com/pieterclaerhout/go-log"
  "github.com/gammazero/deque"
  tc "github.com/vulogov/ThreadComputation"
)


func ExecuteOperator(l *tc.TCExecListener, name string, x interface{}, q *deque.Deque) *tc.TCError {
  var fun TCContextOperatorFunction

  log.Debugf("context-operator name=%v x=%v q=%v", name, x, q.Len())
  for q.Len() > 0 {
    e := q.PopFront()
    switch e.(type) {
    case *tc.TCValue:
      fun = GetContextOperatorCallback(name, x, e.(*tc.TCValue).Value)
    default:
      fun = GetContextOperatorCallback(name, x, e)
    }
    if fun == nil {
      return l.TC.MakeError(fmt.Sprintf("callback for %v(%T %T) not found", name, x, e))
    }
    log.Debugf("context operator: %v %v(%v)", name, x, e)
    res := fun(l, x, e)
    if res == nil {
      //
      // Yes, we do not care if execute function returned nothing
      //
      // return l.TC.MakeError(fmt.Sprintf("callback for %v(%T) returned nil", name, e))
      return nil
    }
    log.Debugf("operator %v(%T %T) = %v", name, x, e, res)
    switch e.(type) {
    case *tc.TCValue:
      tc.ReturnFromFunction(l, name, tc.MakeValue(res, e.(*tc.TCValue).P, e.(*tc.TCValue).TTL))
    default:
      tc.ReturnFromFunction(l, name, res)
    }
  }
  return nil
}
