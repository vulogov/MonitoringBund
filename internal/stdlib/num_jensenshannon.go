package stdlib

import (
  "gonum.org/v1/gonum/stat"
  "github.com/gammazero/deque"
  tc "github.com/vulogov/ThreadComputation"
)



func bundnumJensenShannonMatrix(x interface{}, y interface{}) interface{} {
  switch rx := x.(type) {
  case *tc.TCMatrix:
    switch ry := y.(type) {
    case *tc.TCMatrix:
      m1, n1 := rx.M.Dims()
      m2, n2 := ry.M.Dims()
      if m1 == m2 && n1 == n2 {
        src1 := BUNDMatrixToFloats(rx)
        src2 := BUNDMatrixToFloats(ry)
        res := stat.JensenShannon(src1, src2)
        return res
      }
    }
  }
  return tc.MakeNone()
}

func bundnumJensenShannonNumbers(x interface{}, y interface{}) interface{} {
  switch rx := x.(type) {
  case *tc.TCNumbers:
    switch ry := y.(type) {
    case *tc.TCNumbers:
      if rx.Len() == ry.Len() {
        res := stat.JensenShannon(rx.N, ry.N)
        return res
      }
    }
  }
  return tc.MakeNone()
}

func BUNDnumJensenShannon(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  err := l.ExecuteOperator("numjensenshannon", q)
  if err != nil {
    return nil, err
  }
  return nil, nil
}

func init() {
  tc.SetFunction("num.JensenShannon", BUNDnumJensenShannon)
  tc.RegisterOperatorCallback("numjensenshannon", tc.Matrix, tc.Matrix, bundnumJensenShannonMatrix)
  tc.RegisterOperatorCallback("numjensenshannon", tc.Numbers, tc.Numbers, bundnumJensenShannonNumbers)
}
