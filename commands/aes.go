// MIT License

// Copyright (c) The RAI Authors

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
package commands

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/retail-ai-inc/bean/aes"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// aesEncryptCmd represents the aes:encrypt command.
	aesEncryptCmd = &cobra.Command{
		Use:   "aes:encrypt",
		Short: "Encrypt a plaintext using AES algorithm",
		Long: `This command will encrypt a text using AES algo and print it as a base64 string. The secret a.k.a passphrase is also encoded 
		as a base64 string in env.json file.`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires a plain text to encrypt")
			}

			return nil
		},
		Run: aesEncrypt,
	}
)

var (
	// aesDecryptCmd represents the aes:decrypt command.
	aesDecryptCmd = &cobra.Command{
		Use:   "aes:decrypt",
		Short: "Decrypt a base64 encrypted text using AES algorithm",
		Long: `This command will decrypt an encrypted base64 text using AES algo and print it as a plain text. The secret a.k.a passphrase is also encoded 
		as a base64 string in env.json file.`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires an encrypted text to decrypt")
			}

			return nil
		},
		Run: aesDecrypt,
	}
)

func init() {
	rootCmd.AddCommand(aesEncryptCmd)
	rootCmd.AddCommand(aesDecryptCmd)
}

func aesEncrypt(cmd *cobra.Command, args []string) {

	plaintext := strings.Join(args, "")

	secret := viper.GetString("secret")
	encryptedText, err := aes.BeanAESEncrypt(secret, plaintext)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println(encryptedText)
}

func aesDecrypt(cmd *cobra.Command, args []string) {

	secret := viper.GetString("secret")
	decryptedText, err := aes.BeanAESDecrypt(secret, args[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println(decryptedText)
}
