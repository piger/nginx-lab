-- log the total size of the response's headers.
-- https://github.com/openresty/lua-nginx-module#ngxrespget_headers

local _M = {}

function _M.run()
    local h, err = ngx.resp.get_headers()
    if err == "truncated" then
        ngx.log(ngx.ERR, "truncated headers from lua")
        return
    end

    local size=0
    for k, v in pairs(h) do
        size = size + string.len(v)
    end
    ngx.var.headers_size = size
end

return _M
