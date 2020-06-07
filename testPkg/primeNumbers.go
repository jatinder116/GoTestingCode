package testPkg

func Primenumbers(num int) (a []int) {	
    for i := 1; i <= num; i++ { 
        // skip 0 and 1 as they are niether prime nor composite 
        if i == 1 || i == 0{
			continue
		}
        check := 1; 
        for j := 2; j <= i / 2; j++ { 
            if i % j == 0{ 
                check = 0 
                break 
            } 
        } 
        // check = 1 means i is prime and 0 means not prime 
        if check == 1{
			a = append(a, i)
		}
	} 
	return a;
}
