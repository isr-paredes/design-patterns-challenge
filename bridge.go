package main
import "fmt"

//Interfaces
type Shape interface {
    Draw(renderer Renderer)
}
type Renderer interface {
    RenderCircle()
    RenderSquare()
}

//Bridge between shape and renderer
type AbstractShapeRenderer struct {
    renderer Renderer
}

func (asr *AbstractShapeRenderer) SetRenderer(renderer Renderer) {
    asr.renderer = renderer
}

//Circle
type Circle struct {
    AbstractShapeRenderer
}
func (c *Circle) Draw() {
    c.renderer.RenderCircle()
}

//Square
type Square struct {
    AbstractShapeRenderer
}
func (s *Square) Draw() {
    s.renderer.RenderSquare()
}


type VectorRenderer struct{}

//Methods
func (rr *RasterRenderer) RenderCircle() {
    fmt.Println("Rendering Circle in Raster format")
}
func (rr *RasterRenderer) RenderSquare() {
    fmt.Println("Rendering Square in Raster format")
}


func main() {
	circle := &Circle{}
	square := &Square{}

	vectorRenderer := &VectorRenderer{}
	rasterRenderer := &RasterRenderer{}

	circle.SetRenderer(vectorRenderer)
	square.SetRenderer(rasterRenderer)

	circle.Draw()
	square.Draw()
}