local socket = require("socket")
local coroutine = require("coroutine")

require("world.background")
require("world.character")
require("common.utils")

World = {
    IP = "",
    Port = 0,

    Player = nil,

    PlayerCharacter = nil,
    OtherCharacters = {},

    Background = nil,

    UDPConn = nil,
    Connected = false,

    TimeSinceLastTick = 0,
    TickDuration = 0.15,
}

function World:new(o, player, ip, port)
    o = o or {}
    setmetatable(o, self)
    self.__index = self

    self.IP = ip
    self.Port = port
    self.UDPConn = socket.udp()
    self.UDPConn:settimeout(5)

    self.Background = Background:new()
    self.PlayerCharacter = Character:new(nil, 0, 0)

    local error
    _, error = self.UDPConn:setpeername(self.IP, self.Port)
    if error then
        print("Failed to connect to game server: " .. error)
        return false
    end

    self.Player = player

    Background:Update()

    return o
end

function World:Connect()
    if self.Connected then
        return true
    end

    _, error, _ = self.UDPConn:send("join " .. LocalPlayer.Name .. " \n")
    if error then
        print("Error: " .. error)
        return false
    end

    -- todo: find out what partial means and if I need to be aware of it
    local data, partial
    data, error, partial = self.UDPConn:receive()
    if error then
        print("Error: " .. error)
        return false
    end

    self.UDPConn:settimeout(0.01)
    self.Player.ID = tonumber(data)
    self.Connected = true


    print("..self.Player.ID .. " .. self.Player.ID)
    return true
end

function World:Update(dt)
    if self.Connected then
        self.Background:Update()
        self.PlayerCharacter:Update(dt)
        self.TimeSinceLastTick = self.TimeSinceLastTick + dt

        if self.TimeSinceLastTick >= self.TickDuration then
            self.TimeSinceLastTick = self.TimeSinceLastTick - self.TickDuration
            self:Tick(dt)
        end
    end
end

function World:Tick(dt)
    if not self.Connected then
        return
    end

    if self.PlayerCharacter:HasMoved() then
        self.UDPConn:send("move "..self.Player.ID .." "..self.PlayerCharacter.X.." "..self.PlayerCharacter.Y.." \n")

        self.PlayerCharacter.LastX = self.PlayerCharacter.X
        self.PlayerCharacter.LastY = self.PlayerCharacter.Y
    else
        self.UDPConn:send("tick " .. self.Player.ID .. " \n")
    end

    repeat
        print("Waiting for response...")
        local data, msg = self.UDPConn:receive()

        if data then
            print("Received: " .. (data or "") .. " " .. (msg or ""))
            self:OnMessageReceived(data)
        end
    until not data
end

function World:Draw()
    self.Background:Draw()
    self.PlayerCharacter:Draw()

    for _, otherCharacter in pairs(self.OtherCharacters) do
        otherCharacter:Draw()
    end
end

function World:mousepressed(x, y, button, istouch, presses)
    -- right now, nothing is clickable
end

function World:OnMessageReceived(messageString)
    local msg = SplitString(messageString, " ")

    if msg[1] == "move" then
        local id = tonumber(msg[2])
        local x = tonumber(msg[3])
        local y = tonumber(msg[4])

        print("move id: " .. id .. " x: " .. x .. " y: " .. y)

        if id == self.Player.ID then
            if math.abs(self.PlayerCharacter.X - x) > 100 or math.abs(self.PlayerCharacter.Y - y) > 100 then
                -- TODO - if in debug mode, show a shadow where the server says the player should be
            end
        else
            local character = self.OtherCharacters[id]
            if character == nil then
                character = Character:new(nil, x, y)
                self.OtherCharacters[id] = character
            else
                character.X = x
                character.Y = y
            end
        end
    elseif msg[1] == "join" then
        local id = tonumber(msg[2])
        local name = msg[3]

        print("join id: " .. (id or "") .. " name: " .. (name or ""))

        if id == self.Player.ID then
            self.PlayerCharacter.Name = name
            print("Player joined as " .. name)
        else
            local character = self.OtherCharacters[id]
            if character == nil then
                character = Character:new(nil, 0, 0)
                self.OtherCharacters[id] = character
            end

            character.Name = name
        end
    elseif msg[1] == "tock" then
        print("Tock")
    else
        print("Unknown message: " .. messageString)
    end
end