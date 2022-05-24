package stdlib

import (
  "fmt"
  "time"
  "strings"
  "golang.org/x/exp/rand"
  "gonum.org/v1/gonum/stat/distuv"
  "github.com/gammazero/deque"
  tc "github.com/vulogov/ThreadComputation"
)

type BUND_Random_Distribution_Exponential struct {
  Dist        *distuv.Exponential
}

func MakeRandomDistributionExponential(rate float64) *BUND_Random_Distribution_Exponential {
  res := new(BUND_Random_Distribution_Exponential)
  res.Dist = new(distuv.Exponential)
  res.Dist.Rate = rate
  res.Dist.Src   = rand.NewSource(uint64(time.Now().UnixNano()))
  return res
}

func (d *BUND_Random_Distribution_Exponential) String() string {
  return fmt.Sprintf("random.Exponential[ rate=%v entropy=%v mean=%v ]", d.Dist.Rate, d.Dist.Entropy(), d.Dist.Mean())
}

func (d *BUND_Random_Distribution_Exponential) Entropy() float64 {
  return d.Dist.Entropy()
}

func (d *BUND_Random_Distribution_Exponential) Get() float64 {
  return d.Dist.Rand()
}

func BUND_Random_Distribution_ExponentialConvert(data interface{}, to_type int) interface{} {
  switch e := data.(type) {
  case *BUND_Random_Distribution_Exponential:
    switch to_type {
    case tc.String:
      return e.String()
    }
  }
  return nil
}

func BUNDRDEGet(x interface{}, y interface{}) interface{} {
  switch c := x.(type) {
  case *BUND_Random_Distribution_Exponential:
    switch k := y.(type) {
    case string:
      k = strings.ToLower(k)
      switch k {
      case "entropy":
        return c.Dist.Entropy()
      case "random":
        return c.Dist.Rand()
      case "mean":
        return c.Dist.Mean()
      case "median":
        return c.Dist.Median()
      case "stddev":
        return c.Dist.StdDev()
      }
    }
    return c.Dist.Rand()
  }
  return tc.MakeNone()
}

func BUNDRDEGenerator(v interface{}) *tc.TCIterator {
  switch v.(type) {
  case *BUND_Random_Distribution_Exponential:
    res := new(tc.TCIterator)
    res.Type = RandomExponential
    res.Gen = v.(*BUND_Random_Distribution_Exponential)
    res.Last = nil
    return res
  }
  return nil
}

func BUNDRDEIterGet(res *tc.TCIterator) interface{} {
  out := res.Gen.(*BUND_Random_Distribution_Exponential).Get()
  res.Last = out
  return out
}

func BUNDrandomDistributionExponential(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  rate := EnsureFloat(l.TC.FromContext("rate", 1.0), 1.0)
  for q.Len() > 0 {
    e := q.PopFront()
    switch e.(type) {
    case *tc.TCDict:
      rate = EnsureFloatFromDict(e.(*tc.TCDict), "rate", rate)
    }
  }
  return MakeRandomDistributionExponential(rate), nil
}

func init() {
  tc.SetCommand("random.Distribution.Exponential", BUNDrandomDistributionExponential)
  tc.RegisterConvertCallback(RandomExponential, BUND_Random_Distribution_ExponentialConvert)
  tc.RegisterOperatorCallback("get", RandomNormal, tc.String, BUNDRDNGet)
  tc.RegisterGenCallback(RandomExponential,  BUNDRDEGenerator)
  tc.RegisterNextCallback(RandomExponential, BUNDRDEIterGet)
  tc.RegisterPrevCallback(RandomExponential, BUNDRDEIterGet)
}
