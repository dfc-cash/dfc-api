package dfctoml

import "log"

// ExampleGetTOML gets the dfc.toml file for coins.asia
func ExampleClient_GetStellarToml() {
	_, err := DefaultClient.GetStellarToml("coins.asia")
	if err != nil {
		log.Fatal(err)
	}
}
