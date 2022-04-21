package payment

import (
	"github.com/Flutterwave/Rave-go/rave"
	"os"
)

//initialize API keys
var r = rave.Rave{
	false,
	os.Getenv("FLW_PUBLIC_KEY"),
	os.Getenv("FLW_SECRETE_KEY"),
}

var transfer = rave.Transfer{r}

//payload details
var details = rave.SinglePaymentData{}
