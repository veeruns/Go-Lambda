Commands to create a RootCA
Install Softhsm2
Openssl.

* Create a Slot with Commands
  ```
   softhsm2-util --init-token --slot 1810989539 --label "RootCA"

  ```
