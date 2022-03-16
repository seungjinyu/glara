package glarautils

import (
	"testing"

	"github.com/seungjinyu/glara/settings"
	"k8s.io/client-go/kubernetes"
)

func TestDeleteStatefulSetPod(t *testing.T) {
	type args struct {
		namespace          string
		StatefulSetPodName string
		clientset          *kubernetes.Clientset
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteStatefulSetPod(tt.args.namespace, tt.args.StatefulSetPodName, tt.args.clientset); (err != nil) != tt.wantErr {
				t.Errorf("DeleteStatefulSetPod() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteDaemonSetPod(t *testing.T) {
	type args struct {
		namespace        string
		DaemonSetPodName string
		clientset        *kubernetes.Clientset
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteDaemonSetPod(tt.args.namespace, tt.args.DaemonSetPodName, tt.args.clientset); (err != nil) != tt.wantErr {
				t.Errorf("DeleteDaemonSetPod() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteReplicaSetPod(t *testing.T) {
	type args struct {
		namespace         string
		ReplicaSetPodName string
		clientset         *kubernetes.Clientset
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteReplicaSetPod(tt.args.namespace, tt.args.ReplicaSetPodName, tt.args.clientset); (err != nil) != tt.wantErr {
				t.Errorf("DeleteReplicaSetPod() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInspectPod(t *testing.T) {
	type args struct {
		namespace string
		pod       string
		rStr      string
		kubecli   settings.ClientSetInstance
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InspectPod(tt.args.namespace, tt.args.pod, tt.args.rStr, tt.args.kubecli); (err != nil) != tt.wantErr {
				t.Errorf("InspectPod() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
