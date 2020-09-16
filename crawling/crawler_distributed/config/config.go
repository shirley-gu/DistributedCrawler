package config

const(
	ParseCity = "ParseCity"
	ParseCityList = "ParseCityList"
	ParseProfile = "parseProfile"
	NilParser = "NilParser"

	//Service ports
	// ItemSaverPort = 1234
	// WorkerPort0 = 9000

	//ElasticSearch
	ElasticIndex = "dating_profile"

	//RPC Endpoints
	ItemSaverRpc = "ItmeSaverService.Save"
	CrawlerServiceRpc = "CrawlerService.Process"

	//Rate limiting
	Qps = 20
)