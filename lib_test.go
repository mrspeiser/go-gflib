package gfl_test

import (
  "testing"
  "fmt"
  "github.com/mrspeiser/go-gflib"
)

func TestGreet(t *testing.T){
  gfl.Greet("Matthew")
}

func TestUnsignedInt(t *testing.T){
  gfl.SetSeed()
  a := gfl.UnsignedInt(100)
  b := gfl.UnsignedInt(100)
  fmt.Printf("\nrandom integer: %d\n",a)
  fmt.Printf("random integer: %d\n",b)
}
