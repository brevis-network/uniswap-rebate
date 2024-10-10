package main

import (
	"context"
	"log"
	"time"

	"github.com/brevis-network/uniswap-rebate/dal"
	"github.com/brevis-network/uniswap-rebate/webapi"
)

type Server struct {
	db *dal.DAL
}

func (s *Server) NewProof(ctx context.Context, req *webapi.NewProofReq) (*webapi.NewProofResp, error) {
	log.Println(req)
	ret := &webapi.NewProofResp{
		Reqid: uint64(time.Now().Unix()),
	}
	return ret, nil
}

func (s *Server) GetProof(ctx context.Context, req *webapi.GetProofReq) (*webapi.GetProofResp, error) {
	log.Println(req)
	ret := &webapi.GetProofResp{
		Reqid: req.Reqid,
	}
	return ret, nil
}
