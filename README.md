# LeetCode Solutions in Go

Welcome to my LeetCode solutions repository, where I tackle problems using Go. This repository is organized to help you navigate through the solutions efficiently.

## Repository Structure

### `cli` Directory

The `cli` directory contains a module for creating starter code for daily problems and for solving a specific problem by its title slug.

### `problems` Directory

The `problems` directory stores the problems being solved. Each problem is stored in its own directory with the following convention:

```
ID.title-slug/
└── title-slug.go
```

Where:
- `ID`: The LeetCode problem ID.
- `title-slug`: The URL-friendly title of the problem.

### Getting Started

To use the CLI module for creating starter code:
1. Build the CLI
```bash
go build -o leetcode cli/*.go
```
2. Add the executable in your path and run the following command to see the usage
```bash
leetcode
```

### Contributing

If you'd like to contribute to this repository, you can:
- Add solutions to existing problems.
- Improve existing solutions for better performance or readability.
- Add new problems along with their solutions.

### Feedback

Feedback and suggestions are highly appreciated! If you have any ideas for improving the repository structure or the solutions themselves, feel free to open an issue or submit a pull request.

Happy coding! 🚀



