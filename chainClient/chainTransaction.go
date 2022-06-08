package chainClient

type chainTransaction struct {
	txHash     string
	txValue    string
	txGas      string
	txGasPrice string
	txNonce    string
	txDate     string
	txToHex    string
}

func InitChainTransaction(
	hash string,
	value string,
	gas string,
	gasPrice string,
	nonce string,
	date string,
	toHex string,
) *chainTransaction {

	ct := new(chainTransaction)
	ct.txHash = hash
	ct.txValue = value
	ct.txGas = gas
	ct.txGasPrice = gasPrice
	ct.txNonce = nonce
	ct.txDate = date
	ct.txToHex = toHex

	return ct
}