package main

import (
"fmt"
scp "github.com/bramvdbogaerde/go-scp"
"github.com/bramvdbogaerde/go-scp/auth"
"golang.org/x/crypto/ssh"
"os"
)

func main() {
	// Use SSH key authentication from the auth knapsack
	// we ignore the host key in this example, please change this if you use this library
	//clientConfig, _ := auth.PrivateKey("username", "/path/to/rsa/key", ssh.InsecureIgnoreHostKey())
	clientConfig, _ := auth.PasswordKey("hubert", "asdf@321", ssh.InsecureIgnoreHostKey(),)
	// For other authentication methods see ssh.ClientConfig and ssh.AuthMethod

	// Create a new SCP client
	client := scp.NewClient("192.168.150.120:22", &clientConfig)

	// Connect to the remote server
	err := client.Connect()
	if err != nil {
		fmt.Println("Couldn't establish a connection to the remote server ", err)
		return
	}

	// Open a file
	f, _ := os.Open("/Users/hezh/workspace/Backup/dump/log")

	// Close client connection after the file has been copied
	defer client.Close()

	// Close the file after it has been copied
	defer f.Close()

	// Finaly, copy the file over
	// Usage: CopyFile(fileReader, remotePath, permission)

	err = client.CopyFile(f, "~/b", "0655")

	if err != nil {
		fmt.Println("Error while copying file ", err)
	}
}
