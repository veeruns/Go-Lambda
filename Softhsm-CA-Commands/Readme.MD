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

* Create root Certificate with about key, the directory the root CA cert is stored is readable by root only. Version of openssl in my raspberry pi has a bug, hence had to compile. note the key is slotid:index number. I think you can give slot_name:label_label_of_token also, have to check it output
```
sudo LD_LIBRARY_PATH=/usr/local/lib/ OPENSSL_CONF=openssl.cnf /usr/local/bin/openssl req -new -x509 -days 7300 -sha256 -engine pkcs11 -keyform engine -key 1810989539:10 -out CACERT.pem

```
