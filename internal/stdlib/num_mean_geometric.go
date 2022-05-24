package stdlib

import (
  "gonum.org/v1/gonum/stat"
  "github.com/gammazero/deque"
  tc "github.com/vulogov/ThreadComputation"
)



func bundnumMeanGeometricMatrix(v interface{}) interface{} {
  switch src := v.(type) {
  case *tc.TCMatrix:
    data := BUNDMatrixToFloats(src)
    return stat.GeometricMean(data, nil)
  }
  return nil
}

func bundnumMeanGeometricNumbers(v interface{}) interface{} {
  switch src := v.(type) {
  case *tc.TCNumbers:
    return stat.GeometricMean(src.N, nil)
  }
  return nil
}

func BUNDnumMeanGeometric(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  err := l.ExecuteSingleArgumentFunction("nummeangeometric", q)
  if err != nil {
    return nil, err
  }
  return nil, nil
}

func init() {
  tc.SetFunction("num.Mean.Geometric", BUNDnumMeanGeometric)
  tc.RegisterFunctionCallback("nummeangeometric", tc.Matrix, bundnumMeanGeometricMatrix)
  tc.RegisterFunctionCallback("nummeangeometric", tc.Numbers, bundnumMeanGeometricNumbers)
}
