module github.com/XiaYinchang/harbor/src/core

require (
	github.com/XiaYinchang/harbor v1.7.3
	github.com/XiaYinchang/keystone-go-sdk v0.0.0-20190225053546-63bdfa86437d
	github.com/astaxie/beego v1.11.1
	github.com/beego/i18n v0.0.0-20161101132742-e9308947f407
	github.com/dghubble/sling v1.2.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/docker/distribution v2.7.1+incompatible
	github.com/docker/libtrust v0.0.0-20160708172513-aabc10ec26b7
	github.com/ghodss/yaml v1.0.0
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/smartystreets/goconvey v0.0.0-20190222223459-a17d461953aa // indirect
	github.com/stretchr/testify v1.3.0
	gopkg.in/ldap.v2 v2.5.1
	k8s.io/helm v2.12.3+incompatible
)

replace github.com/XiaYinchang/harbor => ../../../harbor
