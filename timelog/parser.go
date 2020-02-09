package timelog

import (
	"gopkg.in/ini.v1"
)

// Parse parses file data and returns Config.
func Parse(data []byte) (*Config, error) {
	cfg, err := ini.LoadSources(ini.LoadOptions{
		AllowBooleanKeys: true,
	}, data)

	config := &Config{
		Quicklist: make([]string, 0),
	}

	if err != nil {
		return config, err
	}

	if sec, err := cfg.GetSection("quicklist"); err == nil {
		for _, key := range sec.KeyStrings() {
			config.Quicklist = append(config.Quicklist, key)
		}
	}

	return config, nil
}
