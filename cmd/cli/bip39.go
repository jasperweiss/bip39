package main

import (
	"bip39"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"
)

var (
	version   string
	commitID  string
	buildDate string
)

func wordsToEntropyBits(wordCount int) (entropyBits int, err error) {
	var exists bool

	// Map specifying supported word counts and their equivalent entropy bits
	wordToBits := map[int]int{
		12: 128,
		24: 256,
	}

	// Retrieve the entropy bits for the given word count from the map
	entropyBits, exists = wordToBits[wordCount]
	if !exists {
		var allowedWords []string
		for key := range wordToBits {
			allowedWords = append(allowedWords, strconv.Itoa(key))
		}

		sort.Strings(allowedWords)

		// If the word count is not supported, return an error
		return 0, fmt.Errorf("unsupported word count. Allowed words: %s", strings.Join(allowedWords, ", "))
	}

	// Return the corresponding entropy bits for the provided word count
	return entropyBits, nil
}

func generateMnemonicAction(cCtx *cli.Context) error {
	wordCount := cCtx.Int("words")
	if wordCount != 12 && wordCount != 24 {
		return fmt.Errorf("unsupported word count. Allowed words: 12, 24")
	}

	bitSize := 128
	if wordCount == 24 {
		bitSize = 256
	}

	entropy, err := bip39.NewEntropy(bitSize)
	if err != nil {
		return err
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return err
	}

	fmt.Println(mnemonic)

	return nil
}

func existingMnemonicAction(cCtx *cli.Context) error {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return fmt.Errorf("failed to read input")
	}

	mnemonic := scanner.Text()

	if !bip39.IsMnemonicValid(mnemonic) {
		return fmt.Errorf("mnemonic is not valid")
	}

	fmt.Println("mnemonic is valid")

	return nil
}

func main() {
	app := &cli.App{
		Usage: "Generation and verification of mnemonics in BIP39 standard",
		Commands: []*cli.Command{
			{
				Name:  "generate",
				Usage: "BIP39 mnemonic generation\n--words value\tWord count (default: 24)",
				Flags: []cli.Flag{
					&cli.IntFlag{Name: "words", Value: 24},
				},
				Action: func(cCtx *cli.Context) error {
					return generateMnemonicAction(cCtx)
				},
			},
			{
				Name:  "existing",
				Usage: "Check existing BIP39 mnemonic\n",
				Action: func(cCtx *cli.Context) error {
					return existingMnemonicAction(cCtx)
				},
			},
			{
				Name:    "version",
				Usage:   "Print the version\n",
				Aliases: []string{"v"},
				Action: func(cCtx *cli.Context) error {
					fmt.Printf("Version:\t%s\n", version)
					fmt.Printf("Git Commit:\t%s\n", commitID)
					fmt.Printf("Build Date:\t%s\n", buildDate)

					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
}
