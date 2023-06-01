package global

import (
	"gorm.io/gorm"
	"k8s.io/client-go/kubernetes"
)

var (
	GormDb    *gorm.DB
	K8sclient *kubernetes.Clientset
)
