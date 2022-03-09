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
  fmt.Printf("\nrandom integer: %d",a)
  fmt.Printf("\nrandom integer: %d\n",b)
}

func TestRandomAlnum(t *testing.T){
  gfl.SetSeed()
  uia := gfl.UnsignedInt(16)
  uib := gfl.UnsignedInt(16)
  uic := gfl.UnsignedInt(16)

  fmt.Printf("\nrandom integer: %d",uia)
  fmt.Printf("\nrandom integer: %d",uib)
  fmt.Printf("\nrandom integer: %d\n",uic)

  a := gfl.RandomAlnum(uia)
  b := gfl.RandomAlnum(uib)
  c := gfl.RandomAlnum(uic)

  fmt.Printf("\nrandom alnum: %s",a)
  fmt.Printf("\nrandom alnum: %s",b)
  fmt.Printf("\nrandom alnum: %s\n",c)

  x := gfl.RandomAlnum(24)
  y := gfl.RandomAlnum(24)
  z := gfl.RandomAlnum(24)

  fmt.Printf("\nrandom alnum 24 chars: %s",x)
  fmt.Printf("\nrandom alnum 24 chars: %s",y)
  fmt.Printf("\nrandom alnum 24 chars: %s\n",z)
}

func TestArrayMaps(t *testing.T){
  a := gfl.ArrayofStringMaps(10)
  b := gfl.ArrayofIntMaps(10)
  fmt.Printf("\n")
  fmt.Println(a)
  fmt.Println(b)
}
