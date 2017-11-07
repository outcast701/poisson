package poisson

import (
    "time"

    "math"
    "math/rand"
)

type PoissonSample struct {
    Lambda float64
    seeded bool
}

func GeneratePoisson(lambda float64) PoissonSample {
    return PoissonSample{Lambda: lambda, seeded: false}
}

func (p PoissonSample) PMF(k int) float64{
    kf := float64(k)
    lambda := p.Lambda
    return (math.Pow(lambda, kf)*math.Pow(math.E, -lambda))/float64(Fact(k)) 
}

func (p PoissonSample) CMF(k int) float64 {
    mass := 0.0
    for i := 0; i <= k; i++ {
        mass += p.PMF(i)
    }
    return mass
}

func (p PoissonSample) Sample() int {
    if !p.seeded {
        p.seeded = !p.seeded
        rand.Seed(time.Now().UTC().UnixNano())
    }
    sample := rand.Float64()
    i := 0
    for {
        if sample < p.CMF(i) {
            return i
        }
        i++
    }      
}

func Fact(n int) int {
    if n == 0 {
        return 1
    } else {
        return n * Fact(n - 1)
    }
}

