package main

func Burbuja(s []int64) {
	sLength := len(s) - 1
	var aux int64
	for i := 0; i < sLength; i++ {
		for j := 0; j < sLength-i; j++ {
			if s[j] > s[j+1] {
				aux = s[j]
				s[j] = s[j+1]
				s[j+1] = aux
			}
		}
	}
}

func main() {

}
