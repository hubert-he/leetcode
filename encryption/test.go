package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

//实现明文的补码
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	//计算出需要补多少位，即需要padding的数目
	padding := blockSize - len(ciphertext)%blockSize
	//默认的blockSize=16(即采用16*8=128, AES-128长的密钥)
	//最少填充1个byte，如果原文刚好是blocksize的整数倍，则再填充一个blocksize
	//Repeat()函数的功能是把参数一 切片复制 参数二count个,然后合成一个新的字节切片返回
	padtext := bytes.Repeat([]byte{byte(padding)}, padding) //生成填充的文本
	//把补充的内容拼接到明文后面
	return append(ciphertext,padtext...)
}

//去除补码
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	//解密去补码时需取最后一个字节，值为m，则从数据尾部删除m个字节，剩余数据即为加密前的原文
	return origData[:(length - unpadding)]
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)//用0去填充
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}

func main () {
	// 此实例对应的是： echo 'hello backup' | openssl enc -e -aes-256-cbc -a -K '49e084d20fa540076c05a03faa8a4464' -iv '55e033412319e944' -v -p
	// 注意echo 传入的字符串末尾会带有\n，解密后务必清理
	// Load your secret key from a safe place and reuse it across multiple
	// NewCipher calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// knapsack like bcrypt or scrypt.
	key, _ := hex.DecodeString("49e084d20fa540076c05a03faa8a4464")
	iv, _ := hex.DecodeString("55e033412319e944")
	plaintext := PKCS5Padding([]byte("hello backup\n"), aes.BlockSize)
	fmt.Println(string(PKCS5Padding(key, 32)))
	ciphertext := ExampleNewCBCEncrypter(plaintext, ZeroPadding(key, 32), ZeroPadding(iv, aes.BlockSize))
	ExampleNewCBCDecrypter(ciphertext, ZeroPadding(key, 32), ZeroPadding(iv, aes.BlockSize))
}

func ExampleNewCBCEncrypter(plaintext []byte, key []byte, iv []byte) (ciphertext[]byte){



	// CBC mode works on blocks so plaintexts may need to be padded to the
	// next whole block. For an example of such padding, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. Here we'll
	// assume that the plaintext is already of the correct length.
	if len(plaintext)%aes.BlockSize != 0 {
		panic("plaintext is not a multiple of the block size")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.

	ciphertext = make([]byte, aes.BlockSize+len(plaintext))
	/*
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	*/
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.

	fmt.Printf("密文：%x\n", ciphertext)
	fmt.Println(base64.StdEncoding.EncodeToString(ciphertext[aes.BlockSize:]))
	return
}

func ExampleNewCBCDecrypter(ciphertext []byte, key []byte, iv []byte) (plaintext []byte) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	//iv := ciphertext[:aes.BlockSize]
	fmt.Printf("%x\n", ciphertext)
	ciphertext = ciphertext[aes.BlockSize:]

	// CBC mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(ciphertext, ciphertext)

	// If the original plaintext lengths are not a multiple of the block
	// size, padding would have to be added when encrypting, which would be
	// removed at this point. For an example, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. However, it's
	// critical to note that ciphertexts must be authenticated (i.e. by
	// using crypto/hmac) before being decrypted in order to avoid creating
	// a padding oracle.

	fmt.Printf("%s： %x\n", ciphertext, PKCS5UnPadding(ciphertext))
	return ciphertext
}

/*
解包装
*/
func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
