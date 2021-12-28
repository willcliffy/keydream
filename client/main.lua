ConnectHandler = require("connect.connect")

function love.load()
    Button1X = love.graphics.getWidth() / 2
    Button1Y = love.graphics.getHeight() / 2

    Button2X = Button1X
    Button2Y = Button1Y + 100

    ButtonWidth = 50
    ButtonHeight = 50

    Button1State = false
    Button2State = false
end

function love.update(dt)

end

function love.draw()
    if Res and Res['code'] == 200 then
        love.graphics.setColor(255, 255, 255)
        love.graphics.printf("Connected!", 1, 1, 800)
    else
        if Res then
            love.graphics.setColor(255, 0, 0)
            love.graphics.printf("Not connected!\n" .. Res.status, 1, 1, 800)
        end
    end
    
    if Button1State == true then
        love.graphics.setColor(0, 255, 0)
    else
        love.graphics.setColor(255, 0, 0)
    end
    love.graphics.rectangle('fill', Button1X, Button1Y, ButtonWidth, ButtonHeight)

    if Button2State == true then
        love.graphics.setColor(0, 255, 0)
    else
        love.graphics.setColor(255, 0, 0)
    end
    love.graphics.rectangle('fill', Button2X, Button2Y, ButtonWidth, ButtonHeight)

    love.graphics.setColor(255, 255, 255)
    love.graphics.print("Button 1", Button1X, Button1Y)
    love.graphics.print("Button 2", Button2X, Button2Y)
end

function love.mousepressed(x, y, button, istouch, presses)
    if x > Button1X and x < Button1X + ButtonWidth and y > Button1Y and y < Button1Y + ButtonHeight then
        Button1State = not(Button1State)
        Res = ConnectHandler.Connect()
    end

    if x > Button2X and x < Button2X + ButtonWidth and y > Button2Y and y < Button2Y + ButtonHeight then
        Button2State = not(Button2State)
    end
end