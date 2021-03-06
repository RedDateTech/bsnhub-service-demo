module github.com/bianjieai/bsnhub-service-demo/examples/nft-service-provider

require (
	github.com/OneOfOne/xxhash v1.2.5 // indirect
	github.com/ethereum/go-ethereum v1.9.18
	github.com/google/uuid v1.1.1 // indirect
	github.com/irisnet/service-sdk-go v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	github.com/tendermint/tendermint v0.33.8 // indirect
	github.com/tidwall/gjson v1.6.1
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/irisnet/service-sdk-go => github.com/secret2830/service-sdk-go v0.0.0-20200930025908-91ed6ca17b1b
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.33.4-irita-200703.0.20200925112439-d4196a88a285
)

go 1.14
