/*
 * Pipeline API
 *
 * Pipeline v0.3.0 swagger
 *
 * API version: 0.3.0
 * Contact: info@banzaicloud.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

type GenTlsForLogging struct {
	TlsEnabled       bool   `json:"tlsEnabled"`
	GenTlsSecretName string `json:"genTlsSecretName,omitempty"`
	Namespace        string `json:"namespace,omitempty"`
	TlsHost          string `json:"tlsHost,omitempty"`
}