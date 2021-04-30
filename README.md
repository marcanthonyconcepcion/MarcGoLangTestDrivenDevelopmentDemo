# Marc Concepcion's Geometry Demo codes to demonstrate Test-Driven Development in Go Language

This is a Go implementation of the _GeometryDemoPython_ TDD that I had made.

This is a simple Go project that contains _types and functions_ that represent the geometric measurement properties of some well-known two-dimensional and three-dimensional shapes.

This demo generically classifies the shapes into:

## Two-Dimensional Shapes

For each of the Two-Dimensional Shape, its area and perimeter properties are defined.

Here are the two-dimensional shapes described by the project:

1. Circle * 
2. Rectangle
3. Square
4. Triangle

*Please take note that I treat the circumference of a circle as its perimeter based from this statement from Wikipedia: "The perimeter of a circle or ellipse is called its circumference." Source: https://en.wikipedia.org/wiki/Perimeter
Here are the three-dimensional shapes described by the project:

## Three-Dimensional Shapes.

For each of the Three-Dimensional Shape, its surface area and volume properties are defined.

1. Prism
2. Cube
3. Rectangular Prism
4. Triangular Prism
5. Cylinder **
6. Sphere

** Please take note that I treat Cylinder as a Prism based from this statement from Wikipedia: "A solid circular cylinder can be seen as the limiting case of an n-gonal prism where n approaches infinity." Source: https://en.wikipedia.org/wiki/Cylinder#Prisms

## Objective

The objective of this project is to perform test-driven development to define each shape into production.

## About Go Lang

Unlike the similar TDD demo projects I made on Python, PHP and Raku Perl, creating this TDD demo project with Go language requires an almost totally new perspective in coding. Go language does not have well-known _classes_, _exceptions_, _inheritance_ and elementary _object-oriented programming_ that I am very very much accustomed to.

Inspite of the '_culture shock_', I successfully managed to practice the same spirit of TDD and SOLID-style programming using Go language. I can still use the superior power of the simple _interface_. I can still implement an IS-A relationship inspite of not having classes and inheritance (through _composition_ which actually enforces a HAS-A relationship and therefore is a workaround).

Go lang's dual function return (the other for error return) is actually a better error handling technique than using exception which is actually a _goto_ in disguise.

Through this Go lang's TDD demo, I realize that Go lang actually simplifies programming. It challenges and shatters our well-known and well-upheld ideology on good coding making us realize that there could be a significantly different way of coding better. 

By setting aside C++ and the popular languages right after (C#, Java, PHP, Python, Perl etc.) and going back the C language to start a new progressive path on good programming,

Go actually opens my eyes and expands my mind on exploring a different paradigm of coding better without sacrificing basic principles of good programming.

There is a big promise on Go lang. As a programming language made by Google, I expect lots of support and an instant wealth of frameworks and modules that would quickly promote the use of this young language.
