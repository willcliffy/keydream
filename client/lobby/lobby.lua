local cjson = require("cjson")
local http = require("socket.http")
require("components.button")
require("common.constants")
require("components.lobby_worldview")

Lobby = {
    Player = {},

    ConnectButton = nil,
    BackButton = nil,

    Worlds = {},
    WorldViews = {},
}

function Lobby:new(o, player)
    o = o or {}
    setmetatable(o, self)
    self.__index = self

    self.Player = player

    self.ConnectButton = Button:newConnectButton(nil,
        love.graphics.getWidth() / 2 - DefaultButtonWidth / 2,
        love.graphics.getHeight() * 3 / 4 - DefaultButtonHeight / 2)

    self.BackButton = Button:newBackButton(nil,
        love.graphics.getWidth() / 2 - DefaultButtonWidth / 2,
        love.graphics.getHeight() * 3 / 4 - DefaultButtonHeight / 2)

        return o
end

function Lobby:Connect(name)
    local b, c, h = http.request(LobbyURL .. "/api/v1/connect", "{\"name\":\"" .. name .. "\"}")

    if c ~= 200 then
        print("Error connecting to backend: " .. c .. " " .. (b or ""))
        return
    end

    local connectRes = cjson.decode(b)

    self.Worlds = connectRes.Worlds
    for k, v in pairs(self.Worlds) do
        self.WorldViews[k] = WorldView:new(nil, v)
    end

    self.Player:SetState(PlayerState.LOBBY_CONNECTED)
end

function Lobby:JoinWorld(worldNumber)
    local b, c, h = http.request(LobbyURL .. "/api/v1/join", "{\"id\":\"" .. worldNumber .. "\"}")

    if c ~= 200 then
        print("Error joining world: " .. c .. " " .. b)
        return
    end

    local joinRes = cjson.decode(b)

    self.Connect = false
    self.Player:SetState(PlayerState.GAME_CONNECTING)

    -- todo - actually connect to world
end

function Lobby:Draw()
    if self.Player.State == PlayerState.LOBBY_CONNECTED then
        love.graphics.setColor(Color1)
        love.graphics.printf("Connected!", 1, 1, 800)

        local y = love.graphics.getHeight() / 4
        for _, v in pairs(self.WorldViews) do
            v:Draw(y)
            y = y + 100
        end

        Lobby.BackButton:Draw()
    else
        Lobby.ConnectButton:Draw()
    end
end
