package cert

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"log"
	"os"

	"github.com/go-acme/lego/certcrypto"
	"github.com/go-acme/lego/certificate"
	"github.com/go-acme/lego/lego"
	"github.com/go-acme/lego/providers/dns"
	"github.com/go-acme/lego/registration"
)

// MyUser You'll need a user or account type that implements acme.User
type MyUser struct {
	Email        string
	Registration *registration.Resource
	key          crypto.PrivateKey
}

// GetEmail xxx
func (u *MyUser) GetEmail() string {
	return u.Email
}

// GetRegistration xxx
func (u MyUser) GetRegistration() *registration.Resource {
	return u.Registration
}

// GetPrivateKey xxx
func (u *MyUser) GetPrivateKey() crypto.PrivateKey {
	return u.key
}

// CreateCert create a cert
func CreateCert(email, domain, prov string, envs map[string]string) (map[string][]byte, error) {

	for k, v := range envs {
		err := os.Setenv(k, v)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
	}

	// Create a user. New accounts need an email and private key to start.
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	myUser := MyUser{
		Email: email,
		key:   privateKey,
	}

	config := lego.NewConfig(&myUser)

	// This CA URL is configured for a local dev instance of Boulder running in Docker in a VM.
	// config.CADirURL = "http://192.168.99.100:4000/directory"
	config.Certificate.KeyType = certcrypto.RSA2048

	// A client facilitates communication with the CA server.
	client, err := lego.NewClient(config)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	cp, err := dns.NewDNSChallengeProviderByName(prov)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	client.Challenge.SetDNS01Provider(cp)

	// New users will need to register
	reg, err := client.Registration.Register(registration.RegisterOptions{TermsOfServiceAgreed: true})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	myUser.Registration = reg

	request := certificate.ObtainRequest{
		Domains: []string{domain},
		Bundle:  true,
	}
	certificates, err := client.Certificate.Obtain(request)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Each certificate comes back with the cert bytes, the bytes of the client's
	// private key, and a certificate URL. SAVE THESE TO DISK.
	fmt.Println("key:", string(certificates.PrivateKey))
	fmt.Println("cert:", string(certificates.Certificate))

	cert := map[string][]byte{
		"key":  certificates.PrivateKey,
		"cert": certificates.Certificate,
	}
	// ... all done.
	return cert, nil
}
