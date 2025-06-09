package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/jnst/base58"
)

func main() {
	var (
		help     = flag.Bool("h", false, "show help")
		helpLong = flag.Bool("help", false, "show help")
		file     = flag.String("f", "", "input file")
	)
	flag.Parse()

	if *help || *helpLong {
		showHelp()
		return
	}

	args := flag.Args()
	if len(args) == 0 {
		showHelp()
		os.Exit(1)
	}

	command := args[0]
	switch command {
	case "encode":
		if err := encodeCommand(*file, args[1:]); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	case "decode":
		if err := decodeCommand(*file, args[1:]); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	case "help":
		showHelp()
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		showHelp()
		os.Exit(1)
	}
}

func showHelp() {
	fmt.Println("base58 - Base58 encoding and decoding tool")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  base58 encode [data]        Encode data as base58")
	fmt.Println("  base58 encode -f <file>     Encode file contents as base58")
	fmt.Println("  base58 decode [base58]      Decode base58 string")
	fmt.Println("  base58 decode -f <file>     Decode base58 from file")
	fmt.Println("  base58 help                 Show this help")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -f <file>     Read input from file")
	fmt.Println("  -h, --help    Show help")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  echo 'Hello World' | base58 encode")
	fmt.Println("  base58 encode 'Hello World'")
	fmt.Println("  base58 encode -f input.txt")
	fmt.Println("  base58 decode JxF12TrwUP45BMd")
	fmt.Println("  echo 'JxF12TrwUP45BMd' | base58 decode")
}

func encodeCommand(filename string, args []string) error {
	var input []byte
	var err error

	if filename != "" {
		input, err = os.ReadFile(filename)
		if err != nil {
			return fmt.Errorf("reading file: %w", err)
		}
	} else if len(args) > 0 {
		input = []byte(strings.Join(args, " "))
	} else {
		input, err = readStdin()
		if err != nil {
			return fmt.Errorf("reading stdin: %w", err)
		}
	}

	encoded := base58.Encode(input)
	fmt.Println(encoded)
	return nil
}

func decodeCommand(filename string, args []string) error {
	var input string
	var err error

	if filename != "" {
		data, err := os.ReadFile(filename)
		if err != nil {
			return fmt.Errorf("reading file: %w", err)
		}
		input = strings.TrimSpace(string(data))
	} else if len(args) > 0 {
		input = strings.Join(args, " ")
	} else {
		data, err := readStdin()
		if err != nil {
			return fmt.Errorf("reading stdin: %w", err)
		}
		input = strings.TrimSpace(string(data))
	}

	decoded, err := base58.Decode(input)
	if err != nil {
		return fmt.Errorf("decoding: %w", err)
	}

	os.Stdout.Write(decoded)
	return nil
}

func readStdin() ([]byte, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil && err != io.EOF {
		return nil, err
	}
	return []byte(strings.Join(lines, "\n")), nil
}