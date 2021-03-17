package utils

import "fmt"

func Arrays()  {

	var a[5] int ;
	var arr[2][5] int

	for key, value := range a {
		fmt.Printf("key:%d  value:%d\n", key, value)
	}

	for key2,value2 := range arr{
		fmt.Println("key:  value:\n", key2, value2)
	}

	for i:=0 ; i < len(a) ;i++ {
		fmt.Println(a[i])
	}

}
