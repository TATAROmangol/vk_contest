package structs

type Point struct{
    Y int 
    X int
    Weight int 
	Parent *Point
	Cost int
}