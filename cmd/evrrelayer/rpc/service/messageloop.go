package service

import (
	"github.com/Evrynetlabs/evrhub/cmd/evrrelayer/relayer"
	"github.com/Evrynetlabs/evrhub/cmd/evrrelayer/txs"
	"github.com/Evrynetlabs/evrhub/cmd/evrrelayer/types"
	xcommon "github.com/Evrynetlabs/evrhub/x/common/types"
	"github.com/golang/glog"
)

type msgType int

const (
	relayProphecyClaim msgType = iota
)

var mgr *messageMgr

type message struct {
	op    msgType
	param interface{}
}

type messageMgr struct {
	msgChan chan *message
	sub     *relayer.EvrnetSub
}

func startMessageLoop(s *relayer.EvrnetSub) {
	mgr = &messageMgr{
		msgChan: make(chan *message, 10000),
		sub:     s,
	}
	mgr.messageLoop()
}

func (m *messageMgr) messageLoop() {
	for {
		select {
		case msg, open := <-m.msgChan:
			if !open {
				glog.Info("chan closed, message loop exit")
				return
			}
			switch msg.op {
			case relayProphecyClaim:
				glog.Info("message loop msg: relayProphecyClaim")
				if msg.param == nil {
					glog.Info("relayProphecyClaim param nil")
					continue
				}

				pc, ok := msg.param.(*xcommon.EthProphecyClaim)
				if !ok {
					glog.Info("addPubChan param err")
					continue
				}
				m.handleRelayProphecyClaimMsg(pc)
			}
		}
	}
}

func (m *messageMgr) isLoopExit() bool {
	return m.msgChan == nil
}

func RelayProphecyClaim(prophecyClaim *xcommon.EthProphecyClaim) {
	if mgr.isLoopExit() {
		glog.Errorf("channel is close, relayProphecyClaim msg not implement")
		return
	}
	glog.Infof("RelayProphecyClaim prophecyClaim is:%+v", prophecyClaim)

	mgr.msgChan <- &message{op: relayProphecyClaim, param: prophecyClaim}
	return
}

func (m *messageMgr) handleRelayProphecyClaimMsg(prophecyClaim *xcommon.EthProphecyClaim) {
	if m.isLoopExit() {
		glog.Errorf("channel is close, handleRelayProphecyClaimMsg msg not implement")
		return
	}

	var claimType types.Event
	if prophecyClaim.ClaimType == xcommon.LockText {
		claimType = types.MsgLock
	} else if prophecyClaim.ClaimType == xcommon.BurnText {
		claimType = types.MsgBurn
	}

	err := txs.RelayProphecyClaimToEvrnet(m.sub.EvrProvider, m.sub.RegistryContractAddress,
		claimType, *prophecyClaim, m.sub.PrivateKey)
	if err != nil {
		glog.Errorf("handleRelayProphecyClaimMsg err:%+v", err)
		return
	}

	return
}
