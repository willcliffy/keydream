
local connectHandler = {}
local http = require("socket.http")

local backendURL = "http://localhost:8080"

function connectHandler.Connect()
    local http_request = require "http.request"
    local headers, stream = assert(http_request.new_from_uri(backendURL .. "/api/v1/connect"):go())
    local body = assert(stream:get_body_as_string())
    return {
        code = headers:get ":status",
        body = body
    }
end

return connectHandler
