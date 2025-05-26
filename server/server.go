package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/brevis-network/uniswap-rebate/dal"
	"github.com/brevis-network/uniswap-rebate/webapi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
)

type Server struct {
	db *dal.DAL
}

func (s *Server) NewProof(ctx context.Context, req *webapi.NewProofReq) (ret *webapi.NewProofResp, err error) {
	log.Println(req)
	ret = new(webapi.NewProofResp)
	// parse and get tx receipt now even though it may take a long time for this api to return
	// so client knows error immediately instead of later polling
	onec, ok := chainMap[req.ChainId]
	if !ok {
		ret.Errmsg = fmt.Sprintf("unsupported chainid %d", req.ChainId)
		return
	}
	txList := strings.Split(req.TxnHashes, ",")
	receipts, err := onec.FetchTxReceipts(txList)
	if err != nil {
		ret.Errmsg = "get receipts err: " + err.Error()
		return
	}

	reqid := time.Now().Unix()
	// save json file in case need to resume
	fname := fmt.Sprintf("%s/%d-receipts.json", *fdir, reqid)
	raw, _ := json.Marshal(receipts)
	err = os.WriteFile(fname, raw, os.ModePerm)
	if err != nil {
		ret.Errmsg = "save receipts err: " + err.Error()
		return
	}

	// check logs in receipts and prepare proof etc
	prvR, err := onec.ProcessReceipts(receipts, common.HexToAddress(req.Beneficiary))
	if err != nil {
		ret.Errmsg = "ProcessReceipts err: " + err.Error()
		return
	}
	if len(prvR) == 0 {
		ret.Errmsg = "ProcessReceipts returns 0 requests"
		return
	}
	for _, r := range prvR {
		r.ReqId = reqid
	}

	fname = fmt.Sprintf("%s/%d-proveReqs.json", *fdir, reqid)
	raw, _ = json.Marshal(prvR)
	os.WriteFile(fname, raw, os.ModePerm)

	// submit to app prover
	resp, err := http.Post(viper.GetString("prover")+"/prove", "application/json", bytes.NewBuffer(raw))
	if err != nil {
		ret.Errmsg = "post to prover err: " + err.Error()
		return
	}
	defer resp.Body.Close()
	respR, _ := io.ReadAll(resp.Body)
	log.Println(resp.StatusCode, string(respR))

	// save req to db
	err = s.db.ReqAdd(context.Background(), dal.ReqAddParams{
		ID:       reqid,
		Proofreq: req,
	})
	if err != nil {
		ret.Errmsg = "db err: " + err.Error()
		return
	}

	// good to return
	ret.Reqid = uint64(reqid)
	return ret, nil
}

func (s *Server) GetProof(ctx context.Context, req *webapi.GetProofReq) (*webapi.GetProofResp, error) {
	ret := &webapi.GetProofResp{
		Reqid: req.Reqid,
	}
	oneReq, err := s.db.ReqGet(context.Background(), int64(req.Reqid))
	if err != nil {
		return ret, err
	}
	// ret.Status
	ret.Calldata = oneReq.Calldata.ToWebCallData()
	return ret, nil
}
