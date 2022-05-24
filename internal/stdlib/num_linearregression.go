package stdlib

import (
  "gonum.org/v1/gonum/stat"
  "github.com/gammazero/deque"
  tc "github.com/vulogov/ThreadComputation"
)



func bundnumLinearRegressionMatrix(x interface{}, y interface{}) interface{} {
  switch rx := x.(type) {
  case *tc.TCMatrix:
    switch ry := y.(type) {
    case *tc.TCMatrix:
      m1, n1 := rx.M.Dims()
      m2, n2 := ry.M.Dims()
      if m1 == m2 && n1 == n2 {
        res := tc.MakeDict()
        src1 := BUNDMatrixToFloats(rx)
        src2 := BUNDMatrixToFloats(ry)
        alpha, beta := stat.LinearRegression(src1, src2, nil, false)
        r2 := stat.RSquared(src1, src2, nil, alpha, beta)
        res.D.Set("alpha", alpha)
        res.D.Set("beta", beta)
        res.D.Set("rsquare", r2)
        return res
      }
    }
  }
  return tc.MakeNone()
}

func bundnumLinearRegressionNumbers(x interface{}, y interface{}) interface{} {
  switch rx := x.(type) {
  case *tc.TCNumbers:
    switch ry := y.(type) {
    case *tc.TCNumbers:
      if rx.Len() == ry.Len() {
        res := tc.MakeDict()
        alpha, beta := stat.LinearRegression(rx.N, ry.N, nil, false)
        r2 := stat.RSquared(rx.N, ry.N, nil, alpha, beta)
        res.D.Set("alpha", alpha)
        res.D.Set("beta", beta)
        res.D.Set("rsquare", r2)
        return res
      }
    }
  }
  return tc.MakeNone()
}

func BUNDnumLinearRegression(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  err := l.ExecuteOperator("numlinearreg", q)
  if err != nil {
    return nil, err
  }
  return nil, nil
}

func init() {
  tc.SetFunction("num.LinearRegression", BUNDnumLinearRegression)
  tc.RegisterOperatorCallback("numlinearreg", tc.Matrix, tc.Matrix, bundnumLinearRegressionMatrix)
  tc.RegisterOperatorCallback("numlinearreg", tc.Numbers, tc.Numbers, bundnumLinearRegressionNumbers)
}
