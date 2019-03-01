package route

import (
    "github.com/gorilla/mux"
    "github.com/user/stripe/BalanceOperation"
    "github.com/user/stripe/CardOperation"
    "github.com/user/stripe/ChargeOperation"
    "github.com/user/stripe/CustomerOperation"
    "github.com/user/stripe/PaymentIntentOperation"
    "github.com/user/stripe/RefundOperation"
    "github.com/user/stripe/SourceOperation"
    "github.com/user/stripe/TokenOperation"
    "log"
    "net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "CreateCharge",
        "POST",
        "/createcharge",
        ChargeOperation.CreateCharge,
    },
    Route{
        "GetBalance",
        "GET",
        "/getbalance",
        BalanceOperation.GetBalance,
    },
    Route{
        "CreatePaymentIntent",
        "POST",
        "/createpaymentintent",
        PaymentIntentOperation.CreatePaymentIntent,
    },
    Route{
        "RetrievePaymentIntent",
        "GET",
        "/retrievepaymentintent/{paymentintentid}",
        PaymentIntentOperation.RetrievePaymentIntent,
    },
    Route{
        "UpdatePaymentIntent",
        "PUT",
        "/updatepaymentintent/{paymentintentid}",
        PaymentIntentOperation.UpdatePaymentIntent,
    },
    Route{
        "ConfirmPaymentIntent",
        "POST",
        "/confirmpaymentintent",
        PaymentIntentOperation.ConfirmPaymentIntent,
    },
    Route{
        "CapturePaymentIntent",
        "POST",
        "/capturepaymentintent/{paymentintentid}",
        PaymentIntentOperation.CapturePaymentIntent,
    },
    Route{
        "CancelPaymentIntent",
        "POST",
        "/cancelpaymentintent/{paymentintentid}",
        PaymentIntentOperation.CancelPaymentIntent,
    },
    Route{
        "ListAllPaymentIntent",
        "GET",
        "/listallpaymentintent",
        PaymentIntentOperation.ListAllPaymentIntent,
    },
    Route{
        "CreateSource",
        "POST",
        "/createsource/{transfertype}",
        SourceOperation.CreateSource,
    },
    Route{
        "CreateCustomer",
        "POST",
        "/createcustomer",
        CustomerOperation.CreateCustomer,
    },
    Route{
        "CreateCard",
        "POST",
        "/createcard",
        CardOperation.CreateCard,
    },
    Route{
        "CreateToken",
        "POST",
        "/createtoken",
        TokenOperation.CreateToken,
    },
    Route{
        "CreateRefund",
        "POST",
        "/createrefund",
        RefundOperation.CreateRefund,
    },
    Route{
        "CaptureCharge",
        "POST",
        "/capturecharge/{charge}",
        ChargeOperation.CaptureCharge,
    },
}

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        var handler http.Handler
        log.Println(route.Name)
        handler = route.HandlerFunc

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }
    return router
}
