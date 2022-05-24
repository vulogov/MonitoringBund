package stdlib

import (
  "gonum.org/v1/gonum/stat"
  "github.com/gammazero/deque"
  tc "github.com/vulogov/ThreadComputation"
)



func bundnumModeMatrix(v interface{}) interface{} {
  switch src := v.(type) {
  case *tc.TCMatrix:
    data := BUNDMatrixToFloats(src)
    val, _ := stat.Mode(data, nil)
    return val
  }
  return nil
}

func bundnumModeNumbers(v interface{}) interface{} {
  switch src := v.(type) {
  case *tc.TCNumbers:
    val, _ := stat.Mode(src.N, nil)
    return val
  }
  return nil
}

func BUNDnumMode(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  err := l.ExecuteSingleArgumentFunction("nummode", q)
  if err != nil {
    return nil, err
  }
  return nil, nil
}

func init() {
  tc.SetFunction("num.Mode", BUNDnumMode)
  tc.RegisterFunctionCallback("nummode", tc.Matrix, bundnumModeMatrix)
  tc.RegisterFunctionCallback("nummode", tc.Numbers, bundnumModeNumbers)
}
