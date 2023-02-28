package schema

type Unban struct { //Payload
	Player    string `bson:"_id,omitempty"`
	God       string `bson:"email,omitempty"`
	Timestamp string `bson:"timestamp,omitempty"`
}

func (unban Unban) Database() string {
	return "judicial"
}

func (unban Unban) Collection() string {
	return "unbans"
}
