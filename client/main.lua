require("lobby.lobby")
require("lobby.player")
require("world.world")

function love.load()
    love.window.setTitle("Keydream")
    love.graphics.setFont(love.graphics.newFont("assets/fonts/UbuntuMono-Regular.ttf", 42))

    -- todo - make window resizable
    love.window.setMode(1600, 900, {
        fullscreen = false,
        resizable  = false,
        borderless = false
    })

    -- todo - add text box for names
    LocalPlayer = Player:new(nil, "willcliff")
    LocalLobby = Lobby:new(nil, LocalPlayer)
    LocalWorld = World:new(nil, LocalPlayer, "127.0.0.1", 8081)

    love.keyboard.keysPressed = {}
end

function love.update(dt)
    if LocalPlayer:InGame() then
        LocalWorld:Update(dt)
    end
end

function love.draw()
    if LocalPlayer:InLobby() then
        LocalLobby:Draw()
    elseif LocalPlayer:ConnectingToGame() then
        love.graphics.print("Connecting to game...", 0, 0)
        if LocalWorld:Connect() then
            LocalPlayer:SetState(PlayerState.GAME_CONNECTED)
        end
    elseif LocalPlayer:InGame() then
        LocalWorld:Draw()
    end
end

function love.mousepressed(x, y, button, istouch, presses)
    if LocalPlayer:InLobby() then
        LocalLobby:mousepressed(x, y, button, istouch, presses)
    elseif LocalPlayer:InGame() then
        LocalWorld:mousepressed(x, y, button, istouch, presses)
    end
end

function love.keypressed(key)
    love.keyboard.keysPressed[key] = true
end

function love.keyreleased(key)
    love.keyboard.keysPressed[key] = false
end
