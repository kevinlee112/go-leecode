package easy


func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1{
		return ""
	}
	return str1[0:gcd(len(str1),len(str2))]
}

func gcd(a int,b int) int{
	if(a%b == 0){
		return b
	}else{
		return gcd(b,a%b)
	}
}