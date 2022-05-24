package stdlib

import (
  "gonum.org/v1/gonum/stat"
  "github.com/gammazero/deque"
  tc "github.com/vulogov/ThreadComputation"
)



func bundnumHellingerMatrix(x interface{}, y interface{}) interface{} {
  switch rx := x.(type) {
  case *tc.TCMatrix:
    switch ry := y.(type) {
    case *tc.TCMatrix:
      m1, n1 := rx.M.Dims()
      m2, n2 := ry.M.Dims()
      if m1 == m2 && n1 == n2 {
        src1 := BUNDMatrixToFloats(rx)
        src2 := BUNDMatrixToFloats(ry)
        res := stat.Hellinger(src1, src2)
        return res
      }
    }
  }
  return tc.MakeNone()
}

func bundnumHellingerNumbers(x interface{}, y interface{}) interface{} {
  switch rx := x.(type) {
  case *tc.TCNumbers:
    switch ry := y.(type) {
    case *tc.TCNumbers:
      if rx.Len() == ry.Len() {
        res := stat.Hellinger(rx.N, ry.N)
        return res
      }
    }
  }
  return tc.MakeNone()
}

func BUNDnumHellinger(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  err := l.ExecuteOperator("numhellinger", q)
  if err != nil {
    return nil, err
  }
  return nil, nil
}

func init() {
  tc.SetFunction("num.Hellinger", BUNDnumCovariance)
  tc.RegisterOperatorCallback("numhellinger", tc.Matrix, tc.Matrix, bundnumHellingerMatrix)
  tc.RegisterOperatorCallback("numhellinger", tc.Numbers, tc.Numbers, bundnumHellingerNumbers)
}
