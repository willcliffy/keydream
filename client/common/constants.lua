LOCAL = false

HugeFont   = love.graphics.newFont("assets/fonts/lunchds.ttf", 64)
LargeFont  = love.graphics.newFont("assets/fonts/lunchds.ttf", 48)
BigFont    = love.graphics.newFont("assets/fonts/lunchds.ttf", 32)
MediumFont = love.graphics.newFont("assets/fonts/lunchds.ttf", 24)
SmallFont  = love.graphics.newFont("assets/fonts/lunchds.ttf", 18)

MaximumWorldConnectAttempts = 5

Color1 = {190 / 255, 140 / 255, 47 / 255}
Color2 = {84 / 255, 83 / 255, 108 / 255}
Color3 = {170 / 255, 149 / 255, 119 / 255}
Color4 = {171 / 255, 180 / 255, 201 / 255}
Color5 = {211 / 255, 220 / 255, 232 / 255}

DefaultButtonWidth = 250
DefaultButtonHeight = 75

LocalLobbyURL = "http://localhost:8080"
LocalWorldURL = "http://localhost:8081"

BackgroundScale = 4
CharacterScale = 4.5

TileSize = 16
TileSizeScaled = TileSize * BackgroundScale

WindowSizeX = 25 * TileSizeScaled
WindowSizeY = 14 * TileSizeScaled

CharacterMaxSpeed = math.floor(TileSize * 2 / 3)
RemoteCharacterMaxSpeed = 0.97 * CharacterMaxSpeed

-- In pixels, how far the remote character has to be from target location to stop moving
RemoteCharacterAlpha = 1.05 * RemoteCharacterMaxSpeed

TickDuration = 0.2
