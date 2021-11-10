package blivechat

type ConfigOption struct {
	index        int
	Options      []string
	OptionValues []string
}

func NewConfigOption(values map[string]string) ConfigOption {
	options, opvalues := make([]string, 0), make([]string, 0)
	for key, val := range values {
		options = append(options, key)
		opvalues = append(opvalues, val)
	}
	return ConfigOption{
		index:        0,
		Options:      options,
		OptionValues: opvalues,
	}
}

func (c *ConfigOption) Prev() string {
	c.index--
	if c.index < 0 {
		c.index = len(c.Options) - 1
	}
	return c.Current()
}

func (c *ConfigOption) Current() string {
	return c.Options[c.index]
}

func (c *ConfigOption) Next() string {
	c.index = (c.index + 1) % len(c.Options)
	return c.Current()
}

func (c *ConfigOption) Value() string {
	return c.OptionValues[c.index]
}

func (c *ConfigOption) SetIndexToValue(value string) {
	for index, val := range c.OptionValues {
		if val == value {
			c.index = index
		}
	}
}
