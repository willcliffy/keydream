local cjson = require("cjson")
local http = require("socket.http")

require("components.button")
require("common.constants")
require("common.utils")
require("components.lobby_worldview")
require("components.lobby_nameinput")

Lobby = {
    Background = nil,

    Player = {},

    ConnectButton = nil,
    NameInput = nil,
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

    self.NameInput = NameInput:new(nil)

    self.Background = Background:new()

    return o
end

function Lobby:Connect()
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
                "name":"]]..self.Player.Name..[["
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
    self.Player:SetState(PlayerState.WORLD_CONNECTING)
end

function Lobby:Draw()
    if self.Player.State == PlayerState.LOBBY_CONNECTED then
        self.Background:Draw()
        love.graphics.setFont(BigFont)
        love.graphics.setColor(Color1)
        love.graphics.printf("Connected to lobby as "..self.Player.Name, 1, 1, 800)

        local y = love.graphics.getHeight() / 6
        for _, v in pairs(self.WorldViews) do
            v:Draw(y)
            y = y + 100
        end

        Lobby.BackButton:Draw()
    elseif self.Player.State == PlayerState.LOBBY_DISCONNECTED then
        self.Background:Draw()
        love.graphics.setFont(HugeFont)
        -- TODO - properly center things
        love.graphics.printf("Keydream", 10.2 * TileSizeScaled, 4 * TileSizeScaled, 800)
        Lobby.NameInput:Draw()
        Lobby.ConnectButton:Draw()
    end
end

function Lobby:mousepressed(x, y, button, istouch, presses)
    if self.Player.State == PlayerState.LOBBY_CONNECTED then
        for k, v in pairs(self.WorldViews) do
            if v:IsButtonPressed(x, y) then
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
        if self.ConnectButton:IsButtonPressed(x, y) and self.NameInput.Text ~= "" then
            LocalPlayer:SetName(self.NameInput.Text)
            self:Connect()
        end
    end
end

function Lobby:keypressed(key)
    if key == "escape" then
        self.Worlds = {}
        self.WorldViews = {}
        self.Player:SetState(PlayerState.LOBBY_DISCONNECTED)
    elseif Player.State == PlayerState.LOBBY_DISCONNECTED then
        if key == "return" and self.NameInput.Text ~= "" then
            LocalPlayer:SetName(self.NameInput.Text)
            self:Connect()
        else
            self.NameInput:keypressed(key)
        end
    end
end
