package stdlib

import (
  "gonum.org/v1/gonum/stat"
  "github.com/gammazero/deque"
  tc "github.com/vulogov/ThreadComputation"
)



func bundnumMeanHarmonicMatrix(v interface{}) interface{} {
  switch src := v.(type) {
  case *tc.TCMatrix:
    data := BUNDMatrixToFloats(src)
    return stat.GeometricMean(data, nil)
  }
  return nil
}

func bundnumMeanHarmonicNumbers(v interface{}) interface{} {
  switch src := v.(type) {
  case *tc.TCNumbers:
    return stat.GeometricMean(src.N, nil)
  }
  return nil
}

func BUNDnumMeanHarmonic(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  err := l.ExecuteSingleArgumentFunction("nummeanharmonic", q)
  if err != nil {
    return nil, err
  }
  return nil, nil
}

func init() {
  tc.SetFunction("num.Mean.Harmonic", BUNDnumMeanHarmonic)
  tc.RegisterFunctionCallback("nummeanharmonic", tc.Matrix, bundnumMeanHarmonicMatrix)
  tc.RegisterFunctionCallback("nummeanharmonic", tc.Numbers, bundnumMeanHarmonicNumbers)
}
