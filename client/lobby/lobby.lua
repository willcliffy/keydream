Lobby = {
    ConnectButton = nil,
    Connected = false,
    Player = {},
    Worlds = {}
}

local cjson = require("cjson")
local http = require("socket.http")
require("components.button")

local backendURL = "http://localhost:8080"

function Lobby:new(o, player)
    o = o or {}
    setmetatable(o, self)
    self.__index = self
    self.Player = player
    self.ConnectButton = Button:new(nil, love.graphics.getWidth() / 2 - ButtonWidth / 2, love.graphics.getHeight() * 3 / 4 - ButtonHeight / 2, ButtonWidth, ButtonHeight, "Connect", Color1)
    return o
end

function Lobby:Connect(name)
    local b, c, h = http.request(backendURL .. "/api/v1/connect", "{\"name\":\"" .. name .. "\"}")

    if c ~= 200 then
        print("Error connecting to backend: " .. c .. " " .. b)
        return
    end

    self.Worlds = cjson.decode(b)
    self.Connected = true
    self.Player:SetState(PlayerState.LOBBY_CONNECTED)
end

function Lobby:Draw()
    if self.Player.State == PlayerState.Connected then
        local y = World1Y

        love.graphics.setColor(Color1)
        love.graphics.printf("Connected!", 1, 1, 800)

        for k, v in pairs(self.Worlds) do
            love.graphics.setColor(Color3)
            love.graphics.rectangle('fill', World1X, y, 4 * ButtonWidth, ButtonHeight)

            love.graphics.setColor(Color2)
            love.graphics.printf("world " .. v.id, WorldIDOffset + 10, y + 10, 800)
            love.graphics.printf(v.num_players .. "/20 players", WorldNumPlayersOffset + 10, y + 10, 800)
            -- love.graphics.printf(v.region, WorldRegionOffset, i, 800)

            NumWorlds = NumWorlds + 1
            y = y + 100
        end

        love.graphics.setColor(Color3)
        love.graphics.rectangle('fill', Button1X, Button1Y, ButtonWidth, ButtonHeight)

        love.graphics.setColor(Color5)
        love.graphics.print('Back', Button1X + 75, Button1Y + 13)
    else
        Lobby.ConnectButton:Draw()
    end
end
