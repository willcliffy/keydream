require("lobby.lobby")
require("lobby.player")
require("components.colors")

function love.load()
    love.window.setMode(1600, 900, {
        fullscreen = false,
        resizable = true,
        borderless = false
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

    local font = love.graphics.newFont("assets/fonts/UbuntuMono-Regular.ttf", 42)
    love.graphics.setFont(font)

    love.graphics.setBackgroundColor(Color2)

    LocalPlayer = Player:new(nil)
    WorldLobby = Lobby:new(nil, LocalPlayer)
end

function love.update(dt)

end

function love.draw()
    if LocalPlayer:InLobby() then
        WorldLobby:Draw()
    end
end

function love.mousepressed(x, y, button, istouch, presses)
    -- todo - add `IsButtonPressed` helper here. maybe make a button component
    if LocalPlayer:InLobby() then
        if WorldLobby.ConnectButton:IsButtonPressed(x, y) then
            WorldLobby:Connect("Player")
        end
    end
end

function love.resize(w, h)
    Button1X = w / 2 - ButtonWidth / 2
    Button1Y = h * 3 / 4 - ButtonHeight / 2
    -- todo - fix this. maybe have three sizes like league of legends
end
