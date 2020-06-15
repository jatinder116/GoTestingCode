package testpkg

func Sqrandcube(num int) (s, c int) {
	s = num * num
	 c = num * num * num
    return ;
}

// =========================== remove smae elements in ana array eg (["kk","jj","jj","ll"])
// for i := 0; i < len(Dummy); i++ {
// 	if Dummy[i] == Id {
// 		a = append(Dummy[:i], Dummy[i+1:]...)
// 		i--
// 	}
// }