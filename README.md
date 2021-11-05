# Vet Global Stripe function calls

While implementing the [stripe-go](https://github.com/stripe/stripe-go) library I needed to implement it using multiple API keys.
When you want to use the library with multiple API keys you need to make sure that you use the client 
instead of the helper functions.

> If you're dealing with multiple keys, it is recommended you use client.API. This allows you to create as many clients as needed, each with their own individual key.
>
> -- <cite>[Stripe: With a Client][1]</cite>

[1]: https://github.com/stripe/stripe-go#with-a-client

## Disclaimer

This project is a quick and dirty checker to see if helper functions are used.
- The implementation is far from perfect.
- The helper functions are added as a list of strings, and it does not analyze the stripe library for calls to getC()
  for the default client. Which it should, but I've no idea how to write that logic. 
  **If you do: [feel free to let me know](https://github.com/jtorvald/vet-global-stripe/issues/new).** 

## Installation
```
go install
```

## Usage
```
go vet -vettool=$(which vet-global-stripe) package
```

### Report
```
$ go vet -vettool=$(which vet-global-stripe) package
# yourpackage
main.go:50:19: don't use global stripe function: paymentintent.New
```

## LICENSE

MIT Copyright (c) 2021 Jørgen Teunis
