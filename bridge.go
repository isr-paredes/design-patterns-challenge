package main

import "fmt"

// Interfaces
type Shape interface {
    Draw()
}

type Renderer interface {
    RenderCircle()
    RenderSquare()
}

// Bridge between shape and renderer
type AbstractShapeRenderer struct {
    renderer Renderer
}

func (asr *AbstractShapeRenderer) SetRenderer(renderer Renderer) {
    asr.renderer = renderer
}

// Circle
type Circle struct {
    AbstractShapeRenderer
}

func (c *Circle) Draw() {
    c.renderer.RenderCircle()
}

// Square
type Square struct {
    AbstractShapeRenderer
}

func (s *Square) Draw() {
    s.renderer.RenderSquare()
}

// VectorRenderer
type VectorRenderer struct{}

func (vr *VectorRenderer) RenderCircle() {
    fmt.Println("Rendering Circle in Vector format")
}

func (vr *VectorRenderer) RenderSquare() {
    fmt.Println("Rendering Square in Vector format")
}

// RasterRenderer
type RasterRenderer struct{}

func (rr *RasterRenderer) RenderCircle() {
    fmt.Println("Rendering Circle in Raster format")
}

func (rr *RasterRenderer) RenderSquare() {
    fmt.Println("Rendering Square in Raster format")
}

//Drawing
type Drawing struct {
    shapes []Shape
}

func (d *Drawing) AddShape(shape Shape) {
    d.shapes = append(d.shapes, shape)
}

func (d *Drawing) DrawShapes() {
    for _, shape := range d.shapes {
        shape.Draw()
    }
}

func main() {
    drawing := Drawing{}

    circle := &Circle{}
    square := &Square{}

    vectorRenderer := &VectorRenderer{}
    rasterRenderer := &RasterRenderer{}

    circle.SetRenderer(vectorRenderer)
    square.SetRenderer(rasterRenderer)

    drawing.AddShape(circle)
    drawing.AddShape(square)

    drawing.DrawShapes()
}
