package clientCredentials

import (
	crand "crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"log"
	rand "math/rand"

	"github.com/google/uuid"
)

type clientCredentialsCreator struct{}

func NewClientCredentialsCreator() *clientCredentialsCreator {
	return &clientCredentialsCreator{}
}

func (*clientCredentialsCreator) Create() *clientCredentialsOutput {
	clientId := createClientId()
	clientSecret := createClientSecret()

	o := NewClientCredentialsOutput(clientId, clientSecret)
	return o
}

func createClientId() string {
	return uuid.New().String()
}

func createClientSecret() string {
	var src cryptoSource
	rnd := rand.New(src)
	buff := make([]byte, 32)
	rnd.Read(buff)

	secret := base64.StdEncoding.EncodeToString(buff)
	return secret
}

type cryptoSource struct{}

func (s cryptoSource) Seed(seed int64) {}

func (s cryptoSource) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

func (s cryptoSource) Uint64() (v uint64) {
	err := binary.Read(crand.Reader, binary.BigEndian, &v)
	if err != nil {
		log.Fatal(err)
	}
	return v
}
