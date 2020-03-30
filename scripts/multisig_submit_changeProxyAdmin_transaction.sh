#!/usr/bin/env bash
zli contract call -a ${MULTISIG_WALLET_ADDRESS} -t SubmitCustomChangeProxyAdminTransaction -k ${OWNER_KEY} -r "[{\"vname\":\"proxyContract\",\"type\":\"ByStr20\",\"value\":\"0xd3f1dd0d24898123cc35ea7c56138784db915cf6\"},{\"vname\":\"newAdmin\",\"type\":\"ByStr20\",\"value\":\"0xca431681c07a5d06e609a906a881627650f31b69\"}]" -f true