## CLI for generation and verification of mnemonics in BIP39
![BIP39 mnemonics generation tool](assets/images/generate_example_1.png)

### Build
    go build cmd/cli/bip39.go

### Install
    sudo install -t /usr/local/bin bip39

### BIP39 mnemonic generation
    bip39 generate

    --words value   Word count (default: 24)

    Example: bip39 generate --words 12

### Check existing BIP39 mnemonic
    bip39 existing

    Example: echo "word1 word2 ... word12" | bip39 existing
    Outputs: "mnemonic is valid" if valid, error otherwise.

### Thanks
Thanks to Tyler Smith for providing the implementation of [BIP39](https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki) in [Golang](https://github.com/tyler-smith/go-bip39) that allowed us to create this tool!

### License
This BIP39 tool is released under the terms of the MIT license.  
See LICENSE for more information or see https://opensource.org/licenses/MIT.
