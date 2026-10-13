package essentials

// TelegramCoreAddresses are publicly known addresses of Telegram core network.
var TelegramCoreAddresses = map[int][]string{
	1: {
		"149.154.175.50:443",
		"[2001:b28:f23d:f001::a]:443",
	},
	2: {
		"149.154.167.51:443",
		"95.161.76.100:443",
		"[2001:67c:04e8:f002::a]:443",
	},
	3: {
		"149.154.175.100:443",
		"[2001:b28:f23d:f003::a]:443",
	},
	4: {
		"149.154.167.91:443",
		"[2001:67c:04e8:f004::a]:443",
	},
	5: {
		"149.154.171.5:443",
		"[2001:b28:f23f:f005::a]:443",
	},
	203: {
		"91.105.192.100:443",
		"[2a0a:f280:0203:000a:5000:0000:0000:0100]:443",
	},
}

// TelegramFallbackAddresses are secondary addresses of Telegram core network.
// The proxy dials them only after every core address of the DC has failed.
var TelegramFallbackAddresses = map[int][]string{
	1: {"[2001:b28:f23d:f001::e]:443"},
	2: {"[2001:67c:04e8:f002::b]:443"},
	3: {"[2001:b28:f23d:f003::e]:443"},
	4: {"[2001:67c:04e8:f004::b]:443"},
	5: {"[2001:b28:f23f:f005::b]:443"},
}
