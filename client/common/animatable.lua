Animatable = {
    textures = {},
    frame = nil,
    currentTexture = 1,

    -- todo - hardcoded 5 fps animations for now
    frameDuration = 0.20,
}

function Animatable:new(o, textures, frame)
    o = o or {}
    setmetatable(o, self)
    self.__index = self

    o.textures = textures or {}
    o.frame = frame or love.graphics.newQuad(24, 24, 16, 16, textures[1]:getWidth(), textures[1]:getHeight())

    return o
end

function Animatable:Update(dt)
    -- TODO - theres a bug here where the more characters are facing the same direction, the faster the animation will be for all of them
    -- currentTexture appears to be acting global but only when characters face the same direction? I must be goofing some lua/love2d stuff here.
    self.currentTexture = self.currentTexture + (dt / self.frameDuration)
    if self.currentTexture > #self.textures then
        self.currentTexture = 1
    end
end

function Animatable:Draw(x, y, sx, sy)
    love.graphics.draw(self.textures[math.floor(self.currentTexture)], self.frame, x, y, 0, sx, sy)
end
