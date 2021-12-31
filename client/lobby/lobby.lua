local cjson = require("cjson")
local http = require("socket.http")

require("components.button")
require("common.constants")
require("common.utils")
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
    local url
    if LOCAL then
        url = "http://localhost:8080/api/v1/connect"
    else
        url = "http://lobby.keydream.tk/api/v1/connect"
    end

    local b, c, _ = http.request(
        url,
        [[
            {
                "name":"]]..name..[["
            }
        ]])

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

function Lobby:JoinWorld(world)
    local url
    if LOCAL then
        url = "http://localhost:8080/api/v1/join"
    else
        url = "http://lobby.keydream.tk/api/v1/join"
    end

    local b, c, _ = http.request(
        url,
        [[
            {
                "id":]]..world.id..[[
            }
        ]])

    if c ~= 200 then
        -- TODO - show this in the UI somewhere
        print("Error joining world: " .. c .. " " .. (b or ""))
        return
    end

    local joinRes = cjson.decode(b)
    self.Player:SetState(PlayerState.GAME_CONNECTING)
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
    elseif self.Player.State == PlayerState.LOBBY_DISCONNECTED then
        Lobby.ConnectButton:Draw()
    end
end

function Lobby:mousepressed(x, y, button, istouch, presses)
    if self.Player.State == PlayerState.LOBBY_CONNECTED then
        for k, v in pairs(self.WorldViews) do
            if v:IsButtonPressed(x, y) then
                RPrint(v.World)
                self:JoinWorld(v.World)
                return
            end
        end

        if Lobby.BackButton:IsButtonPressed(x, y) then
            self.Worlds = {}
            self.WorldViews = {}
            self.Player:SetState(PlayerState.LOBBY_DISCONNECTED)
        end
    elseif LocalPlayer.State == PlayerState.LOBBY_DISCONNECTED then
        if self.ConnectButton:IsButtonPressed(x, y) then
            self:Connect(LocalPlayer.Name)
        end
    end
end
