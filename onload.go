package main

/*
// The pure C part of the module's initialization callback
#include "common.h"

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
