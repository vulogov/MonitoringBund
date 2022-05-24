package stdlib

import (
  "gonum.org/v1/gonum/mat"
  tc "github.com/vulogov/ThreadComputation"
)

func BUNDMatrixToFloats(m *tc.TCMatrix) []float64 {
  dst := make([]float64, 0)
  m1, _ := m.M.Dims()
  for i := 0; i < m1; i++ {
    dst = append(dst, mat.Row(nil, i, m.M)...)
  }
  return dst
}

func BUNDFloatsToMatrix(m *tc.TCMatrix, data []float64) *tc.TCMatrix {
  m1, n1 := m.M.Dims()
  res := new(tc.TCMatrix)
  res.M = mat.NewDense(m1,n1, data)
  res.P = 0.0
  return res
}
