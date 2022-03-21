package lang

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var (
	bundle    = i18n.NewBundle(language.English)
	localizer = i18n.NewLocalizer(bundle, language.English.String())
)

var (
	Player = struct {
		AskName string
		Greet   func(string) string
	}{
		AskName: localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:          "playerAskName",
				Description: "Ask the player what their name is",
				Other:       "First, what is your name?",
			},
		}),
		Greet: func(name string) string {
			return localizer.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: &i18n.Message{
					ID:          "playerGreet",
					Description: "Greet the player by their name",
					Other:       "Hi {{.Name}}!",
				},
				TemplateData: map[string]interface{}{
					"Name": name,
				},
			})
		},
	}

	Broker = struct {
		IsLocal     string
		AskAddress  string
		AvailableAt func(string, string) string
	}{
		IsLocal: localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:          "brokerIsLocal",
				Description: "Ask if the broker is running on local",
				Other:       "Is the broker running on local?",
			},
		}),
		AskAddress: localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:          "brokerAskAddress",
				Description: "Ask for the address of the broker",
				Other:       "What's the address of the broker, then?",
			},
		}),
		AvailableAt: func(addr string, port string) string {
			return localizer.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: &i18n.Message{
					ID:          "brokerAvailableAt",
					Description: "Broker local address",
					Other:       "Broker available on your network at:\t{{.LocalAddress}}:{{.BrokerPort}}",
				},
				TemplateData: map[string]interface{}{
					"LocalAddress": addr,
					"BrokerPort":   port,
				},
			})
		},
	}

	Server = struct {
		Creating string
		Created  string
	}{
		Creating: localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:          "serverCreating",
				Description: "Creation of a new server instance in progress",
				Other:       "Creating a new server instance...",
			},
		}),
		Created: localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:          "serverCreated",
				Description: "Server instance created",
				Other:       "Server created!",
			},
		}),
	}

	Game = struct {
		AskReset    string
		AskNextTask string
		Started     string
		ListPlayers func(string) string
	}{
		AskReset: localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:          "gameAskReset",
				Description: "Ask for confirmation to reset any existing game",
				Other:       "Do you want to reset any existing game?\nWARNING! This process cannot be undone.",
			},
		}),
		AskNextTask: localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:          "gameAskNextTask",
				Description: "Ask the player for the next task",
				Other:       "Choose your next task",
			},
		}),
		Started: localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:          "gameStarted",
				Description: "Game started",
				Other:       "Game started!",
			},
		}),
		ListPlayers: func(players string) string {
			return localizer.MustLocalize(&i18n.LocalizeConfig{
				DefaultMessage: &i18n.Message{
					ID:          "gameListPlayers",
					Description: "List the player in the game",
					Other:       "Players in the game: {{.Players}}",
				},
				TemplateData: map[string]interface{}{
					"Players": players,
				},
			})
		},
	}

	Task = struct {
		StartGame   string
		ListPlayers string
		DistribuisciCarte string
		ResetGame   string
		GameReset   string
		ExitGame    string
	}{
		StartGame: localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:          "taskStartGame",
				Description: "Start the game",
				Other:       "Start game",
			},
		}),
		DistribuisciCarte: localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:          "taskDistribuisciCarte",
				Description: "DistribuisciCarte ai giocatori",
				Other:       "Distribuisci Carte",
			},
		}),
		ListPlayers: localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:          "taskListPlayers",
				Description: "List players in the game",
				Other:       "List players",
			},
		}),
		ResetGame: localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:          "taskResetGame",
				Description: "Reset any existing game",
				Other:       "Reset",
			},
		}),
		GameReset: localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:          "taskGameReset",
				Description: "Game has been reset",
				Other:       "Reset done ✅",
			},
		}),
		ExitGame: localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:          "taskExitGame",
				Description: "Exit from the game",
				Other:       "Exit",
			},
		}),
	}
)
