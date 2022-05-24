package stdlib

import (
  tc "github.com/vulogov/ThreadComputation"
)

func bundnnNormalizeNumbers(v interface{}) interface{} {
  switch src := v.(type) {
  case *tc.TCNumbers:
    dst := NumNorm(src.N)
    res := tc.MakeNumbers()
    res.Set(dst)
    return res
  }
  return nil
}




func init() {
  tc.RegisterFunctionCallback("mlnormalize", tc.Numbers, bundnnNormalizeNumbers)
}
