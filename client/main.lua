require("lobby.lobby")
require("lobby.player")
require("world.world")
require("common.constants")

function love.load()
    love.window.setTitle("Keydream")

    -- todo - make window resizable
    love.window.setMode(WindowSizeX, WindowSizeY, {
        fullscreen = false,
        resizable  = false,
        borderless = false
    })

    -- todo - add text box for names
    LocalPlayer = Player:new(nil)
    LocalLobby = Lobby:new(nil, LocalPlayer)
    LocalWorld = World:new(nil, LocalPlayer, "127.0.0.1", 8081)

    NumWorldConnectAttempts = 0

    love.keyboard.keysPressed = {}
end

function love.update(dt)
    if LocalPlayer:InWorld() then
        LocalWorld:Update(dt)
    elseif LocalPlayer:ConnectingToWorld() then
        NumWorldConnectAttempts = NumWorldConnectAttempts + 1
        print("Connecting to world: " .. NumWorldConnectAttempts)
        if NumWorldConnectAttempts > MaximumWorldConnectAttempts then
            LocalPlayer:SetState(PlayerState.LOBBY_CONNECTED)
            NumWorldConnectAttempts = 0
        elseif LocalWorld:Connect() then
            LocalPlayer:SetState(PlayerState.WORLD_CONNECTED)
        end
    elseif LocalPlayer:InLobby() then
        -- TODO - add a button to refresh the list of worlds, or refresh on a timer
    end
end

function love.draw()
    if LocalPlayer:InLobby() then
        LocalLobby:Draw()
    elseif LocalPlayer:ConnectingToWorld() then
        love.graphics.setColor(Color1)
        love.graphics.setFont(MediumFont)
        love.graphics.print("Connecting to world...", 0, 0)
    elseif LocalPlayer:InWorld() then
        LocalWorld:Draw()
    end
end

function love.mousepressed(x, y, button, istouch, presses)
    if LocalPlayer:InLobby() then
        LocalLobby:mousepressed(x, y, button, istouch, presses)
    elseif LocalPlayer:InWorld() then
        -- TODO - nothing in the world is clickable yet
        -- LocalWorld:mousepressed(x, y, button, istouch, presses)
    end
end

function love.keypressed(key)
    love.keyboard.keysPressed[key] = true
    if LocalPlayer:InLobby() then
        if key == "return" then
            LocalPlayer:SetName(LocalLobby.NameInput.Text)
            LocalLobby:Connect()
        else
            LocalLobby:keypressed(key)
        end
    end
end

function love.keyreleased(key)
    love.keyboard.keysPressed[key] = false
end
