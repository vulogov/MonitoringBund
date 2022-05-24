package stdlib

import (
  "gonum.org/v1/gonum/stat"
  "github.com/gammazero/deque"
  tc "github.com/vulogov/ThreadComputation"
)



func bundnumCovarianceMatrix(x interface{}, y interface{}) interface{} {
  switch rx := x.(type) {
  case *tc.TCMatrix:
    switch ry := y.(type) {
    case *tc.TCMatrix:
      m1, n1 := rx.M.Dims()
      m2, n2 := ry.M.Dims()
      if m1 == m2 && n1 == n2 {
        src1 := BUNDMatrixToFloats(rx)
        src2 := BUNDMatrixToFloats(ry)
        res := stat.Covariance(src1, src2, nil)
        return res
      }
    }
  }
  return tc.MakeNone()
}

func bundnumCovarianceNumbers(x interface{}, y interface{}) interface{} {
  switch rx := x.(type) {
  case *tc.TCNumbers:
    switch ry := y.(type) {
    case *tc.TCNumbers:
      if rx.Len() == ry.Len() {
        res := stat.Covariance(rx.N, ry.N, nil)
        return res
      }
    }
  }
  return tc.MakeNone()
}

func BUNDnumCovariance(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  err := l.ExecuteOperator("numcovariance", q)
  if err != nil {
    return nil, err
  }
  return nil, nil
}

func init() {
  tc.SetFunction("num.Covariance", BUNDnumCovariance)
  tc.RegisterOperatorCallback("numcovariance", tc.Matrix, tc.Matrix, bundnumCovarianceMatrix)
  tc.RegisterOperatorCallback("numcovariance", tc.Numbers, tc.Numbers, bundnumCovarianceNumbers)
}
