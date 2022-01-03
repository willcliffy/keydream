require("common.constants")
require("common.utils")

RemoteCharacter = Character:new()
RemoteCharacter_mt = { __index = RemoteCharacter }

function RemoteCharacter:new(o, name)
    o = o or {}
    setmetatable(o, RemoteCharacter_mt)

    o.Name = name
    print("RemoteCharacter:new: " .. name)

    o.X = 3 * TileSize
    o.Y = 3 * TileSize

    o.SpeedX = 0
    o.SpeedY = 0

    o.CurrentTarget = 1
    o.TargetLocations = {}

    return o
end

function RemoteCharacter:Update(dt)
    -- TODO - dont be a hacky math major. do this right someday

    if not self.TargetLocations[1] then
        self.State = CharacterState.IDLE
        self.Animations[self.Direction][self.State]:Update(dt)
        return
    end

    local target = self.TargetLocations[1]

    -- if we're close to the target X, just jump there and stop moving in x direction
    if math.abs(self.X - target.X) <= RemoteCharacterAlpha then
        self.X = target.X
        self.SpeedX = 0
    else
        if target.X - self.X < -RemoteCharacterAlpha then
            self.SpeedX = -RemoteCharacterMaxSpeed
        elseif target.X - self.X > RemoteCharacterAlpha then
            self.SpeedX = RemoteCharacterMaxSpeed
        else
            print("legitimately not possible: x")
        end
    end

    -- if we're close to the target Y, just jump there and stop moving in y direction
    if math.abs(self.Y - target.Y) <= 2 * RemoteCharacterAlpha then
        self.Y = target.Y
        self.SpeedY = 0
    else
        if target.Y - self.Y < -RemoteCharacterAlpha then
            self.SpeedY = -RemoteCharacterMaxSpeed
        elseif target.Y - self.Y > RemoteCharacterAlpha then
            self.SpeedY = RemoteCharacterMaxSpeed
        else
            print("legitimately not possible: y")
        end
    end

    -- if we're close to the target X and Y, remove the target from the list
    if self.X == target.X and self.Y == target.Y then
        table.remove(self.TargetLocations, 1)

        self.State = CharacterState.IDLE
        self.SpeedX = 0
        self.SpeedY = 0
        self.Animations[self.Direction][self.State]:Update(dt)
        return
    end

    self.X = self.X + self.SpeedX
    self.Y = self.Y + self.SpeedY

    if self.SpeedX == 0 and self.SpeedY == 0 then
        self.State = CharacterState.IDLE
    else
        self.State = CharacterState.WALK
        if math.abs(self.SpeedX) > math.abs(self.SpeedY) then
            if self.X < target.X then
                self.Direction = WalkingDirections.RIGHT
            elseif self.X > target.X then
                self.Direction = WalkingDirections.LEFT
            end
        else
            if self.Y < target.Y then
                self.Direction = WalkingDirections.DOWN
            elseif self.Y > target.Y then
                self.Direction = WalkingDirections.UP
            end
        end
    end

    self.Animations[self.Direction][self.State]:Update(dt)
end

function RemoteCharacter:Draw()
    love.graphics.setFont(MediumFont)
    love.graphics.print(self.Name, self.X - TileSizeScaled / 2, self.Y - TileSizeScaled / 2)
    self.Animations[self.Direction][self.State]:Draw(self.X, self.Y, CharacterScale, CharacterScale)
end

function RemoteCharacter:OnMove(x, y)
    table.insert(self.TargetLocations, {
        X = x,
        Y = y,
    })
end

function RemoteCharacter:OnJoin()
    print("RemoteCharacter:OnJoin")
end
