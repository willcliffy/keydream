WorldView = {
    X = 0,
    Y = 0,
    Width = 0,
    Height = 0,
    Text = "",
    TextColor = Color5,
    Color = Color3
}

function WorldView:new()
    o = o or {}
    setmetatable(o, self)
    self.__index = self
    
    return o
end

function WorldView:Draw()
    love.graphics.setColor(Color3)
    love.graphics.rectangle('fill', World1X, y, 4 * ButtonWidth, ButtonHeight)

    love.graphics.setColor(Color2)
    love.graphics.printf("world " .. v.id, WorldIDOffset + 10, y + 10, 800)
    love.graphics.printf(v.num_players .. "/20 players", WorldNumPlayersOffset + 10, y + 10, 800)
end

function WorldView:IsButtonPressed(x, y)
    if x > self.X and x < self.X + self.Width and y > self.Y and y < self.Y + self.Height then
        return true
    end
    return false
end
