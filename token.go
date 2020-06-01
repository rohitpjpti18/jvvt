package jvvt

type Token struct{
	raw				string
	header			map[string]string
	payload			map[string]string
}