package utils

import "math/big"

func StringToBigInt(s string) (*big.Int, bool) {
  return new(big.Int).SetString(s, 10)
}
