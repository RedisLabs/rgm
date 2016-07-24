package main

func HandleFoo(ctx *RedisModule, args []string) error {
	ctx.ReplyWithSimpleString("It worked!")
	return nil
}

func main() {}
