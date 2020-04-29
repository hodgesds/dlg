package redis

import (
	"context"
	"time"

	v7 "github.com/go-redis/redis/v7"
)

// Config is a config for a Redis stage.
type Config struct {
	Network            string         `yaml:"network,omitempty"`
	Addr               string         `yaml:"addr"`
	DB                 int            `yaml:"db"`
	Password           string         `yaml:"password,omitempty"`
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
	Commands           []*Command     `yaml:"commands"`
}

// Execute is used to execute the config.
func (c *Config) Execute(ctx context.Context) error {
	client := c.Client()
	var err error
	for _, command := range c.Commands {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if command.Append != nil {
			_, err = client.Append(command.Append.Key, command.Append.Value).Result()
			if err != nil {
				return err
			}
		}
		if command.BLPop != nil {
			_, err = client.BLPop(command.BLPop.Timeout, command.BLPop.Keys...).Result()
			if err != nil {
				return err
			}
		}
		if command.BRPop != nil {
			_, err = client.BRPop(command.BRPop.Timeout, command.BRPop.Keys...).Result()
			if err != nil {
				return err
			}
		}
		if command.BRPopLPush != nil {
			_, err = client.BRPopLPush(
				command.BRPopLPush.Source,
				command.BRPopLPush.Dest,
				command.BRPopLPush.Timeout).
				Result()
			if err != nil {
				return err
			}
		}
		if command.BZPopMax != nil {
			_, err = client.BZPopMax(command.BZPopMax.Timeout, command.BZPopMax.Keys...).Result()
			if err != nil {
				return err
			}
		}
		if command.BZPopMin != nil {
			_, err = client.BZPopMin(command.BZPopMin.Timeout, command.BZPopMin.Keys...).Result()
			if err != nil {
				return err
			}
		}
		if command.BgRewriteAOF != nil {
			_, err = client.BgRewriteAOF().Result()
			if err != nil {
				return err
			}
		}
		if command.BgSave != nil {
			_, err = client.BgSave().Result()
			if err != nil {
				return err
			}
		}
		if command.BitCount != nil {
			_, err = client.BitCount(command.BitCount.Key, command.BitCount.BitCount).Result()
			if err != nil {
				return err
			}
		}
		//if command.BitField != nil {
		//	_, err = client.BitField(command.BitField.Key, command.BitField.Args...).Result()
		//	if err != nil {
		//		return err
		//	}
		//}
		if command.BitOpAnd != nil {
			_, err = client.BitOpAnd(command.BitOpAnd.DestKey, command.BitOpAnd.Keys...).Result()
			if err != nil {
				return err
			}
		}
		if command.BitOpNot != nil {
			_, err = client.BitOpNot(command.BitOpNot.DestKey, command.BitOpNot.Key).Result()
			if err != nil {
				return err
			}
		}
		if command.BitOpOr != nil {
			_, err = client.BitOpOr(command.BitOpOr.DestKey, command.BitOpOr.Keys...).Result()
			if err != nil {
				return err
			}
		}
		if command.BitOpXor != nil {
			_, err = client.BitOpXor(command.BitOpXor.DestKey, command.BitOpXor.Keys...).Result()
			if err != nil {
				return err
			}
		}
		if command.BitPos != nil {
			_, err = client.BitPos(command.BitPos.Key, command.BitPos.Bit, command.BitPos.Pos...).Result()
			if err != nil {
				return err
			}
		}
		if command.ClientGetName != nil {
			_, err = client.ClientGetName().Result()
			if err != nil {
				return err
			}
		}
		if command.ClientID != nil {
			_, err = client.ClientID().Result()
			if err != nil {
				return err
			}
		}
		if command.ClientKill != nil {
			_, err = client.ClientKill(command.ClientKill.IPPort).Result()
			if err != nil {
				return err
			}
		}
		if command.ClientKillByFilter != nil {
			_, err = client.ClientKillByFilter(command.ClientKillByFilter.Keys...).Result()
			if err != nil {
				return err
			}
		}
		if command.ClientList != nil {
			_, err = client.ClientList().Result()
			if err != nil {
				return err
			}
		}
		if command.ClientPause != nil {
			_, err = client.ClientPause(command.ClientPause.Dur).Result()
			if err != nil {
				return err
			}
		}
		if command.ClientUnblock != nil {
			_, err = client.ClientUnblock(command.ClientUnblock.ID).Result()
			if err != nil {
				return err
			}
		}
		if command.ClientUnblockWithError != nil {
			_, err = client.ClientUnblockWithError(command.ClientUnblockWithError.ID).Result()
			if err != nil {
				return err
			}
		}
		if command.Close != nil {
			err = client.Close()
			if err != nil {
				return err
			}
		}
		if command.ClusterAddSlots != nil {
			_, err = client.ClusterAddSlots(command.ClusterAddSlots.Slots...).Result()
			if err != nil {
				return err
			}
		}
		if command.ClusterAddSlotsRange != nil {
			_, err = client.ClusterAddSlotsRange(command.ClusterAddSlotsRange.Min, command.ClusterAddSlotsRange.Max).Result()
			if err != nil {
				return err
			}
		}
		if command.ClusterCountFailureReports != nil {
			_, err = client.ClusterCountFailureReports(command.ClusterCountFailureReports.NodeID).Result()
			if err != nil {
				return err
			}
		}
		if command.ClusterCountKeysInSlot != nil {
			_, err = client.ClusterCountKeysInSlot(command.ClusterCountKeysInSlot.Slot).Result()
			if err != nil {
				return err
			}
		}
		if command.ClusterDelSlots != nil {
			_, err = client.ClusterDelSlots(command.ClusterDelSlots.Slots...).Result()
			if err != nil {
				return err
			}
		}
		if command.ClusterDelSlotsRange != nil {
			_, err = client.ClusterDelSlotsRange(
				command.ClusterDelSlotsRange.Min,
				command.ClusterDelSlotsRange.Max,
			).Result()
			if err != nil {
				return err
			}
		}
		if command.ClusterFailover != nil {
			_, err = client.ClusterFailover().Result()
			if err != nil {
				return err
			}
		}
		if command.ClusterForget != nil {
			_, err = client.ClusterForget(command.ClusterForget.NodeID).Result()
			if err != nil {
				return err
			}
		}
		if command.ClusterGetKeysInSlot != nil {
			_, err = client.ClusterGetKeysInSlot(command.ClusterGetKeysInSlot.Slot, command.ClusterGetKeysInSlot.Count).Result()
			if err != nil {
				return err
			}
		}
		if command.ClusterInfo != nil {
			_, err = client.ClusterInfo().Result()
			if err != nil {
				return err
			}
		}
		if command.ClusterKeySlot != nil {
			_, err = client.ClusterKeySlot(command.ClusterKeySlot.Key).Result()
			if err != nil {
				return err
			}
		}
		if command.ClusterMeet != nil {
			_, err = client.ClusterMeet(command.ClusterMeet.Host, command.ClusterMeet.Port).Result()
			if err != nil {
				return err
			}
		}
		if command.ClusterNodes != nil {
			_, err = client.ClusterNodes().Result()
			if err != nil {
				return err
			}
		}
		if command.ClusterReplicate != nil {
			_, err = client.ClusterReplicate(command.ClusterReplicate.NodeID).Result()
			if err != nil {
				return err
			}
		}
		if command.ClusterResetHard != nil {
			_, err = client.ClusterResetHard().Result()
			if err != nil {
				return err
			}
		}
		if command.ClusterResetSoft != nil {
			_, err = client.ClusterResetSoft().Result()
			if err != nil {
				return err
			}
		}
		if command.ClusterSaveConfig != nil {
			_, err = client.ClusterSaveConfig().Result()
			if err != nil {
				return err
			}
		}
		if command.ClusterSlaves != nil {
			_, err = client.ClusterSlaves(command.ClusterSlaves.NodeID).Result()
			if err != nil {
				return err
			}
		}
		if command.ClusterSlots != nil {
			_, err = client.ClusterSlots().Result()
			if err != nil {
				return err
			}
		}
		if command.ConfigGet != nil {
			_, err = client.ConfigGet(command.ConfigGet.Parameter).Result()
			if err != nil {
				return err
			}
		}
		if command.ConfigResetStat != nil {
			_, err = client.ConfigResetStat().Result()
			if err != nil {
				return err
			}
		}
		if command.ConfigRewrite != nil {
			_, err = client.ConfigRewrite().Result()
			if err != nil {
				return err
			}
		}
		if command.ConfigSet != nil {
			_, err = client.ConfigSet(command.ConfigSet.Parameter, command.ConfigSet.Value).Result()
			if err != nil {
				return err
			}
		}
		if command.DBSize != nil {
			_, err = client.DBSize().Result()
			if err != nil {
				return err
			}
		}
		if command.DebugObject != nil {
			_, err = client.ConfigRewrite().Result()
			if err != nil {
				return err
			}
		}
		if command.Decr != nil {
			_, err = client.Decr(command.Decr.Key).Result()
			if err != nil {
				return err
			}
		}
		if command.DecrBy != nil {
			_, err = client.DecrBy(command.DecrBy.Key, command.DecrBy.Decrement).Result()
			if err != nil {
				return err
			}
		}
		if command.Del != nil {
			_, err = client.Del(command.Del.Keys...).Result()
			if err != nil {
				return err
			}
		}
		if command.Dump != nil {
			_, err = client.Dump(command.Dump.Key).Result()
			if err != nil {
				return err
			}
		}
		if command.Echo != nil {
			_, err = client.Echo(command.Echo.Message).Result()
			if err != nil {
				return err
			}
		}
		if command.Exists != nil {
			_, err = client.Exists(command.Exists.Keys...).Result()
			if err != nil {
				return err
			}
		}
		if command.Expire != nil {
			_, err = client.Expire(command.Expire.Key, command.Expire.Expiration).Result()
			if err != nil {
				return err
			}
		}
		if command.ExpireAt != nil {
			_, err = client.ExpireAt(command.ExpireAt.Key, command.ExpireAt.Time).Result()
			if err != nil {
				return err
			}
		}
		if command.FlushAll != nil {
			_, err = client.FlushAll().Result()
			if err != nil {
				return err
			}
		}
		if command.FlushAllAsync != nil {
			_, err = client.FlushDBAsync().Result()
			if err != nil {
				return err
			}
		}
		if command.FlushDB != nil {
			_, err = client.FlushDB().Result()
			if err != nil {
				return err
			}
		}
		if command.FlushDBAsync != nil {
			_, err = client.FlushDBAsync().Result()
			if err != nil {
				return err
			}
		}
		if command.GeoAdd != nil {
			_, err = client.GeoAdd(command.GeoAdd.Key, command.GeoAdd.GeoLocation...).Result()
			if err != nil {
				return err
			}
		}
		if command.LastSave != nil {
			_, err = client.LastSave().Result()
			if err != nil {
				return err
			}
		}
		if command.GeoDist != nil {
			_, err = client.GeoDist(command.GeoDist.Key, command.GeoDist.Member1, command.GeoDist.Member2, command.GeoDist.Unit).Result()
			if err != nil {
				return err
			}
		}
		if command.GeoHash != nil {
			_, err = client.GeoHash(command.GeoHash.Key, command.GeoHash.Members...).Result()
			if err != nil {
				return err
			}
		}
		if command.GeoPos != nil {
			_, err = client.ConfigRewrite().Result()
			if err != nil {
				return err
			}
		}
		if command.GeoRadius != nil {
			_, err = client.ConfigRewrite().Result()
			if err != nil {
				return err
			}
		}
		if command.GeoRadiusByMember != nil {
			_, err = client.GeoRadiusByMember(command.GeoRadiusByMember.Key, command.GeoRadiusByMember.Member, command.GeoRadiusByMember.Query).Result()
			if err != nil {
				return err
			}
		}
		if command.GeoRadiusByMemberStore != nil {
			_, err = client.GeoRadiusByMemberStore(command.GeoRadiusByMemberStore.Key, command.GeoRadiusByMemberStore.Member, command.GeoRadiusByMemberStore.Query).Result()
			if err != nil {
				return err
			}
		}
		if command.GeoRadiusStore != nil {
			_, err = client.GeoRadiusStore(command.GeoRadiusStore.Key, command.GeoRadiusStore.Longitude, command.GeoRadiusStore.Latitude, command.GeoRadiusStore.Query).Result()
			if err != nil {
				return err
			}
		}
		if command.Get != nil {
			_, err = client.Get(command.Get.Key).Result()
			if err != nil {
				return err
			}
		}
		if command.GetBit != nil {
			_, err = client.GetBit(command.GetBit.Key, command.GetBit.Offset).Result()
			if err != nil {
				return err
			}
		}
		if command.GetRange != nil {
			_, err = client.GetRange(command.GetRange.Key, command.GetRange.Start, command.GetRange.End).Result()
			if err != nil {
				return err
			}
		}
		if command.GetSet != nil {
			_, err = client.GetSet(command.GetSet.Key, command.GetSet.Value).Result()
			if err != nil {
				return err
			}
		}
		if command.HDel != nil {
			_, err = client.HDel(command.HDel.Key, command.HDel.Fields...).Result()
			if err != nil {
				return err
			}
		}
		if command.HExists != nil {
			_, err = client.HExists(command.HExists.Key, command.HExists.Field).Result()
			if err != nil {
				return err
			}
		}
		if command.HGet != nil {
			_, err = client.HGet(command.HGet.Key, command.HGet.Field).Result()
			if err != nil {
				return err
			}
		}
		if command.HGetAll != nil {
			_, err = client.HGetAll(command.HGetAll.Key).Result()
			if err != nil {
				return err
			}
		}
		if command.HIncrBy != nil {
			_, err = client.HIncrBy(command.HIncrBy.Key, command.HIncrBy.Field, command.HIncrBy.Incr).Result()
			if err != nil {
				return err
			}
		}
		if command.HIncrByFloat != nil {
			_, err = client.HIncrByFloat(command.HIncrByFloat.Key, command.HIncrByFloat.Field, command.HIncrByFloat.Incr).Result()
			if err != nil {
				return err
			}
		}
		if command.HKeys != nil {
			_, err = client.HKeys(command.HKeys.Key).Result()
			if err != nil {
				return err
			}
		}
		if command.HLen != nil {
			_, err = client.HLen(command.HLen.Key).Result()
			if err != nil {
				return err
			}
		}
		if command.HMGet != nil {
			_, err = client.HMGet(command.HMGet.Key, command.HMGet.Fields...).Result()
			if err != nil {
				return err
			}
		}
		if command.HScan != nil {
			_, _, err = client.HScan(command.HScan.Key, command.HScan.Cursor, command.HScan.Match, command.HScan.Count).Result()
			if err != nil {
				return err
			}
		}
		if command.HVals != nil {
			_, err = client.HVals(command.HVals.Key).Result()
			if err != nil {
				return err
			}
		}
		if command.Incr != nil {
			_, err = client.Incr(command.Incr.Key).Result()
			if err != nil {
				return err
			}
		}
		if command.IncrBy != nil {
			_, err = client.IncrBy(command.IncrBy.Key, command.IncrBy.Value).Result()
			if err != nil {
				return err
			}
		}
		if command.IncrByFloat != nil {
			_, err = client.IncrByFloat(command.IncrByFloat.Key, command.IncrByFloat.Value).Result()
			if err != nil {
				return err
			}
		}
		if command.Info != nil {
			_, err = client.Info(command.Info.Section...).Result()
			if err != nil {
				return err
			}
		}
		if command.Keys != nil {
			_, err = client.Keys(command.Keys.Pattern).Result()
			if err != nil {
				return err
			}
		}
		if command.LIndex != nil {
			_, err = client.LIndex(command.LIndex.Key, command.LIndex.Index).Result()
			if err != nil {
				return err
			}
		}
		if command.LLen != nil {
			_, err = client.LLen(command.LLen.Key).Result()
			if err != nil {
				return err
			}
		}
		if command.LPop != nil {
			_, err = client.LPop(command.LPop.Key).Result()
			if err != nil {
				return err
			}
		}
		if command.LRange != nil {
			_, err = client.LRange(command.LRange.Key, command.LRange.Start, command.LRange.Stop).Result()
			if err != nil {
				return err
			}
		}
		if command.LTrim != nil {
			_, err = client.LTrim(command.LTrim.Key, command.LTrim.Start, command.LTrim.Stop).Result()
			if err != nil {
				return err
			}
		}
		if command.MGet != nil {
			_, err = client.MGet(command.MGet.Keys...).Result()
			if err != nil {
				return err
			}
		}
		if command.MemoryUsage != nil {
			_, err = client.MemoryUsage(command.MemoryUsage.Key, command.MemoryUsage.Samples...).Result()
			if err != nil {
				return err
			}
		}
		if command.Migrate != nil {
			_, err = client.Migrate(command.Migrate.Host, command.Migrate.Port, command.Migrate.Key, command.Migrate.Db, command.Migrate.Timeout).Result()
			if err != nil {
				return err
			}
		}
		if command.Move != nil {
			_, err = client.Move(command.Move.Key, command.Move.Db).Result()
			if err != nil {
				return err
			}
		}
		if command.ObjectEncoding != nil {
			_, err = client.ObjectEncoding(command.ObjectEncoding.Key).Result()
			if err != nil {
				return err
			}
		}
		if command.ObjectIdleTime != nil {
			_, err = client.ObjectIdleTime(command.ObjectIdleTime.Key).Result()
			if err != nil {
				return err
			}
		}
		if command.ObjectRefCount != nil {
			_, err = client.ObjectRefCount(command.ObjectRefCount.Key).Result()
			if err != nil {
				return err
			}
		}
		if command.PExpire != nil {
			_, err = client.PExpire(command.PExpire.Key, command.PExpire.Expiration).Result()
			if err != nil {
				return err
			}
		}
		if command.PExpireAt != nil {
			_, err = client.PExpireAt(command.PExpireAt.Key, command.PExpireAt.Time).Result()
			if err != nil {
				return err
			}
		}
		if command.PFCount != nil {
			_, err = client.PFCount(command.PFCount.Keys...).Result()
			if err != nil {
				return err
			}
		}
		if command.PFMerge != nil {
			_, err = client.PFMerge(command.PFMerge.Dest, command.PFMerge.Keys...).Result()
			if err != nil {
				return err
			}
		}
		if command.PTTL != nil {
			_, err = client.PTTL(command.PTTL.Key).Result()
			if err != nil {
				return err
			}
		}
		if command.Persist != nil {
			_, err = client.Persist(command.Persist.Key).Result()
			if err != nil {
				return err
			}
		}
		if command.Ping != nil {
			_, err = client.Ping().Result()
			if err != nil {
				return err
			}
		}
		if command.PubSubChannels != nil {
			_, err = client.PubSubChannels(command.PubSubChannels.Pattern).Result()
			if err != nil {
				return err
			}
		}
		if command.PubSubNumPat != nil {
			_, err = client.PubSubNumPat().Result()
			if err != nil {
				return err
			}
		}
		if command.PubSubNumSub != nil {
			_, err = client.PubSubNumSub(command.PubSubNumSub.Channels...).Result()
			if err != nil {
				return err
			}
		}
		if command.Publish != nil {
			_, err = client.Publish(command.Publish.Channel, command.Publish.Message).Result()
			if err != nil {
				return err
			}
		}
		if command.Quit != nil {
			_, err = client.Quit().Result()
			if err != nil {
				return err
			}
		}
		if command.RPop != nil {
			_, err = client.RPop(command.RPop.Key).Result()
			if err != nil {
				return err
			}
		}
		if command.RPopLPush != nil {
			_, err = client.RPopLPush(command.RPopLPush.Source, command.RPopLPush.Destination).Result()
			if err != nil {
				return err
			}
		}
		if command.RandomKey != nil {
			_, err = client.RandomKey().Result()
			if err != nil {
				return err
			}
		}
		if command.ReadOnly != nil {
			_, err = client.ReadOnly().Result()
			if err != nil {
				return err
			}
		}
		if command.ReadWrite != nil {
			_, err = client.ReadWrite().Result()
			if err != nil {
				return err
			}
		}
		if command.Rename != nil {
			_, err = client.Rename(command.Rename.Key, command.Rename.Newkey).Result()
			if err != nil {
				return err
			}
		}
		if command.RenameNX != nil {
			_, err = client.RenameNX(command.RenameNX.Key, command.RenameNX.Newkey).Result()
			if err != nil {
				return err
			}
		}
		if command.Restore != nil {
			_, err = client.Restore(command.Restore.Key, command.Restore.TTL, command.Restore.Value).Result()
			if err != nil {
				return err
			}
		}
		if command.RestoreReplace != nil {
			_, err = client.RestoreReplace(command.RestoreReplace.Key, command.RestoreReplace.TTL, command.RestoreReplace.Value).Result()
			if err != nil {
				return err
			}
		}
		if command.SCard != nil {
			_, err = client.SCard(command.SCard.Key).Result()
			if err != nil {
				return err
			}
		}
		if command.SDiff != nil {
			_, err = client.SDiff(command.SDiff.Keys...).Result()
			if err != nil {
				return err
			}
		}
		if command.SDiffStore != nil {
			_, err = client.SDiffStore(command.SDiffStore.Destination, command.SDiffStore.Keys...).Result()
			if err != nil {
				return err
			}
		}
		if command.SInter != nil {
			_, err = client.SInter(command.SInter.Keys...).Result()
			if err != nil {
				return err
			}
		}
		if command.SInterStore != nil {
			_, err = client.SInterStore(command.SInterStore.Destination, command.SInterStore.Keys...).Result()
			if err != nil {
				return err
			}
		}
		if command.SMembers != nil {
			_, err = client.SMembers(command.SMembers.Key).Result()
			if err != nil {
				return err
			}
		}
		if command.SMembersMap != nil {
			_, err = client.SMembersMap(command.SMembersMap.Key).Result()
			if err != nil {
				return err
			}
		}
		if command.SPop != nil {
			_, err = client.SPop(command.SPop.Key).Result()
			if err != nil {
				return err
			}
		}
		if command.SPopN != nil {
			_, err = client.SPopN(command.SPopN.Key, command.SPopN.Count).Result()
			if err != nil {
				return err
			}
		}
		if command.SRandMember != nil {
			_, err = client.SRandMember(command.SRandMember.Key).Result()
			if err != nil {
				return err
			}
		}
		if command.SRandMemberN != nil {
			_, err = client.SRandMemberN(command.SRandMemberN.Key, command.SRandMemberN.Count).Result()
			if err != nil {
				return err
			}
		}
		if command.SRem != nil {
			_, err = client.SRem(
				command.SRem.Key,
				[]interface{}{command.SRem.Members}...).Result()
			if err != nil {
				return err
			}
		}
		if command.SScan != nil {
			_, _, err = client.SScan(command.SScan.Key, command.SScan.Cursor, command.SScan.Match, command.SScan.Count).Result()
			if err != nil {
				return err
			}
		}
		if command.SUnion != nil {
			_, err = client.SUnion(command.SUnion.Keys...).Result()
			if err != nil {
				return err
			}
		}
		if command.SUnionStore != nil {
			_, err = client.SUnionStore(command.SUnionStore.Destination, command.SUnionStore.Keys...).Result()
			if err != nil {
				return err
			}
		}
		if command.Save != nil {
			_, err = client.Save().Result()
			if err != nil {
				return err
			}
		}
		if command.Scan != nil {
			_, _, err = client.Scan(command.Scan.Cursor, command.Scan.Match, command.Scan.Count).Result()
			if err != nil {
				return err
			}
		}
		if command.ScriptExists != nil {
			_, err = client.ScriptExists(command.ScriptExists.Hashes...).Result()
			if err != nil {
				return err
			}
		}
		if command.ScriptFlush != nil {
			_, err = client.ScriptFlush().Result()
			if err != nil {
				return err
			}
		}
		if command.ScriptKill != nil {
			_, err = client.ScriptKill().Result()
			if err != nil {
				return err
			}
		}
		if command.Set != nil {
			_, err = client.Set(command.Set.Key, command.Set.Value, command.Set.Expiration).Result()
			if err != nil {
				return err
			}
		}
		if command.SetBit != nil {
			_, err = client.SetBit(command.SetBit.Key, command.SetBit.Offset, command.SetBit.Value).Result()
			if err != nil {
				return err
			}
		}
		if command.SetRange != nil {
			_, err = client.SetRange(command.SetRange.Key, command.SetRange.Offset, command.SetRange.Value).Result()
			if err != nil {
				return err
			}
		}
		if command.Shutdown != nil {
			_, err = client.Shutdown().Result()
			if err != nil {
				return err
			}
		}
		if command.ShutdownNoSave != nil {
			_, err = client.ShutdownSave().Result()
			if err != nil {
				return err
			}
		}
		if command.ShutdownSave != nil {
			_, err = client.ShutdownSave().Result()
			if err != nil {
				return err
			}
		}
		if command.SlaveOf != nil {
			_, err = client.SlaveOf(command.SlaveOf.Host, command.SlaveOf.Port).Result()
			if err != nil {
				return err
			}
		}
		if command.Sort != nil {
			_, err = client.Sort(command.Sort.Key, command.Sort.Sort).Result()
			if err != nil {
				return err
			}
		}
		if command.SortInterfaces != nil {
			_, err = client.SortInterfaces(command.SortInterfaces.Key, command.SortInterfaces.Sort).Result()
			if err != nil {
				return err
			}
		}
		if command.SortStore != nil {
			_, err = client.SortStore(command.SortStore.Key, command.SortStore.Store, command.SortStore.Sort).Result()
			if err != nil {
				return err
			}
		}
		if command.TTL != nil {
			_, err = client.TTL(command.TTL.Key).Result()
			if err != nil {
				return err
			}
		}
		if command.Time != nil {
			_, err = client.Time().Result()
			if err != nil {
				return err
			}
		}
		if command.Touch != nil {
			_, err = client.Touch(command.Touch.Keys...).Result()
			if err != nil {
				return err
			}
		}
		if command.Type != nil {
			_, err = client.Type(command.Type.Key).Result()
			if err != nil {
				return err
			}
		}
		if command.Unlink != nil {
			_, err = client.Unlink(command.Unlink.Keys...).Result()
			if err != nil {
				return err
			}
		}
		if command.Wait != nil {
			_, err = client.Wait(command.Wait.NumSlaves, command.Wait.Timeout).Result()
			if err != nil {
				return err
			}
		}
		if command.XAck != nil {
			_, err = client.XAck(command.XAck.Stream, command.XAck.Group, command.XAck.Ids...).Result()
			if err != nil {
				return err
			}
		}
		if command.XAdd != nil {
			_, err = client.XAdd(command.XAdd.A).Result()
			if err != nil {
				return err
			}
		}
		if command.XClaim != nil {
			_, err = client.XClaim(command.XClaim.A).Result()
			if err != nil {
				return err
			}
		}
		if command.XClaimJustID != nil {
			_, err = client.XClaimJustID(command.XClaimJustID.A).Result()
			if err != nil {
				return err
			}
		}
		if command.XDel != nil {
			_, err = client.XDel(command.XDel.Stream, command.XDel.Ids...).Result()
			if err != nil {
				return err
			}
		}
		if command.XGroupCreate != nil {
			_, err = client.XGroupCreate(command.XGroupCreate.Stream, command.XGroupCreate.Group, command.XGroupCreate.Start).Result()
			if err != nil {
				return err
			}
		}
		if command.XGroupCreateMkStream != nil {
			_, err = client.XGroupCreateMkStream(command.XGroupCreateMkStream.Stream, command.XGroupCreateMkStream.Group, command.XGroupCreateMkStream.Start).Result()
			if err != nil {
				return err
			}
		}
		if command.XGroupDelConsumer != nil {
			_, err = client.XGroupDelConsumer(command.XGroupDelConsumer.Stream, command.XGroupDelConsumer.Group, command.XGroupDelConsumer.Consumer).Result()
			if err != nil {
				return err
			}
		}
		if command.XGroupDestroy != nil {
			_, err = client.XGroupDestroy(command.XGroupDestroy.Stream, command.XGroupDestroy.Group).Result()
			if err != nil {
				return err
			}
		}
		if command.XGroupSetID != nil {
			_, err = client.XGroupSetID(command.XGroupSetID.Stream, command.XGroupSetID.Group, command.XGroupSetID.Start).Result()
			if err != nil {
				return err
			}
		}
		if command.XInfoGroups != nil {
			_, err = client.XInfoGroups(command.XInfoGroups.Key).Result()
			if err != nil {
				return err
			}
		}
		if command.XLen != nil {
			_, err = client.XLen(command.XLen.Stream).Result()
			if err != nil {
				return err
			}
		}
		if command.XPending != nil {
			_, err = client.XPending(command.XPending.Stream, command.XPending.Group).Result()
			if err != nil {
				return err
			}
		}
		if command.XPendingExt != nil {
			_, err = client.XPendingExt(command.XPendingExt.A).Result()
			if err != nil {
				return err
			}
		}
		if command.XRange != nil {
			_, err = client.XRange(command.XRange.Stream, command.XRange.Start, command.XRange.Stop).Result()
			if err != nil {
				return err
			}
		}
		if command.XRangeN != nil {
			_, err = client.XRangeN(command.XRangeN.Stream, command.XRangeN.Start, command.XRangeN.Stop, command.XRangeN.Count).Result()
			if err != nil {
				return err
			}
		}
		if command.XRead != nil {
			_, err = client.XRead(command.XRead.A).Result()
			if err != nil {
				return err
			}
		}
		if command.XReadGroup != nil {
			_, err = client.XReadGroup(command.XReadGroup.A).Result()
			if err != nil {
				return err
			}
		}
		if command.XReadStreams != nil {
			_, err = client.XReadStreams(command.XReadStreams.Streams...).Result()
			if err != nil {
				return err
			}
		}
		if command.XRevRange != nil {
			_, err = client.XRevRange(command.XRevRange.Stream, command.XRevRange.Start, command.XRevRange.Stop).Result()
			if err != nil {
				return err
			}
		}
		if command.XRevRangeN != nil {
			_, err = client.XRevRangeN(command.XRevRangeN.Stream, command.XRevRangeN.Start, command.XRevRangeN.Stop, command.XRevRangeN.Count).Result()
			if err != nil {
				return err
			}
		}
		if command.XTrim != nil {
			_, err = client.XTrim(command.XTrim.Key, command.XTrim.MaxLen).Result()
			if err != nil {
				return err
			}
		}
		if command.XTrimApprox != nil {
			_, err = client.XTrimApprox(command.XTrimApprox.Key, command.XTrimApprox.MaxLen).Result()
			if err != nil {
				return err
			}
		}
		if command.ZAdd != nil {
			_, err = client.ZAdd(command.ZAdd.Key, command.ZAdd.Members...).Result()
			if err != nil {
				return err
			}
		}
		if command.ZAddCh != nil {
			_, err = client.ZAddCh(command.ZAddCh.Key, command.ZAddCh.Members...).Result()
			if err != nil {
				return err
			}
		}
		if command.ZAddNX != nil {
			_, err = client.ZAddNX(command.ZAddNX.Key, command.ZAddNX.Members...).Result()
			if err != nil {
				return err
			}
		}
		if command.ZAddNXCh != nil {
			_, err = client.ZAddNXCh(command.ZAddNXCh.Key, command.ZAddNXCh.Members...).Result()
			if err != nil {
				return err
			}
		}
		if command.ZAddXX != nil {
			_, err = client.ZAddXX(command.ZAddXX.Key, command.ZAddXX.Members...).Result()
			if err != nil {
				return err
			}
		}
		if command.ZAddXXCh != nil {
			_, err = client.ZAddXXCh(command.ZAddXXCh.Key, command.ZAddXXCh.Members...).Result()
			if err != nil {
				return err
			}
		}
		if command.ZCard != nil {
			_, err = client.ZCard(command.ZCard.Key).Result()
			if err != nil {
				return err
			}
		}
		if command.ZCount != nil {
			_, err = client.ZCount(command.ZCount.Key, command.ZCount.Min, command.ZCount.Max).Result()
			if err != nil {
				return err
			}
		}
		if command.ZIncr != nil {
			_, err = client.ZIncr(command.ZIncr.Key, command.ZIncr.Member).Result()
			if err != nil {
				return err
			}
		}
		if command.ZIncrBy != nil {
			_, err = client.ZIncrBy(command.ZIncrBy.Key, command.ZIncrBy.Increment, command.ZIncrBy.Member).Result()
			if err != nil {
				return err
			}
		}
		if command.ZIncrNX != nil {
			_, err = client.ZIncrNX(command.ZIncrNX.Key, command.ZIncrNX.Member).Result()
			if err != nil {
				return err
			}
		}
		if command.ZIncrXX != nil {
			_, err = client.ZIncrXX(command.ZIncrXX.Key, command.ZIncrXX.Member).Result()
			if err != nil {
				return err
			}
		}
		if command.ZInterStore != nil {
			_, err = client.ZInterStore(command.ZInterStore.Destination, command.ZInterStore.Store).Result()
			if err != nil {
				return err
			}
		}
		if command.ZLexCount != nil {
			_, err = client.ZLexCount(command.ZLexCount.Key, command.ZLexCount.Min, command.ZLexCount.Max).Result()
			if err != nil {
				return err
			}
		}
		if command.ZPopMax != nil {
			_, err = client.ZPopMax(command.ZPopMax.Key, command.ZPopMax.Count...).Result()
			if err != nil {
				return err
			}
		}
		if command.ZPopMin != nil {
			_, err = client.ZPopMin(command.ZPopMin.Key, command.ZPopMin.Count...).Result()
			if err != nil {
				return err
			}
		}
		if command.ZRange != nil {
			_, err = client.ZRange(command.ZRange.Key, command.ZRange.Start, command.ZRange.Stop).Result()
			if err != nil {
				return err
			}
		}
		if command.ZRangeByLex != nil {
			_, err = client.ZRangeByLex(command.ZRangeByLex.Key, command.ZRangeByLex.Opt).Result()
			if err != nil {
				return err
			}
		}
		if command.ZRangeByScore != nil {
			_, err = client.ZRangeByScore(command.ZRangeByScore.Key, command.ZRangeByScore.Opt).Result()
			if err != nil {
				return err
			}
		}
		if command.ZRangeByScoreWithScores != nil {
			_, err = client.ZRangeByScoreWithScores(command.ZRangeByScoreWithScores.Key, command.ZRangeByScoreWithScores.Opt).Result()
			if err != nil {
				return err
			}
		}
		if command.ZRangeWithScores != nil {
			_, err = client.ZRangeWithScores(command.ZRangeWithScores.Key, command.ZRangeWithScores.Start, command.ZRangeWithScores.Stop).Result()
			if err != nil {
				return err
			}
		}
		if command.ZRank != nil {
			_, err = client.ZRank(command.ZRank.Key, command.ZRank.Member).Result()
			if err != nil {
				return err
			}
		}
		if command.ZRem != nil {
			_, err = client.ZRem(command.ZRem.Key, []interface{}{command.ZRem.Members}...).Result()
			if err != nil {
				return err
			}
		}
		if command.ZRemRangeByLex != nil {
			_, err = client.ZRemRangeByLex(command.ZRemRangeByLex.Key, command.ZRemRangeByLex.Min, command.ZRemRangeByLex.Max).Result()
			if err != nil {
				return err
			}
		}
		if command.ZRemRangeByRank != nil {
			_, err = client.ZRemRangeByRank(command.ZRemRangeByRank.Key, command.ZRemRangeByRank.Start, command.ZRemRangeByRank.Stop).Result()
			if err != nil {
				return err
			}
		}
		if command.ZRemRangeByScore != nil {
			_, err = client.ZRemRangeByScore(command.ZRemRangeByScore.Key, command.ZRemRangeByScore.Min, command.ZRemRangeByScore.Max).Result()
			if err != nil {
				return err
			}
		}
		if command.ZRevRange != nil {
			_, err = client.ZRevRange(command.ZRevRange.Key, command.ZRevRange.Start, command.ZRevRange.Stop).Result()
			if err != nil {
				return err
			}
		}
		if command.ZRevRangeByLex != nil {
			_, err = client.ZRevRangeByLex(command.ZRevRangeByLex.Key, command.ZRevRangeByLex.Opt).Result()
			if err != nil {
				return err
			}
		}
		if command.ZRevRangeByScore != nil {
			_, err = client.ZRevRangeByScore(command.ZRevRangeByScore.Key, command.ZRevRangeByScore.Opt).Result()
			if err != nil {
				return err
			}
		}
		if command.ZRevRangeByScoreWithScores != nil {
			_, err = client.ZRevRangeByScoreWithScores(command.ZRevRangeByScoreWithScores.Key, command.ZRevRangeByScoreWithScores.Opt).Result()
			if err != nil {
				return err
			}
		}
		if command.ZRevRangeWithScores != nil {
			_, err = client.ZRevRangeWithScores(command.ZRevRangeWithScores.Key, command.ZRevRangeWithScores.Start, command.ZRevRangeWithScores.Stop).Result()
			if err != nil {
				return err
			}
		}
		if command.ZRevRank != nil {
			_, err = client.ZRevRank(command.ZRevRank.Key, command.ZRevRank.Member).Result()
			if err != nil {
				return err
			}
		}
		if command.ZScan != nil {
			_, _, err = client.ZScan(command.ZScan.Key, command.ZScan.Cursor, command.ZScan.Match, command.ZScan.Count).Result()
			if err != nil {
				return err
			}
		}
		if command.ZScore != nil {
			_, err = client.ZScore(command.ZScore.Key, command.ZScore.Member).Result()
			if err != nil {
				return err
			}
		}
		if command.ZUnionStore != nil {
			_, err = client.ZUnionStore(command.ZUnionStore.Dest, command.ZUnionStore.Store).Result()
			if err != nil {
				return err
			}
		}

	}

	return nil
}

// Client returns a redis client.
func (c *Config) Client() *v7.Client {
	opts := &v7.Options{
		Network:  c.Network,
		Addr:     c.Addr,
		DB:       c.DB,
		Password: c.Password,
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

// Append ...
type Append struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

// BLPop ...
type BLPop struct {
	Timeout time.Duration `yaml:"timeout"`
	Keys    []string      `yaml:"keys"`
}

// BRPop ...
type BRPop struct {
	Timeout time.Duration `yaml:"timeout"`
	Keys    []string      `yaml:"keys"`
}

// BRPopLPush ...
type BRPopLPush struct {
	Source  string        `yaml:"source"`
	Dest    string        `yaml:"dest"`
	Timeout time.Duration `yaml:"timeout"`
}

// BZPopMax ...
type BZPopMax struct {
	Timeout time.Duration `yaml:"timeout"`
	Keys    []string      `yaml:"keys"`
}

// BZPopMin ...
type BZPopMin struct {
	Timeout time.Duration `yaml:"timeout"`
	Keys    []string      `yaml:"keys"`
}

// BitCount ...
type BitCount struct {
	Key      string       `yaml:"key"`
	BitCount *v7.BitCount `yaml:"bitCount"`
}

// BitField ...
type BitField struct {
	Key  string   `yaml:"key"`
	Args []string `yaml:"args"`
}

// BitOpAnd ...
type BitOpAnd struct {
	DestKey string   `yaml:"destKey"`
	Keys    []string `yaml:"keys"`
}

// BitOpNot ...
type BitOpNot struct {
	DestKey string `yaml:"destKey"`
	Key     string `yaml:"key"`
}

// BitOpOr ...
type BitOpOr struct {
	DestKey string   `yaml:"destKey"`
	Keys    []string `yaml:"keys"`
}

// BitOpXor ...
type BitOpXor struct {
	DestKey string   `yaml:"destKey"`
	Keys    []string `yaml:"keys"`
}

// BitPos ...
type BitPos struct {
	Key string  `yaml:"key"`
	Bit int64   `yaml:"bit"`
	Pos []int64 `yaml:"pos"`
}

// ClientKill ...
type ClientKill struct {
	IPPort string `yaml:"ipPort"`
}

// ClientKillByFilter ...
type ClientKillByFilter struct {
	Keys []string `yaml:"keys"`
}

// ClientPause ...
type ClientPause struct {
	Dur time.Duration `yaml:"dur"`
}

// ClientUnblock ...
type ClientUnblock struct {
	ID int64 `yaml:"id"`
}

// ClientUnblockWithError ...
type ClientUnblockWithError struct {
	ID int64 `yaml:"id"`
}

// ClusterAddSlots ...
type ClusterAddSlots struct {
	Slots []int `yaml:"slots"`
}

// ClusterAddSlotsRange ...
type ClusterAddSlotsRange struct {
	Min int `yaml:"min"`
	Max int `yaml:"max"`
}

// ClusterCountFailureReports ...
type ClusterCountFailureReports struct {
	NodeID string `yaml:"nodeID"`
}

// ClusterCountKeysInSlot ...
type ClusterCountKeysInSlot struct {
	Slot int `yaml:"slot"`
}

// ClusterDelSlots ...
type ClusterDelSlots struct {
	Slots []int `yaml:"slots"`
}

// ClusterDelSlotsRange ...
type ClusterDelSlotsRange struct {
	Min int `yaml:"min"`
	Max int `yaml:"max"`
}

// ClusterForget ...
type ClusterForget struct {
	NodeID string `yaml:"nodeID"`
}

// ClusterGetKeysInSlot ...
type ClusterGetKeysInSlot struct {
	Slot  int `yaml:"slot"`
	Count int `yaml:"count"`
}

// ClusterKeySlot ...
type ClusterKeySlot struct {
	Key string `yaml:"key"`
}

// ClusterMeet ...
type ClusterMeet struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

// ClusterReplicate ...
type ClusterReplicate struct {
	NodeID string `yaml:"nodeID"`
}

// ClusterSlaves ...
type ClusterSlaves struct {
	NodeID string `yaml:"nodeID"`
}

// ConfigGet ...
type ConfigGet struct {
	Parameter string `yaml:"parameter"`
}

// ConfigSet ...
type ConfigSet struct {
	Parameter string `yaml:"parameter"`
	Value     string `yaml:"value"`
}

// DebugObject ...
type DebugObject struct {
	Key string `yaml:"key"`
}

// Decr ...
type Decr struct {
	Key string `yaml:"key"`
}

// DecrBy ...
type DecrBy struct {
	Key       string `yaml:"key"`
	Decrement int64  `yaml:"decrement"`
}

// Del ...
type Del struct {
	Keys []string `yaml:"keys"`
}

// Dump ...
type Dump struct {
	Key string `yaml:"key"`
}

// Echo ...
type Echo struct {
	Message string `yaml:"message"`
}

// Exists ...
type Exists struct {
	Keys []string `yaml:"keys"`
}

// Expire ...
type Expire struct {
	Key        string        `yaml:"key"`
	Expiration time.Duration `yaml:"expiration"`
}

// ExpireAt ...
type ExpireAt struct {
	Key  string    `yaml:"key"`
	Time time.Time `yaml:"time"`
}

// GeoAdd ...
type GeoAdd struct {
	Key         string            `yaml:"key"`
	GeoLocation []*v7.GeoLocation `yaml:"geoLocation"`
}

// GeoDist ...
type GeoDist struct {
	Key     string `yaml:"key"`
	Member1 string `yaml:"member1"`
	Member2 string `yaml:"member2"`
	Unit    string `yaml:"unit"`
}

// GeoHash ...
type GeoHash struct {
	Key     string   `yaml:"key"`
	Members []string `yaml:"members"`
}

// GeoPos ...
type GeoPos struct {
	Key     string   `yaml:"key"`
	Members []string `yaml:"members"`
}

// GeoRadius ...
type GeoRadius struct {
	Key       string            `yaml:"key"`
	Longitude float64           `yaml:"longitude"`
	Latitude  float64           `yaml:"latitude"`
	Query     v7.GeoRadiusQuery `yaml:"query"`
}

// GeoRadiusByMember ...
type GeoRadiusByMember struct {
	Key    string             `yaml:"key"`
	Member string             `yaml:"member"`
	Query  *v7.GeoRadiusQuery `yaml:"query"`
}

// GeoRadiusByMemberStore ...
type GeoRadiusByMemberStore struct {
	Key    string             `yaml:"key"`
	Member string             `yaml:"member"`
	Query  *v7.GeoRadiusQuery `yaml:"query"`
}

// GeoRadiusStore ...
type GeoRadiusStore struct {
	Key       string             `yaml:"key"`
	Longitude float64            `yaml:"longitude"`
	Latitude  float64            `yaml:"latitude"`
	Query     *v7.GeoRadiusQuery `yaml:"query"`
}

// Get ...
type Get struct {
	Key string `yaml:"key"`
}

// GetBit ...
type GetBit struct {
	Key    string `yaml:"key"`
	Offset int64  `yaml:"offset"`
}

// GetRange ...
type GetRange struct {
	Key   string `yaml:"key"`
	Start int64  `yaml:"start"`
	End   int64  `yaml:"end"`
}

// GetSet ...
type GetSet struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

// HDel ...
type HDel struct {
	Key    string   `yaml:"key"`
	Fields []string `yaml:"fields"`
}

// HExists ...
type HExists struct {
	Key   string `yaml:"key"`
	Field string `yaml:"field"`
}

// HGet ...
type HGet struct {
	Key   string `yaml:"key"`
	Field string `yaml:"field"`
}

// HGetAll ...
type HGetAll struct {
	Key string `yaml:"key"`
}

// HIncrBy ...
type HIncrBy struct {
	Key   string `yaml:"key"`
	Field string `yaml:"field"`
	Incr  int64  `yaml:"incr"`
}

// HIncrByFloat ...
type HIncrByFloat struct {
	Key   string  `yaml:"key"`
	Field string  `yaml:"field"`
	Incr  float64 `yaml:"incr"`
}

// HKeys ...
type HKeys struct {
	Key string `yaml:"key"`
}

// HLen ...
type HLen struct {
	Key string `yaml:"key"`
}

// HMGet ...
type HMGet struct {
	Key    string   `yaml:"key"`
	Fields []string `yaml:"fields"`
}

// HScan ...
type HScan struct {
	Key    string `yaml:"key"`
	Cursor uint64 `yaml:"cursor"`
	Match  string `yaml:"match"`
	Count  int64  `yaml:"count"`
}

// HVals ...
type HVals struct {
	Key string `yaml:"key"`
}

// Incr ...
type Incr struct {
	Key string `yaml:"key"`
}

// IncrBy ...
type IncrBy struct {
	Key   string `yaml:"key"`
	Value int64  `yaml:"value"`
}

// IncrByFloat ...
type IncrByFloat struct {
	Key   string  `yaml:"key"`
	Value float64 `yaml:"value"`
}

// Info ...
type Info struct {
	Section []string `yaml:"section"`
}

// Keys ...
type Keys struct {
	Pattern string `yaml:"pattern"`
}

// LIndex ...
type LIndex struct {
	Key   string `yaml:"key"`
	Index int64  `yaml:"index"`
}

// LLen ...
type LLen struct {
	Key string `yaml:"key"`
}

// LPop ...
type LPop struct {
	Key string `yaml:"key"`
}

// LRange ...
type LRange struct {
	Key   string `yaml:"key"`
	Start int64  `yaml:"start"`
	Stop  int64  `yaml:"stop"`
}

// LTrim ...
type LTrim struct {
	Key   string `yaml:"key"`
	Start int64  `yaml:"start"`
	Stop  int64  `yaml:"stop"`
}

// MGet ...
type MGet struct {
	Keys []string `yaml:"keys"`
}

// MemoryUsage ...
type MemoryUsage struct {
	Key     string `yaml:"key"`
	Samples []int  `yaml:"samples"`
}

// Migrate ...
type Migrate struct {
	Host    string        `yaml:"host"`
	Port    string        `yaml:"port"`
	Key     string        `yaml:"key"`
	Db      int           `yaml:"db"`
	Timeout time.Duration `yaml:"timeout"`
}

// Move ...
type Move struct {
	Key string `yaml:"key"`
	Db  int    `yaml:"db"`
}

// ObjectEncoding ...
type ObjectEncoding struct {
	Key string `yaml:"key"`
}

// ObjectIdleTime ...
type ObjectIdleTime struct {
	Key string `yaml:"key"`
}

// ObjectRefCount ...
type ObjectRefCount struct {
	Key string `yaml:"key"`
}

// PExpire ...
type PExpire struct {
	Key        string        `yaml:"key"`
	Expiration time.Duration `yaml:"expiration"`
}

// PExpireAt ...
type PExpireAt struct {
	Key  string    `yaml:"key"`
	Time time.Time `yaml:"time"`
}

// PFCount ...
type PFCount struct {
	Keys []string `yaml:"keys"`
}

// PFMerge ...
type PFMerge struct {
	Dest string   `yaml:"dest"`
	Keys []string `yaml:"keys"`
}

// PTTL ...
type PTTL struct {
	Key string `yaml:"key"`
}

// Persist ...
type Persist struct {
	Key string `yaml:"key"`
}

// PubSubChannels ...
type PubSubChannels struct {
	Pattern string `yaml:"pattern"`
}

// PubSubNumSub ...
type PubSubNumSub struct {
	Channels []string `yaml:"channels"`
}

// Publish ...
type Publish struct {
	Channel string `yaml:"channel"`
	Message string `yaml:"message"`
}

// RPop ...
type RPop struct {
	Key string `yaml:"key"`
}

// RPopLPush ...
type RPopLPush struct {
	Source      string `yaml:"source"`
	Destination string `yaml:"destination"`
}

// Rename ...
type Rename struct {
	Key    string `yaml:"key"`
	Newkey string `yaml:"newkey"`
}

// RenameNX ...
type RenameNX struct {
	Key    string `yaml:"key"`
	Newkey string `yaml:"newkey"`
}

// Restore ...
type Restore struct {
	Key   string        `yaml:"key"`
	TTL   time.Duration `yaml:"ttl"`
	Value string        `yaml:"value"`
}

// RestoreReplace ...
type RestoreReplace struct {
	Key   string        `yaml:"key"`
	TTL   time.Duration `yaml:"ttl"`
	Value string        `yaml:"value"`
}

// SCard ...
type SCard struct {
	Key string `yaml:"key"`
}

// SDiff ...
type SDiff struct {
	Keys []string `yaml:"keys"`
}

// SDiffStore ...
type SDiffStore struct {
	Destination string   `yaml:"destination"`
	Keys        []string `yaml:"keys"`
}

// SInter ...
type SInter struct {
	Keys []string `yaml:"keys"`
}

// SInterStore ...
type SInterStore struct {
	Destination string   `yaml:"destination"`
	Keys        []string `yaml:"keys"`
}

// SMembers ...
type SMembers struct {
	Key string `yaml:"key"`
}

// SMembersMap ...
type SMembersMap struct {
	Key string `yaml:"key"`
}

// SPop ...
type SPop struct {
	Key string `yaml:"key"`
}

// SPopN ...
type SPopN struct {
	Key   string `yaml:"key"`
	Count int64  `yaml:"count"`
}

// SRandMember ...
type SRandMember struct {
	Key string `yaml:"key"`
}

// SRandMemberN ...
type SRandMemberN struct {
	Key   string `yaml:"key"`
	Count int64  `yaml:"count"`
}

// SRem ...
type SRem struct {
	Key     string   `yaml:"key"`
	Members []string `yaml:"members"`
}

// SScan ...
type SScan struct {
	Key    string `yaml:"key"`
	Cursor uint64 `yaml:"cursor"`
	Match  string `yaml:"match"`
	Count  int64  `yaml:"count"`
}

// SUnion ...
type SUnion struct {
	Keys []string `yaml:"keys"`
}

// SUnionStore ...
type SUnionStore struct {
	Destination string   `yaml:"destination"`
	Keys        []string `yaml:"keys"`
}

// Scan ...
type Scan struct {
	Cursor uint64 `yaml:"cursor"`
	Match  string `yaml:"match"`
	Count  int64  `yaml:"count"`
}

// ScriptExists ...
type ScriptExists struct {
	Hashes []string `yaml:"hashes"`
}

// Set ...
type Set struct {
	Key        string        `yaml:"key"`
	Value      string        `yaml:"value"`
	Expiration time.Duration `yaml:"expiration"`
}

// SetBit ...
type SetBit struct {
	Key    string `yaml:"key"`
	Offset int64  `yaml:"offset"`
	Value  int    `yaml:"value"`
}

// SetRange ...
type SetRange struct {
	Key    string `yaml:"key"`
	Offset int64  `yaml:"offset"`
	Value  string `yaml:"value"`
}

// SlaveOf ...
type SlaveOf struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

// Sort ...
type Sort struct {
	Key  string   `yaml:"key"`
	Sort *v7.Sort `yaml:"sort"`
}

// SortInterfaces ...
type SortInterfaces struct {
	Key  string   `yaml:"key"`
	Sort *v7.Sort `yaml:"sort"`
}

// SortStore ...
type SortStore struct {
	Key   string   `yaml:"key"`
	Store string   `yaml:"store"`
	Sort  *v7.Sort `yaml:"sort"`
}

// StrLen ...
type StrLen struct {
	Key string `yaml:"key"`
}

// TTL ...
type TTL struct {
	Key string `yaml:"key"`
}

// Touch ...
type Touch struct {
	Keys []string `yaml:"keys"`
}

// Type ...
type Type struct {
	Key string `yaml:"key"`
}

// Unlink ...
type Unlink struct {
	Keys []string `yaml:"keys"`
}

// Wait ...
type Wait struct {
	NumSlaves int           `yaml:"numSlaves"`
	Timeout   time.Duration `yaml:"timeout"`
}

// XAck ...
type XAck struct {
	Stream string   `yaml:"stream"`
	Group  string   `yaml:"group"`
	Ids    []string `yaml:"ids"`
}

// XAdd ...
type XAdd struct {
	A *v7.XAddArgs `yaml:"a"`
}

// XClaim ...
type XClaim struct {
	A *v7.XClaimArgs `yaml:"a"`
}

// XClaimJustID ...
type XClaimJustID struct {
	A *v7.XClaimArgs `yaml:"a"`
}

// XDel ...
type XDel struct {
	Stream string   `yaml:"stream"`
	Ids    []string `yaml:"ids"`
}

// XGroupCreate ...
type XGroupCreate struct {
	Stream string `yaml:"stream"`
	Group  string `yaml:"group"`
	Start  string `yaml:"start"`
}

// XGroupCreateMkStream ...
type XGroupCreateMkStream struct {
	Stream string `yaml:"stream"`
	Group  string `yaml:"group"`
	Start  string `yaml:"start"`
}

// XGroupDelConsumer ...
type XGroupDelConsumer struct {
	Stream   string `yaml:"stream"`
	Group    string `yaml:"group"`
	Consumer string `yaml:"consumer"`
}

// XGroupDestroy ...
type XGroupDestroy struct {
	Stream string `yaml:"stream"`
	Group  string `yaml:"group"`
}

// XGroupSetID ...
type XGroupSetID struct {
	Stream string `yaml:"stream"`
	Group  string `yaml:"group"`
	Start  string `yaml:"start"`
}

// XInfoGroups ...
type XInfoGroups struct {
	Key string `yaml:"key"`
}

// XLen ...
type XLen struct {
	Stream string `yaml:"stream"`
}

// XPending ...
type XPending struct {
	Stream string `yaml:"stream"`
	Group  string `yaml:"group"`
}

// XPendingExt ...
type XPendingExt struct {
	A *v7.XPendingExtArgs `yaml:"a"`
}

// XRange ...
type XRange struct {
	Stream string `yaml:"stream"`
	Start  string `yaml:"start"`
	Stop   string `yaml:"stop"`
}

// XRangeN ...
type XRangeN struct {
	Stream string `yaml:"stream"`
	Start  string `yaml:"start"`
	Stop   string `yaml:"stop"`
	Count  int64  `yaml:"count"`
}

// XRead ...
type XRead struct {
	A *v7.XReadArgs `yaml:"a"`
}

// XReadGroup ...
type XReadGroup struct {
	A *v7.XReadGroupArgs `yaml:"a"`
}

// XReadStreams ...
type XReadStreams struct {
	Streams []string `yaml:"streams"`
}

// XRevRange ...
type XRevRange struct {
	Stream string `yaml:"stream"`
	Start  string `yaml:"start"`
	Stop   string `yaml:"stop"`
}

// XRevRangeN ...
type XRevRangeN struct {
	Stream string `yaml:"stream"`
	Start  string `yaml:"start"`
	Stop   string `yaml:"stop"`
	Count  int64  `yaml:"count"`
}

// XTrim ...
type XTrim struct {
	Key    string `yaml:"key"`
	MaxLen int64  `yaml:"maxLen"`
}

// XTrimApprox ...
type XTrimApprox struct {
	Key    string `yaml:"key"`
	MaxLen int64  `yaml:"maxLen"`
}

// ZAdd ...
type ZAdd struct {
	Key     string  `yaml:"key"`
	Members []*v7.Z `yaml:"members"`
}

// ZAddCh ...
type ZAddCh struct {
	Key     string  `yaml:"key"`
	Members []*v7.Z `yaml:"members"`
}

// ZAddNX ...
type ZAddNX struct {
	Key     string  `yaml:"key"`
	Members []*v7.Z `yaml:"members"`
}

// ZAddNXCh ...
type ZAddNXCh struct {
	Key     string  `yaml:"key"`
	Members []*v7.Z `yaml:"members"`
}

// ZAddXX ...
type ZAddXX struct {
	Key     string  `yaml:"key"`
	Members []*v7.Z `yaml:"members"`
}

// ZAddXXCh ...
type ZAddXXCh struct {
	Key     string  `yaml:"key"`
	Members []*v7.Z `yaml:"members"`
}

// ZCard ...
type ZCard struct {
	Key string `yaml:"key"`
}

// ZCount ...
type ZCount struct {
	Key string `yaml:"key"`
	Min string `yaml:"min"`
	Max string `yaml:"max"`
}

// ZIncr ...
type ZIncr struct {
	Key    string `yaml:"key"`
	Member *v7.Z  `yaml:"member"`
}

// ZIncrBy ...
type ZIncrBy struct {
	Key       string  `yaml:"key"`
	Increment float64 `yaml:"increment"`
	Member    string  `yaml:"member"`
}

// ZIncrNX ...
type ZIncrNX struct {
	Key    string `yaml:"key"`
	Member *v7.Z  `yaml:"member"`
}

// ZIncrXX ...
type ZIncrXX struct {
	Key    string `yaml:"key"`
	Member *v7.Z  `yaml:"member"`
}

// ZInterStore ...
type ZInterStore struct {
	Destination string     `yaml:"destination"`
	Store       *v7.ZStore `yaml:"store"`
}

// ZLexCount ...
type ZLexCount struct {
	Key string `yaml:"key"`
	Min string `yaml:"min"`
	Max string `yaml:"max"`
}

// ZPopMax ...
type ZPopMax struct {
	Key   string  `yaml:"key"`
	Count []int64 `yaml:"count"`
}

// ZPopMin ...
type ZPopMin struct {
	Key   string  `yaml:"key"`
	Count []int64 `yaml:"count"`
}

// ZRange ...
type ZRange struct {
	Key   string `yaml:"key"`
	Start int64  `yaml:"start"`
	Stop  int64  `yaml:"stop"`
}

// ZRangeByLex ...
type ZRangeByLex struct {
	Key string       `yaml:"key"`
	Opt *v7.ZRangeBy `yaml:"opt"`
}

// ZRangeByScore ...
type ZRangeByScore struct {
	Key string       `yaml:"key"`
	Opt *v7.ZRangeBy `yaml:"opt"`
}

// ZRangeByScoreWithScores ...
type ZRangeByScoreWithScores struct {
	Key string       `yaml:"key"`
	Opt *v7.ZRangeBy `yaml:"opt"`
}

// ZRangeWithScores ...
type ZRangeWithScores struct {
	Key   string `yaml:"key"`
	Start int64  `yaml:"start"`
	Stop  int64  `yaml:"stop"`
}

// ZRank ...
type ZRank struct {
	Key    string `yaml:"key"`
	Member string `yaml:"member"`
}

// ZRem ...
type ZRem struct {
	Key     string   `yaml:"key"`
	Members []string `yaml:"members"`
}

// ZRemRangeByLex ...
type ZRemRangeByLex struct {
	Key string `yaml:"key"`
	Min string `yaml:"min"`
	Max string `yaml:"max"`
}

// ZRemRangeByRank ...
type ZRemRangeByRank struct {
	Key   string `yaml:"key"`
	Start int64  `yaml:"start"`
	Stop  int64  `yaml:"stop"`
}

// ZRemRangeByScore ...
type ZRemRangeByScore struct {
	Key string `yaml:"key"`
	Min string `yaml:"min"`
	Max string `yaml:"max"`
}

// ZRevRange ...
type ZRevRange struct {
	Key   string `yaml:"key"`
	Start int64  `yaml:"start"`
	Stop  int64  `yaml:"stop"`
}

// ZRevRangeByLex ...
type ZRevRangeByLex struct {
	Key string       `yaml:"key"`
	Opt *v7.ZRangeBy `yaml:"opt"`
}

// ZRevRangeByScore ...
type ZRevRangeByScore struct {
	Key string       `yaml:"key"`
	Opt *v7.ZRangeBy `yaml:"opt"`
}

// ZRevRangeByScoreWithScores ...
type ZRevRangeByScoreWithScores struct {
	Key string       `yaml:"key"`
	Opt *v7.ZRangeBy `yaml:"opt"`
}

// ZRevRangeWithScores ...
type ZRevRangeWithScores struct {
	Key   string `yaml:"key"`
	Start int64  `yaml:"start"`
	Stop  int64  `yaml:"stop"`
}

// ZRevRank ...
type ZRevRank struct {
	Key    string `yaml:"key"`
	Member string `yaml:"member"`
}

// ZScan ...
type ZScan struct {
	Key    string `yaml:"key"`
	Cursor uint64 `yaml:"cursor"`
	Match  string `yaml:"match"`
	Count  int64  `yaml:"count"`
}

// ZScore ...
type ZScore struct {
	Key    string `yaml:"key"`
	Member string `yaml:"member"`
}

// ZUnionStore ...
type ZUnionStore struct {
	Dest  string     `yaml:"dest"`
	Store *v7.ZStore `yaml:"store"`
}

// Command is a Redis command.
type Command struct {
	Append                     *Append                     `yaml:"append,omitempty"`
	BLPop                      *BLPop                      `yaml:"blPop,omitempty"`
	BRPop                      *BRPop                      `yaml:"brPop,omitempty"`
	BRPopLPush                 *BRPopLPush                 `yaml:"brPopLPush,omitempty"`
	BZPopMax                   *BZPopMax                   `yaml:"bzPopMax,omitempty"`
	BZPopMin                   *BZPopMin                   `yaml:"bzPopMin,omitempty"`
	BgRewriteAOF               *bool                       `yaml:"bgRewriteAof,omitempty"`
	BgSave                     *bool                       `yaml:"bgSave,omitempty"`
	BitCount                   *BitCount                   `yaml:"bitCount,omitempty"`
	BitField                   *BitField                   `yaml:"bitField,omitempty"`
	BitOpAnd                   *BitOpAnd                   `yaml:"bitOpAnd,omitempty"`
	BitOpNot                   *BitOpNot                   `yaml:"bitOpNot,omitempty"`
	BitOpOr                    *BitOpOr                    `yaml:"bitOpOr,omitempty"`
	BitOpXor                   *BitOpXor                   `yaml:"bitOpXor,omitempty"`
	BitPos                     *BitPos                     `yaml:"bitPos,omitempty,omitempty"`
	ClientGetName              *bool                       `yaml:"clientGetName,omitempty"`
	ClientID                   *bool                       `yaml:"clientId,omitempty"`
	ClientKill                 *ClientKill                 `yaml:"clientKill,omitempty"`
	ClientKillByFilter         *ClientKillByFilter         `yaml:"clientKillByFilter,omitempty"`
	ClientList                 *bool                       `yaml:"clientList,omitempty"`
	ClientPause                *ClientPause                `yaml:"clientPause,omitempty"`
	ClientUnblock              *ClientUnblock              `yaml:"clientUnblock,omitempty"`
	ClientUnblockWithError     *ClientUnblockWithError     `yaml:"clientUnblockWithError"`
	Close                      *bool                       `yaml:"close,omitempty"`
	ClusterAddSlots            *ClusterAddSlots            `yaml:"clusterAddSlots,omitempty"`
	ClusterAddSlotsRange       *ClusterAddSlotsRange       `yaml:"clusterAddSlotsRange,omitempty"`
	ClusterCountFailureReports *ClusterCountFailureReports `yaml:"clusterCountFailureReports,omitempty"`
	ClusterCountKeysInSlot     *ClusterCountKeysInSlot     `yaml:"clusterCountKeysInSlot,omitempty"`
	ClusterDelSlots            *ClusterDelSlots            `yaml:"clusterDelSlots,omitempty"`
	ClusterDelSlotsRange       *ClusterDelSlotsRange       `yaml:"clusterDelSlotsRange,omitempty"`
	ClusterFailover            *bool                       `yaml:"clusterFailover,omitempty"`
	ClusterForget              *ClusterForget              `yaml:"clusterForget,omitempty"`
	ClusterGetKeysInSlot       *ClusterGetKeysInSlot       `yaml:"clusterGetKeysInSlot,omitempty"`
	ClusterInfo                *bool                       `yaml:"clusterInfo,omitempty"`
	ClusterKeySlot             *ClusterKeySlot             `yaml:"clusterKeySlot,omitempty"`
	ClusterMeet                *ClusterMeet                `yaml:"clusterMeet,omitempty"`
	ClusterNodes               *bool                       `yaml:"clusterNodes,omitempty"`
	ClusterReplicate           *ClusterReplicate           `yaml:"clusterReplicate,omitempty"`
	ClusterResetHard           *bool                       `yaml:"clusterResetHard,omitempty"`
	ClusterResetSoft           *bool                       `yaml:"clusterResetSoft,omitempty"`
	ClusterSaveConfig          *bool                       `yaml:"clusterSaveConfig,omitempty"`
	ClusterSlaves              *ClusterSlaves              `yaml:"clusterSlaves,omitempty"`
	ClusterSlots               *bool                       `yaml:"clusterSlots,omitempty"`
	Command                    *bool                       `yaml:"command,omitempty"`
	ConfigGet                  *ConfigGet                  `yaml:"configGet,omitempty"`
	ConfigResetStat            *bool                       `yaml:"configResetStat,omitempty"`
	ConfigRewrite              *bool                       `yaml:"configRewrite,omitempty"`
	ConfigSet                  *ConfigSet                  `yaml:"configSet,omitempty"`
	DBSize                     *bool                       `yaml:"dbSize,omitempty"`
	DebugObject                *DebugObject                `yaml:"debugObject,omitempty"`
	Decr                       *Decr                       `yaml:"decr,omitempty"`
	DecrBy                     *DecrBy                     `yaml:"decrBy,omitempty"`
	Del                        *Del                        `yaml:"del,omitempty"`
	Dump                       *Dump                       `yaml:"dump,omitempty"`
	Echo                       *Echo                       `yaml:"echo,omitempty"`
	Exists                     *Exists                     `yaml:"exists,omitempty"`
	Expire                     *Expire                     `yaml:"expire,omitempty"`
	ExpireAt                   *ExpireAt                   `yaml:"expireAt,omitempty"`
	FlushAll                   *bool                       `yaml:"flushAll,omitempty"`
	FlushAllAsync              *bool                       `yaml:"flushAllAsync,omitempty"`
	FlushDB                    *bool                       `yaml:"flushDB,omitempty"`
	FlushDBAsync               *bool                       `yaml:"flushDBAsync,omitempty"`
	GeoAdd                     *GeoAdd                     `yaml:"geoAdd,omitempty"`
	LastSave                   *bool                       `yaml:"lastSave,omitempty"`
	GeoDist                    *GeoDist                    `yaml:"geoDist,omitempty"`
	GeoHash                    *GeoHash                    `yaml:"geoHash,omitempty"`
	GeoPos                     *GeoPos                     `yaml:"geoPos,omitempty"`
	GeoRadius                  *GeoRadius                  `yaml:"geoRadius,omitempty"`
	GeoRadiusByMember          *GeoRadiusByMember          `yaml:"geoRadiusByMember,omitempty"`
	GeoRadiusByMemberStore     *GeoRadiusByMemberStore     `yaml:"geoRadiusByMemberStore,omitempty"`
	GeoRadiusStore             *GeoRadiusStore             `yaml:"geoRadiusStore,omitempty"`
	Get                        *Get                        `yaml:"get,omitempty"`
	GetBit                     *GetBit                     `yaml:"getBit,omitempty"`
	GetRange                   *GetRange                   `yaml:"getRange,omitempty"`
	GetSet                     *GetSet                     `yaml:"getSet,omitempty"`
	HDel                       *HDel                       `yaml:"hDel,omitempty"`
	HExists                    *HExists                    `yaml:"hExists,omitempty"`
	HGet                       *HGet                       `yaml:"hGet,omitempty"`
	HGetAll                    *HGetAll                    `yaml:"hGetAll,omitempty"`
	HIncrBy                    *HIncrBy                    `yaml:"hIncrBy,omitempty"`
	HIncrByFloat               *HIncrByFloat               `yaml:"hIncrByFloat,omitempty"`
	HKeys                      *HKeys                      `yaml:"hKeys,omitempty"`
	HLen                       *HLen                       `yaml:"hLen,omitempty"`
	HMGet                      *HMGet                      `yaml:"hmGet,omitempty"`
	HScan                      *HScan                      `yaml:"hScan,omitempty"`
	HVals                      *HVals                      `yaml:"hVals,omitempty"`
	Incr                       *Incr                       `yaml:"incr,omitempty"`
	IncrBy                     *IncrBy                     `yaml:"incrBy,omitempty"`
	IncrByFloat                *IncrByFloat                `yaml:"incrByFloat,omitempty"`
	Info                       *Info                       `yaml:"info,omitempty"`
	Keys                       *Keys                       `yaml:"keys,omitempty"`
	LIndex                     *LIndex                     `yaml:"lIndex,omitempty"`
	LLen                       *LLen                       `yaml:"lLen,omitempty"`
	LPop                       *LPop                       `yaml:"lPop,omitempty"`
	LRange                     *LRange                     `yaml:"lRange,omitempty"`
	LTrim                      *LTrim                      `yaml:"lTrim,omitempty"`
	MGet                       *MGet                       `yaml:"mGet,omitempty"`
	MemoryUsage                *MemoryUsage                `yaml:"memoryUsage,omitempty"`
	Migrate                    *Migrate                    `yaml:"migrate,omitempty"`
	Move                       *Move                       `yaml:"move,omitempty"`
	ObjectEncoding             *ObjectEncoding             `yaml:"objectEncoding,omitempty"`
	ObjectIdleTime             *ObjectIdleTime             `yaml:"objectIdleTime,omitempty"`
	ObjectRefCount             *ObjectRefCount             `yaml:"objectRefCount,omitempty"`
	PExpire                    *PExpire                    `yaml:"pExpire,omitempty"`
	PExpireAt                  *PExpireAt                  `yaml:"pExpireAt,omitempty"`
	PFCount                    *PFCount                    `yaml:"pfCount,omitempty"`
	PFMerge                    *PFMerge                    `yaml:"pfMerge,omitempty"`
	PTTL                       *PTTL                       `yaml:"pttl,omitempty"`
	Persist                    *Persist                    `yaml:"persist,omitempty"`
	Ping                       *bool                       `yaml:"ping,omitempty"`
	PubSubChannels             *PubSubChannels             `yaml:"pubSubChannels,omitempty"`
	PubSubNumPat               *bool                       `yaml:"pubSubNumPat,omitempty"`
	PubSubNumSub               *PubSubNumSub               `yaml:"pubSubNumSub,omitempty"`
	Publish                    *Publish                    `yaml:"publish,omitempty"`
	Quit                       *bool                       `yaml:"quit,omitempty"`
	RPop                       *RPop                       `yaml:"rPop,omitempty"`
	RPopLPush                  *RPopLPush                  `yaml:"rPopLPush,omitempty"`
	RandomKey                  *bool                       `yaml:"randomKey,omitempty"`
	ReadOnly                   *bool                       `yaml:"readOnly,omitempty"`
	ReadWrite                  *bool                       `yaml:"readWrite,omitempty"`
	Rename                     *Rename                     `yaml:"rename,omitempty"`
	RenameNX                   *RenameNX                   `yaml:"renameNX,omitempty"`
	Restore                    *Restore                    `yaml:"restore,omitempty"`
	RestoreReplace             *RestoreReplace             `yaml:"restoreReplace,omitempty"`
	SCard                      *SCard                      `yaml:"sCard,omitempty"`
	SDiff                      *SDiff                      `yaml:"sDiff,omitempty"`
	SDiffStore                 *SDiffStore                 `yaml:"sDiffStore,omitempty"`
	SInter                     *SInter                     `yaml:"sInter,omitempty"`
	SInterStore                *SInterStore                `yaml:"sInterStore,omitempty"`
	SMembers                   *SMembers                   `yaml:"sMembers,omitempty"`
	SMembersMap                *SMembersMap                `yaml:"sMembersMap,omitempty"`
	SPop                       *SPop                       `yaml:"sPop,omitempty"`
	SPopN                      *SPopN                      `yaml:"sPopN,omitempty"`
	SRandMember                *SRandMember                `yaml:"sRandMember,omitempty"`
	SRandMemberN               *SRandMemberN               `yaml:"sRandMemberN,omitempty"`
	SRem                       *SRem                       `yaml:"sRem,omitempty"`
	SScan                      *SScan                      `yaml:"sScan,omitempty"`
	SUnion                     *SUnion                     `yaml:"sUnion,omitempty"`
	SUnionStore                *SUnionStore                `yaml:"sUnionStore,omitempty"`
	Save                       *bool                       `yaml:"save,omitempty"`
	Scan                       *Scan                       `yaml:"scan,omitempty"`
	ScriptExists               *ScriptExists               `yaml:"scriptExists,omitempty"`
	ScriptFlush                *bool                       `yaml:"scriptFlush,omitempty"`
	ScriptKill                 *bool                       `yaml:"scriptKill,omitempty"`
	Set                        *Set                        `yaml:"set,omitempty"`
	SetBit                     *SetBit                     `yaml:"setBit,omitempty"`
	SetRange                   *SetRange                   `yaml:"setRange,omitempty"`
	Shutdown                   *bool                       `yaml:"shutdown,omitempty"`
	ShutdownNoSave             *bool                       `yaml:"shutdownNoSave,omitempty"`
	ShutdownSave               *bool                       `yaml:"shutdownSave,omitempty"`
	SlaveOf                    *SlaveOf                    `yaml:"slaveOf,omitempty"`
	Sort                       *Sort                       `yaml:"sort,omitempty"`
	SortInterfaces             *SortInterfaces             `yaml:"sortInterfaces,omitempty"`
	SortStore                  *SortStore                  `yaml:"sortStore,omitempty"`
	TTL                        *TTL                        `yaml:"ttl,omitempty"`
	Time                       *bool                       `yaml:"time,omitempty"`
	Touch                      *Touch                      `yaml:"touch,omitempty"`
	Type                       *Type                       `yaml:"type,omitempty"`
	Unlink                     *Unlink                     `yaml:"unlink,omitempty"`
	Wait                       *Wait                       `yaml:"wait,omitempty"`
	XAck                       *XAck                       `yaml:"xAck,omitempty"`
	XAdd                       *XAdd                       `yaml:"xAdd,omitempty"`
	XClaim                     *XClaim                     `yaml:"xClaim,omitempty"`
	XClaimJustID               *XClaimJustID               `yaml:"xClaimJustID,omitempty"`
	XDel                       *XDel                       `yaml:"xDel,omitempty"`
	XGroupCreate               *XGroupCreate               `yaml:"xGroupCreate,omitempty"`
	XGroupCreateMkStream       *XGroupCreateMkStream       `yaml:"xGroupCreateMkStream,omitempty"`
	XGroupDelConsumer          *XGroupDelConsumer          `yaml:"xGroupDelConsumer,omitempty"`
	XGroupDestroy              *XGroupDestroy              `yaml:"xGroupDestroy,omitempty"`
	XGroupSetID                *XGroupSetID                `yaml:"xGroupSetID,omitempty"`
	XInfoGroups                *XInfoGroups                `yaml:"xInfoGroups,omitempty"`
	XLen                       *XLen                       `yaml:"xLen,omitempty"`
	XPending                   *XPending                   `yaml:"xPending,omitempty"`
	XPendingExt                *XPendingExt                `yaml:"xPendingExt,omitempty"`
	XRange                     *XRange                     `yaml:"xRange,omitempty"`
	XRangeN                    *XRangeN                    `yaml:"xRangeN,omitempty"`
	XRead                      *XRead                      `yaml:"xRead,omitempty"`
	XReadGroup                 *XReadGroup                 `yaml:"xReadGroup,omitempty"`
	XReadStreams               *XReadStreams               `yaml:"xReadStreams,omitempty"`
	XRevRange                  *XRevRange                  `yaml:"xRevRange,omitempty"`
	XRevRangeN                 *XRevRangeN                 `yaml:"xRevRangeN,omitempty"`
	XTrim                      *XTrim                      `yaml:"xTrim,omitempty"`
	XTrimApprox                *XTrimApprox                `yaml:"xTrimApprox,omitempty"`
	ZAdd                       *ZAdd                       `yaml:"zAdd,omitempty"`
	ZAddCh                     *ZAddCh                     `yaml:"zAddCh,omitempty"`
	ZAddNX                     *ZAddNX                     `yaml:"zAddNX,omitempty"`
	ZAddNXCh                   *ZAddNXCh                   `yaml:"zAddNXCh,omitempty"`
	ZAddXX                     *ZAddXX                     `yaml:"zAddXX,omitempty"`
	ZAddXXCh                   *ZAddXXCh                   `yaml:"zAddXXCh,omitempty"`
	ZCard                      *ZCard                      `yaml:"zCard,omitempty"`
	ZCount                     *ZCount                     `yaml:"zCount,omitempty"`
	ZIncr                      *ZIncr                      `yaml:"zIncr,omitempty"`
	ZIncrBy                    *ZIncrBy                    `yaml:"zIncrBy,omitempty"`
	ZIncrNX                    *ZIncrNX                    `yaml:"zIncrNX,omitempty"`
	ZIncrXX                    *ZIncrXX                    `yaml:"zIncrXX,omitempty"`
	ZInterStore                *ZInterStore                `yaml:"zInterStore,omitempty"`
	ZLexCount                  *ZLexCount                  `yaml:"zLexCount,omitempty"`
	ZPopMax                    *ZPopMax                    `yaml:"zPopMax,omitempty"`
	ZPopMin                    *ZPopMin                    `yaml:"zPopMin,omitempty"`
	ZRange                     *ZRange                     `yaml:"zRange,omitempty"`
	ZRangeByLex                *ZRangeByLex                `yaml:"zRangeByLex,omitempty"`
	ZRangeByScore              *ZRangeByScore              `yaml:"zRangeByScore,omitempty"`
	ZRangeByScoreWithScores    *ZRangeByScoreWithScores    `yaml:"zRangeByScoreWithScores,omitempty"`
	ZRangeWithScores           *ZRangeWithScores           `yaml:"zRangeWithScores,omitempty"`
	ZRank                      *ZRank                      `yaml:"zRank,omitempty"`
	ZRem                       *ZRem                       `yaml:"zRem,omitempty"`
	ZRemRangeByLex             *ZRemRangeByLex             `yaml:"zRemRangeByLex,omitempty"`
	ZRemRangeByRank            *ZRemRangeByRank            `yaml:"zRemRangeByRank,omitempty"`
	ZRemRangeByScore           *ZRemRangeByScore           `yaml:"zRemRangeByScore,omitempty"`
	ZRevRange                  *ZRevRange                  `yaml:"zRevRange,omitempty"`
	ZRevRangeByLex             *ZRevRangeByLex             `yaml:"zRevRangeByLex,omitempty"`
	ZRevRangeByScore           *ZRevRangeByScore           `yaml:"zRevRangeByScore,omitempty"`
	ZRevRangeByScoreWithScores *ZRevRangeByScoreWithScores `yaml:"zRevRangeByScoreWithScores,omitempty"`
	ZRevRangeWithScores        *ZRevRangeWithScores        `yaml:"zRevRangeWithScores,omitempty"`
	ZRevRank                   *ZRevRank                   `yaml:"zRevRank,omitempty"`
	ZScan                      *ZScan                      `yaml:"zScan,omitempty"`
	ZScore                     *ZScore                     `yaml:"zScore,omitempty"`
	ZUnionStore                *ZUnionStore                `yaml:"zUnionStore,omitempty"`
}
