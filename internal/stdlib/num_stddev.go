package stdlib

import (
  "gonum.org/v1/gonum/stat"
  "github.com/gammazero/deque"
  tc "github.com/vulogov/ThreadComputation"
)



func bundnumStdDevMatrix(v interface{}) interface{} {
  switch src := v.(type) {
  case *tc.TCMatrix:
    data := BUNDMatrixToFloats(src)
    val := stat.StdDev(data, nil)
    return val
  }
  return nil
}

func bundnumStdDevNumbers(v interface{}) interface{} {
  switch src := v.(type) {
  case *tc.TCNumbers:
    val := stat.StdDev(src.N, nil)
    return val
  }
  return nil
}

func BUNDnumStdDev(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  err := l.ExecuteSingleArgumentFunction("numstddev", q)
  if err != nil {
    return nil, err
  }
  return nil, nil
}

func init() {
  tc.SetFunction("num.StdDev", BUNDnumStdDev)
  tc.RegisterFunctionCallback("numstddev", tc.Matrix, bundnumStdDevMatrix)
  tc.RegisterFunctionCallback("numstddev", tc.Numbers, bundnumStdDevNumbers)
}
