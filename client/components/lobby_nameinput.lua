require("common.constants")

NameInput = {
    Text = "",
}

function NameInput:new(o)
    o = o or {}
    setmetatable(o, self)
    self.__index = self

    o.Text = ""

    o.X = love.graphics.getWidth() / 4
    o.Y = love.graphics.getHeight() / 2

    o.Height = DefaultButtonHeight
    o.Width = love.graphics.getWidth() / 2

    return o
end

function NameInput:Draw()
    love.graphics.setColor(Color3)
    love.graphics.rectangle(
        'fill',
        self.X, self.Y,
        love.graphics.getWidth() / 2, DefaultButtonHeight)

    love.graphics.setColor(Color5)
    love.graphics.setFont(LargeFont)
    love.graphics.print(
        self.Text,
        self.X + DefaultButtonWidth / 4,
        self.Y + DefaultButtonHeight / 4)
end

function NameInput:keypressed(key)
    if key == 'backspace' then
        self.Text = self.Text:sub(1, -2)
    elseif key == 'escape' then
        self.Text = ""
    elseif #self.Text < 12 then
        self.Text = self.Text .. key
    end
end
