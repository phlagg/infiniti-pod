#!/bin/bash
set -e

echo "Flashing firmware..."
tinygo flash -target=nicenano -port=/dev/ttyACM0 .

echo "Firmware flashed!"

echo "Monitoring..."
tinygo monitor -port=/dev/ttyACM0 -baudrate 115200
