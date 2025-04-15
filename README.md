# Mars Rover Technical Challenge Documentation

## Overview

The Mars Rover Technical Challenge is a command-line application that simulates the navigation of rovers on a
plateau surface on Mars. The application processes input to define the plateau's dimensions, the initial positions
of the rovers, and their movement instructions. It then calculates the final positions and orientations of the 
rovers after executing the instructions.

---

## 🏁 How to Use

### 🔧 Prerequisites

- **Go**: Version 1.21 or higher.
- **Make**: Optional but recommended for build and test automation.

---

### 🚀 Running the Application

1. **Build the Application**:
   Use the `make` command to build the application:
   ```sh
   make build
   ```
   This will compile the application into the bin directory:
   ```sh
   bin/mars-rover
   ```

2. **Run the Application**:
   Execute the compiled binary:
   ```sh
   ./bin/mars-rover
   ```
    
    Or just run the make file
    ```sh
    make run
    ```

   Follow the prompts to input the plateau size, rover start positions, and movement instructions.

   Example input:
   ```
   Enter plateau size (e.g. '5 5'): 5 5
   Enter rover start position (e.g. '1 2 N'): 1 2 N
   Enter movement instructions (e.g. 'LMLMLMLMM'): LMLMLMLMM
   Add another rover? (y/n): y
   Enter rover start position (e.g. '1 2 N'): 3 3 E
   Enter movement instructions (e.g. 'LMLMLMLMM'): MMRMMRMRRM
   Add another rover? (y/n): n
   ```
   Example output:
   ```
   1 3 N
   5 1 E
   ```

---

## 📂 Folder Structure

The project is organized as follows:

```
.
├── Makefile                    # Dev/build/test automation
├── cmd/
│   └── mars-rover/
│       └── main.go             # CLI entry point
├── internal/                   # Application packages (non-public)
│   ├── app/                    # Orchestration logic (RunMission, etc.)
│   ├── direction/              # Compass direction enum + logic
│   ├── instruction/            # Movement instruction enum + parsing
│   ├── input/                  # User input handling via stdin
│   ├── parser/                 # Raw string input to domain model parser
│   ├── plateau/                # Plateau size + boundary logic
│   └── rover/                  # Rover state, movement, and execution
└── bin/                        # Compiled binary output
```

---

## 🧪 Running Tests

The project includes a comprehensive suite of unit tests to ensure correctness.

1. **Run All Tests**:
   Use the `make` command to execute all tests:
   ```sh
   make test
   ```

2. **Watch Tests**:
   Automatically re-run tests on file changes:
   ```sh
   make test-watch
   ```

3. **Run Tests Manually**:
   Use the `go test` command to run tests for specific packages:
   ```sh
   go test ./...
   ```

---

## Example Input and Output

### Input:
```
5 5
1 2 N
LMLMLMLMM
3 3 E
MMRMMRMRRM
```

### Output:
```
1 3 N
5 1 E
```
