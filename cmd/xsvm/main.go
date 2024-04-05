// Copyright (C) 2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"encoding/hex"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/vms/platformvm/signer"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/spf13/cobra"
)

func init() {
	cobra.EnablePrefixMatching = true
}

type GetNodeIDReply struct {
	NodeID  ids.NodeID                `json:"nodeID"`
	NodePOP *signer.ProofOfPossession `json:"nodePOP"`
}

func main() {
	//cmd := run.Command()
	//cmd.AddCommand(
	//	account.Command(),
	//	chain.Command(),
	//	issue.Command(),
	//	version.Command(),
	//)
	//ctx := context.Background()
	//if err := cmd.ExecuteContext(ctx); err != nil {
	//	fmt.Fprintf(os.Stderr, "command failed %v\n", err)
	//	os.Exit(1)
	//}
	id, _ := ids.FromString("5")
	cid, _ := ids.FromString("2HRvhEznCtkZxAXYcdcTQ1bZ1xs6EksmemFC5BeHgc3dkEH6GD")
	message, _ := warp.NewUnsignedMessage(
		id,
		cid,
		[]byte("68656c6c6f20776f726c64"),
	)
	//byteHello := []byte("68656c6c6f20776f726c64")

	println(hex.EncodeToString(message.Bytes()))
	keyBytes, _ := hex.DecodeString("0040c137287b7d169e076cb80d69bf98222ea780693b3bb4989c5998508490ff")
	pk, err := bls.SecretKeyFromBytes(keyBytes)
	if err != nil {
		panic(err)
	}
	signature := bls.Sign(pk, message.Bytes())
	pub := bls.PublicFromSecretKey(pk)

	println(hex.EncodeToString(bls.PublicKeyToBytes(pub)))
	//
	println(hex.EncodeToString(bls.SignatureToBytes(signature)))

	var k = "a1b2e995ca49db484093701225ca2ddba2bbf53bcf3c6c137d4913f822ac5a489430bed710640f1cb6aa9c74caafd9a2"
	pubkeyB, _ := hex.DecodeString(k)
	var pubKey, _ = bls.PublicKeyFromBytes(pubkeyB)

	ss, _ := bls.SignatureFromBytes(bls.SignatureToBytes(signature))

	println(bls.Verify(pubKey, ss, message.Bytes()))
}

//88b9363024125bcbdb6b78866e052cb95eb9d73b2261a333df03561093c9e8b74b188a929869b174407cd2163842d27811014807640e02ace08e3bbc3a705a9d5da411626d126ab00443fc815765fa7737d5bdb4e051153df8a9d6f67b5591dd
