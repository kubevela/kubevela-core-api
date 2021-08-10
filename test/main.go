package main

import (
	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	core_oam_dev "github.com/oam-dev/kubevela-core-api/apis/core.oam.dev"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	core_v1beta1 "github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	kubevelaapistandard "github.com/oam-dev/kubevela-core-api/apis/standard.oam.dev/v1alpha1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"
)

var scheme = runtime.NewScheme()

func init() {
	_ = core_oam_dev.AddToScheme(scheme)
	_ = kubevelaapistandard.AddToScheme(scheme)
}

func main() {
	k8sClient, err := client.New(ctrl.GetConfigOrDie(), client.Options{Scheme: scheme})
	if err != nil {
		log.Fatal(err)
	}
	err = k8sClient.Create(context.Background(), &core_v1beta1.ComponentDefinition{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ComponentDefinition",
			APIVersion: "core.oam.dev/v1beta1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-comp",
			Namespace: "vela-system",
		},
		Spec: core_v1beta1.ComponentDefinitionSpec{
			Workload: common.WorkloadTypeDescriptor{
				Definition: common.WorkloadGVK{
					Kind:       "Deployment",
					APIVersion: "apps/v1",
				},
			},
			Schematic: &common.Schematic{
				CUE: &common.CUE{
					Template: webServiceTemplate,
				},
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	err = k8sClient.Create(context.Background(), &core_v1beta1.Application{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Application",
			APIVersion: "core.oam.dev/v1beta1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-app",
			Namespace: "default",
		},
		Spec: core_v1beta1.ApplicationSpec{
			Components: []common.ApplicationComponent{
				{
					Name: "web",
					Type: "webservice",
					Properties: util.Object2RawExtension(map[string]interface{}{
						"image": "nginx:1.14.0",
					}),
					Traits: []common.ApplicationTrait{
						{
							Type: "labels",
							Properties: util.Object2RawExtension(map[string]interface{}{
								"hello": "world",
							}),
						},
					},
				},
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
}

var webServiceTemplate = `output: {
	apiVersion: "apps/v1"
	kind:       "Deployment"
	metadata: labels: {
		"componentdefinition.oam.dev/version": "v1"
	}
	spec: {
		selector: matchLabels: {
			"app.oam.dev/component": context.name
		}
		template: {
			metadata: labels: {
				"app.oam.dev/component": context.name
			}
			spec: {
				containers: [{
					name:  context.name
					image: parameter.image
					if parameter["cmd"] != _|_ {
						command: parameter.cmd
					}
					if context["config"] != _|_ {
						env: context.config
					}
					ports: [{
						containerPort: parameter.port
					}]
				}]
		    }
		}
	}
}
parameter: {
	image: string
	cmd?: [...string]
	port: *80 | int
}
`
