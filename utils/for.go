package utils

import "fmt"

func For()  {
	i := 1
	for i<=3{
		println(i)
		i++
	}

	for i:=1; i<=3;i++ {
		println(i)
	}

	j := 1
	for {
		fmt.Println(j)
		if j == 3 {
			break
		}
		j++;
	}
}
