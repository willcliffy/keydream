WorldView = {
    World = nil,
    Width = 0,
    Height = 0,
    X = 0,
    Y = 0,
    Text = "",
    TextColor = Color5,
    Color = Color3
}

function WorldView:new(o, world)
    o = o or {}
    setmetatable(o, self)
    self.__index = self
    o.World = world
    o.Width = love.graphics.getWidth() / 2
    o.Height = DefaultButtonHeight
    o.X = love.graphics.getWidth() / 4
    return o
end

function WorldView:Draw(y)
    love.graphics.setFont(LargeFont)
    self.Y = y

    love.graphics.setColor(self.Color)
    love.graphics.rectangle('fill', self.X, self.Y, self.Width, self.Height)

    love.graphics.setColor(self.TextColor)
    love.graphics.printf(
        "world " .. self.World.id,
        self.X + DefaultButtonWidth / 4,
        self.Y + DefaultButtonHeight / 4,
        800)
    love.graphics.printf(
        self.World.num_players .. "/20 players",
        2 * self.X + DefaultButtonWidth / 4,
        self.Y + DefaultButtonHeight / 4,
        800)
end

function WorldView:IsButtonPressed(x, y)
    if x > self.X and x < self.X + self.Width and y > self.Y and y < self.Y + self.Height then
        return true
    end
    return false
end
