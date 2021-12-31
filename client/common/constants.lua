LOCAL = true

Color1 = {190 / 255, 140 / 255, 47 / 255}
Color2 = {84 / 255, 83 / 255, 108 / 255}
Color3 = {170 / 255, 149 / 255, 119 / 255}
Color4 = {171 / 255, 180 / 255, 201 / 255}
Color5 = {211 / 255, 220 / 255, 232 / 255}

DefaultButtonWidth = 250
DefaultButtonHeight = 50

LocalLobbyURL = "http://localhost:8080"
LocalWorldURL = "http://localhost:8081"

TestMap = function(x, y)
    local map = {}
    for i = 1, x do
        map[i] = {}
        for j = 1, y do
            map[i][j] = math.random(10)
        end
    end
    return map
end

BackgroundScale = 4
CharacterScale = 4.5

TileSize = 16
TileSizeScaled = TileSize * BackgroundScale
