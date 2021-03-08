package client

import (
	"github.com/Evrynetlabs/evrhub/cmd/ethrelayer/rpc"
	"github.com/Evrynetlabs/evrhub/cmd/util"
	"github.com/Evrynetlabs/evrhub/x/common/types"
	"github.com/golang/glog"
)

func SendProphecyClaimToEvrnet(claim types.EthProphecyClaim) (string, error) {
	cfg := rpc.GetConfig()
	URL := cfg.Url

	URL += util.ETHEREUM_RELAYPROPHECYCLAIM
	connect := util.GetHTTPClient()
	claimString := types.EthProphecyClaimToJsonString(&claim)
	body, _, err := util.HttpGet("POST", URL, claimString, nil, nil, connect)
	if err != nil {
		glog.Error(err)
		return "", err
	}

	return body, nil
}
