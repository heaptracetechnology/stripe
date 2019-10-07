# _Stripe_ OMG Microservice

[![Open Microservice Guide](https://img.shields.io/badge/OMG%20Enabled-👍-green.svg?)](https://microservice.guide)
[![Build Status](https://travis-ci.com/heaptracetechnology/microservice-stripe.svg?branch=master)](https://travis-ci.com/heaptracetechnology/microservice-stripe)
[![codecov](https://codecov.io/gh/heaptracetechnology/microservice-stripe/branch/master/graph/badge.svg)](https://codecov.io/gh/heaptracetechnology/microservice-stripe)

An OMG service for Stripe, it allows individuals and businesses to receive payments over the Internet. Stripe provides the technical, fraud prevention, and banking infrastructure required to operate on-line payment systems.

## Usage in [Storyscript](https://storyscript.io/)

##### Create Charge
```coffee
stripe createCharge amount:'amount' currency:'currency' description:'description' capture:'capture'
```

##### Capture Charge
```coffee
stripe captureCharge charge:'charge'
```

##### Get Balance
```coffee
stripe getBalance
```

##### Create Paymentintent
```coffee
stripe createPaymentIntent amount:'amount' customer:'customer' savePaymentMethod:'savePaymentMethod' captureMethod:'captureMethod' currency:'currency' paymentMethodTypes:'paymentMethodTypes'
```

##### Retrive Paymentintent
```coffee
stripe retrievePaymentIntent paymentIntentID:'paymentIntentID'
```

##### Update Paymentintent
```coffee
stripe updatePaymentIntent amount:'amount' paymentIntentID:'paymentIntentID' shipping:'shipping'
```

##### Capture Paymentintent
```coffee
stripe capturePaymentIntent paymentIntentID:'paymentIntentID' amountToCapture:'amountToCapture'
```

##### Cancel PaymentIntent
```coffee
stripe cancelPaymentIntent paymentIntentID:'paymentIntentID'
```

##### List PaymentIntent
```coffee
stripe listAllPaymentIntent
```

##### Create Customer
```coffee
stripe createCustomer description:'description' email:'email'
```

##### Create Card
```coffee
stripe createCard source:'source' customer:'customer' metadata:'metadata' token:'token'
```

##### Create Refund
```coffee
stripe createRefund charge:'charge'
```

Curious to [learn more](https://docs.storyscript.io/)?

✨🍰✨

## Usage with [OMG CLI](https://www.npmjs.com/package/omg)
##### Create Charge
```shell
$ omg run createCharge -a amount=<AMOUNT> -a currency=<CURRENCY> -a description=<DESCRIPTION> -a capture=<TRUE/FALSE> -e SECRET_KEY=<SECRET_KEY>
```
##### Capture Charge
```shell
$ omg run captureCharge -a charge=<CHARGE_ID> -e SECRET_KEY=<SECRET_KEY>
```
##### Get Balance
```shell
$ omg run getBalance -e SECRET_KEY=<SECRET_KEY>
```
##### Create Paymentintent
```shell
$ omg run createPaymentIntent -a amount=<AMOUNT> -a customer=<CUSTOMER_ID> -a savePaymentMethod=<TRUE/FALSE> -a captureMethod=<AUTOMATIC/MANUAL> -a currency=<CURRENCY> -a paymentMethodTypes=<METHOD_LIST> -e SECRET_KEY=<SECRET_KEY>
```
##### Retrive Paymentintent
```shell
$ omg run retrievePaymentIntent -a paymentIntentID=<PAYMENTINTENT_ID> -e SECRET_KEY=<SECRET_KEY>
```

##### Update Paymentintent
```shell
$ omg run updatePaymentIntent -a amount=<AMOUNT> -a paymentIntentID=<PAYMENTINTENT_ID> -a shipping=<SHIPPING_OBJECT> -e SECRET_KEY=<SECRET_KEY>
```
##### Capture Paymentintent
```shell
$ omg run capturePaymentIntent -a paymentIntentID=<PAYMENTINTENT_ID> -a amountToCapture=<AMOUNT_TO_CAPTURE> -e SECRET_KEY=<SECRET_KEY>
```
##### Cancel PaymentIntent
```shell
$ omg run cancelPaymentIntent -a paymentIntentID=<PAYMENTINTENT_ID> -e SECRET_KEY=<SECRET_KEY>
```
##### List PaymentIntent
```shell
$ omg run listAllPaymentIntent -e SECRET_KEY=<SECRET_KEY>
```
##### Create Customer
```shell
$ omg run createCustomer -a description=<DESCRIPTION> -a email=<CUSTOMER_EMAIL> -e SECRET_KEY=<SECRET_KEY>
```
##### Create Card
```shell
$ omg run createCard -a source=<SOURCE_DETAILS> -a customer=<CUSTOMER_ID> -a metadata=<METADATA> -a token=<CARD_TOKEN>  -e SECRET_KEY=<SECRET_KEY>
```
##### Create Refund
```shell
$ omg run createRefund -a charge=<CHARGE_ID> -e SECRET_KEY=<SECRET_KEY>
```

## License
[MIT License](https://github.com/omg-services/stripe/blob/master/LICENSE).
