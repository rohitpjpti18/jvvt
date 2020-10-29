package jvvt

import (
	"testing"
)

func TestSplitToken1(t *testing.T){
	token := "eyJhbGciOiJub25lIn0.eyJzdWIiOiJ1c2VyMTIzIiwic2Vzc2lvbiI6ImNoNzJnc2IzMjAwMDB1ZG9jbDM2M2VvZnkiLCJuYW1lIjoiUHJldHR5IE5hbWUiLCJsYXN0cGFnZSI6Ii92aWV3cy9zZXR0aW5ncyJ9."
	result := splitToken(token);
	if result[0] != "eyJhbGciOiJub25lIn0"{
		t.Errorf("returned = %s, expected something else", result[0])
	}
}

func TestSplitToken2(t *testing.T){
	token := "eyJhbGciOiJub25lIn0.eyJzdWIiOiJ1c2VyMTIzIiwic2Vzc2lvbiI6ImNoNzJnc2IzMjAwMDB1ZG9jbDM2M2VvZnkiLCJuYW1lIjoiUHJldHR5IE5hbWUiLCJsYXN0cGFnZSI6Ii92aWV3cy9zZXR0aW5ncyJ9."
	result := splitToken(token);
	if result[1] != "eyJzdWIiOiJ1c2VyMTIzIiwic2Vzc2lvbiI6ImNoNzJnc2IzMjAwMDB1ZG9jbDM2M2VvZnkiLCJuYW1lIjoiUHJldHR5IE5hbWUiLCJsYXN0cGFnZSI6Ii92aWV3cy9zZXR0aW5ncyJ9"{
		t.Errorf("returned = %s, expected something else", result[1])
	}
}