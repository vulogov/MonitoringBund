package stdlib

import (
  "github.com/gammazero/deque"
  "github.com/google/uuid"
  tc "github.com/vulogov/ThreadComputation"
)

func BUNDrandomUUID(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  return uuid.NewString(), nil
}


func init() {
  tc.SetCommand("random.UUID", BUNDrandomUUID)
}
