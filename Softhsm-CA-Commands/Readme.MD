Commands to create a CA. These are my experiments. Please point out defects

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
* Create rootCA keypair in softhsm2
```
 pkcs11-tool --module /usr/local/lib/softhsm/libsofthsm2.so --keypairgen --key-type rsa:4096  --id 10 --login --usage-sign --label "ROOT_CA_KEY" --slot 1810989539
```
* Create root CA key. Remember the pin. Key is now stored in softhsm2 (need openssl.cnf changes)
```
pkcs11-tool --module /usr/local/lib/softhsm/libsofthsm2.so --keypairgen --key-type rsa:4096  --id 10 --login --usage-sign --label "ROOT_CA_KEY" --slot 1810989539
```
