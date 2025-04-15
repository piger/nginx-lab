local _M = {}

function _M.run()
    if (ngx.var.http_x_forwarded_for == "" or ngx.var.http_x_forwarded_for == nil) then
        ngx.var.realip_add_x_forwarded_for = ngx.var.realip_remote_addr
    else
        ngx.var.realip_add_x_forwarded_for = ngx.var.http_x_forwarded_for .. ", " .. ngx.var.realip_remote_addr
    end
end

return _M
