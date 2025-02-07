# Even or Odd - A Simple Go Project

This is a simple Go project that implements a deck of even or odd with basic functionalities.

## Features

- Create and shuffle a deck of even or odd.
- Draw even or odd from the deck.
- Save and load decks from files.
- Includes unit tests for deck operations.

## Installation

1. Clone the repository:
   ```sh
   git clone git@github.com:babykittenz/even-or-odd.git
   cd even-or-odd
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```

## Running the Project

To run the project, use the following command:

```sh
go run main.go
```

This will execute the `main.go` file and start the application.

## Project Structure

```
├── deck.go          # Deck logic implementation
├── deck_test.go     # Unit tests for deck functions
├── main.go          # Entry point for the application
├── go.mod           # Go module file
├── LICENSE          # License information
├── .gitignore       # Git ignore file
└── README.md        # Project documentation
```

## Usage

You can use this project to:

- Generate a new deck of even or odd.
- Shuffle the deck.
- Draw specific numbers of even or odd.
- Save and load decks from a file.

## Running Tests

Run the unit tests using:

```sh
go test
```

## License

This project is licensed under the MIT License.

## Contributions

Feel free to submit pull requests and report issues!

## Author

Created by babykittenz.
