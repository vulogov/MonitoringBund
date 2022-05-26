package bund

import (
	"github.com/gammazero/deque"
	tc "github.com/vulogov/ThreadComputation"
)

func bundSetLocalConfig(v interface{}) interface{} {
  switch src := v.(type) {
  case *tc.TCPair:
		switch key := src.X.(type) {
		case string:
			switch val := src.Y.(type) {
			case string:
				Conf.Store(key, val)
			}
		}
  }
  return nil
}

func BUNDSetLocalConf(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
	err := l.ExecuteSingleArgumentFunction("localconfigset", q)
  if err != nil {
    return nil, err
  }
	return nil, nil
}

func init() {
	tc.RegisterFunctionCallback("localconfigset", tc.Pair, bundSetLocalConfig)
	tc.SetCommand("local.SetConfig", BUNDSetLocalConf)
}
