package stdlib

import (
  "gonum.org/v1/gonum/stat"
  "github.com/gammazero/deque"
  tc "github.com/vulogov/ThreadComputation"
)



func bundnumMomentAboutMatrix(v interface{}) interface{} {
  switch src := v.(type) {
  case *tc.TCMatrix:
    data := BUNDMatrixToFloats(src)
    me := stat.Mean(data, nil)
    out := tc.MakeNumbers()
    for _, e := range data {
      out.Add(stat.MomentAbout(e, data, me, nil))
    }
    return BUNDFloatsToMatrix(src, out.N)
  }
  return nil
}

func bundnumMomentAboutNumbers(v interface{}) interface{} {
  switch src := v.(type) {
  case *tc.TCNumbers:
    me := stat.Mean(src.N, nil)
    out := tc.MakeNumbers()
    for i := 0; i < src.Len(); i++ {
      out.Add(stat.MomentAbout(src.N[i], src.N, me, nil))
    }
    return out
  }
  return nil
}

func BUNDnumMomentAbout(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  err := l.ExecuteSingleArgumentFunction("nummomentabout", q)
  if err != nil {
    return nil, err
  }
  return nil, nil
}

func init() {
  tc.SetFunction("num.MomentAbout", BUNDnumMomentAbout)
  tc.RegisterFunctionCallback("nummomentabout", tc.Matrix, bundnumMomentAboutMatrix)
  tc.RegisterFunctionCallback("nummomentabout", tc.Numbers, bundnumMomentAboutNumbers)
}
