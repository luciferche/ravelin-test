package main

import (
	"strconv"
	"fmt"
)



func gh(s string) string {
	var h uint64
	l := len(s)
	fmt.Println("Length %v - sssss - %v",l, s)
  if l > 0 {
    for i := 0;i < l;i++ {
			c := s[i]
      h = ((h<<5) - h) + uint64(c) | 0;
		}
	}
	hs := strconv.FormatUint(h, 16)
	fmt.Println("HASH of string - %v is: \n %v", s,hs)
  return strconv.FormatUint(h, 16)
}