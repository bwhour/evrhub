package service

import (
	"github.com/Evrynetlabs/evrhub/cmd/ethrelayer/relayer"
	"github.com/Evrynetlabs/evrhub/cmd/ethrelayer/rpc"
	"github.com/Evrynetlabs/evrhub/cmd/util"
	"github.com/Evrynetlabs/evrhub/x/common/types"
	"github.com/golang/glog"
	"github.com/labstack/echo/v4"
)

func StartHttpServer(s *relayer.EthereumSub) {
	util.AddHandler("POST", util.EVRNET_RELAYPROPHECYCLAIM, HandleRelayProphecyClaim)

	port := rpc.GetConfig().Port
	util.SetPort(port)

	go startMessageLoop(s)
	go util.Start()
}

func HandleRelayProphecyClaim(c echo.Context) error {
	evrProphecyClaim, err := types.JsonToEvrProphecyClaim(c.Request().Body)
	if err != nil {
		glog.Errorf("jsonToClaimType err: %+v", err)
		return err
	}

	RelayProphecyClaim(evrProphecyClaim)

	return nil
}
