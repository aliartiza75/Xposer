package constants

import (
	"time"
)

const (
	KUBERNETES         = "kubernetes"
	OPENSHIFT          = "openshift"
	ALL_NAMESPACES     = ""
	SERVICES           = "services"
	DOMAIN             = "Domain"
	CERT               = "-cert"
	RESYNC_PERIOD      = 10 * time.Second
	XPOSER_CONFIGMAP   = "xposer"
	EXPOSE_INGRESS_URL = "exposeIngressUrl"
	LOCALLY            = "locally"
	GLOBALLY           = "globally"
	EXPOSE             = "expose"
)
