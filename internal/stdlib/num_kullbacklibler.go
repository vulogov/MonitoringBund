package stdlib

import (
  "gonum.org/v1/gonum/stat"
  "github.com/gammazero/deque"
  tc "github.com/vulogov/ThreadComputation"
)



func bundnumKullbackLeiblerMatrix(x interface{}, y interface{}) interface{} {
  switch rx := x.(type) {
  case *tc.TCMatrix:
    switch ry := y.(type) {
    case *tc.TCMatrix:
      m1, n1 := rx.M.Dims()
      m2, n2 := ry.M.Dims()
      if m1 == m2 && n1 == n2 {
        src1 := BUNDMatrixToFloats(rx)
        src2 := BUNDMatrixToFloats(ry)
        res := stat.KullbackLeibler(src1, src2)
        return res
      }
    }
  }
  return tc.MakeNone()
}

func bundnumKullbackLeiblerNumbers(x interface{}, y interface{}) interface{} {
  switch rx := x.(type) {
  case *tc.TCNumbers:
    switch ry := y.(type) {
    case *tc.TCNumbers:
      if rx.Len() == ry.Len() {
        res := stat.KullbackLeibler(rx.N, ry.N)
        return res
      }
    }
  }
  return tc.MakeNone()
}

func BUNDnumKullbackLeibler(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  err := l.ExecuteOperator("numkullbacklibler", q)
  if err != nil {
    return nil, err
  }
  return nil, nil
}

func init() {
  tc.SetFunction("num.KullbackLeibler", BUNDnumKullbackLeibler)
  tc.RegisterOperatorCallback("numkullbacklibler", tc.Matrix, tc.Matrix, bundnumKullbackLeiblerMatrix)
  tc.RegisterOperatorCallback("numkullbacklibler", tc.Numbers, tc.Numbers, bundnumKullbackLeiblerNumbers)
}
