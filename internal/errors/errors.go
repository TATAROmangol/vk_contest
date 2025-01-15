package errors

import "fmt"

var (
	ErrorUnknownInputSize = fmt.Errorf("Неправильный формат ввода размеров лабиринта")
	ErrorUnknownInputMaze = fmt.Errorf("Неправильный формат ввода лабиринта")
	ErrorUnknownInputStartEnd = fmt.Errorf("Неправильный формат ввода начальной и конечной точек")
)