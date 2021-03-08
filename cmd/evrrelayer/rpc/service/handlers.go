package service

import (
	"github.com/Evrynetlabs/evrhub/cmd/evrrelayer/relayer"
	"github.com/Evrynetlabs/evrhub/cmd/evrrelayer/rpc"
	"github.com/Evrynetlabs/evrhub/cmd/util"
	"github.com/Evrynetlabs/evrhub/x/common/types"
	"github.com/golang/glog"
	"github.com/labstack/echo/v4"
)

func StartHttpServer(s *relayer.EvrnetSub) {
	util.AddHandler("POST", util.ETHEREUM_RELAYPROPHECYCLAIM, HandleRelayProphecyClaim)

	port := rpc.GetConfig().Port
	util.SetPort(port)

	go startMessageLoop(s)
	go util.Start()
}

func HandleRelayProphecyClaim(c echo.Context) error {
	ethProphecyClaim, err := types.JsonToEthProphecyClaim(c.Request().Body)
	if err != nil {
		glog.Errorf("jsonToClaimType err: %+v", err)
		return err
	}

	RelayProphecyClaim(ethProphecyClaim)

	return nil
}
