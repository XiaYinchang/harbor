// Copyright 2018 Project Harbor Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package keystone

import (
	"os"
	"strings"

	"github.com/XiaYinchang/keystone-go-sdk/keystone"
	"github.com/XiaYinchang/harbor/src/common/dao"
	"github.com/XiaYinchang/harbor/src/common/models"
	"github.com/XiaYinchang/harbor/src/common/utils/log"
	"github.com/XiaYinchang/harbor/src/core/auth"
)

var keyConfig *Config

// Auth implements Authenticator interface to authenticate user against DB.
type Auth struct {
	auth.DefaultAuthenticateHelper
}

// Config configures a keystone webhook server
type Config struct {
	KeystoneURL    string
	KeystoneDomain string
	KeystoneUser   string
	KeystonePass   string
}

// NewConfig returns a Config
func NewConfig() *Config {
	return &Config{
		KeystoneURL:    os.Getenv("OS_AUTH_URL"),
		KeystoneDomain: os.Getenv("OS_DOMAIN_NAME"),
	}
}

// Authenticate calls dao to authenticate user.
func (d *Auth) Authenticate(m models.AuthModel) (*models.User, error) {
	p := m.Principal
	if len(strings.TrimSpace(p)) == 0 {
		log.Debugf("Keystone authentication failed for empty user id.")
		return nil, auth.NewErrAuth("Empty user id")
	}
	keyConfig.KeystoneUser = strings.TrimSpace(p)
	keyConfig.KeystonePass = m.Password
	// create new client
	authInfo := keystone.KeystoneAuth{
		AuthURL:    keyConfig.KeystoneURL,
		DomainName: keyConfig.KeystoneDomain,
		UserName:   keyConfig.KeystoneUser,
		Password:   keyConfig.KeystonePass,
	}
	client, err := keystone.NewClient(authInfo)
	if err != nil {
		return nil, err
	}
	u := models.User{}
	u.Username = client.AuthInfo.UserName
	detailedUserInfo, err := dao.GetUser(u)
	if err != nil {
		return nil, err
	}
	detailedUserInfo.Token = client.AuthInfo.Token
	return detailedUserInfo, nil
}

// SearchUser - Check if user exist in local db
func (d *Auth) SearchUser(username string) (*models.User, error) {
	var queryCondition = models.User{
		Username: username,
	}
	return dao.GetUser(queryCondition)
}

// OnBoardUser -
func (d *Auth) OnBoardUser(u *models.User) error {
	return nil
}

func init() {
	auth.Register("keystone_auth", &Auth{})
	keyConfig = NewConfig()
}
