package shops

const (
	EnvCommandTopic     = "COMMAND_TOPIC_NPC_SHOP"
	CommandShopEnter    = "ENTER"
	CommandShopExit     = "EXIT"
	CommandShopBuy      = "BUY"
	CommandShopSell     = "SELL"
	CommandShopRecharge = "RECHARGE"
)

type Command[E any] struct {
	CharacterId uint32 `json:"characterId"`
	Type        string `json:"type"`
	Body        E      `json:"body"`
}

type CommandShopEnterBody struct {
	NpcTemplateId uint32 `json:"npcTemplateId"`
}

type CommandShopExitBody struct {
}

type CommandShopBuyBody struct {
	Slot           uint16 `json:"slot"`
	ItemTemplateId uint32 `json:"itemTemplateId"`
	Quantity       uint32 `json:"quantity"`
	DiscountPrice  uint32 `json:"discountPrice"`
}

type CommandShopSellBody struct {
	Slot           uint16 `json:"slot"`
	ItemTemplateId uint32 `json:"itemTemplateId"`
	Quantity       uint32 `json:"quantity"`
}

type CommandShopRechargeBody struct {
	Slot uint16 `json:"slot"`
}

const (
	EnvStatusEventTopic    = "EVENT_TOPIC_NPC_SHOP_STATUS"
	StatusEventTypeEntered = "ENTERED"
	StatusEventTypeExited  = "EXITED"
	StatusEventTypeError   = "ERROR"
)

type StatusEvent[E any] struct {
	CharacterId uint32 `json:"characterId"`
	Type        string `json:"type"`
	Body        E      `json:"body"`
}

type StatusEventEnteredBody struct {
	NpcTemplateId uint32 `json:"npcTemplateId"`
}

type StatusEventExitedBody struct {
}

type StatusEventErrorBody struct {
	Error string `json:"error"`
}
