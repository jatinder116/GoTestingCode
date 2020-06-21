package assignment1pkg
import (
    "strconv"
)

func Primenum(num int) (a string) {	
    for i := 2; i <= num; i++ { 
        var check int = 1 
        for j := 2; j <= i / 2; j++ { 
            if i % j == 0{ 
                check = 0 
            } 
        } 
        if check == 1{
            a = a + strconv.Itoa(i) + " ";
		}
	} 
	return a;
}
