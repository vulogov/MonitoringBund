package stdlib

import (
  "gonum.org/v1/gonum/stat"
  "github.com/gammazero/deque"
  tc "github.com/vulogov/ThreadComputation"
)



func bundnumCorrelationKendallMatrix(x interface{}, y interface{}) interface{} {
  switch rx := x.(type) {
  case *tc.TCMatrix:
    switch ry := y.(type) {
    case *tc.TCMatrix:
      m1, n1 := rx.M.Dims()
      m2, n2 := ry.M.Dims()
      if m1 == m2 && n1 == n2 {
        src1 := BUNDMatrixToFloats(rx)
        src2 := BUNDMatrixToFloats(ry)
        res := stat.Kendall(src1, src2, nil)
        return res
      }
    }
  }
  return tc.MakeNone()
}

func bundnumCorrelationKendallNumbers(x interface{}, y interface{}) interface{} {
  switch rx := x.(type) {
  case *tc.TCNumbers:
    switch ry := y.(type) {
    case *tc.TCNumbers:
      if rx.Len() == ry.Len() {
        res := stat.Kendall(rx.N, ry.N, nil)
        return res
      }
    }
  }
  return tc.MakeNone()
}

func BUNDnumCorrelationKendall(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  err := l.ExecuteOperator("numcorrelationkendall", q)
  if err != nil {
    return nil, err
  }
  return nil, nil
}

func init() {
  tc.SetFunction("num.Correlation.Kendall", BUNDnumCorrelation)
  tc.RegisterOperatorCallback("numcorrelationkendall", tc.Matrix, tc.Matrix, bundnumCorrelationMatrix)
  tc.RegisterOperatorCallback("numcorrelationkendall", tc.Numbers, tc.Numbers, bundnumCorrelationNumbers)
}
