package stdlib

import (
  floats "gonum.org/v1/gonum/floats"
  mat "gonum.org/v1/gonum/mat"
  "github.com/gammazero/deque"
  tc "github.com/vulogov/ThreadComputation"
)



func bundnumMaxMatrix(v interface{}) interface{} {
  switch src := v.(type) {
  case *tc.TCMatrix:
    return mat.Max(src.M)
  }
  return nil
}

func bundnumMaxNumbers(v interface{}) interface{} {
  switch src := v.(type) {
  case *tc.TCNumbers:
    return floats.Max(src.N)
  }
  return nil
}

func BUNDnumMax(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  err := l.ExecuteSingleArgumentFunction("nummax", q)
  if err != nil {
    return nil, err
  }
  return nil, nil
}

func init() {
  tc.SetFunction("num.Max", BUNDnumMax)
  tc.RegisterFunctionCallback("nummax", tc.Matrix, bundnumMaxMatrix)
  tc.RegisterFunctionCallback("nummax", tc.Numbers, bundnumMaxNumbers)
}
