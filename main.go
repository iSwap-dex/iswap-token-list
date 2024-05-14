package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type TokenInfo struct {
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	Decimals    int64  `json:"decimals"`
	Description string `json:"description"`
	Website     string `json:"website"`
	Twitter     string `json:"twitter"`
	Tokens      map[string]struct {
		Address string `json:"address"`
	} `json:"tokens"`
}

type TokenList struct {
	Name      string   `json:"name"`
	LogoURI   string   `json:"logoURI"`
	Keywords  []string `json:"keywords"`
	Timestamp string   `json:"timestamp"`
	Tokens    []Token  `json:"tokens"`
	Version   Version  `json:"version"`
}

type Token struct {
	ChainID  int64  `json:"chainId"`
	Address  string `json:"address"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals int64  `json:"decimals"`
	LogoURI  string `json:"logoURI"`
}

type Version struct {
	Major int64 `json:"major"`
	Minor int64 `json:"minor"`
	Patch int64 `json:"patch"`
}

var chins map[string]int64
var tokens []Token

var dataDir = "./data"
var chainsFile = "chains.json"
var baseUrl = "https://raw.githubusercontent.com/space-group/token-list/master/"

func main() {
	fileBytes, err := os.ReadFile(chainsFile)
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = json.Unmarshal(fileBytes, &chins)
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = filepath.Walk(dataDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if info.Name() == "data.json" {
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			var tokenInfo TokenInfo
			if err := json.Unmarshal(data, &tokenInfo); err != nil {
				return err
			}

			logoFile := findLogoFile(filepath.Dir(path))

			for chainId, token := range tokenInfo.Tokens {
				outputToken := Token{
					ChainID:  chins[chainId],
					Address:  token.Address,
					Name:     tokenInfo.Name,
					Symbol:   tokenInfo.Symbol,
					Decimals: tokenInfo.Decimals,
					LogoURI:  baseUrl + "data/" + tokenInfo.Symbol + "/" + logoFile,
				}
				tokens = append(tokens, outputToken)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
	}

	output := TokenList{
		Name:      "iswap token list",
		LogoURI:   "",
		Keywords:  nil,
		Timestamp: time.Now().Format(time.RFC3339),
		Version: Version{
			Major: 0,
			Minor: 0,
			Patch: 0,
		},
		Tokens: tokens,
	}

	outputJSON, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = os.WriteFile("token.list.json", outputJSON, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}

	fmt.Println("JSON output written to token.list.json")

}

func findLogoFile(dir string) string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return ""
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".svg" || filepath.Ext(file.Name()) == ".png" {
			return file.Name()
		}
	}
	return ""
}
