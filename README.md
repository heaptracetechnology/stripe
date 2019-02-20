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
$ omg run create_charge -a amount=<AMOUNT> -a currency=<CURRENCY> -a description=<DESCRIPTION> -e SECRET_KEY=<SECRET_KEY>
```
##### Get Balance
```sh
$ omg run get_balance -e SECRET_KEY=<SECRET_KEY>
```
##### Create Paymentintent
```sh
$ omg run create_payment_intent -a amount=<AMOUNT> -a currency=<CURRENCY> -e SECRET_KEY=<SECRET_KEY>
```
##### Retrive Paymentintent
```sh
omg run retrieve_payment_intent -a paymentintent_id=<PAYMENTINTENT_ID> -e SECRET_KEY=<SECRET_KEY>
```
##### Update Paymentintent
```sh
omg run update_payment_intent -a amount=<AMOUNT> -a paymentintent_id=<PAYMENTINTENT_ID> -a shipping=<SHIPPING_OBJECT> -e SECRET_KEY=<SECRET_KEY>
```
##### Cancel PaymentIntent
```sh
omg run cancel_payment_intent -a paymentintent_id=<PAYMENTINTENT_ID> -e SECRET_KEY=<SECRET_KEY>
```
##### List PaymentIntent
```sh
omg run list_all_payment_intent -e SECRET_KEY=<SECRET_KEY>
```
##### Create Source
```sh
omg run create_source -a transfertype=<TRANSFER_TYPE> -a currency=<CURRENCY> -a owner=<OWNER_OBJECT> -e SECRET_KEY=<SECRET_KEY>
```
##### Create Customer
```sh
omg run create_customer -a description=<DESCRIPTION> -a email=<CUSTOMER_EMAIL> -e SECRET_KEY=<SECRET_KEY>
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