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

type UpdateNodePoolsAmazon struct {
	InstanceType string `json:"instanceType,omitempty"`
	SpotPrice    string `json:"spotPrice,omitempty"`
	Autoscaling  bool   `json:"autoscaling,omitempty"`
	Count        int32  `json:"count,omitempty"`
	MinCount     int32  `json:"minCount,omitempty"`
	MaxCount     int32  `json:"maxCount,omitempty"`
	Image        string `json:"image,omitempty"`
}
