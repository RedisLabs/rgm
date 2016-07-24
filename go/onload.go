package main

/*
// The pure C part of the module's initialization callback

#include "../redismodule.h"

extern int goOnLoad(RedisModuleCtx *ctx);

int RedisModule_OnLoad(RedisModuleCtx *ctx) {

    if (RedisModule_Init(ctx, "MODULE_NAME", 1, REDISMODULE_APIVER_1) == REDISMODULE_ERR) {
        return REDISMODULE_ERR;
    }

    if (goOnLoad(ctx) == REDISMODULE_ERR) {
        return REDISMODULE_ERR;
    }

    // if (RedisModule_CreateCommand(
    //             ctx, "go.foo", GoCommand, "readonly", 1, 1, 1) == REDISMODULE_ERR) {
    //     return REDISMODULE_ERR;
    // }

    return REDISMODULE_OK;
}
*/
import "C"

type CommandFlag string

const (
	CommandWrite        CommandFlag = "write"
	CommandReadOnly     CommandFlag = "readonly"
	CommandAdmin        CommandFlag = "admin"
	CommandDenyOOM      CommandFlag = "deny-oom"
	CommandDenyScript   CommandFlag = "deny-script"
	CommandAllowLoading CommandFlag = "allow-loading"
	CommandPubSub       CommandFlag = "pubsub"
	CommandRandom       CommandFlag = "random"
	CommandAllowStale   CommandFlag = "allow-stale"
	CommandNoMonitor    CommandFlag = "no-monitor"
	CommandFast         CommandFlag = "fast"
	CommandGetKeysAPI   CommandFlag = "getkeys-api"
	CommandNoCluster    CommandFlag = "no-cluster"
)

type ModuleInitializer interface {
	InitModule(name string)
	AddCommand(name string, handler RedisHandler, flags ...ModuleFlag)
}

//export goOnLoad
func goOnLoad(ctx *C.RedisModuleCtx) C.int {

	if err := registerCmd(ctx, "go.foo", "readonly", HandleFoo); err != nil {
		return C.REDISMODULE_ERR
	}

	return C.REDISMODULE_OK

}
