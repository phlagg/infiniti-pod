            ANDROID PHONE
        (BLE or USB connection)
                  |
                  v
      +------------------------+
      |   Phone Connectivity   |
      |   transport/ble/,      |
      | transport/ bluetooth/, |
      | transport/usb/         |
      +------------------------+
                  |
                  v
      +------------------------+
      |     iPod Emulation     |
      |   ipod/, playback/     |
      +------------------------+
                  |
                  v
      +------------------------+
      |   iAP Protocol Stack   |
      | packet → lingo → cmd   |
      +------------------------+
                  |
                  v
      +------------------------+
      |   Serial Transport     |
      |   /transport        |
      +------------------------+
                  |
                  v
        INFINITI HEAD UNIT
