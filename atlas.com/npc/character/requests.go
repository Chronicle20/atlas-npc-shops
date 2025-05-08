package character

import (
	"atlas-npc/rest"
	"fmt"
	"github.com/Chronicle20/atlas-rest/requests"
)

const (
	Resource = "characters"
	ById     = Resource + "/%d"
	ByName   = Resource + "?name=%s"
)

func getBaseRequest() string {
	return requests.RootUrl("CHARACTERS")
}

func requestById(id uint32) requests.Request[RestModel] {
	return rest.MakeGetRequest[RestModel](fmt.Sprintf(getBaseRequest()+ById, id))
}

func requestByName(name string) requests.Request[[]RestModel] {
	return rest.MakeGetRequest[[]RestModel](fmt.Sprintf(getBaseRequest()+ByName, name))
}
