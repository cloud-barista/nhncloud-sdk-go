module github.com/cloud-barista/nhncloud-sdk-for-drv

go 1.16

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.3
	github.com/dgrijalva/jwt-go => github.com/golang-jwt/jwt v3.2.1+incompatible
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)

require (
	github.com/kr/text v0.2.0 // indirect
	golang.org/x/crypto v0.0.0-20220817201139-bc19a97f63c8
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v2 v2.4.0
)
