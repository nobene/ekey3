package main

import (
	"fmt"
	"log"
	"strings"
	"io/ioutil"
	"flag"
//	"golang.org/x/term"
	"github.com/tyler-smith/go-bip39"
	"github.com/miguelmota/go-ethereum-hdwallet"
)

func main() {
	from := flag.String("f", "", "filename to read")
	flag.Parse()
	data, _ := ioutil.ReadFile(*from)
	file := string(data)
	line := 0	
	temp := strings.Split(file, "\n")
	for _, item := range temp {
	if len(item) < 2 { break }
	num := strings.Fields(item)[0]
	item = strings.Join(strings.Fields(item)[1:], " ")
//	fmt.Println("[",num,"]\t",item)
	step(num, item)
	line++
	}
}

func step(num string, mnemonic string) {
	if !bip39.IsMnemonicValid(mnemonic) {
		fmt.Println("        " + num + " :  Invalid Mnemonic !!!")
		fmt.Println()
		return
	}
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
		fmt.Println("        " + num + " : hdwallet's creation failed !!!" )
		return
	}
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("      +" + num + " "  + account.Address.Hex())

//	path = hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/1")
//	account, err = wallet.Derive(path, false)
//	if err != nil {
//		log.Fatal(err)
//	}

//	fmt.Println(account.Address.Hex()) // 0x8230645aC28A4EdD1b0B53E7Cd8019744E9dD559
}
