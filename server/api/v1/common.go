package v1

import (
	"fmt"
	"gin-vue-devops/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ClusterID(c *gin.Context) (string, error){
	clusterID := c.DefaultQuery("clusterID", "1")
	clusterIDuint64, err := strconv.ParseUint(clusterID, 10, 32)
	clusterIDuint := uint(clusterIDuint64)
	err, cluster := service.GetK8sCluster(clusterIDuint)
	if err != nil {
		fmt.Println(err)
	}
	return cluster.KubeConfig, nil
}
