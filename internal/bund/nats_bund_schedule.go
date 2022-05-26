package bund

import (
	"github.com/pieterclaerhout/go-log"
	"github.com/bamzi/jobrunner"
	"github.com/vulogov/monitoringbund/internal/stdlib"
	"github.com/gammazero/deque"
	tc "github.com/vulogov/ThreadComputation"
)

func BUNDGetStringFromDictOrContext(l *tc.TCExecListener, d *tc.TCDict, key string) string {
	if d.D.Key(key) {
		v := d.D.Get(key)
		switch v.(type) {
		case string:
			return v.(string)
		}
	} else {
		if l.TC.HaveContext(key) {
			v := l.TC.GetContext(key)
			switch v.(type) {
			case string:
				return v.(string)
			}
		}
	}
	return ""
}

func BUNDGetFromDictOrContext(l *tc.TCExecListener, d *tc.TCDict, key string) interface{} {
	if d.D.Key(key) {
		v := d.D.Get(key)
		return v
	} else {
		if l.TC.HaveContext(key) {
			v := l.TC.GetContext(key)
			return v
		}
	}
	return nil
}

func bundNatsSchedule(l *tc.TCExecListener, v interface{}) interface{} {
	var schedule, name string
	var args     []string

  switch src := v.(type) {
  case *tc.TCDict:
		if src.D.Key("uri") {
			uri := src.D.Get("uri")
			switch uri.(type) {
			case string:
				schedule = BUNDGetStringFromDictOrContext(l, src, "schedule")
				name =   BUNDGetStringFromDictOrContext(l, src, "name")
				if schedule == "" {
					log.Error("No SCHEDULE specified in dictionary")
					return nil
				}
				sjob := TheScript{Name: name, Uri: uri.(string)}
				_args := BUNDGetFromDictOrContext(l, src, "args")
				if _args != nil {
					switch _args.(type) {
					case *tc.TCList:
						for n := 0; n < _args.(*tc.TCList).Len(); n++ {
							arg := _args.(*tc.TCList).Q.At(n)
							switch arg.(type) {
							case string:
								args = append(args, arg.(string))
							}
						}
						sjob.Args = args
					}
				}
				jobrunner.Schedule(schedule, sjob)
				log.Debugf("[ SCHEDULE ] %v at %v", uri, schedule)
			}
		} else {
			log.Error("No URI specified in dictionary")
		}
	default:
		log.Error("Dictionary was expected. %T found", v)
  }
  return nil
}



func BUNDNatsSchedule(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
	err := stdlib.ExecuteSingleArgumentFunction(l, "natsschedule", q)
  if err != nil {
    return nil, err
  }
	return nil, nil
}



func init() {
	stdlib.RegisterContextFunctionCallback("natsschedule", tc.Dict, bundNatsSchedule)
	tc.SetFunction("mbund.Schedule", BUNDNatsSchedule)
}
