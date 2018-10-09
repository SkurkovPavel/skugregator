package main

import (
	sku "./skuParser"
	"fmt"
	"time"
)

func main(){
	ticker := time.Tick(time.Second)
	 i := 0

	for tickTime := range ticker {
		i++

		fmt.Printf("\r step %v time %v",i,tickTime.Format("15:04:05"))
		if i == 10 {
			res := sku.SkuPars()
			if res == "finito" {
				i = 0
			}
		}


	}







}
