package stdlib

import (
  "gonum.org/v1/gonum/stat"
  "github.com/gammazero/deque"
  tc "github.com/vulogov/ThreadComputation"
)



func bundnumSkewMatrix(v interface{}) interface{} {
  switch src := v.(type) {
  case *tc.TCMatrix:
    data := BUNDMatrixToFloats(src)
    val := stat.Skew(data, nil)
    return val
  }
  return nil
}

func bundnumSkewNumbers(v interface{}) interface{} {
  switch src := v.(type) {
  case *tc.TCNumbers:
    val := stat.Skew(src.N, nil)
    return val
  }
  return nil
}

func BUNDnumSkew(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  err := l.ExecuteSingleArgumentFunction("numskew", q)
  if err != nil {
    return nil, err
  }
  return nil, nil
}

func init() {
  tc.SetFunction("num.Skew", BUNDnumMode)
  tc.RegisterFunctionCallback("numskew", tc.Matrix, bundnumSkewMatrix)
  tc.RegisterFunctionCallback("numskew", tc.Numbers, bundnumSkewNumbers)
}
