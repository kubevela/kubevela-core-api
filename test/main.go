package main

import (
	"context"
	"log"

	core_oam_dev "github.com/oam-dev/kubevela-core-api/apis/core.oam.dev"
	core_v1alpha2 "github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1alpha2"
	kubevelaapistandard "github.com/oam-dev/kubevela-core-api/apis/standard.oam.dev/v1alpha1"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
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
	err = k8sClient.Create(context.Background(), &core_v1alpha2.Component{
		ObjectMeta: metav1.ObjectMeta{Name: "test", Namespace: "default"},
		Spec: core_v1alpha2.ComponentSpec{Workload: runtime.RawExtension{
			Object: &appsv1.Deployment{
				TypeMeta: metav1.TypeMeta{
					Kind:       "Deployment",
					APIVersion: "apps/v1",
				},
				Spec: appsv1.DeploymentSpec{
					Selector: &metav1.LabelSelector{
						MatchLabels: map[string]string{"app": "test"},
					},
					Template: v1.PodTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Labels: map[string]string{"app": "test"},
						},
						Spec: v1.PodSpec{
							Containers: []v1.Container{
								{Name: "image1", Image: "nginx"},
							},
						},
					},
				}},
		},
		}})
	if err != nil {
		log.Fatal(err)
	}
	err = k8sClient.Create(context.Background(), &core_v1alpha2.ApplicationConfiguration{
		ObjectMeta: metav1.ObjectMeta{Name: "test-app", Namespace: "default"},
		Spec:       core_v1alpha2.ApplicationConfigurationSpec{Components: []core_v1alpha2.ApplicationConfigurationComponent{{ComponentName: "test"}}}})
	if err != nil {
		log.Fatal(err)
	}
}
