package timelog

// Load reads config and all time entries from HOME dir.
func Load() (*Config, error) {
	config, err := Parse([]byte(`
		[quicklist]
		project
		tag
	`))

	return config, err
}
