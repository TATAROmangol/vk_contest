package dijkstra

import (
	"container/heap"
	"fmt"
	"strings"
	. "vk_contest/internal/structs"
)

var moves = [][]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}

func AddNeighbors(notOpenPoints *PriorityQueue, maze [][]*Point, start *Point){
    lenX := len(maze[0])
    for _, m := range moves {
        newY := start.Y + m[0]
        newX := start.X + m[1]
        if newY < 0 || newY >= len(maze) || newX < 0 || newX >= lenX {
            continue
        }

        nextPoint := maze[newY][newX] 
        if nextPoint.Weight == 0{
            continue
        }

        newCost := start.Cost + nextPoint.Weight
        if nextPoint.Cost == -1 {
            nextPoint.Parent = start 
            nextPoint.Cost = newCost
            heap.Push(notOpenPoints, nextPoint) 
            continue
        }
        if nextPoint.Cost < newCost {
            nextPoint.Cost = newCost
        }
    }

}

func FindPath(maze [][]*Point, start, end *Point) *Point{
    notOpenPoints := &PriorityQueue{}
    heap.Init(notOpenPoints)

    start.Cost = start.Weight
    heap.Push(notOpenPoints, start)
    for notOpenPoints.Len() > 0{
        next := heap.Pop(notOpenPoints).(*Point)
        if next.X == end.X && next.Y == end.Y{
            return next
        }

        AddNeighbors(notOpenPoints, maze, next)
    }

    return nil
}

func GetPathByDijkstra(maze [][]*Point, start, end *Point) string{
	peek := FindPath(maze, start, end)
	if peek == nil{
		return "\nПуть не найден"
	}

	builder := new(strings.Builder)
    builder.WriteString("\n")
	var result []Point
    for p := peek; p != nil; p = p.Parent {
        result = append(result, *p)
    }

    for i := len(result) - 1; i >= 0; i-- {
        builder.WriteString(fmt.Sprintf("(%d, %d)\n", result[i].Y, result[i].X))
    }
	builder.WriteString(".\n")
	
	return builder.String()
}
