package k64dec

import (
	"fmt"

	"github.com/pterm/pterm"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes/scheme"
)

// decode the kubernetes secret from input data
func decode(data []byte) (*corev1.Secret, error) {
	decoder := serializer.NewCodecFactory(scheme.Scheme).UniversalDecoder()
	object := &corev1.Secret{}
	if err := runtime.DecodeInto(decoder, data, object); err != nil {
		return nil, fmt.Errorf("failed to decode secret: %s", err)
	}
	return object, nil
}

// PrintDecodedSecret print decoded the content of secret data
func PrintDecodedSecret(data []byte) error {
	secret, err := decode(data)
	if err != nil {
		return err
	}

	if len(secret.Data) > 0 {
		for k, v := range secret.Data {
			print(k, string(v))
		}
	}

	if len(secret.StringData) > 0 {
		for k, v := range secret.StringData {
			print(k, v)
		}
	}

	return nil
}

// print the key, value to console
func print(k, v string) {
	pterm.Underscore.Printfln(k)
	pterm.Italic.Printf("%s", v)
	fmt.Println()
}
