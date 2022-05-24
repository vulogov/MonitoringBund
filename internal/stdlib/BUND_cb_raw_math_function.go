package stdlib

import (
  "fmt"
  "github.com/pieterclaerhout/go-log"
  tc "github.com/vulogov/ThreadComputation"
)


type TCRawMathFunction func(float64) float64

func bundRawMathFunctionFloat(l *tc.TCExecListener, name interface{}, x interface{}) interface{} {
  log.Debugf("raw-math float handler: %v %v", name, x)
  switch name.(type) {
  case string:
    switch x.(type) {
    case float64:
      fun := GetRawMathFunctionCallback(name.(string))
      if fun != nil {
        res := fun(x.(float64))
        return res
      }
    }
  }
  return nil
}

func bundRawMathFunctionInt(l *tc.TCExecListener, name interface{}, x interface{}) interface{} {
  log.Debugf("raw-math int handler: %v %v", name, x)
  switch name.(type) {
  case string:
    switch x.(type) {
    case int64:
      fun := GetRawMathFunctionCallback(name.(string))
      if fun != nil {
        res := fun(float64(x.(int64)))
        return res
      }
    }
  }
  return nil
}

func bundRawMathFunctionString(l *tc.TCExecListener, name interface{}, x interface{}) interface{} {
  log.Debugf("raw-math int handler: %v %v", name, x)
  switch name.(type) {
  case string:
    switch x.(type) {
    case string:
      fun := GetRawMathFunctionCallback(name.(string))
      if fun != nil {
        convfun := tc.GetConverterCallback(x.(string))
        if convfun == nil {
          l.TC.MakeError("Can not get string converter")
          return nil
        }
        v := convfun(x.(string), tc.Float)
        switch v.(type) {
        case float64:
          res := fun(v.(float64))
          return res
        }
      }
    }
  }
  return nil
}

func bundRawMathFunctionNumbers(l *tc.TCExecListener, name interface{}, x interface{}) interface{} {
  log.Debugf("raw-math numbers handler: %v %v", name, x)
  switch name.(type) {
  case string:
    switch x.(type) {
    case *tc.TCNumbers:
      fun := GetRawMathFunctionCallback(name.(string))
      if fun != nil {
        res := tc.MakeNumbers()
        for i := 0; i < x.(*tc.TCNumbers).Len(); i++ {
          res.Add(fun(x.(*tc.TCNumbers).N[i]))
        }
        return res
      }
    }
  }
  return nil
}

func bundRawMathFunctionMatrix(l *tc.TCExecListener, name interface{}, x interface{}) interface{} {
  log.Debugf("raw-math matrix handler: %v %v", name, x)
  switch name.(type) {
  case string:
    switch x.(type) {
    case *tc.TCMatrix:
      fun := GetRawMathFunctionCallback(name.(string))
      if fun != nil {
        n_res := tc.MakeNumbers()
        m_data := BUNDMatrixToFloats(x.(*tc.TCMatrix))
        for i := 0; i < len(m_data); i++ {
          n_res.Add(fun(m_data[i]))
        }
        return BUNDFloatsToMatrix(x.(*tc.TCMatrix), n_res.N)
      }
    }
  }
  return nil
}

func GetRawMathFunctionCallback(name string) TCRawMathFunction {
  fname := fmt.Sprintf("raw-math.%v", name)
  log.Debugf("Looking for raw math function %v", name)
  if fun, ok := Callbacks.Load(fname); ok {
    log.Debugf("Got: %v", fname)
    return fun.(TCRawMathFunction)
  }
  log.Debugf("No raw function for: %v", name)
  return nil
}

func RegisterRawMathFunctionCallback(name string, fun TCRawMathFunction) {
  fname := fmt.Sprintf("raw-math.%v", name)
  Callbacks.Delete(fname)
  Callbacks.Store(fname, fun)
}

func init() {
  RegisterContextOperatorCallback("rawmath", tc.String, tc.Float, bundRawMathFunctionFloat)
  RegisterContextOperatorCallback("rawmath", tc.String, tc.Int, bundRawMathFunctionInt)
  RegisterContextOperatorCallback("rawmath", tc.String, tc.String, bundRawMathFunctionString)
  RegisterContextOperatorCallback("rawmath", tc.String, tc.Numbers, bundRawMathFunctionNumbers)
  RegisterContextOperatorCallback("rawmath", tc.String, tc.Matrix, bundRawMathFunctionMatrix)
}
