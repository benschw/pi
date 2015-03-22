[![Build Status](https://drone.io/github.com/benschw/pi/status.png)](https://drone.io/github.com/benschw/pi/latest)



- [download latest](https://drone.io/github.com/benschw/pi/files/piled)
- [download latest (arm deb)](https://drone.io/github.com/benschw/pi/files/build/piled.deb)

### Examples

	wget https://drone.io/github.com/benschw/pi/files/build/piled.deb	
	sudo dpkg -i piled.deb
	sudo /etc/init.d/piled start

	curl -X POST http://192.168.0.115:8080/admin/status
	curl -X POST http://192.168.0.115:8080/blink/toggle
	curl -X POST http://192.168.0.115:8080/blink/count-down