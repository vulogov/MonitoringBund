package stdlib

import (
  "fmt"
  "github.com/pieterclaerhout/go-log"
  "github.com/gammazero/deque"
  tc "github.com/vulogov/ThreadComputation"
)

//
// There is a limit on maximum iterations. By default it is 0, means 'unlimited'
//
const MaxSize = 0

func bundGeneratorMakeIterator(l *tc.TCExecListener, i interface{}) interface{} {
  log.Debugf("generator: Attempting to make iterator for %T", i)
  igfun := tc.GetGenCallback(i)
  if igfun == nil {
    l.TC.MakeError(fmt.Sprintf("iter: failed to get generator for %T", i))
    return nil
  }
  iter := igfun(i)
  if iter == nil {
    l.TC.MakeError(fmt.Sprintf("iter: failed to create iterator for %T", i))
    return nil
  }
  iter.SetTC(l.TC)
  return bundGeneratorIterator(l, iter)
}

func bundGeneratorIterator(l *tc.TCExecListener, i interface{}) interface{} {
  var out interface{}
  log.Debugf("generator: got %T as iterator", i)
  switch iter := i.(type) {
  case *tc.TCIterator:
    maxsize := l.TC.FromContext("maxsize", int64(MaxSize)).(int64)
    log.Debugf("generator: maxsize is %v", maxsize)
    out_type := l.TC.FromContext("type", "numbers")
    switch out_type {
    case "numbers":
      out = tc.MakeNumbers()
    default:
      l.TC.MakeError(fmt.Sprintf("generate: unknown output type for generate[]: %v", out_type))
      return nil
    }
    c := int64(0)
    out:
    for {
      v := iter.Next()
      if v == nil {
        break out
      }
      log.Debugf("count=%v, maxsize=%v", c, maxsize)
      if maxsize != 0 && c > (maxsize-1) {
        break out
      }
      switch v.(type) {
      case *tc.TCBreak:
        break out
      case *tc.TCNone:
        break out
      case *tc.TCValue:
        switch v.(*tc.TCValue).Value.(type) {
        case *tc.TCNone:
          break out
        }
      }
      c += 1
      switch out.(type) {
      case *tc.TCNumbers:
        out.(*tc.TCNumbers).Add(v)
      default:
        l.TC.MakeError(fmt.Sprintf("generate: unknown output type for generate[]: %v", out_type))
        return nil
      }
    }
    return out
  }
  log.Debugf("Invalid context for generate[]")
  return nil
}

func BUNDGenerateFunction(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  err := ExecuteSingleArgumentFunction(l, "generator", q)
  if err != nil {
    return nil, err
  }
  return nil, nil
}


func init() {
  tc.SetFunction("generate", BUNDGenerateFunction)
  RegisterContextFunctionCallback("generator", tc.Iterator, bundGeneratorIterator)
  RegisterContextFunctionCallback("generator", tc.Any, bundGeneratorMakeIterator)
}
