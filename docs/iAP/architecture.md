iAP SEND FLOW
--------------

Application / Logic
        |
        v
+---------------------------+
|   Build Command Payload   |
|  (lingoID, commandID,     |
|   payload bytes)          |
+-------------+-------------+
              |
              v
+---------------------------+
|   Protocol Layer          |
|   SendPacket()            |
|                           |
| - Insert start byte       |
| - Insert length           |
| - Insert payload          |
| - Compute checksum        |
| - Build full packet       |
+-------------+-------------+
              |
              v
+---------------------------+
|   Transport Layer         |
|   UART Write()            |
+---------------------------+


iAP RECEIVE FLOW
-----------------

+---------------------------+
|   Transport Layer         |
|   UART ReadByte()         |
+-------------+-------------+
              |
              v
+---------------------------+
|   Protocol Layer          |
|   ParsePacket()           |
|                           |
| - Wait for sync byte      |
| - Wait for start byte     |
| - Read length             |
| - Read payload            |
| - Verify checksum         |
|                           |
| Outputs: Packet{          |
|   LingoID,                |
|   CommandID,              |
|   Payload[]               |
| }                         |
+-------------+-------------+
              |
              v
+---------------------------+
|   Lingo Layer             |
|   (General / Extended)    |
|                           |
| - HandleGeneral()         |
| - HandleExtended()        |
+-------------+-------------+
              |
              v
Application / Logic
