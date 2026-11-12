
// Lab 8 - Task 1
// Demonstrating OOP concepts: Inheritance and Polymorphism

open System

// Base class
type Shape() =
    abstract member Area : unit -> float

    // Default implementation of virtual method
    default _.Area() = 0.0

// Derived class: Circle
type Circle(radius: float) =
    inherit Shape()

    override _.Area() =
        Math.PI * radius * radius

// Derived class: Rectangle
type Rectangle(length: float, width: float) =
    inherit Shape()

    override _.Area() =
        length * width

// Creating objects
let circle = Circle(5.0)
let rectangle = Rectangle(4.0, 6.0)

// Calling methods directly
printfn "Circle Area: %.2f" (circle.Area())
printfn "Rectangle Area: %.2f" (rectangle.Area())

// Demonstrating runtime polymorphism
let shapes: Shape list =
    [ circle :> Shape; rectangle :> Shape ]

printfn "\nDemonstrating Polymorphism:"

for shape in shapes do
    printfn "Area: %.2f" (shape.Area())
