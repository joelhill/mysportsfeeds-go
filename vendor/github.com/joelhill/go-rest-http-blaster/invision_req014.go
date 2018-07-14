package blaster

// req014HeaderCheck will check for the presence of required outgoing
type req014HeaderCheck struct {
	requestIDOK      bool
	requestSourceOK  bool
	callingServiceOK bool
}

// check to see if REQ014 flags are closed
func (c req014HeaderCheck) ok() bool {
	return c.requestIDOK && c.requestSourceOK && c.callingServiceOK
}
