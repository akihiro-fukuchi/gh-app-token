package main

import (
	"testing"
	"time"
)

var _appId = "123456"
var _installationId = "12345678"
var _privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEA7nc8Lb9+4g5L4BrDqHVcAjbygKRKxRbkqUWgx4y6eSG55n9b
6fztSNHZEyLNEOobFOK5b3NWv8gHqai2xTyRwMke5Fi7/uEa/l02mPW7ggu3jqug
rYLwYXf9WKo8957LHf1C+97Fw+L0bp7AAJRr0B8c6SwQTjRYfSbePme8AwPQETyL
HJbs901urcvzrmhmB3Vj9VFHF9IWuhxy5c49YzLO2Mx8Rb5NQyggsAVrRQC9rBon
lXkWYt1BbrKTgF/fJi7FzJSCYunGHIiqTlFEu09cPeaxKedKR2cvxsf3eGG447S6
biv2kySOizbLWrXzEtMiNohk+sw3sbb1wb8pTwIDAQABAoIBADcrSrbiyL4PxKoS
RxXgIOs6PxxX5hx20Wv/+dRw0GtZzCJxcJhPta3VLr1onby+DInmcjRAoN85rdwo
djBndOj/HrCBfuzWs2IJuqnkn/7UKyFMv52k32wNUIWEzRoUmLVVdHvE7EgHZ7l3
7L+1lsPNjui7EwKkxZwes+gII80mUFbeEKfTZiqz1bJwtaIk69wd0fECs551K6ga
0T1qVptVj3yfxgHABqJFGk6BsLJEsmk5KQc5v3kiXm+emf4n1kUBWUb9ed4Qa1k6
5+Lk45b/EgewYVF5PFq0+tbinNVc01xKEgJtAFbil/cMOJorcZpD/3fXvGduH3hs
qZwws7ECgYEA+PWt4urPjS1KqlXpJVWmj5mIzu9ey+Ge+nxgcR7CI7nRdz2AvX07
KQ1UI2ZlEtH1yXk53rvnGxwRWd0r/SSwkZ36U1nhHecsrmET1VPBqIrXy3EedVRQ
xIoNSgWqSDpZ7cbrM6st82lBiLbahT+NbCkCiAd/3UFg1NMlv3r4qG0CgYEA9TWW
BtDlSxZT8KSinFIptblNF4ZkZ/GrdbTctL9SRWXFyj990D/q5+2inCRP42hzHmCJ
3cv7bd9VAzD/WbsM3ln7aDXe5+KGiFKiXfKcvOm+kFB5ECPjWVW5+hXZx9OHqbrv
cYvuqyF2cThHV9yeNuucgHRoU96oSiQ6zCk4+ysCgYBd0HaOI93CXWbdeTI6F2SE
iF69XAZk3ciCq4vMFMMjo0oDnPF+dkps1dD25gcAaI4uNbhQ7o3P8Wu4aVfCNKk0
tks2TZA/LHXx4DMRGFbJpEhdKWtI21T0OvF3C1t3jEWHDIZlGgRezTMcyYre22v3
bhy+FdVhEtniWQ7IcRZyoQKBgFcnysF1cmpz7zXzbpDda1HaIRqhfAKuFWFq/Z+I
+TcNa3Xth0yDy3zQLCIPjg2oTHKZoaciH6X34YGW4swD/hjyJrftneMR0vuVU3zN
BQTomAE3eTBRcTeJjubi6VtrRib/+KeFMznEVRL9C+6gzeN7b08BESvuUia4JeX3
KPetAoGAHyUhRcCaiZnAFJIqTvkuz/5iDG37N53siFnqx2bIhGcopqoUPbJy5Ea+
anervelnYov4OhUKzF6v7DLMeqeMLroMh1B4EihDGKc9vxK8GjDCDLP4EsLTlr7j
zm+MyekFOPgws5cTeLfCJmrVlX7YDN4wHmw19QMaYUzyqZZQLoA=
-----END RSA PRIVATE KEY-----`
var _now = time.Date(2014, time.December, 31, 12, 13, 24, 0, time.UTC)

func TestNewAppConfigOk(t *testing.T) {
	_, err := NewAppConfig(_appId, _installationId, _privateKey, _now)

	if err != nil {
		t.Fatalf("Error result got: error: %+v, want: ok", err)
	}
}

func TestNewAppConfigInvalidPrivateKey(t *testing.T) {
	_, err := NewAppConfig(_appId, _installationId, "wrong-private-key", _now)

	if err == nil {
		t.Fatal("Error result got: ok, want: error")
	}
}
