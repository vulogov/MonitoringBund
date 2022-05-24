package stdlib

import (
  "gonum.org/v1/gonum/stat"
  "github.com/gammazero/deque"
  tc "github.com/vulogov/ThreadComputation"
)



func bundnumStdErrMatrix(v interface{}) interface{} {
  switch src := v.(type) {
  case *tc.TCMatrix:
    data := BUNDMatrixToFloats(src)
    sd := stat.StdDev(data, nil)
    return stat.StdErr(sd, float64(len(data)))
  }
  return nil
}

func bundnumStdErrNumbers(v interface{}) interface{} {
  switch src := v.(type) {
  case *tc.TCNumbers:
    sd := stat.StdDev(src.N, nil)
    return stat.StdErr(sd, float64(src.Len()))
  }
  return nil
}

func BUNDnumStdErr(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  err := l.ExecuteSingleArgumentFunction("numstderr", q)
  if err != nil {
    return nil, err
  }
  return nil, nil
}

func init() {
  tc.SetFunction("num.StdErr", BUNDnumStdErr)
  tc.RegisterFunctionCallback("numstderr", tc.Matrix, bundnumStdErrMatrix)
  tc.RegisterFunctionCallback("numstderr", tc.Numbers, bundnumStdErrNumbers)
}
