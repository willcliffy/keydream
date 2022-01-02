PlayerState = {
    LOBBY_DISCONNECTED = 1,
    LOBBY_CONNECTED    = 2,
    WORLD_CONNECTING    = 3,
    WORLD_CONNECTED     = 4
}

Player = {
    Name = "",
    State = PlayerState.LOBBY_DISCONNECTED,
    ID = 0,
}

function Player:new(o, name)
    o = o or {}
    setmetatable(o, self)
    self.__index = self
    o.Name = name
    return o
end

function Player:SetState(state)
    if state == PlayerState.LOBBY_DISCONNECTED then
        self.State = PlayerState.LOBBY_DISCONNECTED
    elseif state == PlayerState.LOBBY_CONNECTED then
        self.State = PlayerState.LOBBY_CONNECTED
    elseif state == PlayerState.WORLD_CONNECTING then
        self.State = PlayerState.WORLD_CONNECTING
    elseif state == PlayerState.WORLD_CONNECTED then
        self.State = PlayerState.WORLD_CONNECTED
    else
        print("Invalid player state: " .. state)
    end
end

function Player:InLobby()
    return self.State == PlayerState.LOBBY_DISCONNECTED or self.State == PlayerState.LOBBY_CONNECTED
end

function Player:ConnectingToWorld()
    return self.State == PlayerState.WORLD_CONNECTING
end

function Player:InWorld()
    return self.State == PlayerState.WORLD_CONNECTED
end
