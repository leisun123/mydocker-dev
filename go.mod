module docker-demo/mydocker-dev

go 1.13

exclude github.com/Sirupsen/logrus v1.0.5

require (
	github.com/onsi/ginkgo v1.13.0 // indirect
	github.com/sirupsen/logrus v1.0.5
	github.com/stretchr/testify v1.6.1 // indirect
	github.com/urfave/cli v1.19.1
	golang.org/x/crypto v0.0.0-20200604202706-70a84ac30bf9 // indirect
	golang.org/x/sys v0.0.0-20200615200032-f1bc736245b1 // indirect
	gopkg.in/airbrake/gobrake.v2 v2.0.9 // indirect
	gopkg.in/gemnasium/logrus-airbrake-hook.v2 v2.1.2 // indirect
)

replace github.com/Sirupsen/logrus v1.0.5 => github.com/sirupsen/logrus v1.0.5
