package main

import "encoding/json"
import "io/ioutil"
import . "fmt"
import "net/http"

var k string = "9729tnbw4yk5hn9dqyjtmycts5qnxmre"

var r Realm
var a AuctionData
var i Item

func main() {
	realmres, _ := http.Get("https://eu.api.battle.net/wow/auction/data/gordunni?locale=en_GB&apikey=" + k)
	realmdata, _ := ioutil.ReadAll(realmres.Body)
	json.Unmarshal(realmdata, &r)
	Println("REALM DATA LOADED AND UNMARSHALLED")

	auctionres, _ := http.Get(r.Files[0].URL)
	auctiondata, _ := ioutil.ReadAll(auctionres.Body)
	json.Unmarshal(auctiondata, &a)
	Println("AUCTION DATA LOADED AND UNMARSHALLED")

	for _, auc := range a.Auctions {
		if auc.Owner == "Пайак" {
			itemres, _ := http.Get(Sprintf("https://eu.api.battle.net/wow/item/%d?locale=ru_RU&apikey=%s", auc.Item, k))

			itemdata, _ := ioutil.ReadAll(itemres.Body)
			json.Unmarshal(itemdata, &i)

			Printf("%d - %s - %.2f G\n", auc.Item, i.Name, float64(auc.Buyout)/10000)
		}
	}

	Println("ITEMS DATA LOADED AND UNMARSHALLED")
}

type Realm struct {
	Files []struct {
		URL          string `json:"url"`
		LastModified int64  `json:"lastModified"`
	} `json:"files"`
}

type AuctionData struct {
	Realms []struct {
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"realms"`
	Auctions []struct {
		Auc        int    `json:"auc"`
		Item       int    `json:"item"`
		Owner      string `json:"owner"`
		OwnerRealm string `json:"ownerRealm"`
		Bid        int    `json:"bid"`
		Buyout     int    `json:"buyout"`
		Quantity   int    `json:"quantity"`
		TimeLeft   string `json:"timeLeft"`
		Rand       int    `json:"rand"`
		Seed       int    `json:"seed"`
		Context    int    `json:"context"`
		BonusLists []struct {
			BonusListID int `json:"bonusListId"`
		} `json:"bonusLists,omitempty"`
	} `json:"auctions"`
}

type Item struct {
	ID          int           `json:"id"`
	Description string        `json:"description"`
	Name        string        `json:"name"`
	Icon        string        `json:"icon"`
	Stackable   int           `json:"stackable"`
	ItemBind    int           `json:"itemBind"`
	BonusStats  []interface{} `json:"bonusStats"`
	ItemSpells  []struct {
		SpellID int `json:"spellId"`
		Spell   struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Icon        string `json:"icon"`
			Description string `json:"description"`
			CastTime    string `json:"castTime"`
			Cooldown    string `json:"cooldown"`
		} `json:"spell"`
		NCharges   int    `json:"nCharges"`
		Consumable bool   `json:"consumable"`
		CategoryID int    `json:"categoryId"`
		Trigger    string `json:"trigger"`
	} `json:"itemSpells"`
	BuyPrice          int  `json:"buyPrice"`
	ItemClass         int  `json:"itemClass"`
	ItemSubClass      int  `json:"itemSubClass"`
	ContainerSlots    int  `json:"containerSlots"`
	InventoryType     int  `json:"inventoryType"`
	Equippable        bool `json:"equippable"`
	ItemLevel         int  `json:"itemLevel"`
	MaxCount          int  `json:"maxCount"`
	MaxDurability     int  `json:"maxDurability"`
	MinFactionID      int  `json:"minFactionId"`
	MinReputation     int  `json:"minReputation"`
	Quality           int  `json:"quality"`
	SellPrice         int  `json:"sellPrice"`
	RequiredSkill     int  `json:"requiredSkill"`
	RequiredLevel     int  `json:"requiredLevel"`
	RequiredSkillRank int  `json:"requiredSkillRank"`
	ItemSource        struct {
		SourceID   int    `json:"sourceId"`
		SourceType string `json:"sourceType"`
	} `json:"itemSource"`
	BaseArmor            int           `json:"baseArmor"`
	HasSockets           bool          `json:"hasSockets"`
	IsAuctionable        bool          `json:"isAuctionable"`
	Armor                int           `json:"armor"`
	DisplayInfoID        int           `json:"displayInfoId"`
	NameDescription      string        `json:"nameDescription"`
	NameDescriptionColor string        `json:"nameDescriptionColor"`
	Upgradable           bool          `json:"upgradable"`
	HeroicTooltip        bool          `json:"heroicTooltip"`
	Context              string        `json:"context"`
	BonusLists           []interface{} `json:"bonusLists"`
	AvailableContexts    []string      `json:"availableContexts"`
	BonusSummary         struct {
		DefaultBonusLists []interface{} `json:"defaultBonusLists"`
		ChanceBonusLists  []interface{} `json:"chanceBonusLists"`
		BonusChances      []interface{} `json:"bonusChances"`
	} `json:"bonusSummary"`
	ArtifactID int `json:"artifactId"`
}
