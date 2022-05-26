package bund

import (
	"github.com/gammazero/deque"
	"github.com/pieterclaerhout/go-log"
	tc "github.com/vulogov/ThreadComputation"
)

func bundNatsSendSync(v interface{}) interface{} {
  switch src := v.(type) {
  case string:
		pkt, err := MakeSync(src)
		if err == nil {
			log.Debugf("Sending SYNC(%v)", src)
			NatsSendSys(pkt)
		}
  }
  return nil
}

func bundNatsSendMsg(v interface{}) interface{} {
  switch src := v.(type) {
  case string:
		pkt, err := MakeMsg(src)
		if err == nil {
			NatsSendSys(pkt)
			return true
		}
  }
  return false
}

func bundNatsExecScript(v interface{}) interface{} {
  switch script := v.(type) {
  case string:
		pkt, err := MakeScript("mbund.Script[]", "submit", []byte(script), nil, nil)
		if err == nil {
			NatsSend(pkt)
		}
  }
  return nil
}

func BUNDNatsSendSync(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
	err := l.ExecuteSingleArgumentFunction("natssendsync", q)
  if err != nil {
    return nil, err
  }
	return nil, nil
}

func BUNDNatsSendMsg(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
	err := l.ExecuteSingleArgumentFunction("natssendmsg", q)
  if err != nil {
    return nil, err
  }
	return nil, nil
}

func BUNDNatsExecScript(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
	err := l.ExecuteSingleArgumentFunction("natsexec", q)
  if err != nil {
    return nil, err
  }
	return nil, nil
}

func init() {
	tc.RegisterFunctionCallback("natssendsync", tc.String, bundNatsSendSync)
	tc.SetCommand("mbund.Sync", BUNDNatsSendSync)
	tc.RegisterFunctionCallback("natssendmsg", tc.String, bundNatsSendMsg)
	tc.SetFunction("mbund.Message", BUNDNatsSendMsg)
	tc.RegisterFunctionCallback("natsexec", tc.String, bundNatsExecScript)
	tc.SetFunction("mbund.Script", BUNDNatsExecScript)
}
