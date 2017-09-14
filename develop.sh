#!usr/bin/bash
echo "starting the develop script"
( cd client && ng serve --proxy-config proxy.conf.json ) & ( cd api && CompileDaemon -build="go install" -command="server") &
