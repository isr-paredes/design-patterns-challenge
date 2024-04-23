
from abc import ABC, abstractmethod

class Shape(ABC):
    @abstractmethod
    def draw(self):
        pass

class Renderer(ABC):
    @abstractmethod
    def render(self, shape):
        pass

class Circle(Shape):
    def draw(self):
        return "Drawing a circle"

class Square(Shape):
    def draw(self):
        return "Drawing a square"

class VectorRenderer(Renderer):
    def render(self, shape):
        return f"Rendering {shape.draw()} in vector format"

class RasterRenderer(Renderer):
    def render(self, shape):
        return f"Rendering {shape.draw()} in raster format"

class Drawing:
    def __init__(self):
        self.shapes = []
        self.renderer = None

    def add_shape(self, shape):
        self.shapes.append(shape)

    def set_renderer(self, renderer):
        self.renderer = renderer

    def draw_shapes(self):
        if self.renderer is None:
            raise Exception("Renderer not set")
        for shape in self.shapes:
            print(self.renderer.render(shape))

def main():
    drawing = Drawing()
    drawing.add_shape(Circle())
    drawing.add_shape(Square())

    vector_renderer = VectorRenderer()
    drawing.set_renderer(vector_renderer)
    drawing.draw_shapes()

    raster_renderer = RasterRenderer()
    drawing.set_renderer(raster_renderer)
    drawing.draw_shapes()

if __name__ == "__main__":
    main()
