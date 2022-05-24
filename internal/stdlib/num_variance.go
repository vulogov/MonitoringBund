package stdlib

import (
  "gonum.org/v1/gonum/stat"
  "github.com/gammazero/deque"
  tc "github.com/vulogov/ThreadComputation"
)



func bundnumVarianceMatrix(v interface{}) interface{} {
  switch src := v.(type) {
  case *tc.TCMatrix:
    data := BUNDMatrixToFloats(src)
    val := stat.Variance(data, nil)
    return val
  }
  return nil
}

func bundnumVarianceNumbers(v interface{}) interface{} {
  switch src := v.(type) {
  case *tc.TCNumbers:
    val := stat.Variance(src.N, nil)
    return val
  }
  return nil
}

func BUNDnumVariance(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  err := l.ExecuteSingleArgumentFunction("numvariance", q)
  if err != nil {
    return nil, err
  }
  return nil, nil
}

func init() {
  tc.SetFunction("num.Variance", BUNDnumVariance)
  tc.RegisterFunctionCallback("numvariance", tc.Matrix, bundnumVarianceMatrix)
  tc.RegisterFunctionCallback("numvariance", tc.Numbers, bundnumVarianceNumbers)
}
