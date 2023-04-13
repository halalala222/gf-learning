package addressUtils

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/halalala222/gf-learning/internal/model"
)

func Verify(data *model.VerifyModel) bool {
	byteChallenge := []byte(data.Challenge)
	signature, err := hexutil.Decode(data.Signature)
	if err != nil {
		log.Info(err.Error())
		return false
	}
	byteChallenge = accounts.TextHash(byteChallenge)
	signature[crypto.RecoveryIDOffset] -= 27
	recovered, err := crypto.SigToPub(byteChallenge, signature)
	if err != nil {
		log.Info(err.Error())
		return false
	}
	recoveredAddress := crypto.PubkeyToAddress(*recovered).Hex()
	return data.Address == recoveredAddress
}
