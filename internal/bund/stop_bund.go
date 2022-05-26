package bund

import (
	"github.com/gammazero/deque"
	"github.com/pieterclaerhout/go-log"
	tc "github.com/vulogov/ThreadComputation"
)

func bundStopMsg(v interface{}) interface{} {
  switch msg := v.(type) {
  case string:
		log.Infof("[ STOP ] %v", msg)
  }
  return nil
}

func BUNDSendStop(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
	err := l.ExecuteSingleArgumentFunction("bundsendstop", q)
	SendStop()
  if err != nil {
    return nil, err
  }
	return nil, nil
}

func init() {
	tc.RegisterFunctionCallback("bundsendstop", tc.String, bundStopMsg)
	tc.SetCommand("mbund.Stop", BUNDSendStop)
}
