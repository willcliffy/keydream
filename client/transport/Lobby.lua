local lobbyHandler = {}

local cjson = require("cjson")


local backendURL = "http://localhost:8080"

function lobbyHandler.Connect(name)
    local http = require("socket.http")
    local b, c, h = http.request(backendURL .. "/api/v1/connect", "{\"name\":\"" .. name .. "\"}")

    local body = cjson.decode(b)

    if c == 200 then
        return cjson.decode(b)
    end

    print("Error connecting to backend: " .. c .. " " .. b)
    return nil
end

function lobbyHandler.Join(worldID)
    local http = require("socket.http")
    local b, c, h = http.request(backendURL .. "/api/v1/join", "{\"WorldID\":" .. worldID .. "}")

    print (b)

    if c == 200 then
        return cjson.decode(b)
    end

    print("Error connecting to backend: " .. c .. " " .. b)
    return nil
end

return lobbyHandler