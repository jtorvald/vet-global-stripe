package testdata
import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"github.com/stripe/stripe-go/fee"
	"github.com/stripe/stripe-go/paymentintent"
)

func invalid(){
	stripe.Key = "UNKNOWN"
	par := &stripe.ApplicationFeeParams{}
	fee.Get("test", par)

	piPar := &stripe.PaymentIntentParams{}
	paymentintent.New(piPar)

	chPar := &stripe.ChargeParams{}
	charge.New(chPar)
}