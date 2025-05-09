package consumable

import (
	"atlas-npc/rest"
	"fmt"
	"github.com/Chronicle20/atlas-rest/requests"
)

const (
	Resource = "data/consumables"
	ById     = Resource + "/%d"
)

func getBaseRequest() string {
	return requests.RootUrl("DATA")
}

func requestById(id uint32) requests.Request[RestModel] {
	return rest.MakeGetRequest[RestModel](fmt.Sprintf(getBaseRequest()+ById, id))
}
