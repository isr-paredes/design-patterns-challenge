# design-patterns-challenge
Set of Problems

Adapter Pattern: Problem: You have a legacy system that provides data in a format incompatible with a new client application. Implement an adapter that converts the legacy data format into a format compatible with the client application without modifying the existing system.


function getLegacyData() {
  // Simulate fetching data from a legacy system
  legacyData = {
    "employee_name": "John Doe",
    "employee_id": 12345,
    "department_code": "XYZ",
    "salary": 50000
    };
  return legacyData;
}

Adapt it into a new format that uses simpler keys like `"name"`, `"id"`, `"department"`, and `"pay"`.

Bridge Pattern: Problem: You are developing a drawing application that supports different shapes (e.g., circles, squares) and rendering modes (e.g., vector, raster). Use the bridge pattern to decouple the abstraction (shapes) from the implementation (rendering modes) so that new shapes and rendering modes can be added independently.

// Define a Shape interface
interface Shape {
    draw();
}

// Implement Circle and Square classes that implement the Shape interface
class Circle implements Shape {
    draw() {
        // Draw a circle
    }
}

class Square implements Shape {
    draw() {
        // Draw a square
    }
}

// Define a Renderer interface
interface Renderer {
    render(shape);
}

// Implement VectorRenderer and RasterRenderer classes that implement the Renderer interface
class VectorRenderer implements Renderer {
    render(shape) {
        // Render the shape in vector format
    }
}

class RasterRenderer implements Renderer {
    render(shape) {
        // Render the shape in raster format
    }
}

// Define a Drawing class to manage shapes and rendering
class Drawing {
    shapes = new List();

    addShape(shape) {
        shapes.add(shape);
    }

    drawShapes(renderer) {
        foreach (shape in shapes) {
            renderer.render(shape);
        }
    }
}

// Usage example
drawing = new Drawing();
drawing.addShape(new Circle());
drawing.addShape(new Square());

vectorRenderer = new VectorRenderer();
drawing.drawShapes(vectorRenderer);

Introduce the Bridge Pattern by creating an abstract class that connects shapes (Circle and Square) with renderers (VectorRenderer and RasterRenderer).

Composite Pattern: Problem: Create a hierarchical structure for managing a company's organizational chart. Each node in the hierarchy represents an employee or a department. Implement the composite pattern to treat individual employees and departments uniformly as part of the organization.

Facade Pattern: Problem: You are developing a smart home automation system that controls various devices such as lights, thermostats, and security cameras. Each device has its own complex interface and set of commands for operation. Implement a facade that provides a simplified interface for users to interact with multiple devices seamlessly without needing to understand the intricacies of each device's individual commands.

Flyweight Pattern: Problem: You are developing a virtual city simulation game where thousands of buildings need to be rendered on the screen. Each building has unique characteristics such as height, color, and style. Implement a flyweight pattern to optimize memory usage and improve rendering performance by sharing common characteristics (e.g., textures, materials) among similar buildings while allowing each building to retain its unique properties.

Decorator Pattern: Problem: Extend a basic car rental system to support additional features such as GPS navigation, child seat, and insurance options. Apply the decorator pattern to dynamically add and combine these optional features to a base car rental object without altering its core functionality.

// Define a basic Car class representing rental cars
class Car {
    model;
    rentalPrice;
    isAvailable;

    constructor(model, rentalPrice) {
        this.model = model;
        this.rentalPrice = rentalPrice;
        this.isAvailable = true;
    }

    rent() {
        if (isAvailable) {
            isAvailable = false;
            print("Rented " + model + " for $" + rentalPrice);
        } else {
            print("Car " + model + " is not available for rent.");
        }
    }

    returnCar() {
        isAvailable = true;
        print("Returned " + model);
    }
}

// Usage example
car1 = new Car("Toyota Camry", 50);
car2 = new Car("Honda Civic", 40);

car1.rent(); // Rented Toyota Camry for $50
car2.rent(); // Rented Honda Civic for $40
car1.rent(); // Car Toyota Camry is not available for rent.

car1.returnCar(); // Returned Toyota Camry
car2.rent(); // Rented Honda Civic for $40

In this pseudocode, the Car class represents basic rental cars with properties like model, rental price, and availability. Users can rent a car if it's available and return it when they're done. However, this basic implementation doesn't allow for adding additional features or functionalities to cars dynamically, which is where the Decorator Pattern comes in handy.

Proxy Pattern: Problem: You are developing a virtual art gallery application where users can explore and view various artworks. However, some high-resolution artworks consume a significant amount of memory and take time to load. Implement a proxy pattern to improve performance by using proxy objects to lazily load high-resolution artworks only when requested by the user, thereby reducing memory usage and optimizing the user experience.


// Define an Artwork interface for common artwork operations
interface Artwork {
    display();
}

// Implement ConcreteArtwork that represents individual artworks
class ConcreteArtwork implements Artwork {
    name;
    image; // Actual high-resolution image

    constructor(name, image) {
        this.name = name;
        this.image = image;
    }

    display() {
        // Display the artwork
        print("Displaying " + name);
        // Render the high-resolution image
    }
}

// Define an ArtGallery class to manage artworks
class ArtGallery {
    artworks;

    constructor() {
        this.artworks = [];
    }

    addArtwork(artwork) {
        artworks.add(artwork);
    }

    displayAllArtworks() {
        for each (artwork in artworks) {
            artwork.display();
        }
    }
}

// Usage example
gallery = new ArtGallery();

// Create and add artworks to the gallery
artwork1 = new ConcreteArtwork("Mona Lisa", loadHighResolutionImage("Mona Lisa"));
artwork2 = new ConcreteArtwork("Starry Night", loadHighResolutionImage("Starry Night"));

gallery.addArtwork(artwork1);
gallery.addArtwork(artwork2);

// Display all artworks in the gallery
gallery.displayAllArtworks();


In this pseudocode, the ArtGallery class manages individual artworks (ConcreteArtwork) directly without using the Proxy Pattern. Each artwork is responsible for displaying its high-resolution image, and the ArtGallery class provides methods to add artworks and display them all at once.

