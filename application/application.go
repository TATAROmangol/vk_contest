package application

import (
	"bufio"
	"fmt"
	"io"
	. "vk_contest/internal/dijkstra"
	. "vk_contest/internal/errors"
	. "vk_contest/internal/structs"
)

type Application struct{
	in *bufio.Reader
	out *bufio.Writer
	outErr *bufio.Writer
}

func NewApplication(in *bufio.Reader, out *bufio.Writer, outErr *bufio.Writer) *Application{
	return &Application{in, out, outErr}
}

func (a *Application) Run(){
	for {
		maze, start, end, err := readInput(a.in)
		if err == io.EOF{
			a.out.WriteString("\nКонец ввода.")
			return
		}
		if err != nil {
			a.outErr.WriteString(err.Error())
			return
		}

		res := GetPathByDijkstra(maze, start, end)
		a.out.WriteString(res)

		a.out.Flush()
	}
}

func readInput(in *bufio.Reader) ([][]*Point, *Point, *Point, error){
    var ySize, xSize int 
    n, err := fmt.Fscanln(in, &ySize, &xSize)
	if err == io.EOF{
		return make([][]*Point, 0), &Point{}, &Point{}, io.EOF
	}
    if n != 2 || err != nil || ySize < 1 || xSize < 1{
        return make([][]*Point, 0), &Point{}, &Point{}, ErrorUnknownInputSize
    }

    maze := make([][]*Point, ySize)
    for i := range maze {
        maze[i] = make([]*Point, xSize)
    }

    for y := 0; y < ySize; y++{
        for x := 0; x < xSize; x++{
            var n, weight int 
            var err error
            if x < xSize - 1{
                n, err = fmt.Fscan(in, &weight)
            } else {
                n, err = fmt.Fscanln(in, &weight)
            }
            
            if n != 1 || err != nil{
                return maze, &Point{}, &Point{}, ErrorUnknownInputMaze
            }
            maze[y][x] = &Point{y,x,weight, nil, -1}
        }
    }

    var yStart, xStart, yEnd, xEnd int
    n, err = fmt.Fscanln(in, &yStart, &xStart, &yEnd, &xEnd)
    if n != 4 || err != nil{
        return make([][]*Point, 0), &Point{}, &Point{}, ErrorUnknownInputStartEnd
    }

    return maze, maze[yStart][xStart], maze[yEnd][xEnd], nil
}