WorldView = {
    World = nil,
    X = 0,
    Y = 0,
    Width = 0,
    Height = 0,
    Text = "",
    TextColor = Color5,
    Color = Color3
}

function WorldView:new(o, world)
    o = o or {}
    setmetatable(o, self)
    self.__index = self
    o.World = world
    return o
end

function WorldView:Draw(y)
    love.graphics.setColor(Color3)
    love.graphics.rectangle('fill', love.graphics.getWidth() / 4, y, 4 * DefaultButtonWidth, DefaultButtonHeight)

    love.graphics.setColor(Color2)
    love.graphics.printf(
        "world " .. self.World.id,
        love.graphics.getWidth() / 4,
        y,
        800)
    love.graphics.printf(
        self.World.num_players .. "/20 players",
        love.graphics.getWidth() / 2,
        y,
        800)
end

function WorldView:IsButtonPressed(x, y)
    if x > self.X and x < self.X + self.Width and y > self.Y and y < self.Y + self.Height then
        return true
    end
    return false
end
