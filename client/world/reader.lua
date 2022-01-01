UDPReader = {
    UDPConn = nil,
}

function UDPReader:new(o, udpConn)
    o = o or {}
    setmetatable(o, self)
    self.__index = self

    o.UDPConn = udpConn

    return o
end

function UDPReader:Start(onMessageReceived)
    return function()
        local s, status, _ = self.UDPConn:receive()
        if status == "closed" then
            print("connection closed")
            return
        end

        print(s)

        if s then
            onMessageReceived(s)
        end
    end
end
