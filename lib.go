package gfl

import (
  "fmt"
  "math/rand"
  "time"
)

const alnums = "abcdefghijklmnopqrstuvwyxzABCDEFGHIJKLMNOPQRSTUVWYXZ1234567890"
const alphabet = "abcdefghijklmnopqrstuvwyxzABCDEFGHIJKLMNOPQRSTUVWYXZ"

func Greet(name string){
  fmt.Printf("Hello %s!",name)
}

func Getalphabet() string {
  return alphabet;
}

func Getalnums() string {
  return alnums;
}

func SetSeed(){
  /*
    Seed uses the provided seed value to initialize the default Source to a deterministic state. 
    If Seed is not called, the generator behaves as if seeded by Seed(1). 
    Seed values that have the same remainder when divided by 2³¹-1 
    generate the same pseudo-random sequence. 
    Seed, unlike the Rand.Seed method, is safe for concurrent use.
  */
  rand.Seed(time.Now().UTC().UnixNano())
}

func UnsignedInt(length int) int {
  l := length
  if l < 1 {
     l = 1
  }
  Int := rand.Intn(l)
  return Int
}

func RandomAlnum(length int) string {
  s := make([]byte,length)
  for i := range s {
    s[i] = alnums[rand.Intn(len(alnums))]
  }
  return string(s)
}

