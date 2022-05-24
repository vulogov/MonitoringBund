package stdlib

import (
  "gonum.org/v1/gonum/stat"
  "github.com/gammazero/deque"
  tc "github.com/vulogov/ThreadComputation"
)



func bundnumScoreMatrix(v interface{}) interface{} {
  switch src := v.(type) {
  case *tc.TCMatrix:
    data := BUNDMatrixToFloats(src)
    sd := stat.StdDev(data, nil)
    me := stat.Mean(data, nil)
    out := tc.MakeNumbers()
    for _, e := range data {
      out.Add(stat.StdScore(e, me, sd))
    }
    return BUNDFloatsToMatrix(src, out.N)
  }
  return nil
}

func bundnumScoreNumbers(v interface{}) interface{} {
  switch src := v.(type) {
  case *tc.TCNumbers:
    sd := stat.StdDev(src.N, nil)
    me := stat.Mean(src.N, nil)
    out := tc.MakeNumbers()
    for i := 0; i < src.Len(); i++ {
      out.Add(stat.StdScore(src.N[i], me, sd))
    }
    return out
  }
  return nil
}

func BUNDnumScore(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  err := l.ExecuteSingleArgumentFunction("numscore", q)
  if err != nil {
    return nil, err
  }
  return nil, nil
}

func init() {
  tc.SetFunction("num.Score", BUNDnumScore)
  tc.RegisterFunctionCallback("numscore", tc.Matrix, bundnumScoreMatrix)
  tc.RegisterFunctionCallback("numscore", tc.Numbers, bundnumScoreNumbers)
}
