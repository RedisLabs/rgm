package module

/*
// The pure C part of the module's initialization callback
#include "common.h"

int RedisModule_OnLoad(RedisModuleCtx *ctx) {

    if (RedisModule_Init(ctx, (const char*)getModuleName(), 1, REDISMODULE_APIVER_1) == REDISMODULE_ERR) {
        return REDISMODULE_ERR;
    }

    if (goOnLoad(ctx) == REDISMODULE_ERR) {
        return REDISMODULE_ERR;
    }

    return REDISMODULE_OK;
}
*/
import "C"
