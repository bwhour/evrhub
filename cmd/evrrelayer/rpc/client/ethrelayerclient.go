package client

import (
	"github.com/Evrynetlabs/evrhub/cmd/evrrelayer/rpc"
	"github.com/Evrynetlabs/evrhub/cmd/util"
	"github.com/Evrynetlabs/evrhub/x/common/types"
	"github.com/golang/glog"
)

func SendProphecyClaimToEthereum(claim types.EvrProphecyClaim) (string, error) {
	cfg := rpc.GetConfig()
	URL := cfg.Url

	URL += util.EVRNET_RELAYPROPHECYCLAIM
	connect := util.GetHTTPClient()
	claimString := types.EvrProphecyClaimToJsonString(&claim)
	body, _, err := util.HttpGet("POST", URL, claimString, nil, nil, connect)
	if err != nil {
		glog.Error(err)
		return "", err
	}

	return body, nil
}
