+++
title = "Sandro Lain"
#date = "2023-03-20T19:18:19Z"
#author = "Sandro Lain"
authorTwitter = "sandro_lain" #do not include @
#cover = "/photo.jpg"
#tags = ["sandro-lain"]
keywords = ["fullstack", "foundations"]
description = "Full-Stack"
showFullContent = false
readingTime = false
hideComments = false
color = "" #color from the theme settings
+++

# Foundations

## SOLID

- **Single Responsibility Principle (SRP)**: A class should have only one reason to change.
- **Open/Closed Principle (OCP)**: Software entities should be open for extension but closed for modification.
- **Liskov Substitution Principle (LSP)**: Subtypes must be substitutable for their base types.
- **Interface Segregation Principle (ISP)**: Clients should not be forced to depend on interfaces they do not use.
- **Dependency Inversion Principle (DIP)**: High-level modules should not depend on low-level modules. Both should depend on abstractions.

## Clean Code

- **DRY: “Don’t Repeat Yourself”**. It is a principle of software development aimed at reducing repetition of software patterns and code duplication in favor of abstractions and avoiding redundancy. The DRY principle states that "Every piece of knowledge must have a single, unambiguous, authoritative representation within a system". This principle has been formulated by Andy Hunt and Dave Thomas in their book "The Pragmatic Programmer". They apply it quite broadly to include database schemas, test plans, the build system, even documentation.
- **KISS: “Keep it simple, stupid!”** It is a design principle noted by the U.S. Navy in 1960. The KISS principle states that most systems work best if they are kept simple rather than made complicated; therefore, simplicity should be a key goal in design, and unnecessary complexity should be avoided. There are also other variations of this principle with the same meaning such as “Keep it short and simple” and "Keep it simple and straightforward".
- **“Less code, less bugs”** is a common saying in software development. The idea behind it is that the more code you write, the more opportunities there are for mistakes and errors to be introduced. By keeping your codebase small and concise, you can reduce the number of bugs and make it easier to maintain. This is one of the reasons why some developers prefer to use expressive languages like Python or Ruby over more verbose languages like C++ or Java2.

## Two Things

> There are only two hard things in Computer Science: cache invalidation and naming things.
>
> -- Phil Karlton

> There are only two hard problems in distributed systems:  
> 1: Exactly-once delivery  
> 2: Guaranteed order of messages  
> 1: Exactly-once delivery
>
> -- [Mathias Verraes](https://twitter.com/mathiasverraes/status/632260618599403520)

## Design Patterns

- **Creational Design Patterns** deal with object creation and initialization. They provide guidance about which objects are created for a given situation and are used to increase flexibility and to reuse existing code.
  - **Factory Method**: Creates objects with a common interface and lets a class defer instantiation to subclasses.
  - **Abstract Factory**: Creates a family of related objects.
  - **Builder**: A step-by-step pattern for creating complex objects, separating construction and representation.
  - **Prototype**: Supports the copying of existing objects without code becoming dependent on classes.
  - **Singleton**: Restricts object creation for a class to only one instance.
- **Structural Design Patterns** deal with class and object composition. They show how to assemble objects and classes into larger structures.
  - **Adapter**: Allows two incompatible interfaces to work together by wrapping its own interface around that of an already existing class.
  - **Bridge**: Decouples an abstraction from its implementation so that the two can vary independently.
  - **Composite**: Composes objects into tree structures to represent part-whole hierarchies.
  - **Decorator**: Adds behavior to an object dynamically without affecting the behavior of other objects from the same class.
  - **Facade**: Provides a simplified interface to a large body of code.
  - **Flyweight**: Reduces the cost of creating and manipulating a large number of similar objects.
  - **Proxy**: Provides a placeholder for another object to control access to it.
- **Behavioral Design Patterns** deal with the interaction between objects. They define the communication between objects and are concerned with the assignment of responsibilities between objects.
  - **Chain of Responsibility**: Passes a request along a chain of handlers until one of them handles the request.
  - **Command**: Encapsulates a request as an object, thereby allowing for the parameterization of clients with different requests.
  - **Interpreter**: Defines a representation for a grammar along with an interpreter to interpret sentences in the language.
  - **Iterator**: Provides a way to access the elements of an aggregate object sequentially without exposing its underlying representation.
  - **Mediator**: Defines an object that encapsulates how a set of objects interact.
  - **Memento**: Captures and externalizes an object’s internal state so that the object can be restored to this state later.
  - **Observer**: Defines a one-to-many dependency between objects so that when one object changes state, all its dependents are notified and updated automatically.
  - **State**: Allows an object to alter its behavior when its internal state changes.
  - **Strategy**: Defines a family of algorithms, encapsulates each one, and makes them interchangeable.
  - **Template** Method: Defines the skeleton of an algorithm in an operation, deferring some steps to subclasses.
  - **Visitor**: Represents an operation to be performed on the elements of an object structure without changing the classes of the elements on which it operates.

## Tests

- **Acceptance testing**: Verifying whether the whole system works as intended.
- **Integration testing**: Ensuring that software components or functions operate together.
- **Unit testing**: Validating that each software unit performs as expected. A unit is the smallest testable component of an application.
- **Functional testing**: Checking functions by emulating business scenarios, based on functional requirements.
- **Regression testing**: Testing done to verify that the system still works the way it did before.
- **Manual testing**: Human testers execute tests on a software application to identify errors.
- **Automated testing**: Uses software testing tools to execute tests on software applications.
- **Performance testing**: Testing the speed, scalability, and reliability of a system under a specific workload.
- **Security testing**: Testing to ensure that the system is protected against unauthorized access or attacks.
- **Usability testing**: Testing to ensure that the system is easy to use and meets user expectations.
- **Load testing**: Testing the system’s ability to handle a specific amount of load or traffic.
- **Stress testing**: Testing the system’s ability to handle extreme conditions, such as high traffic or limited resources.
- **Compatibility testing**: Testing to ensure that the system works correctly with different hardware, software, and operating systems.

## Microservices

- **Small, independent, and loosely coupled**: Microservices are small, independent, and loosely coupled.
- **Separate codebase**: Each service is a separate codebase, which can be managed by a small development team.
- **Independent deployment**: Services can be deployed independently.
- **Own data persistence**: Services are responsible for persisting their own data or external state.
- **Well-defined APIs**: Services communicate with each other by using well-defined APIs.
- **Polyglot programming**: Supports polyglot programming.
- **Single concern**: Microservice must be engaged in a single concern, making it a lot easier to handle and maintain.
- **Isolation**: Each microservice must be isolated from other microservices during development and testing.
- **Mobility**: It could be able to move.
- **Separate data storage**: It should have its own space where it can store all its data which is separate from other - **microservices.
- **Creation and destruction**: It could be created, destroyed, and recreated as per the needs.
- **Decentralized data management**: Each microservice should own its related domain data model and domain logic.
- **Loosely coupled services**: Create loosely coupled services so you have autonomy of development, deployment, and - **scale for each service.
- **Internal cohesion**: More important than the size of the microservice is the internal cohesion it must have and its independence from other services.
- **Long-term agility**: Microservices enable better maintainability in complex, large, and highly-scalable systems by letting you create applications based on many independently deployable services that each have granular and - **autonomous lifecycles.
- **Independent scaling**: Microservices can scale out independently.
- **Agile changes and rapid iteration**: The microservices approach allows agile changes and rapid iteration of each microservice because you can change specific, small areas of complex, large, and scalable applications.
- **High Cohesion and Low Coupling**: Microservices-based applications should have high cohesion and low coupling. Each service should do one thing and do it well, which means that the services should be highly cohesive. These services should also not depend on each other, which means they should have low coupling.
- **Discrete Boundaries**: In a discrete microservice architecture, each of the microservices is responsible for a specific task. When designing a microservices architecture, you should avoid having cross-functional dependencies between services.


## Bookshelf

- The Pragmatic Programmer -- *Andrew Hunt, David Thomas*
- Clean Code -- *Robert C. Martin*
- Clean Architecture -- *Robert C. Martin*
- Release It! -- *Michael T. Nygard*
- Don't make me think -- *Steve Krug*
- Building Micro-Frontends -- *Luca Mezzalira*
- Micro Frontends in Action -- *Michael Geers*
- Multithreaded Javascript -- *Thomas Hunter, Bryan English*
- Cloud Native Go -- *Matthew A. Titmus*
- Distributed Systems with Node.js -- *Thomas Hunter*
- Good Code, Bad Code -- *Tom Long*
- WebAssembly, The Definitive Guide -- *Brian Sletten*
- Thinking in Systems -- *Donella H. Meadows, Diana Wright*
- Interaction Design -- *Helen Sharp, Yvonne Rogers, Jenny Preece*
- Designing User Interfaces -- *Mike & Diana Malewicz*
