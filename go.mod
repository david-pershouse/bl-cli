module git.mammoth.com.au/github/bl-cli

go 1.15

require (
	git.mammoth.com.au/github/go-binarylane v0.0.0-20201030110224-2755d78d32df
	github.com/blang/semver v3.5.1+incompatible
	github.com/creack/pty v1.1.11
	github.com/digitalocean/doctl v1.49.0
	github.com/digitalocean/godo v1.50.0
	github.com/docker/cli v0.0.0-20200622130859-87db43814b48
	github.com/dustin/go-humanize v1.0.0
	github.com/fatih/color v1.7.0
	github.com/gobwas/glob v0.2.3
	github.com/golang/mock v1.4.4
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/google/uuid v1.1.1
	github.com/mitchellh/copystructure v1.0.0
	github.com/natefinch/pie v0.0.0-20170715172608-9a0d72014007
	github.com/sclevine/spec v1.3.0
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644
	github.com/spf13/cobra v0.0.3
	github.com/spf13/viper v1.4.0
	github.com/stretchr/testify v1.4.0
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	golang.org/x/net v0.0.0-20201027133719-8eef5233e2a1 // indirect
	golang.org/x/oauth2 v0.0.0-20200902213428-5d25da1a8d43
	google.golang.org/appengine v1.6.7 // indirect
	gopkg.in/yaml.v2 v2.2.8
	k8s.io/api v0.17.13
	k8s.io/apimachinery v0.17.13
	k8s.io/client-go v0.17.13
	sigs.k8s.io/yaml v1.2.0
)

replace github.com/stretchr/objx => github.com/stretchr/objx v0.2.0

replace golang.org/x/crypto => golang.org/x/crypto v0.0.0-20200420201142-3c4aac89819a
