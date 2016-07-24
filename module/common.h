#ifndef __GOMODULE_COMMON_H__
#define __GOMODULE_COMMON_H__

#include "redismodule.h"

int goOnLoad(RedisModuleCtx *ctx);

char *getModuleName();

int goDispatch(RedisModuleCtx *p0, RedisModuleString **p1, int p2);

static int rm_CreateCmd(RedisModuleCtx *ctx, char *cmd, char *flags, int i, int j, int k);

#endif
