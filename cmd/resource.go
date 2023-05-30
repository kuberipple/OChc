package cmd

import (
	"context"
	"fmt"
	"github.com/ZhangSetSail/OChc/pkg"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var resourceLong = `helm resource ------`

func newResourceCmd(ochc *pkg.OChcModel) *cobra.Command {
	return &cobra.Command{
		Use:   "resource ",
		Short: "generate a helm chart package with one click through k8s cluster resources",
		Long:  resourceLong,
		Args:  pkg.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			pods, err := ochc.ClientSet.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})
			if err != nil {
				logrus.Errorf("get pods failure: %v", err)
			}
			for _, pod := range pods.Items {
				fmt.Println(pod.Name)
			}
		},
	}
}
