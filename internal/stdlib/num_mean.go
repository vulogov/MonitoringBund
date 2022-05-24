package stdlib

import (
  "gonum.org/v1/gonum/stat"
  "github.com/gammazero/deque"
  tc "github.com/vulogov/ThreadComputation"
)



func bundnumMeanMatrix(v interface{}) interface{} {
  switch src := v.(type) {
  case *tc.TCMatrix:
    data := BUNDMatrixToFloats(src)
    return stat.Mean(data, nil)
  }
  return nil
}

func bundnumMeanNumbers(v interface{}) interface{} {
  switch src := v.(type) {
  case *tc.TCNumbers:
    return stat.Mean(src.N, nil)
  }
  return nil
}

func BUNDnumMean(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  err := l.ExecuteSingleArgumentFunction("nummean", q)
  if err != nil {
    return nil, err
  }
  return nil, nil
}

func init() {
  tc.SetFunction("num.Mean", BUNDnumMean)
  tc.RegisterFunctionCallback("nummean", tc.Matrix, bundnumMeanMatrix)
  tc.RegisterFunctionCallback("nummean", tc.Numbers, bundnumMeanNumbers)
}
