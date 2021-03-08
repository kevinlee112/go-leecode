package _01_400


func findCircleNum(isConnected [][]int) int {
	unionSet := make(map[int]int)
	for i := 0; i < len(isConnected); i++ {
		unionSet[i] = -1
	}
	for i := 0; i <len(isConnected);i++ {
		connectedCities := isConnected[i]
		for j := i + 1; j < len(connectedCities); j++ {

		}

	}
}