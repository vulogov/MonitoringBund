package stdlib

import (
  tc "github.com/vulogov/ThreadComputation"
)

func EnsureFloat(i interface{}, dflt float64) float64 {
  fun := tc.GetConverterCallback(i)
  if fun == nil {
    return dflt
  }
  res := fun(i, tc.Float)
  if res == nil {
    return dflt
  }
  return res.(float64)
}

func EnsureFloatFromDict(d *tc.TCDict, key string, dflt float64) float64 {
  if d.D.Key(key) {
    val := d.D.Get(key)
    return EnsureFloat(val, dflt)
  }
  return dflt
}
