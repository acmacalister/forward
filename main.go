package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

var (
	server = flag.String("server", "localhost:22", "Set the SSH server you are connecting to. Defaults to localhost.")
	local  = flag.String("local", "localhost:8080", "Set your local web server you want share. Defaults to localhost:8080")
	remote = flag.String("remote", "localhost:5000", "Set the remote listener that your ssh connection will use. Defaults to port 5000, as generally it would be behind a reverse proxy like nginx.")
	user   = flag.String("user", "root", "Set the username you wish to connect to your ssh server with. Defaults to root")
)

func main() {
	flag.Parse()
	// This connects to your local ssh agent and uses the rsa keys stored there.
	conn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	ag := agent.NewClient(conn)

	config := &ssh.ClientConfig{
		User: *user, // user from our flags.
		Auth: []ssh.AuthMethod{
			ssh.PublicKeysCallback(ag.Signers), // We are able to get the signers from our ssh agent.
		},
	}

	// dials the remote ssh server you want to connect to.
	sshClientConn, err := ssh.Dial("tcp", *server, config)
	if err != nil {
		log.Fatal(err)
	}

	// sets up a remote address listener from our flags.
	l, err := sshClientConn.Listen("tcp", *remote)

	// a loop that runs forever to forward the remote data to here.
	for {
		sshConn, err := l.Accept() // Gets request made to server address flag.

		localConn, err := net.Dial("tcp", *local) // connects to local address from our flags.
		if err != nil {
			log.Fatal(err)
		}

		// forward the remote request to our local one.
		go func() {
			_, err = io.Copy(localConn, sshConn)
			if err != nil {
				log.Fatal(err)
			}
		}()

		// forward our response to back to the remote address.
		go func() {
			_, err = io.Copy(sshConn, localConn)
			if err != nil {
				log.Fatal(err)
			}
		}()
	}
}
