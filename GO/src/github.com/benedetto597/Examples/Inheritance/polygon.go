package main
 
import "fmt"
 
type Polygon struct {
    Sides int
}
 
func (p *Polygon) NSides() int {
    return p.Sides
}
 
type Triangle struct {
    Polygon // anonymous field
}
 
func main() {
    t := Triangle{
        Polygon{
            Sides: 3,
        },
    }
    fmt.Println(t.NSides()) // 3
}