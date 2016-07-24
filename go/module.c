// #include <stdio.h>

// #include "../redismodule.h"
// #include "module.h"

// //#define MODULE_NAME "foo"
// int GoCommand(RedisModuleCtx *ctx, RedisModuleString **argv, int argc) {
//     return RedisModule_ReplyWithLongLong(ctx, GoDispatch(ctx, argv, argc));
// }

// int RedisModule_OnLoad(RedisModuleCtx *ctx) {
//     if (RedisModule_Init(ctx, MODULE_NAME, 1, REDISMODULE_APIVER_1) == REDISMODULE_ERR) {
//         return REDISMODULE_ERR;
//     }

//     if (GoOnLoad(ctx) == REDISMODULE_ERR) {
//         return REDISMODULE_ERR;
//     }

//     // if (RedisModule_CreateCommand(
//     //             ctx, "go.foo", GoCommand, "readonly", 1, 1, 1) == REDISMODULE_ERR) {
//     //     return REDISMODULE_ERR;
//     // }

//     return REDISMODULE_OK;
// }
