package stdlib

import (
  "gonum.org/v1/gonum/stat"
  "github.com/gammazero/deque"
  tc "github.com/vulogov/ThreadComputation"
)



func bundnumSortMatrix(v interface{}) interface{} {
  switch src := v.(type) {
  case *tc.TCMatrix:
    data := BUNDMatrixToFloats(src)
    stat.SortWeighted(data, nil)
    return BUNDFloatsToMatrix(src, data)
  }
  return nil
}

func bundnumSortNumbers(v interface{}) interface{} {
  switch src := v.(type) {
  case *tc.TCNumbers:
    stat.SortWeighted(src.N, nil)
    return src
  }
  return nil
}

func BUNDnumSort(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  err := l.ExecuteSingleArgumentFunction("numsort", q)
  if err != nil {
    return nil, err
  }
  return nil, nil
}

func init() {
  tc.SetFunction("num.Sort", BUNDnumSort)
  tc.RegisterFunctionCallback("numsort", tc.Matrix, bundnumSortMatrix)
  tc.RegisterFunctionCallback("numsort", tc.Numbers, bundnumSortNumbers)
}
