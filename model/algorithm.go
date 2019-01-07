package model

import jwt "github.com/dgrijalva/jwt-go"

// Algorithms are all the supported algorithms for signing a JWT
var Algorithms = []string{
	"none",
	"es256",
	"es384",
	"es512",
	"hs256",
	"hs384",
	"hs512",
	"ps256",
	"ps384",
	"ps512",
	"rs256",
	"rs384",
	"rs512",
}

// AlgorithmForName returns the algorithm for the specified name
func AlgorithmForName(name string) (jwt.SigningMethod, bool) {
	val, ok := aLookup[name]
	return val, ok
}

var aLookup = map[string]jwt.SigningMethod{
	Algorithms[0]:  jwt.SigningMethodNone,
	Algorithms[1]:  jwt.SigningMethodES256,
	Algorithms[2]:  jwt.SigningMethodES384,
	Algorithms[3]:  jwt.SigningMethodES512,
	Algorithms[4]:  jwt.SigningMethodHS256,
	Algorithms[5]:  jwt.SigningMethodHS384,
	Algorithms[6]:  jwt.SigningMethodHS512,
	Algorithms[7]:  jwt.SigningMethodPS256,
	Algorithms[8]:  jwt.SigningMethodPS384,
	Algorithms[9]:  jwt.SigningMethodPS512,
	Algorithms[10]: jwt.SigningMethodRS256,
	Algorithms[11]: jwt.SigningMethodRS384,
	Algorithms[12]: jwt.SigningMethodRS512,
}
