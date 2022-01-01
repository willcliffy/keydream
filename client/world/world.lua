local socket = require("socket")
local coroutine = require("coroutine")

require("world.background")
require("world.character")
require("world.reader")

World = {
    IP = "",
    Port = 0,

    Player = nil,

    PlayerCharacter = nil,
    OtherCharacters = {},

    Background = nil,

    UPDConn = nil,
    UDPReader = nil,
    UDPReaderThread = nil,
    Connected = false,

    TimeSinceLastTick = 0,
    TickDuration = 0.5,
}

function World:new(o, player, ip, port)
    o = o or {}
    setmetatable(o, self)
    self.__index = self

    self.IP = ip
    self.Port = port
    self.UDPConn = socket.udp()

    self.Background = Background:new()
    self.PlayerCharacter = Character:new(nil, 0, 0)

    self.Player = player

    Background:Update()

    self.UDPReader = UDPReader:new(nil, self.UDPConn)
    self.UDPReaderThread = coroutine.create(self.UDPReader:Start(self.OnMessageReceived))

    return o
end

function World:Connect()
    -- todo: this connect needs to be a coroutine so that
    if self.Connected then
        return true
    end

    local error
    self.UDPConn:settimeout(3)
    _, error = self.UDPConn:setpeername(self.IP, self.Port)
    if error then
        print("Failed to connect to game server: " .. error)
        return false
    end

    _, error, _ = self.UDPConn:send("join " .. LocalPlayer.Name .. "\n")
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

    print("Connected to game server: " .. data)
    self.Player.ID = tonumber(data)
    self.Connected = true
    return true
end

function World:Update(dt)
    if self.Connected then
        self.Background:Update()
        self.PlayerCharacter:Update(dt)
        self.TimeSinceLastTick = self.TimeSinceLastTick + dt

        if self.TimeSinceLastTick >= self.TickDuration then
            self.TimeSinceLastTick = self.TimeSinceLastTick - self.TickDuration
            self:Tick()
        end
    end
end

function World:Tick()
    if self.Connected then
        self.UDPConn:send("tick \n")

        if self.PlayerCharacter:HasMoved() then
            self.UDPConn:send("move "..self.Player.ID .." "..self.PlayerCharacter.X.." "..self.PlayerCharacter.Y.." \n")
        end
    end
end

function World:Draw()
    self.Background:Draw()
    self.PlayerCharacter:Draw()
end

function World:mousepressed(x, y, button, istouch, presses)
    -- right now, nothing is clickable
end

function World:OnMessageReceived(msg)

end