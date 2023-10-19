// package main

// import (
// 	"os"
// )

// func Atoi(ii string) int {
// 	r := []rune(ii)
// 	san := 0
// 	for i := range r {
// 		if r[i] != '-' {
// 			san = san*10 + int(r[i]-'0')
// 		}
// 	}
// 	if r[0] == '-' {
// 		san *= -1
// 	}
// 	return san
// }

// func IsNumeric(ii string) bool {
// 	r := []rune(ii)
// 	tt := true
// 	for i := range r {
// 		if (r[i] < '0' || r[i] > '9') && r[i] != 0 {
// 			tt = false
// 		}
// 	}
// 	return tt
// }

// func Itoi(ii int) string {
// 	s := ""
// 	minus := ""
// 	if ii < 0 {
// 		ii *= -1
// 		minus = "-"
// 	}
// 	for ii%10 != ii {
// 		s = string(rune(ii%10+'0')) + s
// 		ii /= 10
// 	}
// 	return minus + string(rune(ii%10+'0')) + s
// }

// func main() {
// 	if len(os.Args) != 4 {
// 		return
// 	}
// 	arg := os.Args[1:]
// 	bir := 0
// 	eki := 0
// 	max_san := 1<<63 - 1
// 	min_san := -1 << 63
// 	if IsNumeric(arg[0]) {
// 		bir = Atoi(arg[0])
// 	}
// 	if IsNumeric(arg[2]) {
// 		eki = Atoi(arg[2])
// 	}
// 	if bir > max_san || bir < min_san || eki > max_san || eki < min_san {
// 		return
// 	}
// 	oper := arg[1]
// 	switch oper {
// 	case "+":
// 		os.Stdout.WriteString(Itoi(bir+eki) + "\n")
// 	case "-":
// 		os.Stdout.WriteString(Itoi(bir-eki) + "\n")
// 	case "*":
// 		os.Stdout.WriteString(Itoi(bir*eki) + "\n")
// 	case "/":
// 		if eki == 0 {
// 			os.Stdout.WriteString("No division by 0\n")
// 		} else {
// 			os.Stdout.WriteString(Itoi(bir/eki) + "\n")
// 		}
// 	case "%":
// 		if eki == 0 {
// 			os.Stdout.WriteString("No modulo by 0\n")
// 		} else {
// 			os.Stdout.WriteString(Itoi(bir%eki) + "\n")
// 		}
// 	}
// }
