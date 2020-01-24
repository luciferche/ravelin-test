package main

import (
	"strconv"
	"fmt"
)

func gh(s string) string {
	var h uint64 = 5381;
	
	for i:=0;i<len(s);i++ {
		c := s[i]
			fmt.Println(" -- cc - ˜%v", uint8(c))
		// 	h = ((h<<5)-h)
		h = ((h<<5) + h) + uint64(c)
	}

	return strconv.FormatUint(h, 16);
}

