require("common.animatable")
require("common.utils")

CharacterState = {
    IDLE = 1,
    WALK = 2,
}

WalkingDirections = {
    UP = 1,
    DOWN = 2,
    LEFT = 3,
    RIGHT = 4,
}

Character = {
    Animations = {
        [WalkingDirections.DOWN] = {
            [CharacterState.IDLE] = {},
            [CharacterState.WALK] = {},
        },
        [WalkingDirections.UP] = {
            [CharacterState.IDLE] = {},
            [CharacterState.WALK] = {},
        },
        [WalkingDirections.LEFT] = {
            [CharacterState.IDLE] = {},
            [CharacterState.WALK] = {},
        },
        [WalkingDirections.RIGHT] = {
            [CharacterState.IDLE] = {},
            [CharacterState.WALK] = {},
        }
    },

    Direction = WalkingDirections.DOWN,
    State = CharacterState.IDLE,

    LastX = 0,
    LastY = 0,

    X = 0,
    Y = 0,

    Speed = TileSize / 2,
}

function Character:new(o, x, y)
    o = o or {}
    setmetatable(o, self)
    self.__index = self

    o.X = x
    o.Y = y

    for wdStr, wdInt in pairs(WalkingDirections) do
        for csStr, csInt in pairs(CharacterState) do
            local t = {}
            for i = 1, 4 do
                local spriteImage = love.graphics.newImage("assets/sprites/rgs_dev/Character without weapon/"..string.lower(csStr).."/"..string.lower(csStr).." "..string.lower(wdStr)..i..".png")
                spriteImage:setFilter("nearest", "linear")
                t[i] = spriteImage
            end
            o.Animations[wdInt][csInt] = Animatable:new(nil, t)
        end
    end

    return o
end

function Character:Update(dt)
    if love.keyboard.keysPressed["up"] and self.Y - self.Speed > 0 then
        self.Y = self.Y - self.Speed
        self.Direction = WalkingDirections.UP
        self.State = CharacterState.WALK
    elseif love.keyboard.keysPressed["down"] and self.Y + self.Speed < 13 * TileSizeScaled then
        self.Y = self.Y + self.Speed
        self.Direction = WalkingDirections.DOWN
        self.State = CharacterState.WALK
    elseif love.keyboard.keysPressed["left"] and self.X - self.Speed > 0 then
        self.X = self.X - self.Speed
        self.Direction = WalkingDirections.LEFT
        self.State = CharacterState.WALK
    elseif love.keyboard.keysPressed["right"] and self.X + self.Speed < 24 * TileSizeScaled then
        self.X = self.X + self.Speed
        self.Direction = WalkingDirections.RIGHT
        self.State = CharacterState.WALK
    else
        self.State = CharacterState.IDLE
    end

    self.Animations[self.Direction][self.State]:Update(dt)
end

function Character:Draw()
    self.Animations[self.Direction][self.State]:Draw(self.X, self.Y, CharacterScale, CharacterScale)
end

function Character:HasMoved()
    return self.LastX ~= self.X or self.LastY ~= self.Y
end

