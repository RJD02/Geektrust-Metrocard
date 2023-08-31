package main

type locationTypes string

const (
	AIRPORT locationTypes = "AIRPORT"
	CENTRAL               = "CENTRAL"
)

type MetroCard struct {
	Amount       int
	CurrLocation locationTypes
	CitizentType citizenType
	count        int
}

type citizenType string

const UNDEFINED_STRING = ""

const (
	ADULT          citizenType = "ADULT"
	KID                        = "KID"
	SENIOR_CITIZEN             = "SENIOR_CITIZEN"
)
