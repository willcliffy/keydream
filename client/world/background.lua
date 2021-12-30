require("common.utils")

Background = {
    SpriteBatch = {},

    -- todo - movable camera
    MapX = 0,
    MapY = 0,

    TilesDisplayWidth = 25,
    TilesDisplayHeight = 14,

    TileQuads = {},

    -- todo:
    GrassQuads = {},
    StoneQuads = {},
    WallQuads = {},

    -- todo - handle mapbuilding
    map = TestMap(25, 14),
}

function Background:new(o)
    o = o or {}
    setmetatable(o, self)
    self.__index = self

    local tilesetImage = love.graphics.newImage("assets/environment/cainos/TX Tileset Grass.png")
    tilesetImage:setFilter("nearest", "linear")

    -- grass
    local i = 1
    for x = 3, 11 do
        self.TileQuads[i] = love.graphics.newQuad(x * TileSize, 1 * TileSize, TileSize, TileSize, tilesetImage:getWidth(), tilesetImage:getHeight())
        i = i + 1
    end

    self.TileQuads[i] = love.graphics.newQuad(0, 10 * TileSize, TileSize, TileSize, tilesetImage:getWidth(), tilesetImage:getHeight())

    self.SpriteBatch = love.graphics.newSpriteBatch(tilesetImage)

    return o
end

function Background:Update()
    self.SpriteBatch:clear()
    for x = 0, self.TilesDisplayWidth - 1 do
        for y = 0, self.TilesDisplayHeight - 1 do
            self.SpriteBatch:add(
                self.TileQuads[self.map[x + math.floor(self.MapX) + 1][y + math.floor(self.MapY) + 1]],
                x * TileSize,
                y * TileSize)
        end
    end

    self.SpriteBatch:flush()
end

function Background:Draw()
    love.graphics.draw(self.SpriteBatch, 0, 0, 0, BackgroundScale, BackgroundScale)
end

function Background:Move(x, y)
    -- todo - implement camera
end
