require("lobby.lobby")
require("lobby.player")

function love.load()
    love.graphics.setFont(love.graphics.newFont("assets/fonts/UbuntuMono-Regular.ttf", 42))
    love.graphics.setBackgroundColor(Color2)
    love.window.setMode(1600, 900, {
        fullscreen = false,
        resizable = false,
        borderless = false
    })

    LocalPlayer = Player:new(nil, "willcliff")
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
    if LocalPlayer:InLobby() then
        WorldLobby:mousepressed(x, y, button, istouch, presses)
    end

    -- gameserver stuff here
end

function love.resize(w, h)
    -- todo - fix this. maybe have three sizes like league of legends launcher
end
