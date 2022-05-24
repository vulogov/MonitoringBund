package stdlib

import (
  "gonum.org/v1/gonum/stat"
  "github.com/gammazero/deque"
  tc "github.com/vulogov/ThreadComputation"
)



func bundnumEntropyMatrix(v interface{}) interface{} {
  switch src := v.(type) {
  case *tc.TCMatrix:
    data := BUNDMatrixToFloats(src)
    return stat.Entropy(data)
  }
  return nil
}

func bundnumEntropyNumbers(v interface{}) interface{} {
  switch src := v.(type) {
  case *tc.TCNumbers:
    return stat.Entropy(src.N)
  }
  return nil
}

func BUNDnumEntropy(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  err := l.ExecuteSingleArgumentFunction("numentropy", q)
  if err != nil {
    return nil, err
  }
  return nil, nil
}

func init() {
  tc.SetFunction("num.Entropy", BUNDnumEntropy)
  tc.RegisterFunctionCallback("numentropy", tc.Matrix, bundnumEntropyMatrix)
  tc.RegisterFunctionCallback("numentropy", tc.Numbers, bundnumEntropyNumbers)
}
