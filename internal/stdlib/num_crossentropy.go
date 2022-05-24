package stdlib

import (
  "gonum.org/v1/gonum/stat"
  "github.com/gammazero/deque"
  tc "github.com/vulogov/ThreadComputation"
)



func bundnumCrossEntropyMatrix(x interface{}, y interface{}) interface{} {
  switch rx := x.(type) {
  case *tc.TCMatrix:
    switch ry := y.(type) {
    case *tc.TCMatrix:
      m1, n1 := rx.M.Dims()
      m2, n2 := ry.M.Dims()
      if m1 == m2 && n1 == n2 {
        src1 := BUNDMatrixToFloats(rx)
        src2 := BUNDMatrixToFloats(ry)
        res := stat.CrossEntropy(src1, src2)
        return res
      }
    }
  }
  return tc.MakeNone()
}

func bundnumCrossEntropyNumbers(x interface{}, y interface{}) interface{} {
  switch rx := x.(type) {
  case *tc.TCNumbers:
    switch ry := y.(type) {
    case *tc.TCNumbers:
      if rx.Len() == ry.Len() {
        res := stat.CrossEntropy(rx.N, ry.N)
        return res
      }
    }
  }
  return tc.MakeNone()
}

func BUNDnumCrossEntropy(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  err := l.ExecuteOperator("numcovariance", q)
  if err != nil {
    return nil, err
  }
  return nil, nil
}

func init() {
  tc.SetFunction("num.CrossEntropy", BUNDnumCrossEntropy)
  tc.RegisterOperatorCallback("numcentropy", tc.Matrix, tc.Matrix, bundnumCrossEntropyMatrix)
  tc.RegisterOperatorCallback("numcentropy", tc.Numbers, tc.Numbers, bundnumCrossEntropyNumbers)
}
