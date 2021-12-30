require("world.background")
require("world.character")

World = {
    PlayerCharacter = nil,
    OtherCharacters = {},

    Background = nil,
}

function World:new(o)
    o = o or {}
    setmetatable(o, self)
    self.__index = self

    self.Background = Background:new()
    self.PlayerCharacter = Character:new(nil, 0, 0)

    Background:Update()

    return o
end

function World:Update(dt)
    self.Background:Update()
    self.PlayerCharacter:Update(dt)
end

function World:Draw()
    self.Background:Draw()
    self.PlayerCharacter:Draw()
end

function World:mousepressed(x, y, button, istouch, presses)
    -- right now, nothing is clickable
end