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
    o.X = x or 0
    o.Y = y or 0
    o.Width = w or DefaultButtonWidth
    o.Height = h or DefaultButtonWidth
    o.Text = text
    o.Color = color or {1, 1, 1}
    o.TextColor = textColor or {0, 0, 0}
    return o
end

function Button:newConnectButton(o, x, y)
    return Button:new(o, x, y, DefaultButtonWidth, DefaultButtonHeight, "Connect", Color3, Color5)
end

function Button:newBackButton(o, x, y)
    return Button:new(o, x, y, DefaultButtonWidth, DefaultButtonHeight, " Back ", Color3, Color5)
end

function Button:Draw()
    love.graphics.setColor(self.Color)
    love.graphics.rectangle('fill', self.X, self.Y, self.Width, self.Height)
    love.graphics.setColor(self.TextColor)
    -- todo - center text in button
    love.graphics.setFont(BigFont)
    love.graphics.print(self.Text, self.X + 60, self.Y + 20)
end

function Button:IsButtonPressed(x, y)
    if x > self.X and x < self.X + self.Width and y > self.Y and y < self.Y + self.Height then
        return true
    end
    return false
end
