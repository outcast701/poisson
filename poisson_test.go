package poisson

import (
    "testing"
)


func TestPoisson(t *testing.T) {
    p := GeneratePoisson(0.1)
    for i := 0; i <= 10; i++ {
        t.Log(p.PMF(i))
    }
}

func TestFact(t *testing.T) {
    t.Log(Fact(3))
    t.Log(Fact(4))
    t.Log(Fact(5))
    t.Log(Fact(6))
}

func TestPoissonCMF(t *testing.T) {
    p := GeneratePoisson(3.1)
    for i := 0; i <= 10; i++ {
        t.Log(p.CMF(i))
    }
}


func TestSample(t *testing.T) {
    p := GeneratePoisson(3.1)
    for i := 0; i < 5; i++ {
        t.Log(p.Sample())
    }
}
