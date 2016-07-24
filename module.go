package main

/*
// The pure C part of the module's initialization callback
#include "common.h"

static int rm_CreateCmd(RedisModuleCtx *ctx, char *cmd, char *flags, int i, int j, int k) {
	return RedisModule_CreateCommand(ctx, cmd, goDispatch, flags, i,j,k);
}
*/
import "C"

import (
	"errors"
	"strings"
)

var (
	// CommandWrite - The command may modify the data set (it may also read from it).
	CommandWrite = "write"

	// CommandReadOnly - The command returns data from keys but never writes.
	CommandReadOnly = "readonly"

	// CommandAdmin - The command is an administrative command (may change replication or perform similar tasks).
	CommandAdmin = "admin"

	// CommandDenyOOM - The command may use additional memory and should be denied during out of memory conditions.
	CommandDenyOOM = "deny-oom"

	// CommandDenyScript - Don't allow this command in Lua scripts.
	CommandDenyScript = "deny-script"

	// CommandAllowLoading - Allow this command while the server is loading data. Only commands not interacting with the
	// data set should be allowed to run in this mode. If not sure don't use this flag.
	CommandAllowLoading = "allow-loading"

	// CommandPubSub - The command publishes things on Pub/Sub channels.
	CommandPubSub = "pubsub"

	// CommandRandom - The command may have different outputs even starting from the same input arguments and key values.
	CommandRandom = "random"

	// CommandAllowStale - The command is allowed to run on slaves that don't serve stale data. Don't use if you don't know what this means.
	CommandAllowStale = "allow-stale"

	// CommandNoMonitor - Don't propoagate the command on monitor. Use this if the command has sensible data among the arguments.
	CommandNoMonitor = "no-monitor"

	// CommandFast - The command time complexity is not greater than O(log(N)) where N is the size of the collection or anything else
	// representing the normal scalability issue with the command.
	CommandFast = "fast"

	// CommandGetKeysAPI - The command implements the interface to return the arguments that are keys.
	// Used when start/stop/step is not enough because of the command syntax.
	CommandGetKeysAPI = "getkeys-api"

	// CommandNoCluster - The command should not register in Redis Cluster since is not designed to work with it because,
	// for example, is unable to report the position of the keys, programmatically creates key names, or any other reason.
	CommandNoCluster = "no-cluster"
)

// ModuleInitializer is what the module's code gets when it needs to register a module

func InitModule(name string) error {

	moduleName = name
	return nil
}

type commandDescriptor struct {
	name     string
	handler  RedisHandler
	flags    []string
	keyStart int
	keyEnd   int
	keyStep  int
}

var registeredCommands = []commandDescriptor{}

func AddCommand(name string, handler RedisHandler, flags ...string) {

	registeredCommands = append(registeredCommands, commandDescriptor{
		name:     name,
		handler:  handler,
		flags:    flags,
		keyStart: 1,
		keyEnd:   1,
		keyStep:  1,
	})

}

var moduleName string

func registerCmd(ctx *C.RedisModuleCtx, cmd, flags string, handler RedisHandler) error {

	if C.rm_CreateCmd(ctx, C.CString(cmd), C.CString(flags), 1, 1, 1) == C.REDISMODULE_ERR {
		return errors.New("Could not register command")
	}

	handlers[strings.ToLower(cmd)] = handler
	return nil
}

//export goOnLoad
func goOnLoad(ctx *C.RedisModuleCtx) C.int {

	if moduleName == "" {
		panic("No module name is set!")
	}

	for _, cmd := range registeredCommands {
		if err := registerCmd(ctx, cmd.name, strings.Join(cmd.flags, " "), cmd.handler); err != nil {
			return C.REDISMODULE_ERR
		}
	}

	return C.REDISMODULE_OK

}

func MyHandler(m *RedisModule, args []string) error {
	return m.ReplyWithSimpleString("OK!")
}

func init() {
	InitModule("gogo")
	AddCommand("gogo.foo", MyHandler, CommandReadOnly)
	AddCommand("gogo.bar", MyHandler, CommandReadOnly, CommandAdmin)
}

func main() {}
