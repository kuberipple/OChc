package cmd

import (
	"github.com/ZhangSetSail/OChc/pkg"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var cmdLong = `ochc: 这一个测试项目，致力于实现一键导出 helm chart 包，感谢使用。`
var oChc *pkg.OChcModel

func newLogrus() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func NewCMD() *cobra.Command {
	var oCmd = &cobra.Command{
		Use:   "ochc",
		Short: "test",
		Long:  cmdLong,
	}
	oCmd.AddCommand(
		newResourceCmd(oChc),
	)
	return oCmd
}

func NewOChc() {
	oChc = &pkg.OChcModel{}
}

func init() {
	newLogrus()
	NewOChc()
	pkg.SetKubeClient(oChc)
}
