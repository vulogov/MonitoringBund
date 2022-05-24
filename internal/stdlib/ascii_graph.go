package stdlib

import (
  "github.com/guptarohit/asciigraph"
  "github.com/gammazero/deque"
  tc "github.com/vulogov/ThreadComputation"
)

func bundStringGraph(data []float64) string {
  return asciigraph.Plot(data)
}

func bundGenerateGraph(v interface{}) interface{} {
  switch data := v.(type) {
  case *tc.TCNumbers:
    return bundStringGraph(data.N)
  }
  return nil
}

func BUNDstringGraph(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  err := l.ExecuteSingleArgumentFunction("asciigraph", q)
  if err != nil {
    return nil, err
  }
  return nil, nil
}


func init() {
  tc.SetFunction("graph", BUNDstringGraph)
  tc.RegisterFunctionCallback("asciigraph", tc.Numbers, bundGenerateGraph)
}
