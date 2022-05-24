package stdlib

import (
  "math"
)

func init() {
  RegisterRawMathFunctionCallback("math.Abs", math.Abs)
  RegisterRawMathFunctionCallback("math.Sin", math.Sin)
  RegisterRawMathFunctionCallback("math.SinHyper", math.Sinh)
  RegisterRawMathFunctionCallback("math.Cos", math.Cos)
  RegisterRawMathFunctionCallback("math.CosHyper", math.Cosh)
  RegisterRawMathFunctionCallback("math.Tan", math.Tan)
  RegisterRawMathFunctionCallback("math.TanHyper", math.Tanh)
  RegisterRawMathFunctionCallback("math.ArcSin", math.Asin)
  RegisterRawMathFunctionCallback("math.ArcSinHyper", math.Asinh)
  RegisterRawMathFunctionCallback("math.ArcCos", math.Acos)
  RegisterRawMathFunctionCallback("math.ArcCosHyper", math.Acosh)
  RegisterRawMathFunctionCallback("math.ArcTan", math.Atan)
  RegisterRawMathFunctionCallback("math.ArcTanHyper", math.Atanh)
  RegisterRawMathFunctionCallback("math.CubeRoot", math.Cbrt)
  RegisterRawMathFunctionCallback("math.LeastInteger", math.Ceil)
  RegisterRawMathFunctionCallback("math.Error", math.Erf)
  RegisterRawMathFunctionCallback("math.ErrorComplimentary", math.Erfc)
  RegisterRawMathFunctionCallback("math.Exp", math.Exp)
  RegisterRawMathFunctionCallback("math.Exp2", math.Exp2)
  RegisterRawMathFunctionCallback("math.Exp1", math.Expm1)
  RegisterRawMathFunctionCallback("math.Floor", math.Floor)
  RegisterRawMathFunctionCallback("math.Gamma", math.Gamma)
  RegisterRawMathFunctionCallback("math.Bessel0", math.J0)
  RegisterRawMathFunctionCallback("math.Bessel1", math.J1)
  RegisterRawMathFunctionCallback("math.Log", math.Log)
  RegisterRawMathFunctionCallback("math.Log10", math.Log10)
  RegisterRawMathFunctionCallback("math.Round", math.Round)
  RegisterRawMathFunctionCallback("math.RoundToEven", math.RoundToEven)
  RegisterRawMathFunctionCallback("math.Sqrt", math.Sqrt)
  RegisterRawMathFunctionCallback("math.Trunc", math.Trunc)
}
