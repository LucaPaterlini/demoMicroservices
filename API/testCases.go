package API

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"time"
)

type handlerTest struct {
	requestTimeout       time.Duration
	requestPath          string
	requestPathSignature string
	expectedW            *httptest.ResponseRecorder
	expectedBodyBytes    []byte
	description          string
}

var testCasesGetBlockHandler = []handlerTest{

	{
		requestTimeout:       time.Nanosecond,
		requestPath:          "/v1/12",
		requestPathSignature: "/v1/{blockId:[0-9]+}",
		expectedW:            &httptest.ResponseRecorder{Code: 503, HeaderMap: http.Header{}, Body: new(bytes.Buffer)},
		expectedBodyBytes:    []byte("service not available try later"),
		description:          "testing for expired endpoint request",
	},

	{
		requestTimeout:       time.Second,
		requestPath:          "",
		requestPathSignature: "/v1/{blockId:[0-9]+}",
		expectedW:            &httptest.ResponseRecorder{Code: 301, HeaderMap: http.Header{"Location": []string{"/"}}, Body: new(bytes.Buffer)},
		expectedBodyBytes:    []byte{},
		description:          "empty request",
	},
	{
		requestTimeout:       time.Second,
		requestPath:          "some random text",
		requestPathSignature: "/v1/{blockId:[0-9]+}",
		expectedW:            &httptest.ResponseRecorder{Code: 301, HeaderMap: http.Header{"Location": []string{"/some%20random%20text"}}, Body: new(bytes.Buffer)},
		expectedBodyBytes:    []byte{},
		description:          "wrong path request",
	},
	{
		requestTimeout:       time.Second,
		requestPath:          "/v1/",
		requestPathSignature: "/v1/{blockId:[0-9]+}",
		expectedW: &httptest.ResponseRecorder{Code: 404,
			HeaderMap: http.Header{"Content-Type": []string{"text/plain; charset=utf-8"},
				"X-Content-Type-Options": []string{"nosniff"}},
			Body: new(bytes.Buffer),
		},
		expectedBodyBytes: []byte("404 page not found\n"),
		description:       "empty blockId value",
	},
	{
		requestTimeout:       time.Second,
		requestPath:          `/v1/goodMorning`,
		requestPathSignature: "/v1/{blockId:[0-9]+}",
		expectedW: &httptest.ResponseRecorder{Code: 404,
			HeaderMap: http.Header{"Content-Type": []string{"text/plain; charset=utf-8"},
				"X-Content-Type-Options": []string{"nosniff"}},
			Body: new(bytes.Buffer),
		},
		expectedBodyBytes: []byte("404 page not found\n"),
		description:       "wrong blockId parameter",
	},
	{
		requestTimeout:       time.Second,
		requestPath:          `/v1/18446744073709551615`,
		requestPathSignature: "/v1/{blockId:[0-9]+}",
		expectedW:            &httptest.ResponseRecorder{Code: 400, HeaderMap: make(http.Header), Body: new(bytes.Buffer)},
		expectedBodyBytes:    []byte("requested id 18446744073709551615 latest"),
		description:          "request not existent block",
	},
	{
		requestTimeout:       time.Second,
		requestPath:          `/v1/12`,
		requestPathSignature: "/v1/{blockId:[0-9]+}",
		expectedW:            &httptest.ResponseRecorder{Code: 200, HeaderMap: http.Header{"Content-Type": []string{"text/plain; charset=utf-8"}}, Body: new(bytes.Buffer)},
		expectedBodyBytes:    []byte(`{"jsonrpc":"2.0","id":1,"result":{"difficulty":"0x3ffffa000","extraData":"0x476574682f76312e302e302f6c696e75782f676f312e342e32","gasLimit":"0x1388","gasUsed":"0x0","hash":"0xc63f666315fa1eae17e354fab532aeeecf549be93e358737d0648f50d57083a0","logsBloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","miner":"0x0193d941b50d91be6567c7ee1c0fe7af498b4137","mixHash":"0xbe4ba21fe1ecb061e44f178428c772d2a0f59a7aafb5ed4e198eba4df3656e52","nonce":"0x5f6a5cc5c36e6627","number":"0xc","parentHash":"0x3f5e756c3efcb93099361b7ddd0dabfeaa592439437c1c836e443ccb81e93242","receiptsRoot":"0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421","sha3Uncles":"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347","size":"0x219","stateRoot":"0x821c41f30a2fd9580605363784a8a2a6575b255ec37cacf87fe52715b8828d8e","timestamp":"0x55ba42c0","totalDifficulty":"0x33f2ffe033","transactions":[],"transactionsRoot":"0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421","uncles":[]}}`),
		description:          "legit request  block 12",
	},
	{
		requestTimeout:       time.Second,
		requestPath:          `/v1/8373417`,
		requestPathSignature: "/v1/{blockId:[0-9]+}",
		expectedW:            &httptest.ResponseRecorder{Code: 200, HeaderMap: http.Header{"Content-Type": []string{"text/plain; charset=utf-8"}}, Body: new(bytes.Buffer)},
		expectedBodyBytes:    []byte(`{"jsonrpc":"2.0","id":1,"result":{"difficulty":"0x7f3bfd5cbef42","extraData":"0x505059452d65746865726d696e652d7573322d32","gasLimit":"0x7a1200","gasUsed":"0x79d8ca","hash":"0x7277db362335e9ddefe221ee9852a5be3094fce6f37d7042d99779f707481ea2","logsBloom":"0x806042e8805890282040801c22f300a4024a1c3888106b108d002c1025781b92008071330334880048e1546a406e8d6003086241080e2b084a1a130862bc8f0519642c3300230a6068189f6c80c2c420230ea4000b5c69082b08f5a5ab0331020a3c852b02e5c9d9000870af0cc409086d0c997d5180c452a0811117570408100ae3720e812180408624016085504637400848c7481d547f098404084b3061a30b21980422c840917606c48a041e1444d0585c84c8a10220dc842292208410d190019943289c200422aa508c2c00050708a1320ba1bc241604b4c002108265e49550e9a04341601a290206ac082380be8318958690809462c055081ca2808283","miner":"0xea674fdde714fd979de3edf0f56aa9716b898ec8","mixHash":"0x43adb35dce15211ef2d9c57d879633e4aebaf0eafcb6de23a937a6a0b9919b20","nonce":"0x9ddf2bc5b716c6cb","number":"0x7fc4a9","parentHash":"0x762e97f090af51183d5f327f42ac24ce2810936ab53a0d5fe71ccda747fc993f","receiptsRoot":"0x79e6b2fdc28a112da9cdcfedc5afe3d2312bd724224f0e4cf503b33654479ce7","sha3Uncles":"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347","size":"0x52ae","stateRoot":"0x0c5f47e4861491abd9464916ad57c556a8249708e1959a55b7f0fcceb5108caa","timestamp":"0x5d5912c4","totalDifficulty":"0x270553eb7009b2374c2","transactions":["0x1b9066136acfc82169fe1061959ab1b636ced8a8a91271aa074e2d79df1b9f7b","0x726e55dcc0f6a16ab96c8b418673304f88d268f96f6e0acfce5d28863725e8f6","0x16644ac85072c6741fb2a5fcb921a8208ea15967e04f6e412d121ef5584af9d6","0xc36c446ab19d005f1f0efd649cfe0a383055f3c8a2ffc73f8563f656120ba558","0xd3a4a604812b0433ee04c875e650f45262143a11b0fe08bfdb21516fcf26cc78","0x19697decca6aa009c7bb5ef51c4b18aae2ee916ed95ff2190157658b0df8c8e7","0xb9f7f3f9de025344d53db6f1b49b27123c28e43922080264f869e5eac9de6c52","0xade561db7b30086f5a969e5f800c0d63fcbef7d058f92ab28199cb9f7733ed52","0x9abe38460de7370da4821cc0b730eb68b5d4df5905de6520847dbb89aff518e0","0x8aa8d1737592ed976a1f944e03c64d93cf36e23da8d2b37ae73f60e1d4a1c65a","0xdf72d8ebef2f53515d9a83c3da8a04ae594c30aaa491d9ae3582932734ddfa69","0xb87a3b72148ec7872e5cc80d8b62c11c0f038908fd3490dcdc442e6ca6b3b336","0x35fb28d19d6209bfaabb4d83ea14c26768ee6139a4ebcadb0539b3998a854313","0xb26d3bf5ff06d696dd7286d0768e799c053c43e3bc83714c7b393d5e030d6ee0","0xa24d588434be280d55c7e53685d9648ee344831767b20a269303efc35f384cb8","0x529ccf5f9fc89b35c5a438ec50730ca547f14ae6684304e942aa139e0c6c2989","0x7cda7a1c4329d2549bc6d300e99d211cf5ee0d318e3a45e41981f1abefbf28c8","0x620b942a4b405d885a42e17e4c1e265ff3ce903d988be76b077b4a9f5f77a3ff","0x487281649aa6d842c99174a60f6021fc31a22335a81b6a1ebbf2f6d290e8409c","0xa4c8212f1ea9886e2306582e7b540a63fe3875d7e2eec31fddb9dc3884bee781","0x0e600776829ab26b96b0bb60b38eeccd4f50e6349cb9da67934569946c9559b4","0x11d49169bec75e0a18a6d1692f2bbaf47945f2e50bef62b179292e075b5e3b9e","0xbbe652e0b7bca2443bb2fa31011fcf78aad9454f171be5aae2478bccfc8b2e27","0x7373b23d78189ae7e62b7bf994e034a1da1caf4e87b5d866313b8899a3beafe1","0x5be79be9d42ceb7b9790f22bcc7ab69ff0428f231fca057f41aebf7532e6e701","0xc36d65f38b98ac354040522de1043a31cbce61533ed21e204fc74f8005cf0cd9","0x26e359ad195f54920d00a023ca0039019b7ed74bf0951296963222fc7762ac51","0x5db8ea048f55a9889ca199f1f0699febb4f973775feb308ccd3c8d09983664c3","0x7f10e3658fbbb82d6cb33076bbad937a824fe7445ffedf077cc44b34228562db","0xe05b4acfa5734419ac05e18b2ee0468412166ef0a809650030c9885f94a8ee27","0x25bf0d519dd7afcca968a08882b71ec6c38bef1659ec429853c359ed69c69b3c","0x98731095e141e608268c129b3e14996c861c7dd4f1dc45b1e9dbbf47f540106e","0x0a099e0aa633d7b1a74e7b7d4b5778f7ce40856b33ae783900fa34b224144367","0x080ec2005a2321e17d955d34efd0816e268823df08c633ccbef7186141d4b6e4","0xbd979c2b4c52575022625d9db5445199f86350d457fd9dfd9d24aeed9677296c","0x5ad5c1b7526bf369cd4ad2eac5a5dd5097f7b75695bbe333a5894aa9e7c32a98","0xb7a93bd0fbd93d8bef59265a4deb8fc2bbe5b5535ee52f20e3922303e8333409","0x32aab1a9047fae1b35a570aedad411491c5c663542e63cd56286cafeeba14901","0x175aeda1c1989611e633fc3d7ef3a5b641f2840fa9cc06f25f5b29ac0bc6e1b2","0xeb996274a813364b96a12ccedd122348bad80a91d468559c6e93b88f4b12a400","0x8c329344fab5099b960cad17f1a4b8c44419c65248f3ec2422bc7923b9f5e8c6","0x293080fd9e4e2103e065a14ae180e9df48ce7da148d0a71d90df4fc894466a50","0xe1a0c33b8d6d937fc7246797dfc406498eb62f74040a5dbeda23d42195832eec","0x6d44578f73eecf5a6bbff08fa664676536fd3a91bfb5bc173dca70135affd3f0","0x96ceb3da219447c264c0c6cccf53dc19ee92492cac32e91b388fcf38b5816ac5","0x7e99a8d2bb95f07418e9627b0570068969db5fe2b05338d95eda602829556466","0x8df5902d9be7c485f31ea362c931ce20631bd7fcebc681c8e17baafd42b2322d","0xddd7dbc681230d3f830ef16e0f732a8bc81aec9ab8e4d1e14d08c35967fad3c2","0xe0f795cc8bd16894aead90e083b4232b16e1bb8555e91077c2a49b87c97d037f","0x8b14b63b8bb49eb29682459836cf8f133cbd882e1a18fc8e8c89668618ed009a","0xed504e12a1f568ec58b2aae5340652bd91c6e10576729c35e93e7e5487cfa798","0xfe61e05372976d098aa94b4891a675911e5fe0af632cc56100d67709987f010f","0x796ca0786d398eaec57473a273719f608c9ed8ca2b944fb66cec4c2211986789","0xc794b528b0776a3c6ab095a85d1a89ec6d28402468eef2c8140da926bfe3a8f5","0x4a16a9a3206e4d38965d788497f7f3a5b57f2c9c14dc02eba7ef2c6e0eb325d0","0xcb5a33ccd21600b1a3b4632146031b4af548adc2decaa1deaf7baf4ab4e4df58","0x25b65abbccf0b1cd6aed26df90e64a29642c5396d9a601a633317a230067044f","0xe10c400378aad05e3f2a7f576bc6d09ede312118a28914c5a2b221ed748f0515","0x0cf00a7dfde96c5165f68c920ebce9662925ee02a817a3a9c8a2b3a283e2a829","0x59cbebc0553dd8f707038bdb161d73383a1e4f3294cc0c46c36cc00bdf212981","0x4f8aebd96f1e4fc44e46c50d1408025759e314fa5ee85e7231a271d8f16102bc","0x58b46c5c3b918cfff4839057d8a48bc138351490778514d131a145e6f6cb94b5","0x905a3f5dc83edefe431de499e2181b4abf5e338d5b8870720704ec8ba460238a","0xcdeecf3fb6990e7ed854f08cf74959b9dfa1b26ccc271582d6c27247c5e81034","0x457efe64c4568839bedba065efb8c19a937dd4a272923184112fb09931ae0feb","0xc2c8d0d8a841023e13a789577fdde49dbbc67f728d453c08ef32792a73d9e0fa","0xc4bdb7cf4eeb3418f8a42fae6fb69966e2ac01b5c6817677d5f49445c3a2fe82","0xeb6b755c5fef25cf11bebc5fa1d4713ee65acd1912e2f501bd2f1b824646070b","0x6ec02896a7eabca4b761889243e3413a0abdc07147fbabc4bc315b918f368a88","0xe4a309b597f342c765791d90d374407decb51043ec210a970cf2b23a31b89f3a","0xff8ff8ac60f6046fdeb7e307c0b412139b42a111e1c4dc6172415e2250af53b6","0x09d3e898352f16c9d325a5db017b5df5f76e722db80df60f3806aba8d5383a44","0x996ebbc30e62d9b06a2e3716f7dd2f9ced17d95492157cae9bd3d86a57f023f6","0x33f5c353a8bb604bcd01a0ee6083c693189658ed27c36dc754cb87bcf42be31d","0x337571a8ff66e4e5d30bb41c562bbd107baef9ed01d898f42f7225daf9e043e4","0xbef52dff4691a6fd9747941d84ed57833969dfac049d89de077dd3b1a43d38b3","0x3ca93a5f1a9af1beb976fa186b2725070c2ef5cf2b7ec4b1c436f3d6d4026581","0x17890a58227261000cdd0aadb3f2fbf3503722f58f516ab5c7d6a7bcb9e5c02d","0xf7d5ac2c5b80db8ccb46c8d9bb691dee3c983ab9df8e8b4e30b9e1469140ca64","0x14e96d70e5dbd2959f5de9c36db6a1f222ee5eb116141e9c017b7537c96518a9","0xceab422fa936cff16cbed69687bb6b707f94d0a56d3827f5cb7548c9494171cb","0x682c5d49401d26fec0c8d51643cc48df55fa5bd30301435d57b31a36acdcc588","0x80811a539cfa14738589bfaa90acb9900f6055123fb8ad940ad0e17c9fec5338","0x98a021bd24673ef5f82068c4380aa24b6559fa5f231b71cd7be09c1eb66da0a1","0xd7afcb5060de9b8cd1c9437b4eac0bb99832b5350ff8fcc172362716012cf04b","0xa411afa688c85be9d2dff42ec62ce7fe1bc1726a872a973bcb39691cdb607a58","0x02fd55faa33a6824aedc45d843cb4a23ca12495be736023faf5df4ed4d8dbf36","0xb5216a1491e1c9111936d2287d1c0b5ddb0b746ce7d87efc856fac35c738fe85","0x22a4b821f2f58e905f7ed65262c23179cc91a46977a4dadb80402b1e41021c7f","0xff3d0c969d6094eef194bbce5883ada6194190d35557baf541a3c46fcbdd3ffb","0x0ba0c70d55189d4038f805435230def29c0f4f25319d2acdacd611cc17da7da4","0xd2048b059bb00bc3869a3702dcd80f7214221be28c8266b3c771d632819d831c","0x658a06a6366a6877d92f2f71ee454fd6e2c1302074337f95ecd6c38feeee593a","0xafb3c23b3ad1f7f4ede2525e25e9d01fe538bf2fc8c6698303ba6266c47b33f9","0xce76abd3ce771c29e0dd52873d5f6c50a75054e0742d398b4d5d5d97ccd57f15","0xebf1e087a4f1b592f1f78ce9de90ae5e7ddffeb1dce7ea40e692ca6b4bccc9ed","0xabf9eaad5afe444615ac8b4ab6e1f6475c04fdf7fba1824fad2c44d3bfdfaf32","0xc7c368888d24be1b2401e1ee42da3fb7a8388d2ca71ee2165a552f03c03b1a7d","0x9fb78749f22de8071a2a296b76e6f62a3795039b09f47156121724c6cf24dced","0x29baf474f21610d8da37c6b540d90cf86015cea48d57cd55939ccd45369c43bc","0x6e47ca5f5e9c3eef26732373296159f794926752417347092da951e1c8de6da7","0x9219383b4a9bd20feffbb822c68151a008bd02158be55b8561dc9387fffb23cf","0x2dcf4adb7ca3c8b0c624b5afb6a6f7f737ce56e76327bdc8f7f8656ceed47c5c","0xa7c181cf2ecb59b352efb6c664ade6265e3e830b22ac1eda07d392e8ceb09001","0x102375dda045195975ef25804e45be7a5c348ffd8a4296b7e5b5261179b5dd80","0x8050e22bf45a82b56380a38c58752f6ff859fb74cbddf18c77c5ef7b241e11c8","0x50160249ec818dfda4a7ffda2b3e23dfd728d9669d7a055bf93347974751fb10","0xc3eee712a9cc261c19fb94ac52c61ab0695e447e72fb4679d3136f6e4d211b23","0x2a5bf4e14a5849d85a4060240dd41d07f6e90976c980386199498a1932f4405b","0x4af492242204256e6c658361b17da1da3c4741e8dde077098285352e57b6d233","0x132ca88eadef346da000bb5095b0b8cf877dc0175fd27ff58d5099df1967c115","0xef9b51b6fbbf4896eb841d5a591a9fa5b72e5cff33068ed103d9db3a96c47560","0x829a995f62045fd6c78b19a2e6b1b0eed0afb747d40b2a78edcf0ddc697d2299","0xb3f6389cbfa1091e3243a01d4cbeee24b9f9913580f87040e06db92fe7b40ce3","0x3acaae3c1675076888a84c1dde3c0b54c6a5f6f1727abef0cb9c76dc70e46ded","0xfea6af7c8f989373ed671c2058f885bde420a270cda0628fae416ddb117aec30","0xe7204fcd70551fd6594fa87345765f010c255deac56c6f4ef638d73689dc190c","0x0bd2de7121cb1c6e6c64a7c254051e0e7d867ae9a521c3e76115a2de47d71071","0x1066a26f71a8f9d94af30042423135258d26bb294ca5864828e7c9c44c998e98","0x6ac05a7d8642111760b2502917768344cbfe269b28eb0a88705518e8ec84a142","0x37988163abccb32360e68654903e9e44beb5a3dc3bc16bcd98353ba5a0220fa5","0x8dfe3f8248b133a7100e581c9894739764ec320514f801ba3b0ed795ab13d640","0x2b8af18327733914186ac091826a0b6d32fce5a0df2bfaaba0534b629a98a4ca","0xca421ad6603cd0efb7fbacedc6f58a48ad799369d56ee7806906e09ea0a68fac","0x9ec7ff1d8dd758bf16493bf48627c82e7c1d01d8bb3f64b69750cca55fbcf39c","0xb78f35c40f32eed89d47896ce0a92765d3ef7d394d04382df342013ddd556694"],"transactionsRoot":"0x71f135d16639a0b4f787705e61aece2bda79269ec3e7757f3c6b5bc1d7501989","uncles":[]}}`),
		description:          "legit request  of the recent block 8373417",
	},
}

var testCasesGetTransaction = []handlerTest{
	{
		requestTimeout:       time.Nanosecond,
		requestPath:          "/v1/12/1",
		requestPathSignature: "/v1/{blockId:[0-9]+}/{txId:[0-9]+}",
		expectedW:            &httptest.ResponseRecorder{Code: 503, HeaderMap: http.Header{}, Body: new(bytes.Buffer)},
		expectedBodyBytes:    []byte("service not available try later"),
		description:          "testing for expired endpoint request",
	},
	{
		requestTimeout:    time.Second,
		requestPath:       "",
		expectedW:         &httptest.ResponseRecorder{Code: 301, HeaderMap: http.Header{"Location": []string{"/"}}, Body: new(bytes.Buffer)},
		expectedBodyBytes: []byte{},
		description:       "empty request",
	},
	{
		requestTimeout:       time.Second,
		requestPath:          "/v1//1",
		requestPathSignature: "/v1/{blockId:[0-9]+}/{txId:[0-9]+}",
		expectedW:            &httptest.ResponseRecorder{Code: 301, HeaderMap: http.Header{"Location": []string{"/v1/1"}}, Body: new(bytes.Buffer)},
		expectedBodyBytes:    []byte{},
		description:          "empty block index",
	},
	{
		requestTimeout:       time.Second,
		requestPath:          "/v1/1//",
		requestPathSignature: "/v1/{blockId:[0-9]+}/{txId:[0-9]+}",
		expectedW:            &httptest.ResponseRecorder{Code: 301, HeaderMap: http.Header{"Location": []string{"/v1/1/"}}, Body: new(bytes.Buffer)},
		expectedBodyBytes:    []byte{},
		description:          "empty tx index",
	},
	{
		requestTimeout:       time.Second,
		requestPath:          "/v1/837522700/1",
		requestPathSignature: "/v1/{blockId:[0-9]+}/{txId:[0-9]+}",
		expectedW:            &httptest.ResponseRecorder{Code: 400, HeaderMap: http.Header{}, Body: new(bytes.Buffer)},
		expectedBodyBytes:    []byte{},
		description:          "too big block index",
	},
	{
		requestTimeout:       time.Second,
		requestPath:          "/v1/12/0",
		requestPathSignature: "/v1/{blockId:[0-9]+}/{txId:[0-9]+}",
		expectedW:            &httptest.ResponseRecorder{Code: 200, HeaderMap: http.Header{"Content-Type":{"text/plain; charset=utf-8"}}, Body: new(bytes.Buffer)},
		expectedBodyBytes:    []byte(`{"jsonrpc":"2.0","id":12,"result":null}`),
		description:          "legit request block 12 tx 0",
	},
}
