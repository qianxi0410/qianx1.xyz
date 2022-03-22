docker run -p 80:80 -p 443:443 \
    -itd --name=blog \
    -v root/qianxi.me/ssl:/etc/nginx/ssl \
    -v root/qianxi.me/default.conf:/etc/nginx/nginx.conf \
    -v root/qianxi.me/dist:/usr/share/qianxi.me \
    blog