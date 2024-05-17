package utility

import (
	"GoRsyncManager/configs"
	"GoRsyncManager/models"
	"strings"

	"golang.org/x/crypto/ssh"
)

func GetRemoteFileList(hostId int,
	fileFilterCmd string,
	content *models.Content) error {

	var output string
	var err error
	var host models.Host
	var bastionHost1 models.Host
	var bastionHost2 models.Host

	configs.DB.Find(&host, "id=?", hostId)
	b1 := configs.DB.Find(&bastionHost1, "id=?", host.Bastion)
	b2 := configs.DB.Find(&bastionHost2, "id=?", bastionHost1.Bastion)

	if b2.RowsAffected > 0 {
		err = runCommandWithTwoBastion(&bastionHost2, &bastionHost1, &host, fileFilterCmd, &output)
	} else if b1.RowsAffected > 0 {
		err = runCommandWithOneBastion(&bastionHost1, &host, fileFilterCmd, &output)
	} else {
		err = runCommand(&host, fileFilterCmd, &output)
	}

	if err != nil {
		return err
	}
	s := strings.Fields(output)

	content.Pages = GetTotalPage(s[0])
	content.Contents = s[1:]
	return nil
}

func SshConfig(host *models.Host) (*ssh.ClientConfig, error) {
	var auth []ssh.AuthMethod

	if host.Key != "" {
		if host.Password != "" {
			privateKey, err := ssh.ParsePrivateKeyWithPassphrase([]byte(host.Key), []byte(host.Password))
			if err != nil {
				return nil, err
			}
			auth = []ssh.AuthMethod{ssh.PublicKeys(privateKey)}
		} else {
			privateKey, err := ssh.ParsePrivateKey([]byte(host.Key))
			if err != nil {
				return nil, err
			}
			auth = []ssh.AuthMethod{ssh.PublicKeys(privateKey)}
		}
	} else {
		auth = []ssh.AuthMethod{ssh.Password(host.Password)}
	}

	// Set up SSH client configuration for the jump host.
	return &ssh.ClientConfig{
		User:            host.User,
		Auth:            auth,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}, nil
}

func runCommand(host *models.Host, command string, output *string) error {
	// Create SSH client configuration
	hostConfig, err := SshConfig(host)
	if err != nil {
		return err
	}

	// Connect to the SSH server
	client, err := ssh.Dial("tcp", host.Addr, hostConfig)
	if err != nil {
		return err
	}
	defer client.Close()

	// Create a new session
	session, err := client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	// Run the command
	o, err := session.CombinedOutput(command)
	if err != nil {
		return err
	}

	*output = string(o)

	return nil
}

func runCommandWithOneBastion(bastion *models.Host, host *models.Host, command string, output *string) error {
	// Create SSH client configuration for the bastion host
	bastionConfig, err := SshConfig(bastion)
	if err != nil {
		return err
	}

	// Dial to the bastion host
	bastionClient, err := ssh.Dial("tcp", bastion.Addr, bastionConfig)
	if err != nil {
		return err
	}
	defer bastionClient.Close()

	// Create SSH client configuration for the target host
	config, err := SshConfig(host)
	if err != nil {
		return err
	}

	// Dial to the target host through the bastion
	conn, err := bastionClient.Dial("tcp", host.Addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Create a new SSH client connection using the established connection
	ncc, chans, reqs, err := ssh.NewClientConn(conn, host.Addr, config)
	if err != nil {
		return err
	}
	defer ncc.Close()

	// Create a new SSH client using the connection
	client := ssh.NewClient(ncc, chans, reqs)
	defer client.Close()

	// Create a new session
	session, err := client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	// Run the command
	o, err := session.CombinedOutput(command)
	if err != nil {
		return err
	}

	*output = string(o)

	return nil
}

func runCommandWithTwoBastion(bastion1 *models.Host, bastion2 *models.Host, host *models.Host, command string, output *string) error {
	// Create SSH client configuration for the bastion1 host
	bastionConfig1, err := SshConfig(bastion1)
	if err != nil {
		return err
	}

	// Dial to the bastion1 host
	bastionClient1, err := ssh.Dial("tcp", bastion1.Addr, bastionConfig1)
	if err != nil {
		return err
	}
	defer bastionClient1.Close()

	// Create SSH client configuration for the bastion2 host
	bastionConfig2, err := SshConfig(bastion2)
	if err != nil {
		return err
	}

	// Dial to the bastion2 host through the bastion1
	connBastion2, err := bastionClient1.Dial("tcp", bastion2.Addr)
	if err != nil {
		return err
	}
	defer connBastion2.Close()

	// Create a new SSH client connection for bastion2 using the established connection
	nccBastion2, chans, reqs, err := ssh.NewClientConn(connBastion2, bastion2.Addr, bastionConfig2)
	if err != nil {
		return err
	}
	defer nccBastion2.Close()

	// Create a new SSH client of bastion2
	bastionClient2 := ssh.NewClient(nccBastion2, chans, reqs)
	defer bastionClient2.Close()

	// Create SSH client configuration for the target host
	config, err := SshConfig(host)
	if err != nil {
		return err
	}

	// Dial to the target host through the bastion
	conn, err := bastionClient2.Dial("tcp", host.Addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Create a new SSH client connection using the established connection
	ncc, chans, reqs, err := ssh.NewClientConn(conn, host.Addr, config)
	if err != nil {
		return err
	}
	defer ncc.Close()

	// Create a new SSH client using the connection
	client := ssh.NewClient(ncc, chans, reqs)
	defer client.Close()

	// Create a new session
	session, err := client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	// Run the command
	o, err := session.CombinedOutput(command)
	if err != nil {
		return err
	}

	*output = string(o)

	return nil
}
