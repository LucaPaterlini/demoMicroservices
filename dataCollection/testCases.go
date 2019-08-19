package dataCollection

import (
	"fmt"
	"time"
)

type expectedRetrieveTransaction struct {
	Status int
	Header map[string][]string
	Body   []byte
}

var testCasesGetTransaction = []struct {
	blockNumber    uint64
	index          uint64
	requestTimeout time.Duration
	expectedError  error
	expected       expectedRetrieveTransaction
	description    string
}{
	{
		blockNumber:    1,
		index:          1,
		requestTimeout: time.Nanosecond,
		expectedError:  fmt.Errorf("net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)"),
		expected: expectedRetrieveTransaction{
			Status: 0,
			Header: nil,
			Body:   nil,
		},
		description: "testing for expired endpoint request",
	},
	{
		blockNumber:    1000000000000000,
		index:          1,
		requestTimeout: time.Second,
		expectedError:  nil,
		expected: expectedRetrieveTransaction{
			Status: 200,
			Header: map[string][]string{"Content-Length": {"38"},
				"Content-Type": {"application/json"}, "Vary": {"Origin"}},
			Body: []byte(`{"jsonrpc":"2.0","id":1,"result":null}`),
		},
		description: "testing for not mined blocks",
	},
	{
		blockNumber:    1,
		index:          100000000000,
		requestTimeout: time.Second,
		expectedError:  nil,
		expected: expectedRetrieveTransaction{
			Status: 200,
			Header: map[string][]string{"Content-Length": {"49"},
				"Content-Type": {"application/json"}, "Vary": {"Origin"}},
			Body: []byte(`{"jsonrpc":"2.0","id":100000000000,"result":null}`),
		},
		description: "testing for not existing position in an existing block",
	},
	{
		blockNumber:    8368161,
		index:          1,
		requestTimeout: time.Second,
		expectedError:  nil,
		expected: expectedRetrieveTransaction{
			Status: 200,
			Header: map[string][]string{"Content-Length": {"601"},
				"Content-Type": {"application/json"}, "Vary": {"Origin"}},
			Body: []byte(`{"jsonrpc":"2.0","id":1,"result":{"blockHash":"0x4f56d43f13bee11e6ca9739d326e3935428bf1ceaf5b78c211f38709b561e269","blockNumber":"0x7fb021","from":"0x5e032243d507c743b061ef021e2ec7fcc6d3ab89","gas":"0xafc8","gasPrice":"0xd09dc3000","hash":"0x7adfcf6e2947590cb88763af73c923f1d1832c07830a8891199832e506103430","input":"0x","nonce":"0x7aeae","r":"0xe27a9cbd121e2d7aaa6c806591d183c4c5c766c24bb9aace8c2f968fe7805735","s":"0x71df467a247448a2a0c4cf131d1f9e938998aedba43665f0aca241ad5133fed3","to":"0x3b4c009fe957d58626efb439b463fccbe7538ab7","transactionIndex":"0x0","v":"0x26","value":"0x169ffd365951c000"}}`),
		},
		description: "retrieve existing transaction in an existing block",
	},
}

var testCasesGetBlock = []struct {
	blockNumber    uint64
	requestTimeout time.Duration
	expectedError  error
	expected       expectedRetrieveTransaction
	description    string
}{
	{
		blockNumber:    1,
		requestTimeout: time.Nanosecond,
		expectedError:  fmt.Errorf("net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)"),
		expected: expectedRetrieveTransaction{
			Status: 0,
			Header: nil,
			Body:   nil,
		},
		description: "testing for expired endpoint request",
	},
	{
		blockNumber:    1000000000000000,
		requestTimeout: time.Second,
		expectedError:  nil,
		expected: expectedRetrieveTransaction{
			Status: 200,
			Header: map[string][]string{"Content-Length": {"38"},
				"Content-Type": {"application/json"}, "Vary": {"Origin"}},
			Body: []byte(`{"jsonrpc":"2.0","id":1,"result":null}`),
		},
		description: "testing for not mined blocks",
	},
	// TODO: investigate why in this case the header do not have the content length field.
	{
		blockNumber:    8368161,
		requestTimeout: time.Second,
		expectedError:  nil,
		expected: expectedRetrieveTransaction{
			Status: 200,
			Header: map[string][]string{ //"Content-Length": {"800"},
				"Content-Type": {"application/json"}, "Vary": {"Origin"}},
			Body: []byte(`{"jsonrpc":"2.0","id":1,"result":{"difficulty":"0x872add2075a0f","extraData":"0x5050594520737061726b706f6f6c2d6574682d636e2d687a32","gasLimit":"0x7a2125","gasUsed":"0x79e4f0","hash":"0x4f56d43f13bee11e6ca9739d326e3935428bf1ceaf5b78c211f38709b561e269","logsBloom":"0x8c0452a80041904010009a95851140a800080338c062c83b90b33344282b1120510ab0610305000843c321039240532e4ec426242e4021130360162b202e343c4b85608d880f881138061258048b5ba02e58500b0a4d831c015483250e40302c002a5a020a6090101614100000401c4208a6d00aa820c400100c1010d2a00013141533d8ace1d00320802448d148b0168810c0a6400875039d051405e07d396356108e80885250005db607a004d8886261d1742b3be0984069631ca1522a10412028c1aa001003402a4a1c1d4e8937010c00300393fac116249652265b1074506478a05221488a0c01246a80802086240321a0dc109804ce057c698c90000497","miner":"0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c","mixHash":"0x5e2410500963ae68a1fb5cb0766681c9ab7e5b9295e01c7e369f73b1b4b3689c","nonce":"0x4baf65880baa1c96","number":"0x7fb021","parentHash":"0x57e152d17544bec89a9c329cfe875cef5ea407a9039b50a8dcb2564ec8b3866a","receiptsRoot":"0x045ef07fe753a826067c96bf92069a350ef99441ecdaa7dedf4bc2140c1579f8","sha3Uncles":"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347","size":"0x7129","stateRoot":"0x97bf82d43c7bc9f999da536cd1b71cc9e7330578f9e64377227bc64cdcb369c6","timestamp":"0x5d57fb9f","totalDifficulty":"0x26faabe35e84680a6d9","transactions":["0x7adfcf6e2947590cb88763af73c923f1d1832c07830a8891199832e506103430","0xfa3d8d35da7193e21d351935ca47fee233f52b886aae1d7387745661ad12c119","0x2e2c278bb8dbe59105ae24bac4a2c2af368301971ccd7663e45ad6b40286f647","0x9ebf16358167c4b252baa11fe1454833169dd9f117de2b099e301adab284d8c5","0x3d1470b2811590299f4d67d78554e52f9445346cb2f2252c1724f4b096640f17","0xc67086f0191f9befdbfc57f06059b18372f5483fbb0f7a87f16ffc16756f7f94","0x83ccd943a73b51c546d421559b20c56c2b8275341c37588ee89d452c93f60b4b","0x86979c772d623b7b065aa9813437770d03fbad0698b7dff189bb6ec67485b93e","0x84d8596f5c526a84ba0ad8c54e0b3cf573f4f7f31fc31d2392121f556feb9850","0x5ba862c048ea59fb17290564e685b1039788fc304182ea02a2ca7874ce1fddf7","0xc99dacbca36d5bbb5f52a38bc745e3935a09f71bc2fd1e6123f5e2a2fcc1f716","0x52f29c6848677a400a72f3423ad33401b4f600180e58102bd9b5370a485c5443","0xa2ed4b9c6ef6dd98d18932069b5e07b01b5dbd4be5172a75d4e165989ee4d0dc","0x21ca366e4fe3a18167e84a8dd3eb1c1f8b7825b4c092024b8d2b0e604afdd954","0x854d6405c3d2edfb4c0eaa3361aba9c61c3f220cc831122699e32d68b3381a90","0x888da16ed24839f6c90a6d5c820885809e09c57475abc43b180b27f9b6ec0a48","0x13f214321b394bd10bf27df1e397c10c5ae88746e8c1dde22ee89534919a6d17","0x95c22de309d7ef7d02202d028676da5cd67f0aabeed962bd093b84bb32c0193f","0xb64f4fffb424274a79f6bd8742cbdba75eeb30a2f4df9749ad5c36d74e25f0f4","0x08b8cb6425b5f62282ecca5f531f88f68748716c7f4c44817e899e131367e73f","0x8edac80fb427ca37953881e6abf1949c329aea1c25b2181ad3590dc7afc59b2b","0x67086d1d40a9e0d0a8b8b7a8bbe43b7e60b4542db1a95f022b91c495cbbdaa77","0xc60c3fa49c1f0079e76191c86a4283ddbf5a1626cd3478e878450f79105cc0ea","0xa9de3f782ecd0a05b85bd16e97646d8e8829ea779836253bc3368c90f815dbde","0x4fec384fd996ef67f4d957c5836d357190dfddab7c66f364c415b27f63fd191b","0x39dd796fa1bbcb711268af458eb7e4414a1e714188a9ecb704e052a4a5077e1f","0xba2890f6a148f55b6d5a21160fda26163c2459f8b460ebb9f1dee3fc4d629ccc","0x025cde625512bac0df206871ee7ad30d984f1bb14c5c93d9ba0c6008cb6541b4","0xf2c34c4a471c304db5065c23b1ec006eb621c62120c50ef3ac77003e377f651d","0x2e39ccc6307bc18894540b57958f7202cec2570c5e94034408dad8dd36ee6a0a","0x7de7ac0609d21e0b75c4a7d0e139213bd8c78d4cdc94ce89cbb66738526ec958","0x34a2bf02dad2d93d53afab14348c1c1bf2ca863a2707455b38614ed11aa27e45","0x593a7063d6b1ab0a773a1e41ae5b217a10aab7ca1e547d2455e2b22466bb52d7","0x40dc3173c211a4e5c6f33ed6cfb2516a7abee505fb71bb88aaaffd849fc2b2aa","0x72c8010a034f1a854d396dd94fb04354c019d0f02c86435872ec1a365a930df1","0xef424d2d8412965bacb625daed1ef0a6e81b008fd39a7b6b66d48c3fa78edbe5","0x10959b2a0c34ea664d782483ac7166127bb772477ac4fd9990fed1991a8b1109","0xdc1eada7a131e24b8fbea0aa7a3852f19174508a0117adff5d0f73253736c91e","0x37ae8536df7caa2330c40c39da09548e9e42d2ae2c8d5fb43ac1000eaa850f04","0x1f9d8524672fddcb356ff8b4a298d49161987db755ecf11a4eaed1d734a511a8","0xfde5812ebeb06837cdca3b0789fa5cb6a95c5000e98c62f2a48735fb4e9c0c5a","0x371628d11919765a5937b158ab652b0b609b91361b698694d3c854e7307f1f30","0x73dda1ed8e7e55d6b1789b52025d8384f6e0f72cbfad60673805fb0cf47605dc","0x2376e669d37912d78c726ae1f5c416d23d0c86707641c43f57799b7a606c3367","0xe5f456a56d65beaaa5e652b847e70103eb5f922d043553a9a4351c9bc4efb45e","0xaa29cd507ab5dbe5cab665c25f01b1901565d28aa63097448b519410fad2e15d","0xd9edcdb14b685ca2182802e3e16ed1607d8973ec2e357b4853ba8918473cd97a","0xf851ca0a2cfb5ca00fc8030046a8a3a6bfed07ec9066e5f5ae407c83d2060abe","0xf1d9529dfbfa6f0b7910df2d506aaa67b27a8210cf781f87d533ab172f5d40f6","0x89a0ad103e91f68498741b9a07a199b99bb0dc83e749e74107262c7865124b95","0x375d2d94c13f56b84f83b908b701e4546c3c3f078e68122a47ee3e11f5f32061","0x8f789280c9019a9e52b0802cb15d47405da5b645dd7e928649f4ae5b4f4ddce0","0xe7e20b6aa453d22952ff7487bed1da9cdfc5313fc994a47d87fbb14ec613351f","0x1580d4321d96558bfe8196560ce8da4a08e0b3ad176241d23ca8f57b94a37e1d","0xf172848db0b9f267ce97c794150095a5bfdc4b5154a1463f2c37fc75eb068987","0xe7d4dd08bee52785dacd5e73cac7dcb5a38c0d530a344ae8193366f9d5155de4","0x32953057750937963cd805c63eac4368fdd87d5f743b41790773f702f614a46c","0xa241e940c98d96af0bfa3e9d0670ce2c69e699ec5bb1ef193a8ecb387405935f","0x72923c5d64642b14284c381da95c3d2bfd469341199c898e8b8dd5f6c978b35d","0x319890d6d905c82215cb9ee4bdd172a48b116d409b2d71296cdcc18990827309","0x6e0628b7d7fc06ec50bd4f48081f63edccca8beda5949154fd24bfe4c9cb7675","0x035128f08dee5877f74c1e7d909e60349986ee500d2740f456b798021b008510","0x1ddb150d259ccdc0de1a6558b43160d3647c6aa840dd85e7877daf30a0bce9bc","0x207e86b748d5b4dcad08d08d59eee68ace84e07a226964b23e2788aa2a747f8b","0xd2ea017575c513fda6c99c0092d449e569dbe56b7d47bf2535e8cf94f6dee15f","0x65f34c0f76b9c74eddd203fa6da7a0cc5e076474b0cc023d527ea9445617a507","0x15874b4b71ed6b0e76cc216c2411533b08fb18f87767583c9fb7d58630b93a7f","0xcaf75344ba08db2e399c0259effa5df51d0f71404ee2164a44ba0bebda2860b4","0xf107aab6fd997b7e60a47e86694a56b27ab2dbd4db26fee528634441110258f1","0xd76dadfcdac18456d1c82c8b4847825553a1818ebadd1521b7218bbf03a29528","0x337a26259bdbf111ccde199badd847351ababb38c84890c406a2bde192e90231","0xfbc0e08a345d0f09b28acf1a386f0588592914eea6cfcd5ee4bc571891ac98fc","0x7e3dc0f5a4b2f075ec58c0e663e643bdcb8dd90b8c77499e7b2093971ae91f73","0x9e51653acfcbb80e27c5ec7f3d8cf9065b757dffb9f50fffbe7af46eae139197","0xe3860c3ee6cfe08decb8d49e94a9f518b0e9d91a3222d57e9ff287c24543f442","0x6d96ee0e8cdb0bada47dcb5da2bb3639a315f024f78c823d1c2d44ad771cae5e","0xfad531ee6446f41a2faf767ad052c53415e6464c6d56db1758fabcd07f6d5dbc","0x09dbe1ea82b1b355625263d0b225d2f132f60012a6ccbd44718dcc8c6842b13d","0x845e729f872aee93d42650c907f59afcfc7d668166085cf86af5acbf9f25b107","0x405e900467484d26ce3d3ccd9d8e8ae9234be74f06336bd3882b7d50965fe6ad","0x18e874584aebf8808799f80f6ce86bb6a9e2fe5fd09ab6ac2bbd416e7a866839","0x0389af3eaec69e56556b060eb5e7bbe00e8cef47adbf016386e414c9489eac5f","0x0fb202d051ea795a5c3c8ed206209c5e2ea629133e284f68a3451b15da44e609","0x34157aa825138b9326cc5e66c0d33b64edc5bc90531901b8af0618506b96325a","0xfb9db1da8d3615b7a82ad0e3bbf30ff4ccbb4a6b046d8a32dcd29537ec497775","0xa371439ecfff710bcf16e46438b89cce99b8dc29c9667741b68a42429f3ff513","0x752bd181ae7afb9ff10a6300eb1daf12dde99c7c5147d27632b82e5f037c2d26","0xa30bf519a797d5e9e0be7b510135c6ed3341ad80e47be1a3c99429d437ed6825","0x9fd91b4c96ebbed42b00741c4d07e11f7b25b23d104e1f89395fa94279c2121d","0xa86b484ff2bd0e414fc9421e3f4ebb2e707689c5a8989561756f23b8b6379b3a","0x97dff2c4dbe7d2cd48296a1e945eb6a8208e9d726a31133296a137784470241b","0x83152f23ed21b33226699d3cd20efd889d87264427aac666f94d333532921443","0xba2e158c21114fba69f8e4171cd707b45dd13dd8adeb379e317867337dd8e5b7","0x8734bf3cb3efed35b672aeed9c7562ec643274b930cf440a8471b1a46a2a2915","0xdf46fa299edcee30fe70e5088976f75abfec9b893bfe09dc4b1d6ded470c5468","0xd6660bb982c0acb24c9521ffdb05190b25160bdd4d7422193299d20a4efeb4d6","0xc47c9ea2077daacc48ab017beb9902fb26a96e31527fbae93d94528da9540f12","0xc6c5093f439b9d1fb29765e2e5f832624fdc0afaddb58a480ad2626222a728b0","0x3db738b097573c0126a65bb4ebd5b913425219276dd7efe6dded5fb3aeb3f11a","0xb51171fc384b953a2fe5cc2bf4963f8679e6d6df138ecd4d5772128c3fb24e5f","0xa05ec398f5c0bdff286226fa07cf9e5686e4f964e94d82992751fec2802f82e0","0x832b5397c60c70d6522b52077e74b5bc9897f83b97b27a8452098b00193a2235","0x940458204c5afcec277cb684f792c0b2128db8f027d2a0a83504d1aa9683b716","0x1f3f349a17c1b1cd9c5a474714aed102c26cfd2939bc7ef0c42d0675f7613a94","0x1d64dfd9208be27c2c56641c9ac7df41a5f100e41e038046acbaace0a8ed6a46","0x58de497c730913261858a95992964682351d5eaffbb89edf8c5ff184144bb328","0xdf5565706c20c2eafe49afe7ca85eb5af6e75c5c508eb8d8108bb26e8a53bd70","0x54e9c5b3029ff36547d0026c11c06feeeda21f4159ad85f24985ba8ecd5694a2","0x1b8989b269d67dd180f5005e11e26dfc89faa382f8514ac5ad6e734c74dcb9ba","0x8f4b0d18624e260bc4241e508734bbb4195455ea84022e95df884c9269eccf48","0x1dca9f063b67b25b77cfb36d27beebe7211a37538262f069e7477828d7e6faf2","0xac703f061bda67ef18e52eac9c71825f02941b6b9cdb36f3472c17a00199d9fc","0x87199e9d152a080fbfde9a2021ab9544cae3243ded5f67330d76d49b4409a9b0","0x93f8757ff9a411425a08ad4247406fb21bf088601a7cbc58056c10d6ecfc93ee","0xbf180f4eab0072787e154358d6699065d1da7adff60f5f4a53d028d915c53fe8","0x0592c2d510df2ed34cba80f5feadee1af85be4d7a5d9c0df93b0d2242fbc5946","0xed50faf8ad51cad9a50013166a12a64822ecf9c9238611d65e8ab104b2470e76","0x3fe57a65f9659491c0dc405a6cd56434ef8574e8b27aa859dc4f5c70b2ce20cc","0x4db737daa0ab3471f824f8c4540d680dcbf79607454324c0df20996922ea0a5e","0x752d926daae51beca008b784c0fa30fb7413a0181f56682b5264f143a41705bc","0xf09c48d9e981bc6855862edc98074ab71b248b4509411f6c31b830d0f39dd79a","0xf6e0bd4c1d534371cfb321361553ed9885ff6abdda9b55a394a9908db2b01b96","0x92ed21e64a4bb4e1cc65c629a1eb78d1d3b52728d1dfa635086e8c1fad56ece9","0x82588bbb5fcccb217af030c7bb214edd2ce7648b4f09b5667579ba8b1760a2bc","0x065e898ba1cbd48d34ed18df64f67c4cdb07dc8e908db808cab8d288bb6f01a3","0xab0e46a1d7dbe7b7b2f94f8c7ccc5b3466660ba9ece7401276ae94777c7f3f2a","0xa24d0e976e884e76f2b946fce21023c2365caf56b96ac2a2b8a56a943f89cfcd","0xa806a328c53cafcba685a5d794222575c7ca306d5eb841ad186e3b216f68bf7b","0x17ef2934f39c301d6c27590aafbe82627f90f5ce99bf60a87b4d901c4b477400","0xe6e692fb1f9c7519a9e7aca8d01ea7691c6446bf9d8634971a60e6dbcdbc990e","0x10739ddd268884f8161924cd513ad6602159a5318cccfefe9b789f710c5e32f2","0x372a360c52538365819f275138ba888ec3078198959a945a13d59b839e1d2e8c","0xe1ddb8a0a917fc62452fde2f4e119367f210e01b304753b3a1a0ec02d65670ae"],"transactionsRoot":"0xa11f13c5ff70256a7b0f2f6c67600e1db49429c5a1b677a9eaee3419c99ce439","uncles":[]}}`),
		},
		description: "retrieve existing block",
	},
}

var testCasesGetLastBlock = []struct {
	requestTimeout time.Duration
	expectedError  error
	expectedBlock  uint64
	description    string
}{
	{
		requestTimeout: time.Nanosecond,
		expectedError:  fmt.Errorf("net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)"),
		expectedBlock:  0,
		description:    "testing for expired endpoint request",
	},
	{
		requestTimeout: time.Second,
		expectedError:  nil,
		expectedBlock:  8368731,
		description:    "expecting a recent block",
	},
}
