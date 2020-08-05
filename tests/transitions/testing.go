package transitions

import (
	"encoding/json"
	"github.com/Zilliqa/gozilliqa-sdk/transaction"
	"log"
	"strings"
)

const key1 = "e40afdc148c8f169613ba1bb2f9b15186cff6e1f5ad50ddc42aae7e5d51042bb"
const addr1 = "29cf16563fac1ad1596dfe6f333978fece9706ec"
const key2 = "8732034b0c895564d966e3df6968205211c7a2f0140b77c9e13de10c1ce77873"
const addr2 = "e2cd74983c7a3487af3a133a3bf4e7dd76f5d928"

type Testing struct {
}

func NewTesting() *Testing {
	return &Testing{}
}

func (t *Testing) LogStart(tag string) {
	log.Printf("start to test %s\n", tag)
}

func (t *Testing) LogEnd(tag string) {
	log.Printf("end to test %s\n", tag)
}

func (t *Testing) LogError(tag string, err error) {
	log.Fatalf("failed at %s, err = %s\n", tag, err.Error())
}

func (t *Testing) AssertContain(s1, s2 string) {
	if !strings.Contains(s1, s2) {
		log.Fatal("assert failed")
	}
}

func (t *Testing) AssertError(err error) {
	if err == nil {
		log.Fatal("assert error failed")
	}
}

func (t *Testing) GetReceiptString(tnx *transaction.Transaction) string {
	receipt, _ := json.Marshal(tnx.Receipt)
	return string(receipt)
}
