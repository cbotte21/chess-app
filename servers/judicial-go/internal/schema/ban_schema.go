package schema

type Ban struct { //Payload
	Player    string `bson:"_id,omitempty"`
	God       string `bson:"email,omitempty"`
	Reason    string `bson:"password,omitempty"`
	Expiry    string `bson:"expiry,omitempty"`
	Timestamp string `bson:"timestamp,omitempty"`
}

func (ban Ban) Database() string {
	return "judicial"
}

func (ban Ban) Collection() string {
	return "bans"
}
