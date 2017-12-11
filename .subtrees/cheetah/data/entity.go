package data

import (
	"encoding/json"
	"io"
	"time"
)

// JsonMedia is a interface for struct that can be converted to json representation
type JsonMedia interface {
	//ToBytes() ([]byte, error)
	Marshall(io.Writer) error
	Unmarshall(io.Reader) error
}

// User is the entity to interface with users data layer
type User struct {
	ID               string       `json:"id,omitempty" bson:"_id,omitempty"`
	CadunID          int64        `json:"cadun_id,omitempty" bson:"cadun_id,omitempty"`
	Name             string       `json:"name,omitempty" bson:"name,omitempty"`
	Email            string       `json:"email,omitempty" bson:"email,omitempty"`
	Username         string       `json:"username,omitempty" bson:"username,omitempty"`
	Password         string       `json:"password,omitempty" bson:"password,omitempty"`
	Gender           string       `json:"gender,omitempty" bson:"gender,omitempty"`
	DocType          string       `json:"doc_type,omitempty" bson:"doc_type,omitempty"`
	DocNumber        string       `json:"doc_number,omitempty" bson:"doc_number,omitempty"`
	OptinGlobo       bool         `json:"optin_globo,omitempty" bson:"optin_globo,omitempty"`
	OptinPartners    bool         `json:"optin_partners,omitempty" bson:"optin_partners,omitempty"`
	OptinCellPhone   bool         `json:"optin_cellphone,omitempty" bson:"optin_cellphone,omitempty"`
	OriginServiceID  int64        `json:"origin_service_id,omitempty" bson:"origin_service_id,omitempty"`
	UserType         string       `json:"user_type,omitempty" bson:"user_type,omitempty"`
	StatusID         uint8        `json:"status_id" bson:"status_id"`
	Deleted          bool         `json:"deleted,omitempty" bson:"deleted,omitempty"`
	BornedAt         time.Time    `json:"borned_at,omitempty" bson:"borned_at,omitempty"`
	CreatedAt        time.Time    `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt        time.Time    `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	DisabledAt       time.Time    `json:"disabled_at,omitempty" bson:"disabled_at,omitempty"`
	BlockedAt        time.Time    `json:"blocked_at,omitempty" bson:"blocked_at,omitempty"`
	OptinCellPhoneAt time.Time    `json:"optin_cellphone_at,omitempty" bson:"optin_cellphone_at,omitempty"`
	Address          Address      `json:"address,omitempty" bson:"address,omitempty"`
	Phones           []Phone      `json:"phones,omitempty" bson:"phones,omitempty"`
	Facebooks        []FacebookID `json:"facebooks,omitempty" bson:"facebooks,omitempty"`
}

// Marshall writer a json representation of the user struct instance
func (u User) Marshall(writer io.Writer) error {
	return json.NewEncoder(writer).Encode(&u)
}

// Unmarshall reads a json representation into the user struct instance
func (u User) Unmarshall(reader io.Reader) error {
	return json.NewDecoder(reader).Decode(&u)
}

// Address is the entity that holds user address information
type Address struct {
	Localization string `bson:"ddd"`
	Number       string `json:"number" bson:"number"`
	City         string `json:"city" bson:"city"`
	Province     string `json:"province" bson:"province"`
	Country      string `json:"country" bson:"country"`
}

// Phone is the entity for user phones information
type Phone struct {
	DDD    string `json:"ddd" bson:"ddd"`
	Number string `json:"number" bson:"number"`
	Type   string `json:"type" bson:"type"`
}

// FacebookID is the entity to mapping users to facebook accounts
type FacebookID struct {
	AppID      string    `json:"app_id,omitempty" bson:"app_id,omitempty"`
	UserID     string    `json:"user_id,omitempty" bson:"user_id,omitempty"`
	FacebookID string    `json:"facebook_id,omitempty" bson:"facebook_id,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
