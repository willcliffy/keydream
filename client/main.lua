require("lobby.lobby")
require("lobby.player")

function love.load()
    love.window.setTitle("Keydream")
    love.graphics.setFont(love.graphics.newFont("assets/fonts/UbuntuMono-Regular.ttf", 42))

    -- lets try to tile the lobby.
    love.graphics.setBackgroundColor(Color2)

    -- todo - make window resizable
    love.window.setMode(1600, 900, {
        fullscreen = false,
        resizable = false,
        borderless = false
    })

    -- todo - add text box for names
    LocalPlayer = Player:new(nil, "willcliff")
    LocalLobby = Lobby:new(nil, LocalPlayer)
end

function love.update(dt)

end

function love.draw()
    if LocalPlayer:InLobby() then
        LocalLobby:Draw()
    end
end

function love.mousepressed(x, y, button, istouch, presses)
    if LocalPlayer:InLobby() then
        LocalLobby:mousepressed(x, y, button, istouch, presses)
    end

    -- gameserver stuff here
end
