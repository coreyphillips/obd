package service

import (
	"github.com/asdine/storm/q"
	"github.com/tidwall/gjson"
	"log"
	"obd/dao"
	"obd/rpc"
	"testing"
	"time"
)

func TestDemoChannelTreeData(t *testing.T) {
	PathService.CreateDemoChannelNetwork("A", "F", 6, nil, true)

	for index, node := range PathService.openList {
		log.Println(index, node)
	}
	log.Println()

	for index, node := range PathService.openList {
		if node.IsTarget {
			log.Println("findPath:", index, node)
		}
	}

	//for key, node := range branchMap {
	//	log.Println(key, node)
	//}

}

func TestDelDemoChannelInfoData(t *testing.T) {

	info, err := rpcClient.OmniGetinfo()
	log.Println(info)
	log.Println(err)
	return

	//multiAddr, err := rpcClient.CreateMultiSig(2, []string{"03b2e7ecc5ff62feb342943a1364f555e8302f507f78c6392c82b9e12c95ccb40b", "03f1603966fc3986d7681a7bf7a1e6b8b44c6009939c28da21f065c1b991aeff12"})
	//if err != nil {
	//}
	//var  RSMCMultiAddress = gjson.Get(multiAddr, "address").String()
	//var RSMCRedeemScript = gjson.Get(multiAddr, "redeemScript").String()
	//log.Println(RSMCMultiAddress)
	//log.Println(RSMCRedeemScript)

	hex := "0200000001b196f5b373c6b8d1f3b688d14d865bdd33366dae7cc4e12e1f2691ad710aad7c000000004a000047522103f1603966fc3986d7681a7bf7a1e6b8b44c6009939c28da21f065c1b991aeff1221023cd1ad5e232182d51ef75bc9fb4821a86e3b057a34f7ec30de0ed138681d133a52aeffffffff03241d00000000000017a9140b727d44cc8eba2f9e81ac040c54e68376905685870000000000000000166a146f6d6e6900000000000000790000000017d784001c0200000000000017a9140b727d44cc8eba2f9e81ac040c54e683769056858700000000"
	hex = "0200000001b196f5b373c6b8d1f3b688d14d865bdd33366dae7cc4e12e1f2691ad710aad7c00000000d9004730440220156c3ef2449fcca3f5df395bcd5ada67c3fe3cd025b68713b0178803bb91543d02203c57a136c58d167a65a7bbf5eda5d2902e0eca284eac34454daa6d0fd5784c2201473044022054d6c37e39ceed3bcf693201b942193fab6784217869108374caada7ffabca6f02205cc212ed336263524bebd1419ce04e8fc19940865a24a6d430ebd87efbe934480147522103f1603966fc3986d7681a7bf7a1e6b8b44c6009939c28da21f065c1b991aeff1221023cd1ad5e232182d51ef75bc9fb4821a86e3b057a34f7ec30de0ed138681d133a52aeffffffff03241d00000000000017a914eb11c3cf8ccae06e28ed5d081a94c53487991e9e870000000000000000166a146f6d6e690000000000000079000000005f5e10001c0200000000000017a914eb11c3cf8ccae06e28ed5d081a94c53487991e9e8700000000"
	//priv := "cSyJChW5xdhhWCwdSFvWuNC5SgwYqgpWwbVXXKE5gs293BwnM4q5"
	//prevtxts :="[{\"txid\":\"770099f1b24bc17783852932c7f92badcdb0dee83b495bc3c4e71e0624172aaa\",\"scriptPubkey\":\"a914eb11c3cf8ccae06e28ed5d081a94c53487991e9e87\",\"vout\":0,\"value\":0.0000746},{\"txid\":\"770099f1b24bc17783852932c7f92badcdb0dee83b495bc3c4e71e0624172aaa\",\"scriptPubkey\":\"a914eb11c3cf8ccae06e28ed5d081a94c53487991e9e87\",\"vout\":2,\"value\":0.0000054}]"
	inputs := make([]rpc.TransactionInputItem, 0)
	var node = rpc.TransactionInputItem{
		Txid:         "770099f1b24bc17783852932c7f92badcdb0dee83b495bc3c4e71e0624172aaa",
		ScriptPubKey: "a914eb11c3cf8ccae06e28ed5d081a94c53487991e9e87",
		Vout:         0,
		Amount:       0.0000746,
	}
	inputs = append(inputs, node)
	node = rpc.TransactionInputItem{
		Txid:         "770099f1b24bc17783852932c7f92badcdb0dee83b495bc3c4e71e0624172aaa",
		ScriptPubKey: "a914eb11c3cf8ccae06e28ed5d081a94c53487991e9e87",
		Vout:         2,
		Amount:       0.0000054,
	}
	inputs = append(inputs, node)
	result, err := rpcClient.OmniDecodeTransactionWithPrevTxs(hex, inputs)
	log.Println(result)
	log.Println(err)

}

func TestDelDemoChannelInfoOne(t *testing.T) {
	err := db.Delete("DemoChannelInfo", 4)
	log.Println(err)
	err = db.Delete("DemoChannelInfo", 2)
	log.Println(err)
}

func TestGetDemoChannelInfoData(t *testing.T) {
	var nodes []dao.DemoChannelInfo
	db.All(&nodes)
	for _, value := range nodes {
		log.Println(value)
	}
}

func TestCreateDemoChannelInfoDataSingle(t *testing.T) {
	node := &dao.DemoChannelInfo{
		PeerIdA: "D",
		AmountA: 10,
		PeerIdB: "F",
		AmountB: 18,
	}
	db.Save(node)
}

func TestCreateDemoChannelInfoData(t *testing.T) {
	node := &dao.DemoChannelInfo{
		PeerIdA: "A",
		AmountA: 18,
		PeerIdB: "B",
		AmountB: 0,
	}
	db.Save(node)
	node = &dao.DemoChannelInfo{
		PeerIdA: "B",
		AmountA: 16,
		PeerIdB: "C",
		AmountB: 0,
	}
	db.Save(node)
	node = &dao.DemoChannelInfo{
		PeerIdA: "C",
		AmountA: 14,
		PeerIdB: "D",
		AmountB: 0,
	}
	db.Save(node)
	node = &dao.DemoChannelInfo{
		PeerIdA: "D",
		AmountA: 12,
		PeerIdB: "E",
		AmountB: 0,
	}
	db.Save(node)
	node = &dao.DemoChannelInfo{
		PeerIdA: "E",
		AmountA: 10,
		PeerIdB: "F",
		AmountB: 0,
	}
	db.Save(node)
	node = &dao.DemoChannelInfo{
		PeerIdA: "F",
		AmountA: 8,
		PeerIdB: "G",
		AmountB: 0,
	}
	db.Save(node)
	node = &dao.DemoChannelInfo{
		PeerIdA: "G",
		AmountA: 6,
		PeerIdB: "H",
		AmountB: 0,
	}
	db.Save(node)

}

func TestPathManager_GetPath(t *testing.T) {
	multiAddr, err := rpcClient.CreateMultiSig(2, []string{"039ee94a8131ef437059383bd0bb9ca3b7fd9cae0554f9df8b7d786ebf173f1c20", "0216847047b926a1ff88e97fb0ebed8d0482c69521e9f8bc499c06b108a4972b82"})
	rsmcMultiAddress := gjson.Get(multiAddr, "address").String()
	rsmcRedeemScript := gjson.Get(multiAddr, "redeemScript").String()
	json, err := rpcClient.GetAddressInfo(rsmcMultiAddress)
	rsmcMultiAddressScriptPubKey := gjson.Get(json, "scriptPubKey").String()
	log.Println(err)
	log.Println(rsmcMultiAddress)
	log.Println(rsmcRedeemScript)
	log.Println(rsmcMultiAddressScriptPubKey)
}

func TestChannelManager_AliceOpenChannel(t *testing.T) {
	hex := "0200000002aa2a1724061ee7c4c35b493be8deb0cdad2bf9c73229858377c14bb2f199007700000000d90047304402202bcbbd9b87621fcc73773504340af87614c33aadabea0270d7e4cac30b38ad74022015659132e1d6007d27a14c3cc5a45a18daa1faf0d08e6a1a507068f65477b4ff014730440220055d1e8422451f510c0e3f94fd86ae9cc6d8dcc37d90ac875ad10f5e8e47b83702201c27c9990826cd70a0e9a8054680d6c7e13d94bddecfcbfa098ef392a9bd5c060147522103985e8880628058da7c49b0968e4e7d2819240b60255a1c9b5f2407a4056b5f5421023cd1ad5e232182d51ef75bc9fb4821a86e3b057a34f7ec30de0ed138681d133a52aee8030000aa2a1724061ee7c4c35b493be8deb0cdad2bf9c73229858377c14bb2f199007702000000d900473044022078ee3af2fc5465cfda6e54843a47c84a8b17cc6be65bfdd5f901ac25af33e5b00220162e156ac04bfc270731dab7b68e1c753eba9ee565b137f9854a8003c65a2e9001473044022079ba94d2591d885d8d815476a88467ec13c589ed8518c03fbcee0e7ec4bf39410220236f3925c958a825ea1dadd20e7111152acd4ceb8d03bf6e94fbcca0a47c30e90147522103985e8880628058da7c49b0968e4e7d2819240b60255a1c9b5f2407a4056b5f5421023cd1ad5e232182d51ef75bc9fb4821a86e3b057a34f7ec30de0ed138681d133a52aee80300000366110000000000001976a91484902564ba3ce47952d86a0d53c17402b3cce96588ac0000000000000000166a146f6d6e690000000000000079000000005f5e100022020000000000001976a91484902564ba3ce47952d86a0d53c17402b3cce96588ac00000000"
	result, err := rpcClient.TestMemPoolAccept(hex)
	log.Println(err)
	log.Println(result)
}

func TestDecodeTx(t *testing.T) {
	hex := "01000000000101534de8c19bcf4e889da37be7ee1b331394f02fbca89ec0c8ab36b90405123bf012000000232200200816d6760892870b603700429562e9736cf02e6d2992a370d12475a3ae476134ffffffff0878ac0100000000001976a91437842906dc811ea031fab008cbda1e64ff8f79fc88ac5fac0100000000001976a914ae616938464bc8b20e223d77faa85b786236bf6588ac069d0300000000001976a914c921b3c4b35d8c6563a22bf70aa3b9e148c3c56c88ac68d06a000000000017a91459dce0de30fb800bf7a0edbd5fdd4d55ee6f671387f4ee1e000000000017a914917a7bd69821981b6a47476500e3d124ac1a971587881c15000000000017a914df4d0bbe607ba16b55b19acad4fab2e50688256587cc771900000000001600144be986e60ace60b0ca514ef5268d0f10a68a29bbd395ae000000000017a9146e396886d8e981923bcf9627212c68f8c831ee0587040047304402207db7b7a1cd9242cab6b7a566090f01088475095a15e0d2e211b5f8d8e4517b46022069c1768b50965c57d7d38c995bf469c4d7c378423fc9348e8901f05e668ac28e014730440220445965b4b3afdc5d2c65d282c5313539f47cf9291bb3adb08977c7666e800b5c022067e04187a8edbb2be17a87f5165b08844ca9d541f12318cb02c5d4019d432764014752210351a7ded3bffafc00f74bbb94d7749d2460c328721e4afd6e5edd7d760b87441d21039215c3a527e0d59335709a3d675687bdac618c40d1a914f68d1609972566f82052ae00000000"
	result, err := rpcClient.OmniDecodeTransaction(hex)
	log.Println(err)
	log.Println(result)
	result, err = rpcClient.DecodeRawTransaction(hex)
	log.Println(err)
	log.Println(result)
	//txid, hex, err := rpcClient.BtcSignRawTransaction(hex, "cTPtw7uhNXWeBroEzFur3WZQr8WgPojE4WipsxTNBqbGsruMJG33")
	//log.Println(err)
	//log.Println(txid)
	//log.Println(result)
	//hex = "02000000012ecd34ce812f36a876d6f5b3ab2ccb3478eea69e6af4a337fb1941ae8b8a625100000000d9004730440220514b3ed6d636c69b2c936f9a57ecc248f00618c46e61bee5e8408192c4a25570022045d1c1b191a9b6c4ee56a51129baf9d79a26852d5ccbe572d8fcd961b360e8c80147304402200cad5e9707d489534823a5e100e6e020d555d6ac87cab285b23c4fc706eb689402204c7d3962ad0cb27c81295dbfaed216325c26a1b099ed328eefff81212501903a0147522103f1603966fc3986d7681a7bf7a1e6b8b44c6009939c28da21f065c1b991aeff12210216847047b926a1ff88e97fb0ebed8d0482c69521e9f8bc499c06b108a4972b8252aeffffffff033c1900000000000017a9140ff6b304e80589566854573a3c528ee0cb7dfbe4870000000000000000166a146f6d6e69000000000000007900000000773594001c0200000000000017a9140ff6b304e80589566854573a3c528ee0cb7dfbe48700000000"
	//result, err = rpcClient.TestMemPoolAccept(hex)
	//log.Println(err)
	//log.Println(result)

}

func TestGetBalance(t *testing.T) {

	address := "n362wgEWVqbymxYjVkkq7jLQjQdeW93Ncc"
	//address := "n362wgEWVqbymxYjVkkq7jLQjQdeW93Ncc"
	s, _ := rpcClient.OmniGetbalance(address, 121)
	log.Println(s)
	balance, _ := rpcClient.OmniGetAllBalancesForAddress(address)
	log.Println(balance)
	balance1, _ := rpcClient.GetBalanceByAddress(address)
	log.Println(balance1)
	result, _ := rpcClient.ListUnspent(address)
	log.Println(result)
}

func TestSend(t *testing.T) {

}

func TestCreateAddress(t *testing.T) {
	address, _ := rpcClient.GetNewAddress("newAddress")
	log.Println(address)
	result, _ := rpcClient.DumpPrivKey(address)
	log.Println(result)
	rpcClient.ValidateAddress(address)
}

func TestChannelManager_Test(t *testing.T) {

	hexC := "02000000012d6e0d747667b159ef24d731a3f1f6bdc14d83c241087dc898aad85cc4a14f4400000000d900473044022065deb2b5c9ed9a41383c8300a81a95cf416bdb16616f35a12a99a38b14d02e98022054af1192e8e6c9a9aeaa5595a4659f93bd1011d8723fe57d448d5c3793ff26240147304402206d628f7405999bbc21a8cfe1a530b6969d96f7b648302e419f22440b79412759022044a30a9aa8bacb3371fc5952eabf50265115b12d36511062cc1089268e17c4ae0147522103b4df7d3026a437f537dcc0a9e681ffdfb000c9f1223189adf18364588d46e5592103c57bea53afd7c3c2d75653ca35ca968c8e9610b6448f822cfb006730870ee96152aeffffffff033c1900000000000017a914faafd20558ca1529b96fb1cdd40e4ef1915ed1e4870000000000000000166a146f6d6e690000000000000079000000001dcd65001c0200000000000017a914faafd20558ca1529b96fb1cdd40e4ef1915ed1e48700000000"
	result, err := rpcClient.DecodeRawTransaction(hexC)
	log.Println(err)
	log.Println(result)
	//hexRD := "0200000002952b9f91f48d53637685c88791ea4ea046c637fdca988662f322f3bbab88cf8600000000d90047304402204a33088ade500d6d0a40051231b4223186809b855cd8e534a0f4ccda2b90a8e302200657eea4dbdf52573402156d30125b6048e75e5d4376af16a262187aa99a0d6d01473044022042cea8d1389fb2c844f0e59ca9992943af962c973f424ffafef89f415934225202206ee20104936fbdb0fbc25a4bf6f6629b4b89a672e6ea4cb1070e50679920f6450147522102b8302d22a50fd84f34d528ff98998a6959bc7fb8f45b5f3fb44e23101aa5d8f22103647e81480a71989ee5c3391763d6aac445bb104f0dce688002c18a3bba6ed42b52aee8030000952b9f91f48d53637685c88791ea4ea046c637fdca988662f322f3bbab88cf8602000000d90047304402206dd99daaa88f50c3403bcba536e50f8d3288e2ca5e9fdcd97ba8d3bf9520384a0220158f12b5c43fec31c51ca7597d9dabaade62eec5798c7197db163691177865be0147304402205fe693c8ea0dcf48692fc51dec95913abde3c64084fb47bfade8291ffd9a4bef02202d571fe14754347c9153582ea6f744581fd5a30d147bbc63b2aa2704f89a29ee0147522102b8302d22a50fd84f34d528ff98998a6959bc7fb8f45b5f3fb44e23101aa5d8f22103647e81480a71989ee5c3391763d6aac445bb104f0dce688002c18a3bba6ed42b52aee803000003620b0000000000001976a914ec9c3fabfa57c7862ba594b70988a7b4f485744188ac0000000000000000166a146f6d6e6900000000000000790000000011e1a30022020000000000001976a914ec9c3fabfa57c7862ba594b70988a7b4f485744188ac00000000"
	hexRD := "02000000023cbcb3fe0254dd2a4ac81172e1de5520310fa0cea134ec827e6dbee153b10d7000000000d900473044022028c6c2f4de0baba040904b266b548020ee64f266aef4e00fa18f209ffecdd23e02207c96c5986a15a99b5f66c5e9d80b0e32140fc14d72dee23782ac47a6258da3d40147304402203edbc9b302eb3296090727472a35430da57cce8f2c8a2b10da69638e211df9360220117203aaa265440d5b61faed5926e92c5230753bc2348c528f553ab802d32fc60147522103ea01f8b137df5744ec2b0b91bc46139cabf228403264df65f6233bd7f0cbd17d2103c57bea53afd7c3c2d75653ca35ca968c8e9610b6448f822cfb006730870ee96152aee80300003cbcb3fe0254dd2a4ac81172e1de5520310fa0cea134ec827e6dbee153b10d7002000000d900473044022029bb9797f6e4030120e614ed1153f8649f3aef6cc7c94fbd1a97c095ba55d882022013575c341f7cc69269407ac027097b2e9997631e6bad6df6508d60649fa1c4ff014730440220078e235854f316419ddd7801bc8737817bc51b25a8977709d0eae3c0d88eb63d02205af3887519aafe49a4acd833566f00f01b95d4fed363cc46df5eb0f5dc28a2270147522103ea01f8b137df5744ec2b0b91bc46139cabf228403264df65f6233bd7f0cbd17d2103c57bea53afd7c3c2d75653ca35ca968c8e9610b6448f822cfb006730870ee96152aee803000003620b0000000000001976a914ec9c3fabfa57c7862ba594b70988a7b4f485744188ac0000000000000000166a146f6d6e690000000000000079000000001dcd650022020000000000001976a914ec9c3fabfa57c7862ba594b70988a7b4f485744188ac00000000"
	result, err = rpcClient.DecodeRawTransaction(hexRD)
	log.Println(err)
	log.Println(result)
	hexBR := "02000000023cbcb3fe0254dd2a4ac81172e1de5520310fa0cea134ec827e6dbee153b10d7000000000d9004730440220051b1a5d236d41efcd5b40613c6f0621b5870b61cb6cf0e7fcf1eb11ca25210902203280ac5f7f195e2a2f411bbf7c8ce09b64097af1d3361c19985916109edf363a0147304402200d1e967ee3105f2e5fcd2ecf81244b37c9e93362c2d55506a9dadcd9282f2c7602202fa88cd04286f2fea5336b03fb4238b06f955810384dff09d1858560456ee5090147522103ea01f8b137df5744ec2b0b91bc46139cabf228403264df65f6233bd7f0cbd17d2103c57bea53afd7c3c2d75653ca35ca968c8e9610b6448f822cfb006730870ee96152aeffffffff3cbcb3fe0254dd2a4ac81172e1de5520310fa0cea134ec827e6dbee153b10d7002000000d90047304402201df6b76e21bcae0ed0883c06e7f46e1f4e930fe6eae85394ecfb5b29e0e1e0bc022008272bd938edf2a0bfcd10446431dfcc3452bcbb2fef0cc436aeb91bf44732f001473044022029988dfeca46b9a4592922953e56bfbc43d84cedc2ba21c1d4206c190f8735fc02200479e5f45ad6ddbb0ff5b8dceb1491379ca5b1580ef1e80401656d60e422e62a0147522103ea01f8b137df5744ec2b0b91bc46139cabf228403264df65f6233bd7f0cbd17d2103c57bea53afd7c3c2d75653ca35ca968c8e9610b6448f822cfb006730870ee96152aeffffffff03620b0000000000001976a914ec9c3fabfa57c7862ba594b70988a7b4f485744188ac0000000000000000166a146f6d6e690000000000000079000000001dcd650022020000000000001976a914744846d33d79479478c2858c008ad93f77c2259d88ac00000000"
	result, err = rpcClient.DecodeRawTransaction(hexBR)
	log.Println(err)
	log.Println(result)
}

func TestTask(t *testing.T) {
	log.Println("aaa")
	node := &dao.RDTxWaitingSend{}
	node.TransactionHex = "111"
	node.IsEnable = true
	node.CreateAt = time.Now()
	db.Save(node)

	var nodes []dao.RDTxWaitingSend
	err := db.Select().Find(&nodes)
	if err != nil {
		return
	}

	for _, item := range nodes {
		item.IsEnable = false
		item.TransactionHex = "33333"
		item.FinishAt = time.Now()
		err := db.Update(&item)
		log.Println(err)
		db.UpdateField(&item, "IsEnable", false)
	}
	var nodes2 []dao.RDTxWaitingSend

	db.Select(q.Eq("IsEnable", true)).Find(&nodes2)
	if err != nil {
		return
	}

}
