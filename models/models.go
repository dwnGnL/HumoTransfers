package models

type DbData struct {
	DSN string `yaml:"DSN"`
}

type Countries struct {
	ID     int64  `gorm:"column:id;primary_key;autoIncrement"`
	Name   string `gorm:"column:name"`
	Icon   string `gorm:"column:icon"`
	Active bool   `gorm:"column:active;"`
}

type Languages struct {
	ID     int64  `gorm:"column:id;primary_key;autoIncrement"`
	Name   string `gorm:"column:name"`
	Icon   string `gorm:"column:icon"`
	Active bool   `gorm:"column:active;default true"`
}

type SysMessage struct {
	ID     int64  `gorm:"column:id;primary_key;autoIncrement"`
	Name   string `gorm:"column:name"`
	Active bool   `gorm:"column:active;default true"`
}

type Currency struct {
	ID   int64  `gorm:"column:id;primary_key;autoIncrement"`
	Name string `gorm:"column:name;"`
	Icon string `gorm:"column:icon"`
}

type Test struct {
	ID       int64  `gorm:"column:id;primary_key;autoIncrement"`
	Entity   string `gorm:"column:entity"`
	EntityId int64  `gorm:"column:entity_id"`
	LangId   int64  `gorm:"column:lang_id;references:languages"`
	Value    string `gorm:"column:value"`
}

type Agents struct {
	ID        int64  `gorm:"column:id;primary_key;autoIncrement"`
	Name      string `gorm:"column:name"`
	LegalName string `gorm:"column:legal_name"`
	Active    bool   `gorm:"column:active;default true"`
}

type AccountAgent struct {
	ID         int64  `gorm:"column:id;primary_key;autoIncrement"`
	AgentId    int64  `gorm:"column:agent_id;references_agents"`
	CurrencyId int64  `gorm:"column:curr_id;references_currency"`
	Active     bool   `gorm:"column:active;default true"`
	IsDefault  bool   `gorm:"column:is_default;default true"`
	Type       string `gorm:"column:type"`
}

// *bool - null bool