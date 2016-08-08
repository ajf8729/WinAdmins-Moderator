package bot

//go:generate go run ../cmd/gen/bot_wrappers.go -o wrappers.go

import (
	"github.com/bwmarrin/discordgo"
	"github.com/fzzy/radix/extra/pool"
	"github.com/jonas747/dutil/commandsystem"
	"github.com/jonas747/yagpdb/common"
	"log"
	"strings"
	"time"
)

const (
	VERSION = "0.7 __Interesting__ ALPHA"
)

var (
	Config        *common.Config
	Session       *discordgo.Session
	RedisPool     *pool.Pool
	CommandSystem *commandsystem.System

	// When the bot was started
	Started = time.Now()
)

func Run() {

	Config = common.Conf
	Session = common.BotSession
	RedisPool = common.RedisPool

	Session.State.MaxMessageCount = 100

	CommandSystem = commandsystem.NewSystem(Session, "")
	CommandSystem.SendError = true
	CommandSystem.CensorError = CensorError

	log.Println("Running bot...")
	for _, plugin := range plugins {
		plugin.InitBot()
		log.Println("Initialized bot plugin", plugin.Name())
	}

	Session.AddHandler(HandleReady)
	Session.AddHandler(CustomGuildCreate(HandleGuildCreate))
	Session.AddHandler(CustomGuildDelete(HandleGuildDelete))

	// Session.Debug = true
	// Session.LogLevel = discordgo.LogDebug

	err := Session.Open()
	if err != nil {
		log.Println("Failed opening bot connection", err)
	}
}

// Keys and other sensitive information shouldnt be sent in error messages, but just in case it is
func CensorError(err error) string {
	toCensor := []string{
		common.BotSession.Token,
		common.Conf.ClientSecret,
		common.Conf.PastebinDevKey,
	}

	out := err.Error()
	for _, c := range toCensor {
		out = strings.Replace(out, c, "", -1)
	}

	return out
}
