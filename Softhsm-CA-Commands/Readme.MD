Commands to create a RootCA
Install Softhsm2
Openssl.

  ```
    softhsm2-util --show-slots
  ```
   Shows the slot number, use one for RootCA
* Create a Slot with Commands
  ```
   softhsm2-util --init-token --slot 1810989539 --label "RootCA"

  ```
