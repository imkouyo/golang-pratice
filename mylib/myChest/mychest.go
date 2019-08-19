package myChest

type Move struct {
	x int
	y int
}
var MoveList = [8]Move {
	{x: -2, y: 1},
	{x: -1, y: 2},
	{x: 1, y: 2},
	{x: 2, y: 1},
	{x: 2, y: -1},
	{x: 1, y: -2},
	{x: -1, y: -2},
	{x: -2, y: -1},
}


func KnightGoCheckerboard(x, y int) {
	var mask = createMask()
	var currentPosition = setPosition(x, y)
	mask = setMask(currentPosition.x, currentPosition.y, mask)

}

func createMask()[][] bool {
	mask := make( [][]bool, 8)
	for i := range mask {
		mask[i] = make( []bool, 8)
	}
	return mask
}

func setPosition(x, y int) Move {
	return Move{ x: x, y: y }
}

func setMask(x, y int, mask [][]bool)[][] bool {
	mask[x][y] = true
	return mask
}

func checkNext(mask [][]bool, position Move) []bool {
	nextStep := make([]bool, 8)
	for i:= range MoveList{
		if position.x + MoveList[i].x < 8 && position.x + MoveList[i].x >= 0 && position.y + MoveList[i].y < 8 && position.y + MoveList[i].y >= 0 {
			nextStep[i] = true
		}
	}
	return nextStep
}

func checkNextNext(pathList []bool, mask [][]bool, position Move) {

}

