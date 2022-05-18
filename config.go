package ion

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config holds the basic structure of Ions
// config file.
type Config struct {
	// ApiKey is the token of which Ion
	// will use along with the path supported by
	// the prefix name to grab server statistics
	// within Sapling.
	//
	ApiKey string `json:"api_key"`

	// AdminUsername is the username that the admin will
	// use to login to the dashboard.
	//
	AdminUsername string `json:"admin_username"`

	// AdminPassword is the password the admin
	// will use to login to the dashboard.
	//
	AdminPassword string `json:"admin_password"`

	// Host holds the host IONs webserver will run on.
	// Ex: :8080 -> localhost:8080
	//
	Host string `json:"host"`
}

// New creates a new instance of the ion config file.
func (c *Config) New(apiToken string, adminUsername string, adminPassword string) *Config {
	return &Config{
		ApiKey:        apiToken,
		AdminUsername: adminUsername,
		AdminPassword: adminPassword,
	}
}

// RetrieveApiKey returns the API key.
func (c *Config) RetrieveApiKey() string {
	return c.Get().ApiKey
}

// RetrieveAdminUsername returns the login username. Primarily used by the
// admin to access the panel.
func (c *Config) RetrieveAdminUsername() string {
	return c.Get().AdminUsername
}

// RetrieveAdminPassword returns the login username. Primarily used by the
// admin to access the panel.
func (c *Config) RetrieveAdminPassword() string {
	return c.Get().AdminPassword
}

func (c *Config) GetHost() string {
	return c.Get().Host
}

// Exists checks if the config file named `ion_config.json` exists. Returns true
// if it does, returns false if it doesn't.
func (c *Config) Exists() bool {
	if _, err := os.Stat("ion_config.json"); os.IsNotExist(err) {
		return false
	}

	return true
}

// Create creates a ion config file and writes the defaults to said file. Throws a panic
// if anything goes wrong during file creation.
func (c *Config) Create() os.File {
	file, err := os.Create("ion_config.json")

	file.Write([]byte(`{
    "api_key": "api-key",
    "admin_username": "username",
    "admin_password": "password",
	"host": ":8080"
}`))
	if err != nil {
		panic(err)
	}

	return *file
}

func (c *Config) Get() Config {
	var configBuffer Config

	file, err := os.Open("ion_config.json")

	if err != nil {
		fmt.Println(err.Error())
	}

	decodedJson := json.NewDecoder(file)
	err = decodedJson.Decode(&configBuffer)
	if err != nil {
		fmt.Println("JSON Decode Err: ", err)
	}

	return configBuffer
}
