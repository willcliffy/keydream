require("common.constants")

Button = {
    X = 0,
    Y = 0,
    Width = 0,
    Height = 0,
    Text = "",
    TextColor = Color5,
    Color = Color3
}

function Button:new(o, x, y, w, h, text, color, textColor)
    o = o or {}
    setmetatable(o, self)
    self.__index = self
    o.X = x
    o.Y = y
    o.Width = w
    o.Height = h
    o.Text = text
    o.Color = color or {1, 1, 1}
    o.TextColor = textColor or {0, 0, 0}
    return o
end

function Button:newConnectButton(o, x, y)
    return Button:new(o, x, y, 250, 75, "Connect", Color3, Color5)
end

function Button:newBackButton(o, x, y)
    return Button:new(o, x, y, 250, 75, "  Back  ", Color3, Color5)
end

function Button:Draw()
    love.graphics.setColor(self.Color)
    love.graphics.rectangle('fill', self.X, self.Y, self.Width, self.Height)
    love.graphics.setColor(self.TextColor)
    -- todo - center text in button
    love.graphics.print(self.Text, self.X + 50, self.Y + 13)
end

function Button:IsButtonPressed(x, y)
    if x > self.X and x < self.X + self.Width and y > self.Y and y < self.Y + self.Height then
        return true
    end
    return false
end
