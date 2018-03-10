# gopriceoptions ![build status](https://travis-ci.org/jasonmerecki/gopriceoptions.svg?branch=master)
Calculations for theoretical option pricing, using Go. To my knowledge, this is the first self-contained open source implementation of Black-Scholes and Bjerksund-Stensland option pricing models, plus Greek calculations, in Go. 

The implementation intentionally avoids structs; I didn't want any garbage collection, so users of this software are welcome to wrap it in a struct if they find it helpful.
