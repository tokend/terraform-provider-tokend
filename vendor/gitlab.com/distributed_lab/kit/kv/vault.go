package kv

//import (
//	"path"
//	"strings"
//
//	"github.com/davecgh/go-spew/spew"
//	vault "github.com/hashicorp/vault/api"
//	"gitlab.com/distributed_lab/logan/v3/errors"
//)
//
//type VaultConfig struct {
//	Host    string
//	Token   string
//	Prefix  string
//	Backend string
//}
//
//type vaultbackend struct {
//	config VaultConfig
//	client *vault.Client
//}
//
//func NewVault(config VaultConfig) GetterSetter {
//	return &vaultbackend{
//		config: config,
//	}
//}
//
//func (b *vaultbackend) ensureClient() error {
//	if b.client == nil {
//		config := vault.DefaultConfig()
//		if b.config.Host != "" {
//			config.Address = b.config.Host
//		}
//		client, err := vault.NewClient(config)
//		if err != nil {
//			return errors.Wrap(err, "failed to init vault client")
//		}
//		if b.config.Token != "" {
//			client.SetToken(b.config.Token)
//		}
//		b.client = client
//	}
//	return nil
//}
//
//func (b vaultbackend) buildPath(key string) string {
//	return path.Join(b.config.Backend, "data", b.config.Prefix, key)
//}
//
//func (b vaultbackend) Put(key string, value interface{}) error {
//	if err := b.ensureClient(); err != nil {
//		return err
//	}
//	i := strings.LastIndex(key, "/")
//	//fmt.Println(key[:i], key[i+1:], value)
//	_, err := b.client.Logical().Write(b.buildPath(key[:i]), map[string]interface{}{
//		"data": map[string]interface{}{
//			key[i+1:]: value,
//		},
//	})
//	if err != nil {
//		return errors.Wrap(err, "failed to write vault secret")
//	}
//	return nil
//}
//
//func (b vaultbackend) Get(key string) (interface{}, error) {
//	if err := b.ensureClient(); err != nil {
//		return nil, err
//	}
//	i := strings.LastIndex(key, "/")
//	//fmt.Println(key[:i], key[i+1:], value)
//	secret, err := b.client.Logical().Read(b.buildPath(key[:i]))
//	if err != nil {
//		return nil, errors.Wrap(err, "failed to write vault secret")
//	}
//	return secret.Data, nil
//}
//
//func (b vaultbackend) SetStringMap(key string, value map[string]interface{}) error {
//	if err := b.ensureClient(); err != nil {
//		return err
//	}
//	secret, err := b.client.Logical().Write(b.buildPath(key), map[string]interface{}{
//		"data": value,
//	})
//	if err != nil {
//		return errors.Wrap(err, "failed to write vault secret")
//	}
//	spew.Dump(secret)
//	return nil
//}
//
//func (b vaultbackend) GetStringMap(key string) (map[string]interface{}, error) {
//	if err := b.ensureClient(); err != nil {
//		return nil, err
//	}
//	secret, err := b.client.Logical().Read(b.buildPath(key))
//	if err != nil {
//		panic(errors.Wrap(err, "failed to access vault"))
//	}
//	if secret != nil && secret.Data != nil {
//		// TODO check for metadata and is destroyed
//		return secret.Data["data"].(map[string]interface{}), nil
//	}
//	return nil, nil
//}
