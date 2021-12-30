Animatable = {
    Textures = {},
    Frame = nil,
    CurrentTexture = 1,
    -- todo - 20 fps for now
    FrameDuration = 0.25,
}

function Animatable:new(o, textures)
    o = o or {}
    setmetatable(o, self)
    self.__index = self

    o.Textures = textures or {}
    o.Frame = love.graphics.newQuad(24, 24, 16, 16, textures[1]:getWidth(), textures[1]:getHeight())

    return o
end

function Animatable:Update(dt)
    self.CurrentTexture = self.CurrentTexture + (dt / self.FrameDuration)
    if self.CurrentTexture > #self.Textures then
        self.CurrentTexture = 1
    end
end

function Animatable:Draw(x, y, sx, sy)
    love.graphics.draw(self.Textures[math.floor(self.CurrentTexture)], self.Frame, x, y, 0, sx, sy)
end
