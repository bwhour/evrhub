package types

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/golang/glog"
)

func EvrProphecyClaimToJsonString(e *EvrProphecyClaim) string {
	bytes, err := json.Marshal(e)
	if err != nil {
		glog.Errorf("EvrProphecyClaimToJsonString err:%+v", err)
		return ""
	}
	return string(bytes)
}

func JsonToEvrProphecyClaim(r io.Reader) (*EvrProphecyClaim, error) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		glog.Errorf("read evrprophecyclaim request body err: %+v", err)
		return nil, err
	}

	epc := EvrProphecyClaim{}
	err = json.Unmarshal(body, &epc)
	if err != nil {
		glog.Errorf("convert string to EvrProphecyClaim err: %+v", err)
		return nil, err
	}

	return &epc, nil
}

func EthProphecyClaimToJsonString(e *EthProphecyClaim) string {
	bytes, err := json.Marshal(e)
	if err != nil {
		glog.Errorf("EthProphecyClaimToJsonString err:%+v", err)
		return ""
	}
	return string(bytes)
}

func JsonToEthProphecyClaim(r io.Reader) (*EthProphecyClaim, error) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		glog.Errorf("read ethprophecyclaim request body err: %+v", err)
		return nil, err
	}

	epc := EthProphecyClaim{}
	err = json.Unmarshal(body, &epc)
	if err != nil {
		glog.Errorf("convert string to EthProphecyClaim err: %+v", err)
		return nil, err
	}

	return &epc, nil
}
