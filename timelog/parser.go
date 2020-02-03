package timelog

import (
	"gopkg.in/ini.v1"
)

// Parse parses file data and returns Config.
func Parse(data []byte) (*Config, error) {
	cfg, err := ini.LoadSources(ini.LoadOptions{
		AllowBooleanKeys: true,
	}, data)

	if err != nil {
		return &Config{}, err
	}

	quicklist := make([]string, 0)
	if sec, err := cfg.GetSection("quicklist"); err == nil {
		for _, key := range sec.KeyStrings() {
			quicklist = append(quicklist, key)
		}
	}

	config := &Config{
		Quicklist: quicklist,
	}

	return config, nil
}
