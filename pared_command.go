package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

//SSHClient Type needs ssh.ClientConfig
type SSHClient struct {
	Config *ssh.ClientConfig
	Host   string
	Port   int
}

type sshcommand struct {
	command string
	stdout  io.Writer
	stderr  io.Writer
}

//YinitAgent Auth methods reads from SSH_AUTH_SOCK.
func YinitAgent() ssh.AuthMethod {
	if sshAgent, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
		return ssh.PublicKeysCallback(agent.NewClient(sshAgent).Signers)
	}
	return nil
}

func (client *SSHClient) runonhost(cmd string) (op, er string, error error) {
	//Create a new session first
	session, err := client.newsession()
	if err != nil {
		return "", "", err
	}
	defer session.Close()

	stdoutBuf := new(bytes.Buffer)
	stderrBuf := new(bytes.Buffer)
	
	session.Stdout = stdoutBuf
	session.Stderr = stderrBuf

	session.Run("hostname")

	stdoutput := stdoutBuf.String()
	stderror := stderrBuf.String()

	fmt.Printf("stdoutput was %s\n", stdoutBuf.String())
	fmt.Printf("stderror was %s\n", stderror)

	return stdoutput, stderror, nil

}

// If sudo is there in the command make sure yinit is run. so that we have
// the sudo tokens
// Need to understand how that works.

func (client *SSHClient) newsession() (*ssh.Session, error) {
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", client.Host, client.Port), client.Config)
	if err != nil {
		return nil, fmt.Errorf("Failed to dial: %s", err)
	}

	session, err := conn.NewSession()
	if err != nil {
		return nil, fmt.Errorf("Failed to create session: %s", err)
	}

	return session, nil

}

func SshRun(hostname string, command string) (stdoutput, stderror string) {
	user := os.Getenv("User")
	fmt.Printf("User is %s\n",user)
	sshconfig := &ssh.ClientConfig{
		User: "veeru",
		Auth: []ssh.AuthMethod{
			YinitAgent(),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	fmt.Printf("Run command %s on host %s", command, hostname)
	client := &SSHClient{
		Config: sshconfig,
		Host:   hostname,
		Port:   22,
	}
	op, er, _ := client.runonhost("hostname")
	fmt.Printf("Output from run is %s %s\n",op,er)
	return op, er
}

func main(){
SshRun("kubenode3.test1.ensemble.ne1.yahoo.com","ls -lrt /;hostname")
}
