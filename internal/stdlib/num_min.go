package stdlib

import (
  floats "gonum.org/v1/gonum/floats"
  mat "gonum.org/v1/gonum/mat"
  "github.com/gammazero/deque"
  tc "github.com/vulogov/ThreadComputation"
)



func bundnumMinMatrix(v interface{}) interface{} {
  switch src := v.(type) {
  case *tc.TCMatrix:
    return mat.Min(src.M)
  }
  return nil
}

func bundnumMinNumbers(v interface{}) interface{} {
  switch src := v.(type) {
  case *tc.TCNumbers:
    return floats.Min(src.N)
  }
  return nil
}

func BUNDnumMin(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  err := l.ExecuteSingleArgumentFunction("nummin", q)
  if err != nil {
    return nil, err
  }
  return nil, nil
}

func init() {
  tc.SetFunction("num.Min", BUNDnumMin)
  tc.RegisterFunctionCallback("nummin", tc.Matrix, bundnumMinMatrix)
  tc.RegisterFunctionCallback("nummin", tc.Numbers, bundnumMinNumbers)
}
