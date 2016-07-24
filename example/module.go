package main

import "github.com/RedisLabs/rgm/module"

func FooHandler(m *module.Redis, args []string) error {
	return m.ReplyWithSimpleString("OK!")
}

func BarHandler(m *module.Redis, args []string) error {

	vargs := []interface{}{"foo", "bar", 1, 2.3, 4}
	return m.Reply(vargs)
}

func init() {

	mod := module.NewModule("gogo")

	mod.AddCommand("gogo.foo", FooHandler, []string{module.CommandReadOnly}, 1, 1, 1)
	mod.AddCommand("gogo.bar", BarHandler, []string{module.CommandReadOnly}, 1, 1, 1)

	module.InitModule(mod)
}

func main() {}
