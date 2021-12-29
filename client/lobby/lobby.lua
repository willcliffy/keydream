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

local defaultButtonWidth = 200
local defaultButtonHeight = 50

function Lobby:new(o, player)
    o = o or {}
    setmetatable(o, self)
    self.__index = self
    self.Player = player
    self.ConnectButton = Button:newConnectButton(nil,
        love.graphics.getWidth() / 2 - defaultButtonWidth / 2,
        love.graphics.getHeight() * 3 / 4 - defaultButtonHeight / 2)
    return o
end

function Lobby:Connect(name)
    local b, c, h = http.request(backendURL .. "/api/v1/connect", "{\"name\":\"" .. name .. "\"}")

    if c ~= 200 then
        print("Error connecting to backend: " .. c .. " " .. b)
        return
    end

    local connectRes = cjson.decode(b)
    self.Worlds = connectRes.Worlds

    self.Connected = true
    self.Player:SetState(PlayerState.LOBBY_CONNECTED)
end

function Lobby:JoinWorld(worldNumber)
    local b, c, h = http.request(backendURL .. "/api/v1/join", "{\"id\":\"" .. worldNumber .. "\"}")

    if c ~= 200 then
        print("Error joining world: " .. c .. " " .. b)
        return
    end

    local joinRes = cjson.decode(b)

    self.Connect = false
    self.Player:SetState(PlayerState.GAME_CONNECTING)
end

function Lobby:Draw()
    if self.Player.State == PlayerState.LOBBY_CONNECTED then
        local y = World1Y

        love.graphics.setColor(Color1)
        love.graphics.printf("Connected!", 1, 1, 800)

        for k, v in pairs(self.Worlds.Worlds) do
            love.graphics.setColor(Color3)
            love.graphics.rectangle('fill', World1X, y, 4 * defaultButtonWidth, defaultButtonHeight)

            love.graphics.setColor(Color2)
            love.graphics.printf("world " .. v.id, WorldIDOffset + 10, y + 10, 800)
            love.graphics.printf(v.num_players .. "/20 players", WorldNumPlayersOffset + 10, y + 10, 800)

            y = y + 100
        end

        love.graphics.setColor(Color3)
        love.graphics.rectangle('fill', Button1X, Button1Y, defaultButtonWidth, defaultButtonHeight)

        love.graphics.setColor(Color5)
        love.graphics.print('Back', Button1X + 75, Button1Y + 13)
    else
        Lobby.ConnectButton:Draw()
    end
end
