package main

import "Zilliqa/stake-test/transitions"

func main() {
	t := transitions.NewTesting()
	//t.UpgradeTo()
	//t.ChangeProxyAdmin()
	//t.Pause()
	//t.Unpause()
	//t.UpdateAdmin()
	//t.UpdateVerifier()
	//t.AddFunds()
	//t.DrainContractBalance()
	//t.UpdateStakingParameters()
	//t.AddSSN()
	//t.AddDelegator()
	//t.RemoveSSN()
	//t.DelegateStake()
	//t.UpdateReceiveAddr()
	//t.AssignStakeReward()
	//t.AssignStakeReward2()
	//t.WithDrawStakeAmount()
	//t.AddSSNAfterUpgrade()
	//t.UpdateComm()
	t.WithdrawComm()
}
