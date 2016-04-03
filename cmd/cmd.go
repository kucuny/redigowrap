package cmd

type RedisCommandCluster interface {

}

type RedisCommandConnection interface {
	Auth(password string) (bool, error)
	Echo(message string) (string, error)
	Ping() (string, error)
	Select(index int) (bool, error)
	Quit() (string, error)
}

type RedisCommandGeo interface {
	GeoAdd() (error)
	GeoHash() (error)
	GeoPos() (error)
	GeoDist() (error)
	GeoRadius() (error)
	GeoRadiusByMember() (error)
}

type RedisCommandHashes interface {
	HDel(hashKey string, fields []string) (int, error)
	HExists(hashKey, field string) (bool, error)
	HGet(hashKey, field string) (string, error)
	HGetFloat64(hashKey, field string) (float64, error)
	HGetAll(hashKey string) (map[string]string, error)
	HGetAllInterface(hashKey string) (map[string]interface{}, error)
	HGetAllFloat64(hashKey string) (map[string]float64, error)
	HIncrBy(hashKey, field string, increment int) (int, error)
	HIncrByFloat(hashKey, field string, increment float64) (float64, error)
	HKeys(hashKey string) ([]string, error)
	HLen(hashKey string) (int, error)
	HMGet(hashKey string, fields []string) ([]string, error)
	HMGetFloat64(hashKey string, fields []string) ([]float64, error)
	HMSet(hashKey string, fieldValue map[string]string) (bool, error)
	HMSetInterface(hashKey string, fieldValue map[string]interface{}) (bool, error)
	HMSetFloat64(hashKey string, fieldValue map[string]float64) (bool, error)
	// HScan() ()
	HSet(hashKey, field, value string) (int, error)
	HSetFloat64(hashKey, field string, value float64) (int, error)
	HSetNX(hashKey, field, value string) (int, error)
	HStrLen(hashKey, field string) (int, error)
	HVals(hashKey string) ([]string, error)
}

type RedisCommandHyperLogLog interface {

}

type RedisCommandKeys interface {

}

type RedisCommandLists interface {

}

type RedisCommandPubSub interface {

}

type RedisCommandScripting interface {

}

type RedisCommandServer interface {

}

type RedisCommandSets interface {

}

type RedisCommandSortedSets interface {

}

type RedisCommandStrings interface {

}

type RedisCommandTransactions interface {

}
