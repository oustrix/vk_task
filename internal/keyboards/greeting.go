package keyboards

import "vk_bot/pkg/botapi"

func NewGreetingKeyboard() botapi.Keyboard {
	return botapi.Keyboard{
		OneTime: false,
		Buttons: [][]botapi.KeyboardButton{
			{
				{
					Action: botapi.KeyboardButtonAction{
						Type:  "text",
						Label: "Самому",
					},
				},
				{
					Action: botapi.KeyboardButtonAction{
						Type:  "text",
						Label: "лучшему",
					},
				},
				{
					Action: botapi.KeyboardButtonAction{
						Type:  "text",
						Label: "ревьюверу",
					},
				},
				{
					Action: botapi.KeyboardButtonAction{
						Type:  "text",
						Label: "<3",
					},
				},
			},
			{
				{
					Action: botapi.KeyboardButtonAction{
						Type:  "text",
						Label: "Хорошего",
					},
				},
				{
					Action: botapi.KeyboardButtonAction{
						Type:  "text",
						Label: "дня)",
					},
				},
			},
		},
	}
}
