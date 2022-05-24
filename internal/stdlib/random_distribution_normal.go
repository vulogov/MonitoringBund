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

type BUND_Random_Distribution_Normal struct {
  Dist        *distuv.Normal
}

func MakeRandomDistributionNormal(mu float64, sigma float64) *BUND_Random_Distribution_Normal {
  res := new(BUND_Random_Distribution_Normal)
  res.Dist = new(distuv.Normal)
  res.Dist.Mu = mu
  res.Dist.Sigma = sigma
  res.Dist.Src   = rand.NewSource(uint64(time.Now().UnixNano()))
  return res
}

func (d *BUND_Random_Distribution_Normal) String() string {
  return fmt.Sprintf("random.Normal[ mu=%v sigma=%v entropy=%v mean=%v ]", d.Dist.Mu, d.Dist.Sigma, d.Dist.Entropy(), d.Dist.Mean())
}

func (d *BUND_Random_Distribution_Normal) Entropy() float64 {
  return d.Dist.Entropy()
}

func (d *BUND_Random_Distribution_Normal) Get() float64 {
  return d.Dist.Rand()
}

func BUNDRandomDistributionNormalConvert(data interface{}, to_type int) interface{} {
  switch e := data.(type) {
  case *BUND_Random_Distribution_Normal:
    switch to_type {
    case tc.String:
      return e.String()
    }
  }
  return nil
}

func BUNDRDNGet(x interface{}, y interface{}) interface{} {
  switch c := x.(type) {
  case *BUND_Random_Distribution_Normal:
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

func BUNDRDNGenerator(v interface{}) *tc.TCIterator {
  switch v.(type) {
  case *BUND_Random_Distribution_Normal:
    res := new(tc.TCIterator)
    res.Type = RandomNormal
    res.Gen = v.(*BUND_Random_Distribution_Normal)
    res.Last = nil
    return res
  }
  return nil
}

func BUNDRDNIterGet(res *tc.TCIterator) interface{} {
  out := res.Gen.(*BUND_Random_Distribution_Normal).Get()
  res.Last = out
  return out
}

func BUNDrandomDistributionNormal(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  mu := EnsureFloat(l.TC.FromContext("mu", 1.0), 1.0)
  sigma := EnsureFloat(l.TC.FromContext("sigma", 2.0), 2.0)
  for q.Len() > 0 {
    e := q.PopFront()
    switch e.(type) {
    case *tc.TCDict:
      mu = EnsureFloatFromDict(e.(*tc.TCDict), "mu", mu)
      sigma = EnsureFloatFromDict(e.(*tc.TCDict), "sigma", sigma)
    }
  }
  return MakeRandomDistributionNormal(mu, sigma), nil
}

func init() {
  tc.SetCommand("random.Distribution.Normal", BUNDrandomDistributionNormal)
  tc.RegisterConvertCallback(RandomNormal, BUNDRandomDistributionNormalConvert)
  tc.RegisterOperatorCallback("get", RandomNormal, tc.String, BUNDRDNGet)
  tc.RegisterGenCallback(RandomNormal,  BUNDRDNGenerator)
  tc.RegisterNextCallback(RandomNormal, BUNDRDNIterGet)
  tc.RegisterPrevCallback(RandomNormal, BUNDRDNIterGet)
}
