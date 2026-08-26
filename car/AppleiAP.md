# Infiniti-Pod Test Payloads

### ⏭️ Next Track
55 AA 03 02 00 00 01 FA

### ⏮️ Previous Track
55 AA 03 02 00 00 08 F3

### ⏯️ Play / Pause
55 AA 03 02 00 00 02 F9

### ⏹️ Release Button
55 AA 03 02 00 00 00 FB

### ❌ Bad Header (Error Test)
99 AA 03 02 00 00 01 FA

### ❌ Bad Checksum (Error Test)
55 AA 03 02 00 00 01 00

Action: Next Track
Hex Payload Buffer: 55 AA 03 02 00 00 01 FA
Expected Signal Byte: 0x01
Resulting Action: Triggers ble.KeyNext

Action: Previous Track
Hex Payload Buffer: 55 AA 03 02 00 00 08 F3
Expected Signal Byte: 0x08
Resulting Action: Triggers ble.KeyPrevious

Action: Play / Pause
Hex Payload Buffer: 55 AA 03 02 00 00 02 F9
Expected Signal Byte: 0x02
Resulting Action: Triggers ble.KeyPlayPause

Action: Release Button
Hex Payload Buffer: 55 AA 03 02 00 00 00 FB
Expected Signal Byte: 0x00
Resulting Action: Safely Ignored
