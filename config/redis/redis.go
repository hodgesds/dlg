package redis

import (
	"time"

	v7 "github.com/go-redis/redis/v7"
)

// Config is a config for a Redis stage.
type Config struct {
	Network            string         `yaml:"network,omitempty"`
	Addr               string         `yaml:"addr"`
	DB                 int            `yaml:"db"`
	PoolSize           *int           `yaml:"poolSize,omitempty"`
	MaxRetries         *int           `yaml:"maxRetries,omitempty"`
	MinRetryBackoff    *time.Duration `yaml:"minRetryBackoff,omitempty"`
	MaxRetryBackoff    *time.Duration `yaml:"maxRetryBackoff,omitempty"`
	DialTimeout        *time.Duration `yaml:"dialTimeout,omitempty"`
	ReadTimeout        *time.Duration `yaml:"readTimeout,omitempty"`
	WriteTimeout       *time.Duration `yaml:"writeTimeout,omitempty"`
	MaxConnAge         *time.Duration `yaml:"maxConnAge,omitempty"`
	PoolTimeout        *time.Duration `yaml:"poolTimeout,omitempty"`
	IdleTimeout        *time.Duration `yaml:"idleTimeout,omitempty"`
	IdleCheckFrequency *time.Duration `yaml:"idleCheckFrequency,omitempty"`
	Commands           []Command      `yaml:"commands"`
}

// Client returns a redis client.
func (c *Config) Client() *v7.Client {
	opts := &v7.Options{
		Network: c.Network,
		Addr:    c.Addr,
		DB:      c.DB,
	}
	if c.PoolSize != nil {
		opts.PoolSize = *c.PoolSize
	}
	if c.MaxRetries != nil {
		opts.MaxRetries = *c.MaxRetries
	}
	if c.MinRetryBackoff != nil {
		opts.MinRetryBackoff = *c.MinRetryBackoff
	}
	if c.MaxRetryBackoff != nil {
		opts.MaxRetryBackoff = *c.MaxRetryBackoff
	}
	if c.DialTimeout != nil {
		opts.DialTimeout = *c.DialTimeout
	}
	if c.ReadTimeout != nil {
		opts.ReadTimeout = *c.ReadTimeout
	}
	if c.WriteTimeout != nil {
		opts.WriteTimeout = *c.WriteTimeout
	}
	if c.MaxConnAge != nil {
		opts.MaxConnAge = *c.MaxConnAge
	}
	if c.PoolTimeout != nil {
		opts.PoolTimeout = *c.PoolTimeout
	}
	if c.IdleTimeout != nil {
		opts.IdleTimeout = *c.IdleTimeout
	}
	if c.IdleCheckFrequency != nil {
		opts.IdleCheckFrequency = *c.IdleCheckFrequency
	}

	return v7.NewClient(opts)
}

// Command is a Redis command.
type Command struct {
	Append *struct {
		Key   string `yaml:"key"`
		Value string `yaml:"value"`
	} `yaml:"append,omitempty"`
	BLPop *struct {
		Timeout time.Duration `yaml:"timeout"`
		Keys    []string      `yaml:"keys"`
	} `yaml:"blPop,omitempty"`
	BRPop *struct {
		Timeout time.Duration `yaml:"timeout"`
		Keys    []string      `yaml:"keys"`
	} `yaml:"brPop,omitempty"`
	BRPopLPush *struct {
		Source  string        `yaml:"source"`
		Dest    string        `yaml:"dest"`
		Timeout time.Duration `yaml:"timeout"`
	} `yaml:"brPopLPush,omitempty"`
	BZPopMax *struct {
		Timeout time.Duration `yaml:"timeout"`
		Keys    []string      `yaml:"keys"`
	} `yaml:"bzPopMax,omitempty"`
	BZPopMin *struct {
		Timeout time.Duration `yaml:"timeout"`
		Keys    []string      `yaml:"keys"`
	} `yaml:"bzPopMin,omitempty"`
	BgRewriteAOF *bool `yaml:"bgRewriteAof"`
	BgSave       *bool `yaml:"bgSave"`
	BitCount     *struct {
		Key      string       `yaml:"key"`
		BitCount *v7.BitCount `yaml:"bitCount"`
	} `yaml:"bitCount,omitempty"`
	BitField *struct {
		Key string `yaml:"key"`
		//Args []interface{} `yaml:"args"`
	} `yaml:"bitField,omitempty"`
	BitOpAnd *struct {
		DestKey string   `yaml:"destKey"`
		Keys    []string `yaml:"keys"`
	} `yaml:"bitOpAnd,omitempty"`
	BitOpNot *struct {
		DestKey string `yaml:"destKey"`
		Key     string `yaml:"key"`
	} `yaml:"bitOpNot,omitempty"`
	BitOpOr *struct {
		DestKey string   `yaml:"destKey"`
		Keys    []string `yaml:"keys"`
	} `yaml:"bitOpOr,omitempty"`
	BitOpXor *struct {
		DestKey string   `yaml:"destKey"`
		Keys    []string `yaml:"keys"`
	} `yaml:"bitOpXor,omitempty"`
	BitPos *struct {
		Key string  `yaml:"key"`
		Bit int64   `yaml:"bit"`
		Pos []int64 `yaml:"pos"`
	} `yaml:"bitPos,omitempty,omitempty"`
	ClientGetName *bool `yaml:"clientGetName,omitempty"`
	ClientID      *bool `yaml:"clientId,omitempty"`
	ClientKill    *struct {
		IPPort string `yaml:"ipPort"`
	} `yaml:"clientKill,omitempty"`
	ClientKillByFilter *struct {
		Keys []string `yaml:"keys"`
	} `yaml:"clientKillByFilter,omitempty"`
	ClientList  *bool `yaml:"clientList,omitempty"`
	ClientPause *struct {
		Dur time.Duration `yaml:"dur"`
	} `yaml:"clientPause,omitempty"`
	ClientUnblock *struct {
		ID int64 `yaml:"id"`
	} `yaml:"clientUnblock,omitempty"`
	ClientUnblockWithError *struct {
		ID int64 `yaml:"id"`
	} `yaml:"clientUnblockWithError"`
	Close           *bool `yaml:"close,omitempty"`
	ClusterAddSlots *struct {
		Slots []int `yaml:"slots"`
	} `yaml:"clusterAddSlots"`
	ClusterAddSlotsRange *struct {
		Min int `yaml:"min"`
		Max int `yaml:"max"`
	} `yaml:"clusterAddSlotsRange"`
	ClusterCountFailureReports *struct {
		NodeID string `yaml:"nodeID"`
	} `yaml:"clusterCountFailureReports"`
	ClusterCountKeysInSlot *struct {
		Slot int `yaml:"slot"`
	} `yaml:"clusterCountKeysInSlot"`
	ClusterDelSlots *struct {
		Slots []int `yaml:"slots"`
	} `yaml:"clusterDelSlots"`
	ClusterDelSlotsRange *struct {
		Min int `yaml:"min"`
		Max int `yaml:"max"`
	} `yaml:"clusterDelSlotsRange"`
	ClusterFailover *bool `yaml:"clusterFailover"`
	ClusterForget   *struct {
		NodeID string `yaml:"nodeID"`
	} `yaml:"clusterForget"`
	ClusterGetKeysInSlot *struct {
		Slot  int `yaml:"slot"`
		Count int `yaml:"count"`
	} `yaml:"clusterGetKeysInSlot"`
	ClusterInfo    *bool `yaml:"clusterInfo"`
	ClusterKeySlot *struct {
		Key string `yaml:"key"`
	} `yaml:"clusterKeySlot"`
	ClusterMeet *struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"clusterMeet"`
	ClusterNodes     *bool `yaml:"clusterNodes"`
	ClusterReplicate *struct {
		NodeID string `yaml:"nodeID"`
	} `yaml:"clusterReplicate"`
	ClusterResetHard  *bool `yaml:"clusterResetHard"`
	ClusterResetSoft  *bool `yaml:"clusterResetSoft"`
	ClusterSaveConfig *bool `yaml:"clusterSaveConfig"`
	ClusterSlaves     *struct {
		NodeID string `yaml:"nodeID"`
	} `yaml:"clusterSlaves"`
	ClusterSlots *bool `yaml:"clusterSlots"`
	Command      *bool `yaml:"command"`
	ConfigGet    *struct {
		Parameter string `yaml:"parameter"`
	} `yaml:"configGet"`
	ConfigResetStat *bool `yaml:"configResetStat"`
	ConfigRewrite   *bool `yaml:"configRewrite"`
	ConfigSet       *struct {
		Parameter string `yaml:"parameter"`
		Value     string `yaml:"value"`
	} `yaml:"configSet"`
	DBSize      *bool `yaml:"dbSize"`
	DbSize      *bool `yaml:"dbSize"`
	DebugObject *struct {
		Key string `yaml:"key"`
	} `yaml:"debugObject"`
	Decr *struct {
		Key string `yaml:"key"`
	} `yaml:"decr"`
	DecrBy *struct {
		Key       string `yaml:"key"`
		Decrement int64  `yaml:"decrement"`
	} `yaml:"decrBy"`
	Del *struct {
		keys []string `yaml:"keys"`
	} `yaml:"del"`

	Dump *struct {
		Key string `yaml:"key"`
	} `yaml:"dump"`
	Echo *struct {
		//message interface{} `yaml:"message"`
	} `yaml:"echo"`
	Eval *struct {
		Script string   `yaml:"script"`
		Keys   []string `yaml:"keys"`
		//args   []interface{} `yaml:"args"`
	} `yaml:"eval"`
	EvalSha *struct {
		Sha1 string   `yaml:"sha1"`
		Keys []string `yaml:"keys"`
		//args []interface{} `yaml:"args"`
	} `yaml:"evalSha"`
	Exists *struct {
		Keys []string `yaml:"keys"`
	} `yaml:"exists"`
	Expire *struct {
		Key        string        `yaml:"key"`
		Expiration time.Duration `yaml:"expiration"`
	} `yaml:"expire"`
	ExpireAt *struct {
		Key string    `yaml:"key"`
		Tm  time.Time `yaml:"tm"`
	} `yaml:"expireAt"`
	FlushAll      *bool `yaml:"flushAll"`
	FlushAllAsync *bool `yaml:"flushAllAsync"`
	FlushDB       *bool `yaml:"flushDB"`
	FlushDBAsync  *bool `yaml:"flushDBAsync"`
	GeoAdd        *struct {
		Key         string            `yaml:"key"`
		GeoLocation []*v7.GeoLocation `yaml:"geoLocation"`
	} `yaml:"geoAdd"`
	GeoDist *struct {
		Key     string `yaml:"key"`
		Member1 string `yaml:"member1"`
		Member2 string `yaml:"member2"`
		Unit    string `yaml:"unit"`
	} `yaml:"geoDist"`
	GeoHash *struct {
		Key     string   `yaml:"key"`
		Members []string `yaml:"members"`
	} `yaml:"geoHash"`
	GeoPos *struct {
		Key     string   `yaml:"key"`
		Members []string `yaml:"members"`
	} `yaml:"geoPos"`
	GeoRadius *struct {
		Key       string             `yaml:"key"`
		Longitude float64            `yaml:"longitude"`
		Latitude  float64            `yaml:"latitude"`
		Query     *v7.GeoRadiusQuery `yaml:"query"`
	} `yaml:"geoRadius"`
	GeoRadiusByMember *struct {
		Key    string             `yaml:"key"`
		Member string             `yaml:"member"`
		Query  *v7.GeoRadiusQuery `yaml:"query"`
	} `yaml:"geoRadiusByMember"`
	GeoRadiusByMemberStore *struct {
		Key    string             `yaml:"key"`
		Member string             `yaml:"member"`
		Query  *v7.GeoRadiusQuery `yaml:"query"`
	} `yaml:"geoRadiusByMemberStore"`
	GeoRadiusStore *struct {
		Key       string             `yaml:"key"`
		Longitude float64            `yaml:"longitude"`
		Latitude  float64            `yaml:"latitude"`
		Query     *v7.GeoRadiusQuery `yaml:"query"`
	} `yaml:"geoRadiusStore"`
	Get *struct {
		Key string `yaml:"key"`
	} `yaml:"get"`
	GetBit *struct {
		Key    string `yaml:"key"`
		Offset int64  `yaml:"offset"`
	} `yaml:"getBit"`
	GetRange *struct {
		Key   string `yaml:"key"`
		Start int64  `yaml:"start"`
		End   int64  `yaml:"end"`
	} `yaml:"getRange"`
	GetSet *struct {
		Key string `yaml:"key"`
		//value interface{} `yaml:"value"`
	} `yaml:"getSet"`
	HDel *struct {
		Key    string   `yaml:"key"`
		Fields []string `yaml:"fields"`
	} `yaml:"hDel"`
	HExists *struct {
		Key   string `yaml:"key"`
		Field string `yaml:"field"`
	} `yaml:"hExists"`
	HGet *struct {
		Key   string `yaml:"key"`
		Field string `yaml:"field"`
	} `yaml:"hGet"`
	HGetAll *struct {
		Key string `yaml:"key"`
	} `yaml:"hGetAll"`
	HIncrBy *struct {
		Key   string `yaml:"key"`
		Field string `yaml:"field"`
		Incr  int64  `yaml:"incr"`
	} `yaml:"hIncrBy"`
	HIncrByFloat *struct {
		Key   string  `yaml:"key"`
		Field string  `yaml:"field"`
		Incr  float64 `yaml:"incr"`
	} `yaml:"hIncrByFloat"`
	HKeys *struct {
		Key string `yaml:"key"`
	} `yaml:"hKeys"`
	HLen *struct {
		Key string `yaml:"key"`
	} `yaml:"hLen"`
	HMGet *struct {
		Key    string   `yaml:"key"`
		Fields []string `yaml:"fields"`
	} `yaml:"hmGet"`
	HMSet *struct {
		Key string `yaml:"key"`
		//values []interface{} `yaml:"values"`
	} `yaml:"hmSet"`
	HScan *struct {
		Key    string `yaml:"key"`
		Cursor uint64 `yaml:"cursor"`
		Match  string `yaml:"match"`
		Count  int64  `yaml:"count"`
	} `yaml:"hScan"`
	HSet *struct {
		Key string `yaml:"key"`
		//values []interface{} `yaml:"values"`
	} `yaml:"hSet"`
	HSetNX *struct {
		Key   string `yaml:"key"`
		Field string `yaml:"field"`
		//value interface{} `yaml:"value"`
	} `yaml:"hSetNX"`
	HVals *struct {
		Key string `yaml:"key"`
	} `yaml:"hVals"`
	Incr *struct {
		Key string `yaml:"key"`
	} `yaml:"incr"`
	IncrBy *struct {
		Key   string `yaml:"key"`
		Value int64  `yaml:"value"`
	} `yaml:"incrBy"`
	IncrByFloat *struct {
		Key   string  `yaml:"key"`
		Value float64 `yaml:"value"`
	} `yaml:"incrByFloat"`
	Info *struct {
		Section []string `yaml:"section"`
	} `yaml:"info"`
	Keys *struct {
		Pattern string `yaml:"pattern"`
	} `yaml:"keys"`
	LIndex *struct {
		Key   string `yaml:"key"`
		Index int64  `yaml:"index"`
	} `yaml:"lIndex"`
	//LInsert *struct {
	//	Key   string`yaml:"key"`
	//	Op    string `yaml:"op"`
	//	Pivot `yaml:"pivot"`
	//	Value interface{} `yaml:"value"`
	//} `yaml:"lInsert"`
	//LInsertAfter *struct {
	//	key   string `yaml:"key"`
	//	pivot `yaml:"pivot"`
	//	value interface{} `yaml:"value"`
	//} `yaml:"lInsertAfter"`
	//LInsertBefore *struct {
	//	key   string `yaml:"key"`
	//	pivot `yaml:"pivot"`
	//	value interface{} `yaml:"value"`
	//} `yaml:"lInsertBefore"`
	LLen *struct {
		Key string `yaml:"key"`
	} `yaml:"lLen"`
	LPop *struct {
		Key string `yaml:"key"`
	} `yaml:"lPop"`
	//LPush *struct {
	//	Key    string        `yaml:"key"`
	//	Values []interface{} `yaml:"values"`
	//} `yaml:"lPush"`
	//LPushX *struct {
	//	Key    string        `yaml:"key"`
	//	Values []interface{} `yaml:"values"`
	//} `yaml:"lPushX"`
	LRange *struct {
		Key   string `yaml:"key"`
		Start int64  `yaml:"start"`
		Stop  int64  `yaml:"stop"`
	} `yaml:"lRange"`
	//LRem *struct {
	//	Key   string      `yaml:"key"`
	//	Count int64       `yaml:"count"`
	//	Value interface{} `yaml:"value"`
	//} `yaml:"lRem"`
	//LSet *struct {
	//	Key   string      `yaml:"key"`
	//	Index int64       `yaml:"index"`
	//	Value interface{} `yaml:"value"`
	//} `yaml:"lSet"`
	LTrim *struct {
		Key   string `yaml:"key"`
		Start int64  `yaml:"start"`
		Stop  int64  `yaml:"stop"`
	} `yaml:"lTrim"`
	LastSave *bool `yaml:"lastSave"`
	MGet     *struct {
		Keys []string `yaml:"keys"`
	} `yaml:"mGet"`
	//MSet *struct {
	//	Values []interface{} `yaml:"values"`
	//} `yaml:"mSet"`
	//MSetNX *struct {
	//	Values []interface{} `yaml:"values"`
	//} `yaml:"mSetNX"`
	MemoryUsage *struct {
		Key     string `yaml:"key"`
		Samples []int  `yaml:"samples"`
	} `yaml:"memoryUsage"`
	Migrate *struct {
		Host    string        `yaml:"host"`
		Port    string        `yaml:"port"`
		Key     string        `yaml:"key"`
		Db      int           `yaml:"db"`
		Timeout time.Duration `yaml:"timeout"`
	} `yaml:"migrate"`
	Move *struct {
		Key string `yaml:"key"`
		Db  int    `yaml:"db"`
	} `yaml:"move"`
	ObjectEncoding *struct {
		Key string `yaml:"key"`
	} `yaml:"objectEncoding"`
	ObjectIdleTime *struct {
		Key string `yaml:"key"`
	} `yaml:"objectIdleTime"`
	ObjectRefCount *struct {
		Key string `yaml:"key"`
	} `yaml:"objectRefCount"`
	PExpire *struct {
		Key        string        `yaml:"key"`
		Expiration time.Duration `yaml:"expiration"`
	} `yaml:"pExpire"`
	PExpireAt *struct {
		Key string    `yaml:"key"`
		Tm  time.Time `yaml:"tm"`
	} `yaml:"pExpireAt"`
	//PFAdd *struct {
	//	key string        `yaml:"key"`
	//	els []interface{} `yaml:"els"`
	//} `yaml:"pfAdd"`
	PFCount *struct {
		Keys []string `yaml:"keys"`
	} `yaml:"pfCount"`
	PFMerge *struct {
		Dest string   `yaml:"dest"`
		Keys []string `yaml:"keys"`
	} `yaml:"pfMerge"`

	PTTL *struct {
		Key string `yaml:"key"`
	} `yaml:"pttl"`
	Persist *struct {
		Key string `yaml:"key"`
	} `yaml:"persist"`
	Ping           *bool `yaml:"ping"`
	PubSubChannels *struct {
		Pattern string `yaml:"pattern"`
	} `yaml:"pubSubChannels"`
	PubSubNumPat *bool `yaml:"pubSubNumPat"`
	PubSubNumSub *struct {
		Channels []string `yaml:"channels"`
	} `yaml:"pubSubNumSub"`
	Publish *struct {
		Channel string      `yaml:"channel"`
		Message interface{} `yaml:"message"`
	} `yaml:"publish"`
	Quit *bool `yaml:"quit"`
	RPop *struct {
		Key string `yaml:"key"`
	} `yaml:"rPop"`
	RPopLPush *struct {
		Source      string `yaml:"source"`
		Destination string `yaml:"destination"`
	} `yaml:"rPopLPush"`
	//RPush *struct {
	//	key    string        `yaml:"key"`
	//	values []interface{} `yaml:"values"`
	//} `yaml:"rPush"`
	//RPushX *struct {
	//	key    string        `yaml:"key"`
	//	values []interface{} `yaml:"values"`
	//} `yaml:"rPushX"`
	RandomKey *bool `yaml:"randomKey"`
	ReadOnly  *bool `yaml:"readOnly"`
	ReadWrite *bool `yaml:"readWrite"`
	Rename    *struct {
		Key    string `yaml:"key"`
		Newkey string `yaml:"newkey"`
	} `yaml:"rename"`
	RenameNX *struct {
		Key    string `yaml:"key"`
		Newkey string `yaml:"newkey"`
	} `yaml:"renameNX"`
	Restore *struct {
		Key   string        `yaml:"key"`
		TTL   time.Duration `yaml:"ttl"`
		Value string        `yaml:"value"`
	} `yaml:"restore"`
	RestoreReplace *struct {
		Key   string        `yaml:"key"`
		TTL   time.Duration `yaml:"ttl"`
		Value string        `yaml:"value"`
	} `yaml:"restoreReplace"`
	//SAdd *struct {
	//	Key     string        `yaml:"key"`
	//	Members []interface{} `yaml:"members"`
	//} `yaml:"sAdd"`
	SCard *struct {
		Key string `yaml:"key"`
	} `yaml:"sCard"`
	SDiff *struct {
		Keys []string `yaml:"keys"`
	} `yaml:"sDiff"`
	SDiffStore *struct {
		Destination string   `yaml:"destination"`
		Keys        []string `yaml:"keys"`
	} `yaml:"sDiffStore"`
	SInter *struct {
		Keys []string `yaml:"keys"`
	} `yaml:"sInter"`
	SInterStore *struct {
		Destination string   `yaml:"destination"`
		Keys        []string `yaml:"keys"`
	} `yaml:"sInterStore"`
	//SIsMember *struct {
	//	Key    string      `yaml:"key"`
	//	Member interface{} `yaml:"member"`
	//} `yaml:"sIsMember"`
	SMembers *struct {
		Key string `yaml:"key"`
	} `yaml:"sMembers"`
	SMembersMap *struct {
		Key string `yaml:"key"`
	} `yaml:"sMembersMap"`
	//SMove *struct {
	//	Source      `yaml:"source"`
	//	Destination string      `yaml:"destination"`
	//	Member      interface{} `yaml:"member"`
	//} `yaml:"sMove"`
	SPop *struct {
		Key string `yaml:"key"`
	} `yaml:"sPop"`
	SPopN *struct {
		Key   string `yaml:"key"`
		Count int64  `yaml:"count"`
	} `yaml:"sPopN"`
	SRandMember *struct {
		Key string `yaml:"key"`
	} `yaml:"sRandMember"`
	SRandMemberN *struct {
		Key   string `yaml:"key"`
		Count int64  `yaml:"count"`
	} `yaml:"sRandMemberN"`
	SRem *struct {
		Key     string        `yaml:"key"`
		Members []interface{} `yaml:"members"`
	} `yaml:"sRem"`
	SScan *struct {
		Key    string `yaml:"key"`
		Cursor uint64 `yaml:"cursor"`
		Match  string `yaml:"match"`
		Count  int64  `yaml:"count"`
	} `yaml:"sScan"`
	SUnion *struct {
		Keys []string `yaml:"keys"`
	} `yaml:"sUnion"`
	SUnionStore *struct {
		Destination string   `yaml:"destination"`
		Keys        []string `yaml:"keys"`
	} `yaml:"sUnionStore"`
	Save *bool `yaml:"save"`
	Scan *struct {
		Cursor uint64 `yaml:"cursor"`
		Match  string `yaml:"match"`
		Count  int64  `yaml:"count"`
	} `yaml:"scan"`
	ScriptExists *struct {
		Hashes []string `yaml:"hashes"`
	} `yaml:"scriptExists"`
	ScriptFlush *bool `yaml:"scriptFlush"`
	ScriptKill  *bool `yaml:"scriptKill"`
	ScriptLoad  *struct {
	} `yaml:"scriptLoad"`
	//Set *struct {
	//	Key        string        `yaml:"key"`
	//	Value      interface{}   `yaml:"value"`
	//	Expiration time.Duration `yaml:"expiration"`
	//} `yaml:"set"`
	SetBit *struct {
		Key    string `yaml:"key"`
		Offset int64  `yaml:"offset"`
		Value  int    `yaml:"value"`
	} `yaml:"setBit"`
	//SetNX *struct {
	//	Key        string        `yaml:"key"`
	//	Value      interface{}   `yaml:"value"`
	//	Expiration time.Duration `yaml:"expiration"`
	//} `yaml:"setNX"`
	SetRange *struct {
		Key    string `yaml:"key"`
		Offset int64  `yaml:"offset"`
		Value  string `yaml:"value"`
	} `yaml:"setRange"`
	//SetXX *struct {
	//	Key        string        `yaml:"key"`
	//	Value      interface{}   `yaml:"value"`
	//	Expiration time.Duration `yaml:"expiration"`
	//} `yaml:"setXX"`
	Shutdown       *bool `yaml:"shutdown"`
	ShutdownNoSave *bool `yaml:"shutdownNoSave"`
	ShutdownSave   *bool `yaml:"shutdownSave"`
	SlaveOf        *struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"slaveOf"`
	SlowLog *bool `yaml:"slowLog"`
	Sort    *struct {
		Key  string   `yaml:"key"`
		Sort *v7.Sort `yaml:"sort"`
	} `yaml:"sort"`
	SortInterfaces *struct {
		Key  string   `yaml:"key"`
		Sort *v7.Sort `yaml:"sort"`
	} `yaml:"sortInterfaces"`
	SortStore *struct {
		Key   string   `yaml:"key"`
		Store string   `yaml:"store"`
		Sort  *v7.Sort `yaml:"sort"`
	} `yaml:"sortStore"`
	StrLen *struct {
		Key string `yaml:"key"`
	} `yaml:"strLen"`
	String *bool `yaml:"string"`
	TTL    *struct {
		Key string `yaml:"key"`
	} `yaml:"ttl"`
	Time  *bool `yaml:"time"`
	Touch *struct {
		Keys []string `yaml:"keys"`
	} `yaml:"touch"`
	Type *struct {
		Key string `yaml:"key"`
	} `yaml:"type"`
	Unlink *struct {
		Keys []string `yaml:"keys"`
	} `yaml:"unlink"`
	Wait *struct {
		NumSlaves int           `yaml:"numSlaves"`
		Timeout   time.Duration `yaml:"timeout"`
	} `yaml:"wait"`
	XAck *struct {
		Stream string   `yaml:"stream"`
		Group  string   `yaml:"group"`
		Ids    []string `yaml:"ids"`
	} `yaml:"xAck"`
	XAdd *struct {
		A *v7.XAddArgs `yaml:"a"`
	} `yaml:"xAdd"`
	XClaim *struct {
		A *v7.XClaimArgs `yaml:"a"`
	} `yaml:"xClaim"`
	XClaimJustID *struct {
		A *v7.XClaimArgs `yaml:"a"`
	} `yaml:"xClaimJustID"`
	XDel *struct {
		Stream string   `yaml:"stream"`
		Ids    []string `yaml:"ids"`
	} `yaml:"xDel"`
	XGroupCreate *struct {
		Stream string `yaml:"stream"`
		Group  string `yaml:"group"`
		Start  string `yaml:"start"`
	} `yaml:"xGroupCreate"`
	XGroupCreateMkStream *struct {
		Stream string `yaml:"stream"`
		Group  string `yaml:"group"`
		Start  string `yaml:"start"`
	} `yaml:"xGroupCreateMkStream"`
	XGroupDelConsumer *struct {
		Stream   string `yaml:"stream"`
		Group    string `yaml:"group"`
		Consumer string `yaml:"consumer"`
	} `yaml:"xGroupDelConsumer"`
	XGroupDestroy *struct {
		Stream string `yaml:"stream"`
		Group  string `yaml:"group"`
	} `yaml:"xGroupDestroy"`
	XGroupSetID *struct {
		Stream string `yaml:"stream"`
		Group  string `yaml:"group"`
		Start  string `yaml:"start"`
	} `yaml:"xGroupSetID"`
	XInfoGroups *struct {
		Key string `yaml:"key"`
	} `yaml:"xInfoGroups"`
	XLen *struct {
		Stream string `yaml:"stream"`
	} `yaml:"xLen"`
	XPending *struct {
		Stream string `yaml:"stream"`
		Group  string `yaml:"group"`
	} `yaml:"xPending"`
	XPendingExt *struct {
		A *v7.XPendingExtArgs `yaml:"a"`
	} `yaml:"xPendingExt"`
	XRange *struct {
		Stream string `yaml:"stream"`
		Start  string `yaml:"start"`
		Stop   string `yaml:"stop"`
	} `yaml:"xRange"`
	XRangeN *struct {
		Stream string `yaml:"stream"`
		Start  string `yaml:"start"`
		Stop   string `yaml:"stop"`
		Count  int64  `yaml:"count"`
	} `yaml:"xRangeN"`
	XRead *struct {
		A *v7.XReadArgs `yaml:"a"`
	} `yaml:"xRead"`
	XReadGroup *struct {
		A *v7.XReadGroupArgs `yaml:"a"`
	} `yaml:"xReadGroup"`
	XReadStreams *struct {
		Streams []string `yaml:"streams"`
	} `yaml:"xReadStreams"`
	XRevRange *struct {
		Stream string `yaml:"stream"`
		Start  string `yaml:"start"`
		Stop   string `yaml:"stop"`
	} `yaml:"xRevRange"`
	XRevRangeN *struct {
		Stream string `yaml:"stream"`
		Start  string `yaml:"start"`
		Stop   string `yaml:"stop"`
		Count  int64  `yaml:"count"`
	} `yaml:"xRevRangeN"`
	XTrim *struct {
		Key    string `yaml:"key"`
		MaxLen int64  `yaml:"maxLen"`
	} `yaml:"xTrim"`
	XTrimApprox *struct {
		Key    string `yaml:"key"`
		MaxLen int64  `yaml:"maxLen"`
	} `yaml:"xTrimApprox"`
	ZAdd *struct {
		Key     string  `yaml:"key"`
		Members []*v7.Z `yaml:"members"`
	} `yaml:"zAdd"`
	ZAddCh *struct {
		Key     string  `yaml:"key"`
		Members []*v7.Z `yaml:"members"`
	} `yaml:"zAddCh"`
	ZAddNX *struct {
		Key     string  `yaml:"key"`
		Members []*v7.Z `yaml:"members"`
	} `yaml:"zAddNX"`
	ZAddNXCh *struct {
		Key     string  `yaml:"key"`
		Members []*v7.Z `yaml:"members"`
	} `yaml:"zAddNXCh"`
	ZAddXX *struct {
		Key     string  `yaml:"key"`
		Members []*v7.Z `yaml:"members"`
	} `yaml:"zAddXX"`
	ZAddXXCh *struct {
		Key     string  `yaml:"key"`
		Members []*v7.Z `yaml:"members"`
	} `yaml:"zAddXXCh"`
	ZCard *struct {
		Key string `yaml:"key"`
	} `yaml:"zCard"`
	ZCount *struct {
		Key string `yaml:"key"`
		Min string `yaml:"min"`
		Max string `yaml:"max"`
	} `yaml:"zCount"`
	ZIncr *struct {
		Key    string `yaml:"key"`
		Member *v7.Z  `yaml:"member"`
	} `yaml:"zIncr"`
	ZIncrBy *struct {
		Key       string  `yaml:"key"`
		Increment float64 `yaml:"increment"`
		Member    string  `yaml:"member"`
	} `yaml:"zIncrBy"`
	ZIncrNX *struct {
		Key    string `yaml:"key"`
		Member *v7.Z  `yaml:"member"`
	} `yaml:"zIncrNX"`
	ZIncrXX *struct {
		Key    string `yaml:"key"`
		Member *v7.Z  `yaml:"member"`
	} `yaml:"zIncrXX"`
	ZInterStore *struct {
		Destination string     `yaml:"destination"`
		Store       *v7.ZStore `yaml:"store"`
	} `yaml:"zInterStore"`
	ZLexCount *struct {
		Key string `yaml:"key"`
		Min string `yaml:"min"`
		Max string `yaml:"max"`
	} `yaml:"zLexCount"`
	ZPopMax *struct {
		Key   string  `yaml:"key"`
		Count []int64 `yaml:"count"`
	} `yaml:"zPopMax"`
	ZPopMin *struct {
		Key   string  `yaml:"key"`
		Count []int64 `yaml:"count"`
	} `yaml:"zPopMin"`
	ZRange *struct {
		Key   string `yaml:"key"`
		Start int64  `yaml:"start"`
		Stop  int64  `yaml:"stop"`
	} `yaml:"zRange"`
	ZRangeByLex *struct {
		Key string       `yaml:"key"`
		Opt *v7.ZRangeBy `yaml:"opt"`
	} `yaml:"zRangeByLex"`
	ZRangeByScore *struct {
		Key string       `yaml:"key"`
		Opt *v7.ZRangeBy `yaml:"opt"`
	} `yaml:"zRangeByScore"`
	ZRangeByScoreWithScores *struct {
		Key string       `yaml:"key"`
		Opt *v7.ZRangeBy `yaml:"opt"`
	} `yaml:"zRangeByScoreWithScores"`
	ZRangeWithScores *struct {
		Key   string `yaml:"key"`
		Start int64  `yaml:"start"`
		Stop  int64  `yaml:"stop"`
	} `yaml:"zRangeWithScores"`
	ZRank *struct {
		Key    string `yaml:"key"`
		Member string `yaml:"member"`
	} `yaml:"zRank"`
	//ZRem *struct {
	//	key     string        `yaml:"key"`
	//	members []interface{} `yaml:"members"`
	//} `yaml:"zRem"`
	ZRemRangeByLex *struct {
		Key string `yaml:"key"`
		Min string `yaml:"min"`
		Max string `yaml:"max"`
	} `yaml:"zRemRangeByLex"`
	ZRemRangeByRank *struct {
		Key   string `yaml:"key"`
		Start int64  `yaml:"start"`
		Stop  int64  `yaml:"stop"`
	} `yaml:"zRemRangeByRank"`
	ZRemRangeByScore *struct {
		Key string `yaml:"key"`
		Min string `yaml:"min"`
		Max string `yaml:"max"`
	} `yaml:"zRemRangeByScore"`
	ZRevRange *struct {
		Key   string `yaml:"key"`
		Start string `yaml:"start"`
		Stop  int64  `yaml:"stop"`
	} `yaml:"zRevRange"`
	ZRevRangeByLex *struct {
		Key string       `yaml:"key"`
		Opt *v7.ZRangeBy `yaml:"opt"`
	} `yaml:"zRevRangeByLex"`
	ZRevRangeByScore *struct {
		Key string       `yaml:"key"`
		Opt *v7.ZRangeBy `yaml:"opt"`
	} `yaml:"zRevRangeByScore"`
	ZRevRangeByScoreWithScores *struct {
		Key string       `yaml:"key"`
		Opt *v7.ZRangeBy `yaml:"opt"`
	} `yaml:"zRevRangeByScoreWithScores"`
	ZRevRangeWithScores *struct {
		Key   string `yaml:"key"`
		Start int64  `yaml:"start"`
		Stop  int64  `yaml:"stop"`
	} `yaml:"zRevRangeWithScores"`
	ZRevRank *struct {
		Key    string `yaml:"key"`
		Member string `yaml:"member"`
	} `yaml:"zRevRank"`
	ZScan *struct {
		Key    string `yaml:"key"`
		Cursor uint64 `yaml:"cursor"`
		Match  string `yaml:"match"`
		Count  int64  `yaml:"count"`
	} `yaml:"zScan"`
	ZScore *struct {
		Key    string `yaml:"key"`
		Member string `yaml:"member"`
	} `yaml:"zScore"`
	ZUnionStore *struct {
		Dest  string     `yaml:"dest"`
		Store *v7.ZStore `yaml:"store"`
	} `yaml:"zUnionStore"`
}
