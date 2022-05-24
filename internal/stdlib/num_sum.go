package stdlib

import (
  floats "gonum.org/v1/gonum/floats"
  mat "gonum.org/v1/gonum/mat"
  "github.com/gammazero/deque"
  tc "github.com/vulogov/ThreadComputation"
)



func bundnumSumMatrix(v interface{}) interface{} {
  switch src := v.(type) {
  case *tc.TCMatrix:
    return mat.Sum(src.M)
  }
  return nil
}

func bundnumSumNumbers(v interface{}) interface{} {
  switch src := v.(type) {
  case *tc.TCNumbers:
    return floats.SumCompensated(src.N)
  }
  return nil
}

func BUNDnumSum(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  err := l.ExecuteSingleArgumentFunction("numsum", q)
  if err != nil {
    return nil, err
  }
  return nil, nil
}

func init() {
  tc.SetFunction("num.Sum", BUNDnumSum)
  tc.RegisterFunctionCallback("numsum", tc.Matrix, bundnumSumMatrix)
  tc.RegisterFunctionCallback("numsum", tc.Numbers, bundnumSumNumbers)
}
