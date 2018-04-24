package utils

import (
	"encoding/json"
	"os"
	"strings"
)

const (
	RuleFromFile   = "file"
	RuleFromString = "string"

	FormatHostList  = "hostlist"
	FormatBlackList = "blacklist"
)

type RuleEntry struct {
	Type    string   `json:"type"`
	Format  string   `json:"format"`
	Sources []string `json:"sources"`
}

type ServerConfig struct {
	Address string `json:"address"`
}

type RestConfig struct {
	Address string `json:"address`
}

type Config struct {
	Server ServerConfig `json"server"`
	Rest   *RestConfig  `json:"rest"`

	RuleFolder string      `json:"ruledir"`
	Rules      []RuleEntry `json:"rules"`
}

func (r *RuleEntry) IsRawList() bool {
	return strings.EqualFold(r.Type, RuleFromString)
}

func (r *RuleEntry) Parse(s string) (name, addr string) {

	switch r.Format {

	case FormatHostList:
		spl := strings.Fields(s)
		return strings.TrimSpace(spl[1]), strings.TrimSpace(spl[0])
		break

	case FormatBlackList:
	default:
		break
	}

	return s, "0.0.0.0"
}

func DefaultConfig() *Config {
	return &Config{
		Rest: &RestConfig{
			Address: ":8080",
		},
		Server: ServerConfig{
			Address: ":53",
		},
		RuleFolder: "rules",
		Rules: []RuleEntry{
			{
				Type:   "string",
				Format: "hostlist",
				Sources: []string{
					"127.0.0.1 localhost",
					"0.0.0.0 0.0.0.0",
				},
			}, {
				Type:   "file",
				Format: "blacklist",
				Sources: []string{
					"simple_tracking.txt",
				},
			},
		},
	}
}

func ResetConfig(filename string) *Config {
	f, err := os.Create(filename)
	if err != nil {
		return nil
	}
	defer f.Close()

	c := DefaultConfig()

	j := json.NewEncoder(f)
	j.SetIndent("", "\t")
	j.Encode(c)
	return c
}

func ConfigFromFile(filename string) *Config {
	if f, err := os.Open(filename); err == nil {
		defer f.Close()

		j := json.NewDecoder(f)

		c := &Config{}
		err = j.Decode(c)

		if err != nil {
			return nil
		}

		return c

	}
	return nil
}
