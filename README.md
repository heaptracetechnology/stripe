# Stripe as a microservice
An OMG service for Stripe, it allows individuals and businesses to receive payments over the Internet. Stripe provides the technical, fraud prevention, and banking infrastructure required to operate on-line payment systems.

[![Open Microservice Guide](https://img.shields.io/badge/OMG-enabled-brightgreen.svg?style=for-the-badge)](https://microservice.guide)

## [OMG](hhttps://microservice.guide) CLI

### OMG

* omg validate
```
omg validate
```
* omg build
```
omg build
```
### Test Service

* Test the service by following OMG commands

### CLI

##### Create Charge
```sh
$ omg run createcharge -a amount=<AMOUNT> -a currency=<CURRENCY> -a description=<DESCRIPTION> -e SECRET_KEY=<SECRET_KEY>
```
##### Get Balance
```sh
$ omg run getbalance -e SECRET_KEY=<SECRET_KEY>
```
##### Create Paymentintent
```sh
$ omg run createpaymentintent -a amount=<AMOUNT> -a currency=<CURRENCY> -e SECRET_KEY=<SECRET_KEY>
```
## License
### [MIT](https://choosealicense.com/licenses/mit/)

## Docker
### Build
```
docker build -t microservice-stripe .
```
### RUN
```
docker run -p 3000:3000 microservice-stripe
```