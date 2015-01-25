#! /bin/sh
# /etc/init.d/goncierge.sh

# Goncierge start/stop startup script
#     1. Copy this file to /etc/init.d on Raspberry Pi
#     2. Make script executable:
#            sudo chmod 755 /etc/init.d/goncierge.sh
#     3. Test starting and stopping script:
#            sudo /etc/init.d/goncierge.sh start
#            sudo /etc/init.d/goncierge.sh stop
#     4. Register script to run at startup:
#            sudo update-rc.d goncierge.sh defaults

### BEGIN INIT INFO
# Provides:          goncierge.sh
# Required-Start:    $remote_fs $syslog
# Required-Stop:     $remote_fs $syslog
# Default-Start:     2 3 4 5
# Default-Stop:      0 1 6
# Short-Description: Simple script to start goncierge at boot
# Description:       A simple script that will start/stop goncierge at boot/shutdown.
### END INIT INFO

case "$1" in
  start)
    echo "Starting Goncierge"
    /home/pi/goncierge
    ;;
  stop)
    echo "Stopping Goncierge"
    killall goncierge 
    ;;
  *)
    echo "Usage: /etc/init.d/goncierge.sh {start|stop}"
    exit 1
    ;;
esac

exit 0
