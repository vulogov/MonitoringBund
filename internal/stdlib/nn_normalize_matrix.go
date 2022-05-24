package stdlib

import (
  tc "github.com/vulogov/ThreadComputation"
)



func bundnnNormalizeMatrix(v interface{}) interface{} {
  switch src := v.(type) {
  case *tc.TCMatrix:
    n := BUNDMatrixToFloats(src)
    dst := NumNorm(n)
    res := BUNDFloatsToMatrix(src, dst)
    return res
  }
  return nil
}




func init() {
  tc.RegisterFunctionCallback("mlnormalize", tc.Matrix, bundnnNormalizeMatrix)
}
