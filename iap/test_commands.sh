#!bash

# --- CONFIGURATION ---
# Replace this with your actual serial port (e.g., /dev/ttyUSB0, /dev/ttyACM0, /dev/tty.usbserial-xxx)
PORT="${1:-/dev/ttyUSB0}"
BAUD="19200"

# --- SYSTEM CHECK ---
if [ ! -e "$PORT" ]; then
    echo -e "\033[0;31m[ERROR] Serial port '$PORT' not found.\033[0m"
    echo "Usage: $0 [/dev/your-serial-port]"
    exit 1
fi

# Configure serial port parameters using stty
# 19200 baud, 8 data bits, no parity, 1 stop bit, disable echo/canonical processing
stty -F "$PORT" $BAUD cs8 -cstopb -parenb -echo -icanon min 1 time 0

echo "============================================="
echo " Starting Infiniti-Pod Protocol Test Suite   "
echo " Target Port: $PORT @ $BAUD baud             "
echo "============================================="
echo ""

# Helper function to send hex payloads and print status
send_test() {
    local action="$1"
    local hex_payload="$2"
    local expected="$3"
    local resulting_action="$4"

    echo -e "\033[1;34m[TEST] Sending Command:\033[0m $action"
    echo "  Payload (Hex):  $hex_payload"
    echo "  Expected Byte:  $expected"
    echo "  Expected Logic: $resulting_action"

    # Convert hex string into raw binary bytes and blast down the pipeline
    echo -ne "$hex_payload" | xxd -r -p > "$PORT"

    # Read raw bytes into a temp file
    dd if="$PORT" bs=1 count=32 of=resp.bin 2>/dev/null

    # Convert to hex
    actual_hex=$(xxd -p resp.bin | tr -d '\n')
    echo "Actual Response (hex): $actual_hex"

    if [ "$actual_hex" != "$expected" ]; then
        echo -e "\033[0;31m  [✘] Expected '$expected', got '$actual_hex'.\033[0m"
    else
        echo -e "\033[0;32m  [✔] Expected '$expected', got '$actual_hex'.\033[0m"
    fi
    echo "---------------------------------------------"

    # Give the MCU time to process the transmission before the next payload drops
    sleep 1.5
}


send_test "SendLingoSupport" \
          "FF FF 55 03 00 01 04 F8" \
          "ff550400020001f9" \
          "Sends Ack"

send_test "RequestVersion" \
          "FF 55 03 04 00 12 E7" \
          "ff55050400130114cf" \
          "Starts Handshake"

send_test "RequestiPodName" \
          "FF 55 03 04 00 14 E5" \
          "ff55120400154d69636861656c27732050686f6e656e" \
          "Returns 'Michael's Phone'"

send_test "GetPlayStatus" \
          "FF 55 03 04 00 1C DD" \
          "ff550c04001d000249f00000753002f1" \
          "Returns Play Status"

# # 1. Next Track
# send_test "Next Track" \
#           "FF55AA0302000001FA" \
#           "0x01" \
#           "Triggers ble.KeyNext"

# # 2. Previous Track
# send_test "Previous Track" \
#           "FF55AA0302000008F3" \
#           "0x08" \
#           "Triggers ble.KeyPrevious"

# # 3. Play / Pause
# send_test "Play / Pause" \
#           "FF55AA0302000002F9" \
#           "0x02" \
#           "Triggers ble.KeyPlayPause"

# # 4. Release Button
# send_test "Release Button" \
#           "FF55AA0302000000FB" \
#           "0x00" \
#           "Safely Ignored"

echo -e "\033[1;32mAll test sequences sent successfully!\033[0m"
