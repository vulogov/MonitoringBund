package stdlib

import (
  "github.com/gammazero/deque"
  "github.com/zach-klippenstein/goregen"
  tc "github.com/vulogov/ThreadComputation"
)

func BUNDrandomString(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  res, err := regen.Generate("[A-Za-z]{16,64}")
  if err != nil {
    return nil, err
  }
  return res, nil
}

func bundGenerateRandomStringFromPattern(v interface{}) interface{} {
  switch v.(type) {
  case string:
    res, err := regen.Generate(v.(string))
    if err != nil {
      return nil
    }
    return res
  }
  return nil
}

func BUNDrandomStringFromPattern(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  err := l.ExecuteSingleArgumentFunction("randomstring", q)
  if err != nil {
    return nil, err
  }
  return nil, nil
}


func init() {
  tc.SetCommand("random.String", BUNDrandomString)
  tc.SetCommand("random.StringFromPattern", BUNDrandomStringFromPattern)
  tc.RegisterFunctionCallback("randomstring", tc.String, bundGenerateRandomStringFromPattern)
}
