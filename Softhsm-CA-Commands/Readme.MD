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
* Create Slot for Intermediate CA keys.

```
softhsm2-util --init-token --slot 318488403 --label "IntermediateCA"

```

* Create intermediate key

```

pkcs11-tool --module /usr/local/lib/softhsm/libsofthsm2.so --keypairgen --key-type rsa:4096  --id 10 --login --usage-sign --label "INTERMEDIATE_CA_KEY" --slot 318488403

```

* Create Intermediate CA CSR
```
sudo LD_LIBRARY_PATH=/usr/local/lib/ OPENSSL_CONF=openssl.cnf /usr/local/bin/openssl req -new -sha256 -engine pkcs11 -keyform engine -key 318488403:10 -out csr/intermediate.csr.pem

```

* Sign the CSR with rootCA
```
sudo LD_LIBRARY_PATH=/usr/local/lib/ OPENSSL_CONF=openssl.cnf /usr/local/bin/openssl x509 -req  -CAkeyform engine -engine pkcs11 -in ../IntermediateCA/csr/intermediate.csr.pem -CA ./CACERT.pem  -CAkey 1810989539:10 -days 3650 -sha256 -out ../IntermediateCA/certs/Intermediate.cert.pem -set_serial 1001
```

* Generate Server Cert Key and CSR the normal always

* Sign using key from pkcs11 store
```
sudo LD_LIBRARY_PATH=/usr/local/lib/ OPENSSL_CONF=openssl.cnf /usr/local/bin/openssl ca  -extensions server_cert -keyform engine -engine pkcs11 -keyfile 318488403:10 -in csr/server.csr -out certs/server.cert.pem

```
