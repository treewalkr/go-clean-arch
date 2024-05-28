# Clean Architecture with Fiber

This repository demonstrates the step-by-step development of a REST API using [Go Fiber](https://gofiber.io/) while progressively adopting the principles of Clean Architecture. The commits are ordered to show the evolution of the project, starting from a very simple Fiber application and gradually refactoring to a robust clean architecture implementation.

## Features

- **Progressive Refactoring**:
  - The repository starts with a basic Fiber setup.
  - Each commit introduces a new layer or feature, following clean architecture principles.
- **Clean Architecture Principles**:
  - Separation of concerns with layers such as `domain`, `application`, `interface` (`adapter`), and `infrastructure`.
  - Dependency inversion for testability and flexibility.

## Commit Progression

The project evolves through the following stages, with each commit reflecting a milestone:

1. **Simple Fiber Application**:
   - Basic setup with a single endpoint (`Hello, Fiber!`).
2. **Introducing Clean Architecture**:
   - Adding a `domain` layer with core models and interfaces.
   - Application layer (`service`) for business logic.
3. **Repository Abstraction**:
   - Implementing in-memory repository.
4. **Dependency Injection**:
   - Manual DI setup.
   - Automated DI using Google’s [Wire](https://github.com/google/wire).

## How to Run the Application

### Prerequisites

- [Go](https://golang.org/doc/install) (1.22 or later)
