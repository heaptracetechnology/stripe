package route

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/heaptracetechnology/microservice-stripe/balanceoperation"
	"github.com/heaptracetechnology/microservice-stripe/cardoperation"
	"github.com/heaptracetechnology/microservice-stripe/chargeoperation"
	"github.com/heaptracetechnology/microservice-stripe/customeroperation"
	"github.com/heaptracetechnology/microservice-stripe/paymentintentoperation"
	"github.com/heaptracetechnology/microservice-stripe/refundoperation"
	"github.com/heaptracetechnology/microservice-stripe/sourceoperation"
	"github.com/heaptracetechnology/microservice-stripe/tokenoperation"
)

//Route struct
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes stripe
type Routes []Route

var routes = Routes{
	Route{
		"CreateCharge",
		"POST",
		"/createCharge",
		chargeoperation.CreateCharge,
	},
	Route{
		"GetBalance",
		"GET",
		"/getBalance",
		balanceoperation.GetBalance,
	},
	Route{
		"CreatePaymentIntent",
		"POST",
		"/createPaymentIntent",
		paymentintentoperation.CreatePaymentIntent,
	},
	Route{
		"RetrievePaymentIntent",
		"GET",
		"/retrievePaymentIntent/{paymentIntentID}",
		paymentintentoperation.RetrievePaymentIntent,
	},
	Route{
		"UpdatePaymentIntent",
		"PUT",
		"/updatePaymentIntent/{paymentIntentID}",
		paymentintentoperation.UpdatePaymentIntent,
	},
	Route{
		"CapturePaymentIntent",
		"POST",
		"/capturePaymentIntent/{paymentIntentID}",
		paymentintentoperation.CapturePaymentIntent,
	},
	// Route{
	// 	"ConfirmPaymentIntent",
	// 	"POST",
	// 	"/confirmPaymentIntent/{paymentIntentID}",
	// 	paymentintentoperation.ConfirmPaymentIntent,
	// },
	Route{
		"CancelPaymentIntent",
		"POST",
		"/cancelPaymentIntent/{paymentIntentID}",
		paymentintentoperation.CancelPaymentIntent,
	},
	Route{
		"ListAllPaymentIntent",
		"GET",
		"/listAllPaymentIntent",
		paymentintentoperation.ListAllPaymentIntent,
	},
	Route{
		"CreateSource",
		"POST",
		"/createSource/{transferType}",
		sourceoperation.CreateSource,
	},
	Route{
		"CreateCustomer",
		"POST",
		"/createCustomer",
		customeroperation.CreateCustomer,
	},
	Route{
		"CreateCard",
		"POST",
		"/createCard",
		cardoperation.CreateCard,
	},
	Route{
		"CreateToken",
		"POST",
		"/createToken",
		tokenoperation.CreateToken,
	},
	Route{
		"CreateRefund",
		"POST",
		"/createRefund",
		refundoperation.CreateRefund,
	},
	Route{
		"CaptureCharge",
		"POST",
		"/captureCharge/{charge}",
		chargeoperation.CaptureCharge,
	},
	Route{
		"HealthCheck",
		"GET",
		"/health",
		chargeoperation.HealthCheck,
	},
}

//NewRouter stripe
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
