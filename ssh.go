package main

//
//import (
//	"golang.org/x/crypto/ssh"
//	"github.com/mitchellh/go-homedir"
//	"os"
//)
//
//type SignerContainer struct {
//	signers []ssh.Signer
//}
//
//func (t *SignerContainer) Key(i int) (key ssh.PublicKey, err error) {
//	if i >= len(t.signers) {
//		return
//	}
//	key = t.signers[i].PublicKey()
//	return
//}
//
//func (t *SignerContainer) Sign(i int, rand io.Reader, data []byte) (sig []byte, err error) {
//	if i >= len(t.signers) {
//		return
//	}
//	sig, err = t.signers[i].Sign(rand, data)
//	return
//}
//
//func makeSigner(keyname string) (signer ssh.Signer, err error) {
//	fp, err := os.Open(keyname)
//	if err != nil {
//		return
//	}
//	defer fp.Close()
//
//	buf, _ := ioutil.ReadAll(fp)
//	signer, _ = ssh.ParsePrivateKey(buf)
//	return
//}
//
//func makeKeyring() ssh.ClientAuth {
//	signers := []ssh.Signer{}
//	keys := []string{os.Getenv("HOME") + "/.ssh/id_rsa", os.Getenv("HOME") + "/.ssh/id_dsa"}
//
//	for _, keyname := range keys {
//		signer, err := makeSigner(keyname)
//		if err == nil {
//			signers = append(signers, signer)
//		}
//	}
//
//	return ssh.ClientAuthKeyring(&SignerContainer{signers})
//}
//
//func sshAuth(keyPaths ...string) {
//	signers := []ssh.Signer{}
//	for _, keyname := range keyPaths{
//		signer, err := makeSigner(keyname)
//		if err == nil {
//			signers = append(signers, signer)
//		}
//	}
//
//	return ssh.ClientAuthKeyring(&SignerContainer{signers})
//}
//
//func AddLineToAuthorizedKeys()
