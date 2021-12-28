Lobby = require("transport.Lobby")

function love.load()
    love.window.setMode( 1600, 900, {
        fullscreen = false,
        resizable = true,
        borderless = false,
    })

    
    WorldIDOffset = love.graphics.getWidth() / 4
    WorldNumPlayersOffset = love.graphics.getWidth() / 2
    WorldRegionOffset = love.graphics.getWidth() * 3 / 4

    ButtonWidth = 250
    ButtonHeight = 75

    World1X = WorldIDOffset
    World1Y = 150

    Button1X = love.graphics.getWidth() / 2 - ButtonWidth / 2
    Button1Y = love.graphics.getHeight() * 3 / 4 - ButtonHeight / 2

    Color1 = {
        190 / 255,
        140 / 255,
         47 / 255,
    }

    Color2 = {
        84 / 255,
        83 / 255,
        108 / 255,
    }

    Color3 = {
        170 / 255,
        149 / 255,
        119 / 255,
    }

    Color4 = {
        171 / 255,
        180 / 255,
        201 / 255,
    }

    Color5 = {
        211 / 255,
        220 / 255,
        232 / 255,
    }

    local font = love.graphics.newFont("assets/fonts/UbuntuMono-Regular.ttf", 42)
    love.graphics.setFont(font)

    love.graphics.setBackgroundColor(Color2)

    NumWorlds = 0

    PlayerState = 0
end

function love.update(dt)

end

function love.draw()
    if ConnectLobbyRes then
        local y = World1Y

        love.graphics.setColor(Color1)
        love.graphics.printf("Connected!", 1, 1, 800)

        for k, v in pairs(ConnectLobbyRes.Worlds) do
            love.graphics.setColor(Color3)
            love.graphics.rectangle('fill', World1X, y, 4 * ButtonWidth, ButtonHeight)

            love.graphics.setColor(Color2)
            love.graphics.printf("world " .. v.id, WorldIDOffset + 10, y + 10, 800)
            love.graphics.printf(v.num_players .. "/20 players", WorldNumPlayersOffset + 10, y + 10, 800)
            -- love.graphics.printf(v.region, WorldRegionOffset, i, 800)

            NumWorlds = NumWorlds + 1
            y = y + 100
        end

        love.graphics.setColor(Color3)
        love.graphics.rectangle('fill', Button1X, Button1Y, ButtonWidth, ButtonHeight)
    
        love.graphics.setColor(Color5)
        love.graphics.print('Back', Button1X + 75, Button1Y + 13)
    else
        love.graphics.setColor(Color3)
        love.graphics.rectangle('fill', Button1X, Button1Y, ButtonWidth, ButtonHeight)
    
        love.graphics.setColor(Color5)
        love.graphics.print('Connect', Button1X + 45, Button1Y + 13)
    end
end

function love.mousepressed(x, y, button, istouch, presses)
    if x > Button1X and x < Button1X + ButtonWidth and y > Button1Y and y < Button1Y + ButtonHeight then
        Button1State = not(Button1State)
        local res = Lobby.Connect("will")
        if res then
            ConnectLobbyRes = res
        end
    end

    -- todo - make this work with multiple worlds
    if x > World1X and x < World1X + 4 * ButtonWidth and y > World1Y and y < World1Y + ButtonHeight then
        local res = Lobby.Join(ConnectLobbyRes.Worlds[1].id)
        if res then
            JoinWorldRes = res
        end
    end
end

function love.resize(w, h)
    Button1X = w / 2 - ButtonWidth / 2
    Button1Y = h * 3 / 4 - ButtonHeight / 2
    -- todo - fix this. maybe have three sizes like league of legends
end
