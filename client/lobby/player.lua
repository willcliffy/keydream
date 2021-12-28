PlayerState = {
    LOBBY_DISCONNECTED = 1,
    LOBBY_CONNECTED = 2,
    GAME_CONNECTING = 3,
    GAME_CONNECTED = 4
}

Player = {
    State = PlayerState.LOBBY_DISCONNECTED
}

function Player:new(o)
    o = o or {}
    setmetatable(o, self)
    self.__index = self
    return o
end

function Player:SetState(state)
    if state == PlayerState.LOBBY_DISCONNECTED then
        self.State = PlayerState.LOBBY_DISCONNECTED
    elseif state == PlayerState.LOBBY_CONNECTED then
        self.State = PlayerState.LOBBY_CONNECTED
    elseif state == PlayerState.GAME_CONNECTING then
        self.State = PlayerState.GAME_CONNECTING
    elseif state == PlayerState.GAME_CONNECTED then
        self.State = PlayerState.GAME_CONNECTED
    else
        print("Invalid player state: " .. state)
    end
end

function Player:InLobby()
    return self.State == PlayerState.LOBBY_DISCONNECTED or self.State == PlayerState.LOBBY_CONNECTED
end
