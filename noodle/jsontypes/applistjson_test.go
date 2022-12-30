package jsontypes_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mwinters-stuff/noodle/noodle/jsontypes"
)

func TestDecodeString(t *testing.T) {
	json := `
	{
		"appcount": 2,
		"apps": [
			{
				"appid": "666a52d2389b1864d0c376ef7d3a84e9cc54edb8",
				"name": "Ackee",
				"website": "https://github.com/electerious/Ackee",
				"license": "MIT License",
				"description": "Self-hosted, Node.js based analytics tool for those who care about privacy. Ackee runs on your own server, analyzes the traffic of your websites and provides useful statistics in a minimal interface.",
				"enhanced": false,
				"tile_background": "light",
				"icon": "ackee.png",
				"sha": "d3683e09f5dba0a2c5acb3b4dcae055f6837d23d"
			},
			{
				"appid": "140902edbcc424c09736af28ab2de604c3bde936",
				"name": "AdGuard Home",
				"website": "https://github.com/AdguardTeam/AdGuardHome",
				"license": "GNU General Public License v3.0 only",
				"description": "AdGuard Home is a network-wide software for blocking ads & tracking. After you set it up, it'll cover ALL your home devices, and you don't need any client-side software for that.\r\n\r\nIt operates as a DNS server that re-routes tracking domains to a \"black hole,\" thus preventing your devices from connecting to those servers. It's based on software we use for our public AdGuard DNS servers -- both share a lot of common code.",
				"enhanced": true,
				"tile_background": "light",
				"icon": "adguardhome.png",
				"sha": "ed488a0993be8bff0c59e9bf6fe4fbc2f21cffb7"
			}
	]
	}
	`
	data, err := jsontypes.UnmarshalAppList([]byte(json))
	require.NoError(t, err)

	assert.Equal(t, int64(2), data.Appcount)

	assert.Equal(t, "666a52d2389b1864d0c376ef7d3a84e9cc54edb8", data.Apps[0].Appid)
	assert.Equal(t, "Ackee", data.Apps[0].Name)
	assert.Equal(t, "https://github.com/electerious/Ackee", data.Apps[0].Website)
	assert.Equal(t, "MIT License", data.Apps[0].License)
	assert.Equal(t, "Self-hosted, Node.js based analytics tool for those who care about privacy. Ackee runs on your own server, analyzes the traffic of your websites and provides useful statistics in a minimal interface.", data.Apps[0].Description)
	assert.Equal(t, false, data.Apps[0].Enhanced)
	assert.Equal(t, "light", data.Apps[0].TileBackground)
	assert.Equal(t, "ackee.png", data.Apps[0].Icon)
	assert.Equal(t, "d3683e09f5dba0a2c5acb3b4dcae055f6837d23d", data.Apps[0].SHA)

	assert.Equal(t, "140902edbcc424c09736af28ab2de604c3bde936", data.Apps[1].Appid)
	assert.Equal(t, "AdGuard Home", data.Apps[1].Name)
	assert.Equal(t, "https://github.com/AdguardTeam/AdGuardHome", data.Apps[1].Website)
	assert.Equal(t, "GNU General Public License v3.0 only", data.Apps[1].License)
	assert.Equal(t, "AdGuard Home is a network-wide software for blocking ads & tracking. After you set it up, it'll cover ALL your home devices, and you don't need any client-side software for that.\r\n\r\nIt operates as a DNS server that re-routes tracking domains to a \"black hole,\" thus preventing your devices from connecting to those servers. It's based on software we use for our public AdGuard DNS servers -- both share a lot of common code.", data.Apps[1].Description)
	assert.Equal(t, true, data.Apps[1].Enhanced)
	assert.Equal(t, "light", data.Apps[1].TileBackground)
	assert.Equal(t, "adguardhome.png", data.Apps[1].Icon)
	assert.Equal(t, "ed488a0993be8bff0c59e9bf6fe4fbc2f21cffb7", data.Apps[1].SHA)

}
