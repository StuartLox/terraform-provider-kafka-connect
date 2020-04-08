package connect

import (
	"log"
	"crypto/tls"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	kc "github.com/StuartLox/go-kafka-connect/lib/connectors"
)

func Provider() terraform.ResourceProvider {
	log.Printf("[INFO] Creating Provider")
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("KAFKA_CONNECT_URL", ""),
			},
			"client_cert": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLIENT_CERT", ""),
			},
			"client_key": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLIENT_KEY", ""),
			},
		},
		ConfigureFunc: providerConfigure,
		ResourcesMap: map[string]*schema.Resource{
			"kafka-connect_connector": kafkaConnectorResource(),
		},
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	log.Printf("[INFO] Initializing KafkaConnect client")
	addr := d.Get("url").(string)
	certFile := d.Get("client_cert").(string)
	keyFile := d.Get("client_key").(string)

	c := kc.NewClient(addr)
	if len(certFile) > 0 && len(keyFile) > 0 {
		cert, err := tls.LoadX509KeyPair(certFile, keyFile)
		if err != nil {
			return nil, err
		}
		c.SetClientCertificates(cert)
	}

	return c, nil
}
