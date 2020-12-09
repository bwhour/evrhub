package rest
/*
import "github.com/Evrynetlabs/evrsdk/types/rest"


import (
	"fmt"
	"github.com/Evrynetlabs/evrsdk/types/rest"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/Evrynetlabs/evrhub/x/ethbridge/types"
)

const (
	restEthereumChainID = "ethereumChainID"
	restBridgeContract  = "bridgeContract"
	restNonce           = "nonce"
	restSymbol          = "symbol"
	restTokenContract   = "tokenContract"
	restEthereumSender  = "ethereumSender"
)

type createEthClaimReq struct {
	BaseReq               rest.BaseReq `json:"base_req"`
	EthereumChainID       int          `json:"ethereum_chain_id"`
	BridgeContractAddress string       `json:"bridge_registry_contract_address"`
	Nonce                 int          `json:"nonce"`
	Symbol                string       `json:"symbol"`
	TokenContractAddress  string       `json:"token_contract_address"`
	EthereumSender        string       `json:"ethereum_sender"`
	EvrnetReceiver        string       `json:"evrnet_Receiver"`
	Validator             string       `json:"validator"`
	Amount                int64        `json:"amount"`
	ClaimType             string       `json:"claim_type"`
}

type burnOrLockEthReq struct {
	BaseReq          rest.BaseReq `json:"base_req"`
	EthereumChainID  string       `json:"ethereum_chain_id"`
	TokenContract    string       `json:"token_contract_address"`
	EvrnetSender     string       `json:"evrnet_sender"`
	EthereumReceiver string       `json:"ethereum_receiver"`
	Amount           int64        `json:"amount"`
	Symbol           string       `json:"symbol"`
}

func createClaimHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createEthClaimReq

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		bridgeContractAddress := types.NewEthereumAddress(req.BridgeContractAddress)

		tokenContractAddress := types.NewEthereumAddress(req.TokenContractAddress)

		ethereumSender := types.NewEthereumAddress(req.EthereumSender)

		evrnetReceiver, err := sdk.AccAddressFromBech32(req.EvrnetReceiver)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		validator, err := sdk.ValAddressFromBech32(req.Validator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		claimType, err := types.StringToClaimType(req.ClaimType)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, types.ErrInvalidClaimType.Error())
			return
		}

		// create the message
		ethBridgeClaim := types.NewEthBridgeClaim(
			req.EthereumChainID, bridgeContractAddress, req.Nonce, req.Symbol,
			tokenContractAddress, ethereumSender, evrnetReceiver, validator, req.Amount, claimType)
		msg := types.NewMsgCreateEthBridgeClaim(ethBridgeClaim)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

func getProphecyHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		ethereumChainID := vars[restEthereumChainID]
		ethereumChainIDString, err := strconv.Atoi(ethereumChainID)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		bridgeContract := types.NewEthereumAddress(vars[restBridgeContract])

		nonce := vars[restNonce]
		nonceString, err := strconv.Atoi(nonce)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		tokenContract := types.NewEthereumAddress(vars[restTokenContract])

		symbol := vars[restSymbol]
		if strings.TrimSpace(symbol) == "" {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, "symbol is empty")
			return
		}

		ethereumSender := types.NewEthereumAddress(vars[restEthereumSender])

		bz, err := cliCtx.Codec.MarshalJSON(
			types.NewQueryEthProphecyParams(
				ethereumChainIDString, bridgeContract, nonceString, symbol, tokenContract, ethereumSender))
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		route := fmt.Sprintf("custom/%s/%s", storeName, types.QueryEthProphecy)
		res, _, err := cliCtx.QueryWithData(route, bz)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func burnOrLockHandler(cliCtx context.CLIContext, lockOrBurn string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req burnOrLockEthReq

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		ethereumChainID, err := strconv.Atoi(req.EthereumChainID)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		evrnetSender, err := sdk.AccAddressFromBech32(req.EvrnetSender)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		ethereumReceiver := types.NewEthereumAddress(req.EthereumReceiver)

		// create the message
		var msg sdk.Msg
		switch lockOrBurn {
		case "lock":
			msg = types.NewMsgLock(ethereumChainID, evrnetSender, ethereumReceiver, req.Amount, req.Symbol)
		case "burn":
			msg = types.NewMsgBurn(ethereumChainID, evrnetSender, ethereumReceiver, req.Amount, req.Symbol)
		}
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
*/