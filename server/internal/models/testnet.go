package models

type Testnet struct {
	Name        string `bson:"nameTestnet json:"nameTestnet"`
	Description string `bson:"descriptionTestnet json:"descriptionTestnet"`
	Link        string `bson:"linkTestnet json:"linkTestnet"`
	DataAirdrop string `bson:"dataAirdropTestnet json:"dataAirdropTestnet"`
}

type UpdateTestnet struct {
	TestnetName    string `bson:"testnetName json:"testnetName"`
	NewName        string `bson:"new_nameTestnet json:"new_nameTestnet"`
	NewDescription string `bson:"new_descriptionTestnet json:"new_descriptionTestnet"`
	NewLink        string `bson:"new_linkTestnet json:"new_linkTestnet"`
	NewDataAirdrop string `bson:"new_dataAirdropTestnet json:"new_dataAirdropTestnet"`
}
